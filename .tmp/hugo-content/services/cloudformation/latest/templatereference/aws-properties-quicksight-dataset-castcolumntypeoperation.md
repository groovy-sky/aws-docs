This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet CastColumnTypeOperation

A transform operation that casts a column to a different type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "Format" : String,
  "NewColumnType" : String,
  "SubType" : String
}

```

### YAML

```yaml

  ColumnName: String
  Format: String
  NewColumnType: String
  SubType: String

```

## Properties

`ColumnName`

Column name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

When casting a column from string to datetime type, you can supply a string in a
format supported by Quick Sight to denote the source data format.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NewColumnType`

New column data type.

_Required_: Yes

_Type_: String

_Allowed values_: `STRING | INTEGER | DECIMAL | DATETIME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubType`

The sub data type of the new column. Sub types are only available for decimal columns that are part of a SPICE dataset.

_Required_: No

_Type_: String

_Allowed values_: `FLOAT | FIXED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CalculatedColumn

CastColumnTypesOperation

All content copied from https://docs.aws.amazon.com/.
