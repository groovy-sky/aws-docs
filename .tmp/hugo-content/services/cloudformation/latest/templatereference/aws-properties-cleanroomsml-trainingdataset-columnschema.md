This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::TrainingDataset ColumnSchema

Metadata for a column.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "ColumnTypes" : [ String, ... ]
}

```

### YAML

```yaml

  ColumnName: String
  ColumnTypes:
    - String

```

## Properties

`ColumnName`

The name of a column.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_](([a-zA-Z0-9_ ]+-)*([a-zA-Z0-9_ ]+))?$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ColumnTypes`

The data type of column.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CleanRoomsML::TrainingDataset

Dataset

All content copied from https://docs.aws.amazon.com/.
