package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"

	informerv1 "github.com/couchbase/service-broker/generated/informers/externalversions/servicebroker/v1alpha1"
	v1 "github.com/couchbase/service-broker/pkg/apis/servicebroker/v1alpha1"
	"github.com/couchbase/service-broker/pkg/client"
	"github.com/couchbase/service-broker/pkg/log"
	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

const (
	// ConfigurationName is the configuration resource name.
	ConfigurationName = "couchbase-service-broker"
)

type configuration struct {
	// clients is the set of clients this instance of the broker uses, by default
	// this will use in-cluster Kubernetes, however may be replaced by fake clients
	// by a test framework.
	clients client.Clients

	// config is the user supplied configuration custom resource.
	config *v1.ServiceBrokerConfig

	// token is the API access token.
	token string

	// namespace is the default namespace the broker is running in.
	namespace string

	// lock is used to remove races around the use of the context.
	// The context can be read by many, but can only be written
	// by one when there are no readers.
	// Updates must appear atomic so handlers should hold the read
	// lock while processing a request.
	lock sync.RWMutex
}

// c is the global configuration struct.
var c *configuration

// createHandler add the service broker configuration when the underlying
// resource is created.
func createHandler(obj interface{}) {
	brokerConfiguration, ok := obj.(*v1.ServiceBrokerConfig)
	if !ok {
		glog.Error("unexpected object type in config add")
		return
	}

	if brokerConfiguration.Name != ConfigurationName {
		glog.V(log.LevelDebug).Info("unexpected object name in config delete:", brokerConfiguration.Name)
		return
	}

	if err := updateStatus(brokerConfiguration); err != nil {
		glog.Info("service broker configuration invalid, see resource status for details")
		glog.V(1).Info(err)

		c.lock.Lock()
		defer c.lock.Unlock()

		c.config = nil

		return
	}

	glog.Info("service broker configuration created, service ready")

	if glog.V(1) {
		object, err := json.Marshal(brokerConfiguration)
		if err == nil {
			glog.V(1).Info(string(object))
		}
	}

	c.lock.Lock()
	c.config = brokerConfiguration
	c.lock.Unlock()
}

// updateHandler modifies the service broker configuration when the underlying
// resource updates.
func updateHandler(oldObj, newObj interface{}) {
	brokerConfiguration, ok := newObj.(*v1.ServiceBrokerConfig)
	if !ok {
		glog.Error("unexpected object type in config update")
		return
	}

	if brokerConfiguration.Name != ConfigurationName {
		glog.V(log.LevelDebug).Info("unexpected object name in config update:", brokerConfiguration.Name)
		return
	}

	if err := updateStatus(brokerConfiguration); err != nil {
		glog.Info("service broker configuration invalid, see resource status for details")
		glog.V(1).Info(err)

		c.lock.Lock()
		defer c.lock.Unlock()

		c.config = nil

		return
	}

	glog.Info("service broker configuration updated")

	if glog.V(1) {
		object, err := json.Marshal(brokerConfiguration)
		if err == nil {
			glog.V(1).Info(string(object))
		}
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	c.config = brokerConfiguration
}

// deleteHandler deletes the service broker configuration when the underlying
// resource is deleted.
func deleteHandler(obj interface{}) {
	brokerConfiguration, ok := obj.(*v1.ServiceBrokerConfig)
	if !ok {
		glog.Error("unexpected object type in config delete")
		return
	}

	if brokerConfiguration.Name != ConfigurationName {
		glog.V(log.LevelDebug).Info("unexpected object name in config delete:", brokerConfiguration.Name)
		return
	}

	glog.Info("service broker configuration deleted, service unready")

	c.lock.Lock()
	c.config = nil
	c.lock.Unlock()
}

// Configure initializes global configuration and must be called before starting
// the API service.
func Configure(clients client.Clients, namespace, token string) error {
	glog.Info("configuring service broker")

	// Create the global configuration structure.
	c = &configuration{
		clients:   clients,
		token:     token,
		namespace: namespace,
	}

	handlers := &cache.ResourceEventHandlerFuncs{
		AddFunc:    createHandler,
		UpdateFunc: updateHandler,
		DeleteFunc: deleteHandler,
	}

	informer := informerv1.NewServiceBrokerConfigInformer(clients.Broker(), namespace, time.Minute, nil)
	informer.AddEventHandler(handlers)

	stop := make(chan struct{})

	go informer.Run(stop)

	if !cache.WaitForCacheSync(stop, informer.HasSynced) {
		return fmt.Errorf("service broker config shared informer failed to syncronize")
	}

	return nil
}

// Lock puts a read lock on the configuration during the lifetime
// of a request.
func Lock() {
	c.lock.RLock()
}

// Unlock releases the read lock on the configuration after a
// request has completed.
func Unlock() {
	c.lock.RUnlock()
}

// Clients returns a set of Kubernetes clients.
func Clients() client.Clients {
	return c.clients
}

// Config returns the user specified custom resource.
func Config() *v1.ServiceBrokerConfig {
	return c.config
}

// Token returns the API bearer token.
func Token() string {
	return c.token
}

// Namespace returns the broker namespace.
func Namespace() string {
	return c.namespace
}

func getBindingForServicePlan(config *v1.ServiceBrokerConfig, serviceName, planName string) *v1.ConfigurationBinding {
	for index, binding := range config.Spec.Bindings {
		if binding.Service == serviceName && binding.Plan == planName {
			return &config.Spec.Bindings[index]
		}
	}

	return nil
}

func getTemplateByName(config *v1.ServiceBrokerConfig, templateName string) *v1.ConfigurationTemplate {
	for index, template := range config.Spec.Templates {
		if template.Name == templateName {
			return &config.Spec.Templates[index]
		}
	}

	return nil
}

// updateStatus runs any analysis on the confiuration, makes and commits any modifications.
// In particular this allows the status to say you have made a configuration error.
// A returned error means don't accept the configuration, set to nil so the service broker
// reports unready and doesn't serve any API requests.
func updateStatus(config *v1.ServiceBrokerConfig) error {
	var rerr error

	// Assume the configuration is valid, then modify if an error
	// has occurred, finally retain the transition time if an existing
	// condition exists and it has the same status.
	validCondition := v1.ServiceBrokerConfigCondition{
		Type:   v1.ConfigurationValid,
		Status: v1.ConditionTrue,
		LastTransitionTime: metav1.Time{
			Time: time.Now(),
		},
		Reason: "ValidationSucceeded",
	}

	if err := validate(config); err != nil {
		validCondition.Status = v1.ConditionFalse
		validCondition.Reason = "ValidationFailed"
		validCondition.Message = err.Error()

		rerr = err
	}

	for _, condition := range config.Status.Conditions {
		if condition.Type == v1.ConfigurationValid {
			if condition.Status == validCondition.Status {
				validCondition.LastTransitionTime = condition.LastTransitionTime
			}

			break
		}
	}

	// Update the status if it has been modified.
	status := v1.ServiceBrokerConfigStatus{
		Conditions: []v1.ServiceBrokerConfigCondition{
			validCondition,
		},
	}

	if reflect.DeepEqual(config.Status, status) {
		return rerr
	}

	newConfig := config.DeepCopy()
	newConfig.Status = status

	if _, err := c.clients.Broker().ServicebrokerV1alpha1().ServiceBrokerConfigs(c.namespace).Update(newConfig); err != nil {
		glog.Infof("failed to update service broker configuration status: %v", err)
		return rerr
	}

	return rerr
}

// validate does any validation that cannot be performed by the JSON schema
// included in the CRD.
func validate(config *v1.ServiceBrokerConfig) error {
	// Check that service offerings and plans are bound properly to configuration.
	for _, service := range config.Spec.Catalog.Services {
		for _, plan := range service.Plans {
			// Each service plan must have a service binding.
			binding := getBindingForServicePlan(config, service.Name, plan.Name)
			if binding == nil {
				return fmt.Errorf("service plan %s for offering %s does not have a configuration binding", plan.Name, service.Name)
			}

			// Only bindable service plans may have templates for bindings.
			bindable := service.Bindable
			if plan.Bindable != nil {
				bindable = *plan.Bindable
			}

			if !bindable && binding.ServiceBinding != nil {
				return fmt.Errorf("service plan %s for offering %s not bindable, but configuration binding %s defines service binding configuarion", plan.Name, service.Name, binding.Name)
			}

			if bindable && binding.ServiceBinding == nil {
				return fmt.Errorf("service plan %s for offering %s bindable, but configuration binding %s does not define service binding configuarion", plan.Name, service.Name, binding.Name)
			}
		}
	}

	// Check that configuration bindings are properly configured.
	for _, binding := range config.Spec.Bindings {
		// Bindings cannot do nothing.
		if len(binding.ServiceInstance.Parameters) == 0 && len(binding.ServiceInstance.Templates) == 0 {
			return fmt.Errorf("configuration binding %s does nothing for service instances", binding.Name)
		}

		if binding.ServiceBinding != nil {
			if len(binding.ServiceBinding.Parameters) == 0 && len(binding.ServiceBinding.Templates) == 0 {
				return fmt.Errorf("configuration binding %s does nothing for service bindings", binding.Name)
			}
		}

		// Binding templates must exist.
		for _, template := range binding.ServiceInstance.Templates {
			if getTemplateByName(config, template) == nil {
				return fmt.Errorf("template %s referenced by configuration %s service instance must exist", template, binding.Name)
			}
		}

		if binding.ServiceBinding != nil {
			for _, template := range binding.ServiceBinding.Templates {
				if getTemplateByName(config, template) == nil {
					return fmt.Errorf("template %s referenced by configuration %s service binding must exist", template, binding.Name)
				}
			}
		}
	}

	return nil
}
