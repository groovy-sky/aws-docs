This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service ServiceConnectClientAlias

Each alias ("endpoint") is a fully-qualified name and port number that other tasks
("clients") can use to connect to this service.

Each name and port mapping must be unique within the namespace.

Tasks that run in a namespace can use short names to connect to services in the
namespace. Tasks can connect to services across all of the clusters in the namespace.
Tasks connect through a managed proxy container that collects logs and metrics for
increased visibility. Only the tasks that Amazon ECS services create are supported with
Service Connect. For more information, see [Service Connect](../../../amazonecs/latest/developerguide/service-connect.md)
in the _Amazon Elastic Container Service Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DnsName" : String,
  "Port" : Integer,
  "TestTrafficRules" : ServiceConnectTestTrafficRules
}

```

### YAML

```yaml

  DnsName: String
  Port: Integer
  TestTrafficRules:
    ServiceConnectTestTrafficRules

```

## Properties

`DnsName`

The `dnsName` is the name that you use in the applications of client tasks
to connect to this service. The name must be a valid DNS name but doesn't need to be
fully-qualified. The name can include up to 127 characters. The name can include
lowercase letters, numbers, underscores (\_), hyphens (-), and periods (.). The name
can't start with a hyphen.

If this parameter isn't specified, the default value of
`discoveryName.namespace` is used. If the `discoveryName`
isn't specified, the port mapping name from the task definition is used in
`portName.namespace`.

To avoid changing your applications in client Amazon ECS services, set this to the
same name that the client application uses by default. For example, a few common names
are `database`, `db`, or the lowercase name of a database, such as
`mysql` or `redis`. For more information, see [Service Connect](../../../amazonecs/latest/developerguide/service-connect.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The listening port number for the Service Connect proxy. This port is available inside
of all of the tasks within the same namespace.

To avoid changing your applications in client Amazon ECS services, set this to the
same port that the client application uses by default. For more information, see [Service Connect](../../../amazonecs/latest/developerguide/service-connect.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TestTrafficRules`

The configuration for test traffic routing rules used during blue/green deployments with Amazon ECS Service Connect. This allows you to route a portion of traffic to the new service revision of your service for testing before shifting all production traffic.

_Required_: No

_Type_: [ServiceConnectTestTrafficRules](aws-properties-ecs-service-serviceconnecttesttrafficrules.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](../userguide/aws-resource-ecs-service.md#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceConnectAccessLogConfiguration

ServiceConnectConfiguration

All content copied from https://docs.aws.amazon.com/.
