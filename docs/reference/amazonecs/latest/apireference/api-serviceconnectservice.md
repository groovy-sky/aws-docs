# ServiceConnectService

The Service Connect service object configuration. For more information, see [Service Connect](../../../../services/amazonecs/latest/developerguide/service-connect.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Contents

**portName**

The `portName` must match the name of one of the `portMappings`
from all the containers in the task definition of this Amazon ECS service.

Type: String

Required: Yes

**clientAliases**

The list of client aliases for this Service Connect service. You use these to assign
names that can be used by client applications. The maximum number of client aliases that
you can have in this list is 1.

Each alias ("endpoint") is a fully-qualified name and port number that other Amazon
ECS tasks ("clients") can use to connect to this service.

Each name and port mapping must be unique within the namespace.

For each `ServiceConnectService`, you must provide at least one
`clientAlias` with one `port`.

Type: Array of [ServiceConnectClientAlias](api-serviceconnectclientalias.md) objects

Required: No

**discoveryName**

The `discoveryName` is the name of the new AWS Cloud Map service that Amazon ECS
creates for this Amazon ECS service. This must be unique within the AWS Cloud Map namespace. The
name can contain up to 64 characters. The name can include lowercase letters, numbers,
underscores (\_), and hyphens (-). The name can't start with a hyphen.

If the `discoveryName` isn't specified, the port mapping name from the task
definition is used in `portName.namespace`.

Type: String

Required: No

**ingressPortOverride**

The port number for the Service Connect proxy to listen on.

Use the value of this field to bypass the proxy for traffic on the port number
specified in the named `portMapping` in the task definition of this
application, and then use it in your VPC security groups to allow traffic into the proxy
for this Amazon ECS service.

In `awsvpc` mode and Fargate, the default value is the container port
number. The container port number is in the `portMapping` in the task
definition. In bridge mode, the default value is the ephemeral port of the Service
Connect proxy.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 65535.

Required: No

**timeout**

A reference to an object that represents the configured timeouts for Service
Connect.

Type: [TimeoutConfiguration](api-timeoutconfiguration.md) object

Required: No

**tls**

A reference to an object that represents a Transport Layer Security (TLS)
configuration.

Type: [ServiceConnectTlsConfiguration](api-serviceconnecttlsconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/serviceconnectservice.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/serviceconnectservice.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/serviceconnectservice.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ServiceConnectConfiguration

ServiceConnectServiceResource
