This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule EcsParameters

The custom parameters to be used when the target is an Amazon ECS task.

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
  "PlacementStrategies" : [ PlacementStrategy, ... ],
  "PlatformVersion" : String,
  "PropagateTags" : String,
  "ReferenceId" : String,
  "TagList" : [ Tag, ... ],
  "TaskCount" : Integer,
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
  PlacementStrategies:
    - PlacementStrategy
  PlatformVersion: String
  PropagateTags: String
  ReferenceId: String
  TagList:
    - Tag
  TaskCount: Integer
  TaskDefinitionArn: String

```

## Properties

`CapacityProviderStrategy`

The capacity provider strategy to use for the task.

If a `capacityProviderStrategy` is specified, the `launchType`
parameter must be omitted. If no `capacityProviderStrategy` or launchType is
specified, the `defaultCapacityProviderStrategy` for the cluster is used.

_Required_: No

_Type_: Array of [CapacityProviderStrategyItem](aws-properties-events-rule-capacityproviderstrategyitem.md)

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableECSManagedTags`

Specifies whether to enable Amazon ECS managed tags for the task. For more
information, see [Tagging Your Amazon ECS\
Resources](../../../amazonecs/latest/developerguide/ecs-using-tags.md) in the Amazon Elastic Container Service Developer Guide.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableExecuteCommand`

Whether or not to enable the execute command functionality for the containers in this
task. If true, this enables execute command functionality on all containers in the
task.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Group`

Specifies an ECS task group for the task. The maximum length is 255 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchType`

Specifies the launch type on which your task is running. The launch type that you specify
here must match one of the launch type (compatibilities) of the target task. The
`FARGATE` value is supported only in the Regions where AWS Fargate
with Amazon ECS is supported. For more information, see [AWS Fargate on Amazon ECS](../../../amazonecs/latest/developerguide/aws-fargate.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

_Required_: No

_Type_: String

_Allowed values_: `EC2 | FARGATE | EXTERNAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

Use this structure if the Amazon ECS task uses the `awsvpc` network
mode. This structure specifies the VPC subnets and security groups associated with the task,
and whether a public IP address is to be used. This structure is required if
`LaunchType` is `FARGATE` because the `awsvpc` mode is
required for Fargate tasks.

If you specify `NetworkConfiguration` when the target ECS task does not use the
`awsvpc` network mode, the task fails.

_Required_: No

_Type_: [NetworkConfiguration](aws-properties-events-rule-networkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlacementConstraints`

An array of placement constraint objects to use for the task. You can specify up to 10
constraints per task (including constraints in the task definition and those specified at
runtime).

_Required_: No

_Type_: Array of [PlacementConstraint](aws-properties-events-rule-placementconstraint.md)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlacementStrategies`

The placement strategy objects to use for the task. You can specify a maximum of five
strategy rules per task.

_Required_: No

_Type_: Array of [PlacementStrategy](aws-properties-events-rule-placementstrategy.md)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlatformVersion`

Specifies the platform version for the task. Specify only the numeric portion of the
platform version, such as `1.1.0`.

This structure is used only if `LaunchType` is `FARGATE`. For more
information about valid platform versions, see [AWS Fargate\
Platform Versions](../../../amazonecs/latest/developerguide/platform-versions.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropagateTags`

Specifies whether to propagate the tags from the task definition to the task. If no value
is specified, the tags are not propagated. Tags can only be propagated to the task during task
creation. To add tags to a task after task creation, use the TagResource API action.

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

`TagList`

The metadata that you apply to the task to help you categorize and organize them. Each tag
consists of a key and an optional value, both of which you define. To learn more, see [RunTask](../../../../reference/amazonecs/latest/apireference/api-runtask.md#ECS-RunTask-request-tags) in the Amazon ECS API Reference.

_Required_: No

_Type_: Array of [Tag](aws-properties-events-rule-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskCount`

The number of tasks to create based on `TaskDefinition`. The default is
1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskDefinitionArn`

The ARN of the task definition to use if the event target is an Amazon ECS task.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### ECS example

The following example defines the ECS parameters.

#### JSON

```json

{
  "LaunchType": "FARGATE",
  "NetworkConfiguration": {
    "AwsVpcConfiguration": {
      "AssignPublicIp": "DISABLED",
      "SecurityGroups": [
        {
          "Fn: : GetAtt": [
            "ScheduledFargateTaskScheduledTaskDefSecurityGroupE075BC19",
            "GroupId"
          ]
        }
      ],
      "Subnets": [
        {
          "Ref": "Vpc01"
        }
      ]
    }
  },
  "TaskCount": 2,
  "TaskDefinitionArn": {
     "Ref": "ScheduledFargateTaskScheduledTaskDef521FA675"
  },
  "enableECSManagedTags": true,
  "placementConstraints": [
    {
        "expression": "attribute:ecs.instance-type == t2.small",
        "type": "memberOf"
    }
  ],
  "placementStrategy": [
  {
     "field": "instanceId",
     "type ": "binpack"
  }
  ],
  "propagateTags": "TASK_DEFINITION",
  "referenceId": "idopsks",
  "startedBy": "user1",
  "tags": [
      {
          "key": "maintask",
          "value ": "taskvalue"
      }
  ]
}
```

#### YAML

```yaml

LaunchType: "FARGATE"
NetworkConfiguration:
  AwsVpcConfiguration:
    AssignPublicIp: "DISABLED"
    SecurityGroups:
      Fn: : GetAtt:
        "ScheduledFargateTaskScheduledTaskDefSecurityGroupE075BC19",
        "GroupId"
    Subnets:
      Ref:
        "Vpc01"
TaskCount: 2
TaskDefinitionArn:
  Ref:
    "ScheduledFargateTaskScheduledTaskDef521FA675"
enableECSManagedTags: true
placementConstraints:
  expression:
    "attribute:ecs.instance-type == t2.small"
  type:
    "memberOf"
placementStrategy:
  field:
    "instanceId"
  type:
    "binpack"
propagateTags:
  "TASK_DEFINITION"
referenceId:
  "idopsks"
startedBy:
  "user1"
tags:
  key:
    "maintask"
  value:
    "taskvalue"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeadLetterConfig

HttpParameters

All content copied from https://docs.aws.amazon.com/.
