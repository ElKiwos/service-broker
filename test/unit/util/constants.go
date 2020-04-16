// Copyright 2020 Couchbase, Inc.
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

package util

const (
	// CA is the CA for the server certificate.
	CA = `-----BEGIN CERTIFICATE-----
MIIDSzCCAjOgAwIBAgIUEFGQGyBv9Asf9XMjL6Zcs2kUCfcwDQYJKoZIhvcNAQEL
BQAwFjEUMBIGA1UEAwwLRWFzeS1SU0EgQ0EwHhcNMjAwMjIxMTQyNDA5WhcNMzAw
MjE4MTQyNDA5WjAWMRQwEgYDVQQDDAtFYXN5LVJTQSBDQTCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBAM7xKnvpJmAi9MK7E0dQcGFoSmHj8r6WRpRfi/20
aTaQ6ukkxzMW6Gcsl/qpYpqJHn3tem0ryRNt23dnl0p7ZTMRnzlgfbZs0HEk8O+B
s+FwZ+gGWcMiR+9rMTCz2trIuQEqu4Hipr2WF2NDq3nzS7z553Iimk1fdHDSUBBe
5D/JUNPG1YtWYVYs5VUIA4m96gtVIsalaaEnoDDKiDo+tcD8ThD4B/mo5y/6L3An
XRyUQD/5cdKxurcDqu8dvZGWW9CIkKNke3XnPZxenGjdYOs1t837oHmof2jWqCOo
7rCv14rXnmQyeMnFQYR2kEI2+PxSv7utcdTU6YKePbS5pXkCAwEAAaOBkDCBjTAd
BgNVHQ4EFgQUwMFny4aBejuuWtoaG9BfO03QTggwUQYDVR0jBEowSIAUwMFny4aB
ejuuWtoaG9BfO03QTgihGqQYMBYxFDASBgNVBAMMC0Vhc3ktUlNBIENBghQQUZAb
IG/0Cx/1cyMvplyzaRQJ9zAMBgNVHRMEBTADAQH/MAsGA1UdDwQEAwIBBjANBgkq
hkiG9w0BAQsFAAOCAQEAhcqeu5LTt5OY1H9bcjihJTm6iG42liiZ9yQXAHLOI1yE
QdmmQgLh2+dwEKbZv+F/7u0OnFdAGZdU7tQRWfsTDVRCC3rF+jWdBnaYkcVT8LCo
fTkMXXda2F6HatS1mdBQmg44Lq9OkjJoQjGRC9xW4ZR+ELy02bTogIgEuHrMLot+
cbK7HX+1LI0yQ0r09bnaZUQCjQUASxiEvamPMOXktlYfFAmMP0HpCO02tIertM4C
tH+ZvXKsdJuWaR4gfkvJGkx1XyHtguWtSHoDGdfECw/vSxRyfLBUusla48S5lF2P
/Vj5Th0XvZDEkS1nb4ywGUKapse+ujSmrr80BK/P/Q==
-----END CERTIFICATE-----`

	// Cert is a valid for "DNS:localhost" until 2022-05-26.
	Cert = `-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIRAOcrzz4scdBTQozXBqz3WgAwDQYJKoZIhvcNAQELBQAw
FjEUMBIGA1UEAwwLRWFzeS1SU0EgQ0EwHhcNMjAwMjIxMTQyNjQ5WhcNMjIwNTI2
MTQyNjQ5WjARMQ8wDQYDVQQDDAZzZXJ2ZXIwggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQCmA3N0Gwt/FNgmLoSAjE2Hrp7KbTKczjpuTV39tIj5oyJV+PBu
5LHGcs8A9KNf0ppzzW3F6plYjwZTNBjMQr6+Cu3637tAJHKhOP+kozJu+G+lTui+
N4fJaY4HHPZ4YeLwGJuYQM2dEuSM6uyWsdJpIQTEsAkZHYRcmeJi9gJ2k0ELIJp2
TeYkMqWiqBW8M7mJggPNNmG7CLlGhJ2sC0PbwbveQUNG62aUhEJO05RyyXaYtIlz
b8qohNEGa5g/6KAUwZr/zibswzIrNhQaRvkv/lXutUdYHz/GediI7w+gHj9i84I5
om3D+CPftCVATut4Ws6DhZevXj4sLUcp/SppAgMBAAGjgbgwgbUwCQYDVR0TBAIw
ADAdBgNVHQ4EFgQUr/5/0DyO1AKXhRdiftxfHzfB33EwUQYDVR0jBEowSIAUwMFn
y4aBejuuWtoaG9BfO03QTgihGqQYMBYxFDASBgNVBAMMC0Vhc3ktUlNBIENBghQQ
UZAbIG/0Cx/1cyMvplyzaRQJ9zATBgNVHSUEDDAKBggrBgEFBQcDATALBgNVHQ8E
BAMCBaAwFAYDVR0RBA0wC4IJbG9jYWxob3N0MA0GCSqGSIb3DQEBCwUAA4IBAQCs
mnGTb+tZddnFcnrLBKZeyo6TSpQeufnB2PjEBEIHrxgnWsySmhkvdO8mXDToa9L6
NuapvbqLzwin1wc9/OS5NRM74BtGWPFtAt4oaBASLo59Fivdo1JtmL359o+lztlT
NgRAKydyRu2DrnyuBUlHXWalPbHPHcl4JnOqrLBz+R/g+nT2/BXGprZ+x3+xLII/
+wcItBeSmFKw9k08d3qYBfRC6wuBLnrUhxdmxy2V7x+TxXMhVBsirNYtr51SeZ0I
fgwCyu99qIaSpQzhamfapow3qF5rjY6kg6366XFr2wYlw45WMNpPmtGXKSXlS48t
aD0AkDedZzi4p3X6Af5S
-----END CERTIFICATE-----`

	// Key is the private key for cert.
	Key = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCmA3N0Gwt/FNgm
LoSAjE2Hrp7KbTKczjpuTV39tIj5oyJV+PBu5LHGcs8A9KNf0ppzzW3F6plYjwZT
NBjMQr6+Cu3637tAJHKhOP+kozJu+G+lTui+N4fJaY4HHPZ4YeLwGJuYQM2dEuSM
6uyWsdJpIQTEsAkZHYRcmeJi9gJ2k0ELIJp2TeYkMqWiqBW8M7mJggPNNmG7CLlG
hJ2sC0PbwbveQUNG62aUhEJO05RyyXaYtIlzb8qohNEGa5g/6KAUwZr/zibswzIr
NhQaRvkv/lXutUdYHz/GediI7w+gHj9i84I5om3D+CPftCVATut4Ws6DhZevXj4s
LUcp/SppAgMBAAECggEADVMBlTQGfDCkGIxrPhYEsvsk64JQKZ2zut6iyJYQ2Fhr
jRLp2TypuA/G5YC9DBfPJbQ7N0NZA26XR15LDzncLUybRSNn2AIU4TO98OzYQ2fV
LiNjMsEqONQr/g4pCghxOmv/MP0ig5TcmyLP0lh7Vsy7oT2vvUBNO1FuhhrQn8uY
uBh6wQopL3mT+GiOkJVhXzKyCs/ZjJFMt5icCGxewOcRY6iIrBUUlmj9mTM9ysIW
/7J2q37ISCOCyC1JElugCtv7R3ZI6ZwBMMbhINfFLiaq0+UJXU9YxFCKlBOtDZ1E
OKHRPDpPThWHYADnFkq/79fRbW0rtNttTkMgUlx6wQKBgQDRhEH1JEGLONCQcWDG
71Xx9vhg45UpLX+i1QLC/eRst9bbtnjXrRBdDWRb3njHZWi5DE1r3VUMLsQt4hRx
PFM/wo1fqlpXad+iNbJFEEUk16/DwbO8XEc4UIxVYnLVL8HcxEZLRoMstqjD5kPL
eNrwaBLcqBr8Qm/T6wy/5fDXDQKBgQDK2GGGSv7QbLVuc9UVHtPzWLchYRG5/h1e
biIXXoJg6Nd9tf9/bZQkawOr7eFfKD/BKQvkQYtHMPQVoIp8fpXX7Spk07i1PC01
ZXNc+NT0kVrJuorsZxN/yZM8c9tex9pAdYAnRqkbp6XvYGjNdNIYhz2Qpx3ZbcCP
jNXXZneJzQKBgHEh9WullC7VEumsDxHckpABR8UpnpWJl+4ZD1CzP/DkpAQn18C+
FqPoY1SoIJeqzo19cyDXduEJL62G8nrilCFNsIEDv5yL8tHoJMbeLjfir2oI4kDH
oz1pYR2J92/eRdQrt7lV7ebrCt4dLGZmb/J4gBbePxQP28qWlV/Zjd7RAoGAevaQ
qAfuUAqWMU6mbRczBOFSojlltYoF46h/ogr4niaH+vzI1UZn92un3iFl7XlIrJ9l
Rgk1lQJn9HRNfwp9a1epy1VNMxA3l5bYSBPPhDJZBtC+RnB1sZFQX+UbpmkgNNMF
zMlY2hrWzDV1UpbuhU/2Uin8PkH56QtG1jyXWkECgYBahkRVEh0zJ29YmYL648df
MQbdqODqE5k2LKq5aUNje1Mn5yQoFxmLbDXKqYCHT5RGwILw2vLqP3j4xQG/NVtx
gsjgUnOUm2N9H6SjMxFMbfX6fEmA9npxmXnGtULq3Vrj1MSXiP2Gch9chPooD66B
vFKYb6v/2ENK34iu3AHVow==
-----END PRIVATE KEY-----`

	// Token is the default OAuth bearer token.
	Token = "HeMan"

	// Namespace is the default namespace, that isn't default.
	Namespace = "Skeletor"
)
