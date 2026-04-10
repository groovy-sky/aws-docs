This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::Schedule EcsParameters

The templated target type for the Amazon ECS [`RunTask`](../../../../reference/amazonecs/latest/apireference/api-runtask.md) API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityProviderStrategy" : [ CapacityProviderStrategyItem, ... ],
  "EnableECSManagedTags" : Boolean,
  "EnableExecuteCommand" : Boolean,
  "Group" : String,
  "LaunchType" : String,
  "NetworkConfiguration" : NetworkConfiguration,
  "PlacementConstraints" : [ PlacementConstraint, ... ],
  "PlacementStrategy" : [ PlacementStrategy, ... ],
  "PlatformVersion" : String,
  "PropagateTags" : String,
  "ReferenceId" : String,
  "Tags" : [ {Key: Value, ...}, ... ],
  "TaskCount" : Number,
  "TaskDefinitionArn" : String
}

```

### YAML

```yaml

  CapacityProviderStrategy:
    - CapacityProviderStrategyItem
  EnableECSManagedTags: Boolean
  EnableExecuteCommand: Boolean
  Group: String
  LaunchType: String
  NetworkConfiguration:
    NetworkConfiguration
  PlacementConstraints:
    - PlacementConstraint
  PlacementStrategy:
    - PlacementStrategy
  PlatformVersion: String
  PropagateTags: String
  ReferenceId: String
  Tags:
    -
    Key: Value
  TaskCount: Number
  TaskDefinitionArn: String

```

## Properties

`CapacityProviderStrategy`

The capacity provider strategy to use for the task.

_Required_: No

_Type_: Array of [CapacityProviderStrategyItem](aws-properties-scheduler-schedule-capacityproviderstrategyitem.md)

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableECSManagedTags`

Specifies whether to enable Amazon ECS managed tags for the task. For more information, see [Tagging Your Amazon ECS Resources](../../../amazonecs/latest/developerguide/ecs-using-tags.md)
in the _Amazon ECS Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableExecuteCommand`

Whether or not to enable the execute command functionality for the containers in this task. If true, this enables execute command functionality on all containers in the task.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Group`

Specifies an Amazon ECS task group for the task. The maximum length is 255 characters.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchType`

Specifies the launch type on which your task is running. The launch type that you specify here must match one of the launch type (compatibilities) of the target task.
The `FARGATE` value is supported only in the Regions where Fargate with Amazon ECS is supported.
For more information, see [AWS Fargate on Amazon ECS](../../../amazonecs/latest/developerguide/aws-fargate.md) in the _Amazon ECS Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `EC2 | FARGATE | EXTERNAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

This structure specifies the network configuration for an ECS task.

_Required_: No

_Type_: [NetworkConfiguration](aws-properties-scheduler-schedule-networkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlacementConstraints`

An array of placement constraint objects to use for the task. You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime).

_Required_: No

_Type_: Array of [PlacementConstraint](aws-properties-scheduler-schedule-placementconstraint.md)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlacementStrategy`

The task placement strategy for a task or service.

_Required_: No

_Type_: [Array](aws-properties-scheduler-schedule-placementstrategy.md) of [PlacementStrategy](aws-properties-scheduler-schedule-placementstrategy.md)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlatformVersion`

Specifies the platform version for the task. Specify only the numeric portion of the platform version, such as `1.1.0`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropagateTags`

Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags are not propagated.
Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the Amazon ECS [`TagResource`](../../../../reference/amazonecs/latest/apireference/api-tagresource.md)
API action.

_Required_: No

_Type_: String

_Allowed values_: `TASK_DEFINITION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReferenceId`

The reference ID to use for the task.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The metadata that you apply to the task to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define.
For more information, see [`RunTask`](../../../../reference/amazonecs/latest/apireference/api-runtask.md) in the _Amazon ECS API Reference_.

_Required_: No

_Type_: Array of Object

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskCount`

The number of tasks to create based on `TaskDefinition`. The default is `1`.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskDefinitionArn`

The Amazon Resource Name (ARN) of the task definition to use if the event target is an Amazon ECS task.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeadLetterConfig

EventBridgeParameters

All content copied from https://docs.aws.amazon.com/.
