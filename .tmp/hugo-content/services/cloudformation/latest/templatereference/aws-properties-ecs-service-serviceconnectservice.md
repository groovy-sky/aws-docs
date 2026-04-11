This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service ServiceConnectService

The Service Connect service object configuration. For more information, see [Service Connect](../../../amazonecs/latest/developerguide/service-connect.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientAliases" : [ ServiceConnectClientAlias, ... ],
  "DiscoveryName" : String,
  "IngressPortOverride" : Integer,
  "PortName" : String,
  "Timeout" : TimeoutConfiguration,
  "Tls" : ServiceConnectTlsConfiguration
}

```

### YAML

```yaml

  ClientAliases:
    - ServiceConnectClientAlias
  DiscoveryName: String
  IngressPortOverride: Integer
  PortName: String
  Timeout:
    TimeoutConfiguration
  Tls:
    ServiceConnectTlsConfiguration

```

## Properties

`ClientAliases`

The list of client aliases for this Service Connect service. You use these to assign
names that can be used by client applications. The maximum number of client aliases that
you can have in this list is 1.

Each alias ("endpoint") is a fully-qualified name and port number that other Amazon
ECS tasks ("clients") can use to connect to this service.

Each name and port mapping must be unique within the namespace.

For each `ServiceConnectService`, you must provide at least one
`clientAlias` with one `port`.

_Required_: No

_Type_: Array of [ServiceConnectClientAlias](aws-properties-ecs-service-serviceconnectclientalias.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DiscoveryName`

The `discoveryName` is the name of the new AWS Cloud Map service that Amazon ECS
creates for this Amazon ECS service. This must be unique within the AWS Cloud Map namespace. The
name can contain up to 64 characters. The name can include lowercase letters, numbers,
underscores (\_), and hyphens (-). The name can't start with a hyphen.

If the `discoveryName` isn't specified, the port mapping name from the task
definition is used in `portName.namespace`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngressPortOverride`

The port number for the Service Connect proxy to listen on.

Use the value of this field to bypass the proxy for traffic on the port number
specified in the named `portMapping` in the task definition of this
application, and then use it in your VPC security groups to allow traffic into the proxy
for this Amazon ECS service.

In `awsvpc` mode and Fargate, the default value is the container port
number. The container port number is in the `portMapping` in the task
definition. In bridge mode, the default value is the ephemeral port of the Service
Connect proxy.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortName`

The `portName` must match the name of one of the `portMappings`
from all the containers in the task definition of this Amazon ECS service.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

A reference to an object that represents the configured timeouts for Service
Connect.

_Required_: No

_Type_: [TimeoutConfiguration](aws-properties-ecs-service-timeoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tls`

A reference to an object that represents a Transport Layer Security (TLS)
configuration.

_Required_: No

_Type_: [ServiceConnectTlsConfiguration](aws-properties-ecs-service-serviceconnecttlsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](../userguide/aws-resource-ecs-service.md#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceConnectConfiguration

ServiceConnectTestTrafficRules

All content copied from https://docs.aws.amazon.com/.
