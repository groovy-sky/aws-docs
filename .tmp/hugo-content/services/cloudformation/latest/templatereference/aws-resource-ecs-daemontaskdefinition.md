This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::DaemonTaskDefinition

The details of a daemon task definition. A daemon task definition is a template that
describes the containers that form a daemon. Daemons deploy cross-cutting software
agents independently across your Amazon ECS infrastructure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECS::DaemonTaskDefinition",
  "Properties" : {
      "ContainerDefinitions" : [ DaemonContainerDefinition, ... ],
      "Cpu" : String,
      "ExecutionRoleArn" : String,
      "Family" : String,
      "Memory" : String,
      "Tags" : [ Tag, ... ],
      "TaskRoleArn" : String,
      "Volumes" : [ Volume, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ECS::DaemonTaskDefinition
Properties:
  ContainerDefinitions:
    - DaemonContainerDefinition
  Cpu: String
  ExecutionRoleArn: String
  Family: String
  Memory: String
  Tags:
    - Tag
  TaskRoleArn: String
  Volumes:
    - Volume

```

## Properties

`ContainerDefinitions`

A list of container definitions in JSON format that describe the containers that make
up the daemon task.

_Required_: No

_Type_: Array of [DaemonContainerDefinition](aws-properties-ecs-daemontaskdefinition-daemoncontainerdefinition.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Cpu`

The number of CPU units used by the daemon task.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS
container agent permission to make Amazon Web Services API calls on your behalf.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Family`

The name of a family that this daemon task definition is registered to.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Memory`

The amount of memory (in MiB) used by the daemon task.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-ecs-daemontaskdefinition-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskRoleArn`

The short name or full Amazon Resource Name (ARN) of the IAM role that grants
containers in the daemon task permission to call Amazon Web Services APIs on your
behalf.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Volumes`

The list of data volume definitions for the daemon task.

_Required_: No

_Type_: Array of [Volume](aws-properties-ecs-daemontaskdefinition-volume.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`DaemonTaskDefinitionArn`

The full Amazon Resource Name (ARN) of the daemon task definition.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ContainerDependency

All content copied from https://docs.aws.amazon.com/.
