This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service ServiceConnectTestTrafficRules

The test traffic routing configuration for Amazon ECS blue/green deployments. This
configuration allows you to define rules for routing specific traffic to the new service
revision during the deployment process, allowing for safe testing before full production
traffic shift.

For more information, see [Service Connect for Amazon ECS blue/green deployments](../../../amazonecs/latest/developerguide/service-connect-blue-green.md) in the _Amazon Elastic Container Service Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Header" : ServiceConnectTestTrafficRulesHeader
}

```

### YAML

```yaml

  Header:
    ServiceConnectTestTrafficRulesHeader

```

## Properties

`Header`

The HTTP header-based routing rules that determine which requests should be routed to the new service version during blue/green deployment testing. These rules provide fine-grained control over test traffic routing based on request headers.

_Required_: Yes

_Type_: [ServiceConnectTestTrafficRulesHeader](aws-properties-ecs-service-serviceconnecttesttrafficrulesheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceConnectService

ServiceConnectTestTrafficRulesHeader

All content copied from https://docs.aws.amazon.com/.
