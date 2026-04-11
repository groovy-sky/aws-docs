This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service DeploymentAlarms

One of the methods which provide a way for you to quickly identify when a deployment
has failed, and then to optionally roll back the failure to the last working
deployment.

When the alarms are generated, Amazon ECS sets the service deployment to failed. Set
the rollback parameter to have Amazon ECS to roll back your service to the last
completed deployment after a failure.

You can only use the `DeploymentAlarms` method to detect failures when the
`DeploymentController` is set to `ECS`.

For more information, see [Rolling\
update](../../../amazonecs/latest/developerguide/deployment-type-ecs.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmNames" : [ String, ... ],
  "Enable" : Boolean,
  "Rollback" : Boolean
}

```

### YAML

```yaml

  AlarmNames:
    - String
  Enable: Boolean
  Rollback: Boolean

```

## Properties

`AlarmNames`

One or more CloudWatch alarm names. Use a "," to separate the alarms.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enable`

Determines whether to use the CloudWatch alarm option in the service deployment
process.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rollback`

Determines whether to configure Amazon ECS to roll back the service if a service
deployment fails. If rollback is used, when a service deployment fails, the service is
rolled back to the last deployment that completed successfully.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](../userguide/aws-resource-ecs-service.md#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityProviderStrategyItem

DeploymentCircuitBreaker

All content copied from https://docs.aws.amazon.com/.
