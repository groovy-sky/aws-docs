This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DataPathSort

Allows data paths to be sorted by a specific data value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Direction" : String,
  "SortPaths" : [ DataPathValue, ... ]
}

```

### YAML

```yaml

  Direction: String
  SortPaths:
    - DataPathValue

```

## Properties

`Direction`

Determines the sort direction.

_Required_: Yes

_Type_: String

_Allowed values_: `ASC | DESC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortPaths`

The list of data paths that need to be sorted.

_Required_: Yes

_Type_: Array of [DataPathValue](aws-properties-quicksight-analysis-datapathvalue.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPathLabelType

DataPathType

All content copied from https://docs.aws.amazon.com/.
