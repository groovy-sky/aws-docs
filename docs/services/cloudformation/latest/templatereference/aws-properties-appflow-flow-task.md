This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow Task

A class for modeling different type of tasks. Task implementation varies based on the
`TaskType`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectorOperator" : ConnectorOperator,
  "DestinationField" : String,
  "SourceFields" : [ String, ... ],
  "TaskProperties" : [ TaskPropertiesObject, ... ],
  "TaskType" : String
}

```

### YAML

```yaml

  ConnectorOperator:
    ConnectorOperator
  DestinationField: String
  SourceFields:
    - String
  TaskProperties:
    - TaskPropertiesObject
  TaskType: String

```

## Properties

`ConnectorOperator`

The operation to be performed on the provided source fields.

_Required_: No

_Type_: [ConnectorOperator](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-connectoroperator.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationField`

A field in a destination connector, or a field value against which Amazon AppFlow
validates a source field.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFields`

The source fields to which a particular task is applied.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskProperties`

A map used to store task-related information. The execution service looks for particular
information based on the `TaskType`.

_Required_: No

_Type_: Array of [TaskPropertiesObject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-taskpropertiesobject.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskType`

Specifies the particular task implementation that Amazon AppFlow performs.

_Allowed values_: `Arithmetic` \| `Filter` \|
`Map` \| `Map_all` \| `Mask` \| `Merge` \|
`Truncate` \| `Validate`

_Required_: Yes

_Type_: String

_Allowed values_: `Arithmetic | Filter | Map | Map_all | Mask | Merge | Passthrough | Truncate | Validate | Partition`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Task](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_Task.html) in the _Amazon AppFlow API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

TaskPropertiesObject
