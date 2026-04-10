This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service DeploymentCircuitBreaker

###### Note

The deployment circuit breaker can only be used for services using the rolling
update ( `ECS`) deployment type.

The **deployment circuit breaker** determines whether a
service deployment will fail if the service can't reach a steady state. If it is turned
on, a service deployment will transition to a failed state and stop launching new tasks.
You can also configure Amazon ECS to roll back your service to the last completed
deployment after a failure. For more information, see [Rolling\
update](../../../amazonecs/latest/developerguide/deployment-type-ecs.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

For more information about API failure reasons, see [API failure\
reasons](../../../amazonecs/latest/developerguide/api-failures-messages.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enable" : Boolean,
  "Rollback" : Boolean
}

```

### YAML

```yaml

  Enable: Boolean
  Rollback: Boolean

```

## Properties

`Enable`

Determines whether to use the deployment circuit breaker logic for the service.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rollback`

Determines whether to configure Amazon ECS to roll back the service if a service
deployment fails. If rollback is on, when a service deployment fails, the service is
rolled back to the last deployment that completed successfully.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](../userguide/aws-resource-ecs-service.md#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentAlarms

DeploymentConfiguration

All content copied from https://docs.aws.amazon.com/.
