This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe EcsTaskOverride

The overrides that are associated with a task.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerOverrides" : [ EcsContainerOverride, ... ],
  "Cpu" : String,
  "EphemeralStorage" : EcsEphemeralStorage,
  "ExecutionRoleArn" : String,
  "InferenceAcceleratorOverrides" : [ EcsInferenceAcceleratorOverride, ... ],
  "Memory" : String,
  "TaskRoleArn" : String
}

```

### YAML

```yaml

  ContainerOverrides:
    - EcsContainerOverride
  Cpu: String
  EphemeralStorage:
    EcsEphemeralStorage
  ExecutionRoleArn: String
  InferenceAcceleratorOverrides:
    - EcsInferenceAcceleratorOverride
  Memory: String
  TaskRoleArn: String

```

## Properties

`ContainerOverrides`

One or more container overrides that are sent to a task.

_Required_: No

_Type_: Array of [EcsContainerOverride](aws-properties-pipes-pipe-ecscontaineroverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cpu`

The cpu override for the task.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EphemeralStorage`

The ephemeral storage setting override for the task.

###### Note

This parameter is only supported for tasks hosted on Fargate that use
the following platform versions:

- Linux platform version `1.4.0` or later.

- Windows platform version `1.0.0` or later.

_Required_: No

_Type_: [EcsEphemeralStorage](aws-properties-pipes-pipe-ecsephemeralstorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the task execution IAM role override for the task. For
more information, see [Amazon ECS\
task execution IAM role](../../../amazonecs/latest/developerguide/task-execution-iam-role.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-]+):([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.+)|(\$(\.[\w/_-]+(\[(\d+|\*)\])*)*)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceAcceleratorOverrides`

The Elastic Inference accelerator override for the task.

_Required_: No

_Type_: Array of [EcsInferenceAcceleratorOverride](aws-properties-pipes-pipe-ecsinferenceacceleratoroverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Memory`

The memory override for the task.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskRoleArn`

The Amazon Resource Name (ARN) of the IAM role that containers in this task can assume.
All containers in this task are granted the permissions that are specified in this role.
For more information, see [IAM Role for Tasks](../../../amazonecs/latest/developerguide/task-iam-roles.md) in
the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-]+):([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.+)|(\$(\.[\w/_-]+(\[(\d+|\*)\])*)*)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcsResourceRequirement

Filter

All content copied from https://docs.aws.amazon.com/.
