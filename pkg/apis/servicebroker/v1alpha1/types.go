// Copyright 2020-2021 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file  except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the  License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint:godot
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	// labelBase is the root of all labels and annotations.
	labelBase = "servicebroker.couchbase.com"

	// VersionAnnotaiton records the broker version for upgrades.
	VersionAnnotaiton = labelBase + "/version"

	// ResourceAnnotation records the resource for updates.
	ResourceAnnotation = labelBase + "/resource"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=all;couchbase
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="valid",type="string",JSONPath=".status.conditions[?(@.type==\"ConfigurationValid\")].status",description="whether the configuration is valid"
// +kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"
type ServiceBrokerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceBrokerConfigSpec   `json:"spec"`
	Status            ServiceBrokerConfigStatus `json:"status,omitempty"`
}

// ServiceBrokerConfigSpec defines the top level service broker configuration
// data structure.
type ServiceBrokerConfigSpec struct {
	// Catalog is the Open Service Broker service catalog definition. More info:
	// https://github.com/couchbase/service-broker/tree/master/documentation/modules/ROOT/pages/concepts/catalog.adoc
	Catalog ServiceCatalog `json:"catalog"`

	// Templates is a set of resource templates that can be rendered by the service broker. More info:
	// https://github.com/couchbase/service-broker/tree/master/documentation/modules/ROOT/pages/concepts/templates.adoc
	// +listType=map
	// +listMapKey=name
	Templates []ConfigurationTemplate `json:"templates"`

	// Bindings is a set of bindings that link service plans to resource templates. More info:
	// https://github.com/couchbase/service-broker/tree/master/documentation/modules/ROOT/pages/concepts/bindings.adoc
	// +kubebuilder:validation:MinItems=1
	// +listType=map
	// +listMapKey=name
	Bindings []ConfigurationBinding `json:"bindings"`
}

// ServiceCatalog is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type ServiceCatalog struct {
	// Services is an array of Service Offering objects. More info:
	// https://github.com/couchbase/service-broker/tree/master/documentation/modules/ROOT/pages/concepts/catalog.adoc#service-offerings
	// +kubebuilder:validation:MinItems=1
	// +listType=map
	// +listMapKey=name
	Services []ServiceOffering `json:"services"`
}

// ServiceOffering is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type ServiceOffering struct {
	// Name is the name of the Service Offering. MUST be unique across all Service Offering
	// objects returned in this response. MUST be a non-empty string. Using a CLI-friendly name
	// is RECOMMENDED.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// ID is an identifier used to correlate this Service Offering in future requests to the
	// Service Broker. This MUST be globally unique such that Platforms (and their users) MUST
	// be able to assume that seeing the same value (no matter what Service Broker uses it) will
	// always refer to this Service Offering. MUST be a non-empty string. Using a GUID is RECOMMENDED.
	// +kubebuilder:validation:MinLength=1
	ID string `json:"id"`

	// Descriptions is a short description of the service. MUST be a non-empty string.
	// +kubebuilder:validation:MinLength=1
	Description string `json:"description"`

	// Tags provide a flexible mechanism to expose a classification, attribute, or base
	// technology of a service, enabling equivalent services to be swapped out without changes
	// to dependent logic in applications, buildpacks, or other services. E.g. mysql, relational,
	// redis, key-value, caching, messaging, amqp.
	// +listType=set
	Tags []string `json:"tags,omitempty"`

	// Requires is a list of permissions that the user would have to give the service, if they provision
	// it. The only permissions currently supported are syslog_drain, route_forwarding and volume_mount.
	// +kubebuilder:validation:Enum=syslog_drain;route_forwarding;volume_mount
	// +listType=set
	Requires []string `json:"requires,omitempty"`

	// Bindable specifies whether Service Instances of the service can be bound to applications. This
	// specifies the default for all Service Plans of this Service Offering. Service Plans can override
	// this field (see Service Plan Object).
	Bindable bool `json:"bindable"`

	// Metadata is an opaque object of metadata for a Service Offering. It is expected that Platforms will
	// treat this as a blob. Note that there are conventions in existing Service Brokers and Platforms for
	// fields that aid in the display of catalog data.
	// +kubebuilder:pruning:PreserveUnknownFields
	Metadata *runtime.RawExtension `json:"metadata,omitempty"`

	// Dashboard is a Cloud Foundry extension described in Catalog Extensions. Contains the data necessary
	// to activate the Dashboard SSO feature for this service.
	DashboardClient *DashboardClient `json:"dashboardClient,omitempty"`

	// PlanUpdatable is whether the Service Offering supports upgrade/downgrade for Service Plans by default.
	// Service Plans can override this field (see Service Plan). Please note that the misspelling of the
	// attribute plan_updatable as plan_updateable was done by mistake. We have opted to keep that misspelling
	// instead of fixing it and thus breaking backward compatibility. Defaults to false.
	PlanUpdatable bool `json:"planUpdatable,omitempty"`

	// ServicePlan is a list of Service Plans for this Service Offering, schema is defined below. MUST
	// contain at least one Service Plan. More info:
	// https://github.com/couchbase/service-broker/tree/master/documentation/modules/ROOT/pages/concepts/catalog.adoc#service-plans
	// +kubebuilder:validation:MinItems=1
	// +listType=map
	// +listMapKey=name
	Plans []ServicePlan `json:"plans"`
}

// DashboardClient is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type DashboardClient struct {
	// ID is the id of the OAuth client that the dashboard will use. If present, MUST be a non-empty string.
	// +kubebuilder:validation:MinLength=1
	ID string `json:"id"`

	// Secret is a secret for the dashboard client. If present, MUST be a non-empty string.
	// +kubebuilder:validation:MinLength=1
	Secret string `json:"secret"`

	// RedirectedURI is a URI for the service dashboard. Validated by the OAuth token server when the dashboard
	// requests a token.
	RedirectedURI string `json:"redirectedURI,omitempty"`
}

// ServicePlan is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type ServicePlan struct {
	// ID is an identifier used to correlate this Service Offering in future requests to the
	// Service Broker. This MUST be globally unique such that Platforms (and their users) MUST
	// be able to assume that seeing the same value (no matter what Service Broker uses it) will
	// always refer to this Service Offering. MUST be a non-empty string. Using a GUID is RECOMMENDED.
	// +kubebuilder:validation:MinLength=1
	ID string `json:"id"`

	// Name is the name of the Service Plan. MUST be unique within the Service Offering. MUST be
	// a non-empty string. Using a CLI-friendly name is RECOMMENDED.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Description is a short description of the Service Plan. MUST be a non-empty string.
	// +kubebuilder:validation:MinLength=1
	Description string `json:"description"`

	// Metadata is an opaque object of metadata for a Service Plan. It is expected that Platforms
	// will treat this as a blob. Note that there are conventions in existing Service Brokers and
	// Platforms for fields that aid in the display of catalog data.
	// +kubebuilder:pruning:PreserveUnknownFields
	Metadata *runtime.RawExtension `json:"metadata,omitempty"`

	// Free, when false, Service Instances of this Service Plan have a cost. The default is true.
	Free bool `json:"free,omitempty"`

	// Bindable specifies whether Service Instances of the Service Plan can be bound to applications.
	// This field is OPTIONAL. If specified, this takes precedence over the bindable attribute of
	// the Service Offering. If not specified, the default is derived from the Service Offering.
	Bindable *bool `json:"bindable,omitempty"`

	// Schemas are schema definitions for Service Instances and Service Bindings for the Service
	// Plan. More info:
	// https://github.com/couchbase/service-broker/tree/master/documentation/modules/ROOT/pages/concepts/catalog.adoc#json-schemas
	Schemas *Schemas `json:"schemas,omitempty"`
}

// Schemas is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type Schemas struct {
	// ServiceInstance is the schema definitions for creating and updating a Service Instance.
	ServiceInstance *ServiceInstanceSchema `json:"serviceInstance,omitempty"`

	// ServiceBinding is the schema definition for creating a Service Binding. Used only if the
	// Service Plan is bindable.
	ServiceBinding *ServiceBindingSchema `json:"serviceBinding,omitempty"`
}

// ServiceInstanceSchema is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type ServiceInstanceSchema struct {
	// Create is the schema definition for creating a Service Instance.
	Create *InputParamtersSchema `json:"create,omitempty"`

	// Update is the chema definition for updating a Service Instance.
	Update *InputParamtersSchema `json:"update,omitempty"`
}

// ServiceBindingSchema is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type ServiceBindingSchema struct {
	// Create is the schema definition for creating a Service Binding.
	Create *InputParamtersSchema `json:"create,omitempty"`
}

// InputParamtersSchema is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type InputParamtersSchema struct {
	// Parameters is the schema definition for the input parameters. Each input parameter is
	// expressed as a property within a JSON object.
	// +kubebuilder:pruning:PreserveUnknownFields
	Parameters *runtime.RawExtension `json:"parameters,omitempty"`
}

// MaintenanceInfo is defined by:
// https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md#body
type MaintenanceInfo struct {
	Version string `json:"version,omitempty"`
}

// ConfigurationTemplate defines a resource template for use when either
// creating a service instance or service binding.
type ConfigurationTemplate struct {
	// Name is the name of the template
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Template defines the resource template, it can be any kind of resource
	// supported by client-go or couchbase.
	// +kubebuilder:pruning:PreserveUnknownFields
	Template *runtime.RawExtension `json:"template"`

	// Singleton alters the behaviour of resource creation.  Typically we will
	// create a resource and use parameters to alter it's name, ensuring it
	// doesn't already exist.  Singleton resources will first check to see
	// whether they exist before attempting creation.
	Singleton bool `json:"singleton,omitempty"`
}

// RegistryValue sets a registry key using a template.
type RegistryValue struct {
	// Name is the name of the registry key to set.
	Name string `json:"name"`

	// Value is the templated string value to calculate. More info:
	// https://github.com/couchbase/service-broker/tree/master/documentation/modules/ROOT/pages/concepts/dynamic-attributes.adoc
	Value string `json:"value"`
}

// RegistryScope allows the user to configure where the registry will be provisioned.
// +kubebuilder:validation:Enum=Explicit;BrokerLocal;InstanceLocal;Prefixed
type RegistryScope string

const (
	// RegistryScopeTenantPrefixed provisions the registry in a "tnt-" prefixed namespace according to the origin of the request.
	RegistryScopeTenantPrefixed RegistryScope = "Prefixed"

	// RegistryScopeExplicit provisions the registry where you tell it to.
	RegistryScopeExplicit RegistryScope = "Explicit"

	// RegistryScopeBrokerLocal provisions the registry in the same namespace
	// as the broker is running in.
	RegistryScopeBrokerLocal RegistryScope = "BrokerLocal"

	// RegistryScopeInstanceLocal provisions the registry in the same namespace
	// as the service instance.
	RegistryScopeInstanceLocal RegistryScope = "InstanceLocal"
)

// ConfigurationBinding binds a service plan to a set of templates
// required to realize that plan.
type ConfigurationBinding struct {
	// Name is a unique identifier for the binding.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// RegistryScope controls where the registry for a service instance
	// or binding is located.  The service broker makes all generated
	// resources owned by the relevant registry, so deleting a service
	// instance means deleting the registry and letting garbage collection
	// do the rest.  What is particularly important is that resources
	// must be located in the same namespace as their owners, or they will
	// be garbage collected.  "BrokerLocal", the default provisions service
	// registries in the same namespace as the service broker.  "Explicit"
	// allows service registries to be hard coded to a specific namespace.
	// "InstanceLocal" will provision service registries in the same
	// namespace as the service instance was provisioned in.
	// +kubebuilder:default="BrokerLocal"
	RegistryScope RegistryScope `json:"registryScope,omitempty"`

	// RegistryNamespace is only relevant when used with RegistryScope in the
	// "Explicit" mode, and specifies the exact namespace a service instance
	// registry will be generated in.
	// +kubebuilder:validation:MinLength=1
	RegistryNamespace string `json:"registryNamespace,omitempty"`

	// RegistryEnabledOrganizations is only relevant when used with RegistryScope in the
	// "Prefixed" mode, and specifies the organizations which are enabled for createion of
	// the registry in dedicated namespaces.
	RegistryEnabledOrganizations []string `json:"registryEnabledOrganizations,omitempty"`

	// RegistryPrefix is only relevant when used with RegistryScope in the
	// "Prefix" mode, and specifies the prefix to the namespace where the resources for the
	// enabled organizations should be created in.
	// +kubebuilder:validation:MinLength=1
	RegistryPrefix string `json:"registryPrefix,omitempty"`

	// Service is the name of the service offering to bind to.
	// +kubebuilder:validation:MinLength=1
	Service string `json:"service"`

	// Plan is the name of the service plan to bind to.
	// +kubebuilder:validation:MinLength=1
	Plan string `json:"plan"`

	// ServiceInstance defines the set of templates to render and create when
	// a new service instance is created.
	ServiceInstance ServiceBrokerTemplateList `json:"serviceInstance"`

	// ServiceBinding defines the set of templates to render and create when
	// a new service binding is created.  This attribute is optional based on
	// whether the service plan allows binding.
	ServiceBinding *ServiceBrokerTemplateList `json:"serviceBinding,omitempty"`
}

// ServiceBrokerTemplateList is an ordered list of templates to use
// when performing a specific operation.
type ServiceBrokerTemplateList struct {
	// Registry allows the pre-calculation of dynamic configuration from
	// request inputs i.e. registry or parameters, or generated e.g. passwords.
	// +listType=map
	// +listMapKey=name
	Registry []RegistryValue `json:"registry,omitempty"`

	// Templates defines all the templates that will be created, in order,
	// by the service broker for this operation.
	// This field is deprecated, use steps instead.
	// +listType=set
	Templates []string `json:"templates,omitempty"`

	// ReadinessChecks defines a set of tests that define whether a service instance
	// or service binding is actually ready as reported by the service broker polling
	// API.
	// +listType=map
	// +listMapKey=name
	ReadinessChecks []ConfigurationReadinessCheck `json:"readinessChecks,omitempty"`

	// Steps allows a service instance or binding deployment to be split into steps.
	// A steps will block until the readiness check, if defined, passes, before
	// continuing on to the next one.  Steps cannot be used at the same time as
	// templates and readiness checks.
	// +listType=map
	// +listMapKey=name
	Steps []ServiceBrokerTemplateListStep `json:"steps,omitempty"`
}

// ServiceBrokerTemplateListStep allows a service instance to be provisioned in steps
// blocking until a readiness check has completed before moving on to the next one.
type ServiceBrokerTemplateListStep struct {
	// Name of the step for logging and debugging purposes.
	Name string `json:"name"`

	// Templates defines all the templates that will be created, in order,
	// by the service broker for this operation.
	// +listType=set
	Templates []string `json:"templates,omitempty"`

	// ReadinessChecks defines a set of tests that define whether a step is complete.
	// These checks have no affect on the aysnchronous polling at the service broker
	// API level, as such it's common to define these between steps only, and have a
	// top level readiness check for service availability.
	// +listType=map
	// +listMapKey=name
	ReadinessChecks []ConfigurationReadinessCheck `json:"readinessChecks,omitempty"`
}

// ConfigurationReadinessCheck is a readiness check to perform on a service instance
// or binding before declaring it ready and provisioning has completed.
type ConfigurationReadinessCheck struct {
	// Name is a unique name for the readiness check for debugging purposes.
	Name string `json:"name"`

	// Condition allows the service broker to poll well-formed status conditions
	// in order to determine whether a specific resource is ready.
	Condition *ConfigurationReadinessCheckCondition `json:"condition,omitempty"`

	// Timeout is the timeout durations for this check.
	// +kubebuilder:default="1m"
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// ConfigurationReadinessCheckCondition allows the service broker to poll well-formed
// status conditions in order to determine whether a specific resource is ready.
// This can be thought of a `kubectl wait` but done properly.
type ConfigurationReadinessCheckCondition struct {
	// APIVersion is the resource api version e.g. "apps/v1"
	APIVersion string `json:"apiVersion"`

	// Kind is the resource kind to poll e.g. "Deployment"
	Kind string `json:"kind"`

	// Namespace is the namespace the resource resides in.
	Namespace string `json:"namespace"`

	// Name is the resource name to poll.
	Name string `json:"name"`

	// Type is the type of the condition to look for e.g. "Available"
	Type string `json:"type"`

	// Status is the status of the condition that must match e.g. "True"
	Status string `json:"status"`
}

// ServiceBrokerConfigStatus records status information about a configuration
// as the Service Broker processes it.
type ServiceBrokerConfigStatus struct {
	// Conditions indicate state of particular aspects of a configuration.
	Conditions []ServiceBrokerConfigCondition `json:"conditions,omitempty"`
}

// ServiceBrokerConfigConditionType is the type of condition being described.
type ServiceBrokerConfigConditionType string

const (
	// ConfigurationValid records whether the configuration is valid or
	// not.
	ConfigurationValid ServiceBrokerConfigConditionType = "ConfigurationValid"
)

// ConditionStatus is used to define what state the condition is in.
type ConditionStatus string

const (
	// ConditionTrue means that the resource meets the condition.
	ConditionTrue ConditionStatus = "True"

	// ConditionFalse means that the resource does not meet the condition.
	ConditionFalse ConditionStatus = "False"
)

// ServiceBrokerConfigCondition represents a condition associated with the configuration.
type ServiceBrokerConfigCondition struct {
	// Type is the type of condition.
	Type ServiceBrokerConfigConditionType `json:"type"`

	// Status is the status of the condition, whether it is true or false.
	Status ConditionStatus `json:"status"`

	// LastTransitionTime records the last time the status changed from one value
	// to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// Reason is a unique one word camel case reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Message is a human readable message indicating details about the last transition.
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ServiceBrokerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceBrokerConfig `json:"items"`
}
