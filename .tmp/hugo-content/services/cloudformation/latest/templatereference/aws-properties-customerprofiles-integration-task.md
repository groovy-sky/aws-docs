This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration Task

The `Task` property type specifies the class for modeling different type of
tasks. Task implementation varies based on the TaskType.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectorOperator" : ConnectorOperator,
  "DestinationField" : String,
  "SourceFields" : [ String, ... ],
  "TaskProperties" : [ TaskPropertiesMap, ... ],
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
    - TaskPropertiesMap
  TaskType: String

```

## Properties

`ConnectorOperator`

The operation to be performed on the provided source fields.

_Required_: No

_Type_: [ConnectorOperator](aws-properties-customerprofiles-integration-connectoroperator.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationField`

A field in a destination connector, or a field value against which Amazon AppFlow
validates a source field.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFields`

The source fields to which a particular task is applied.

_Required_: Yes

_Type_: Array of String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskProperties`

A map used to store task-related information. The service looks for particular
information based on the TaskType.

_Required_: No

_Type_: Array of [TaskPropertiesMap](aws-properties-customerprofiles-integration-taskpropertiesmap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskType`

Specifies the particular task implementation that Amazon AppFlow performs.

_Required_: Yes

_Type_: String

_Allowed values_: `Arithmetic | Filter | Map | Mask | Merge | Truncate | Validate`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TaskPropertiesMap

All content copied from https://docs.aws.amazon.com/.
