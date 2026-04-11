This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet InputColumn

Metadata for a column that is used as the input of a transform operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Name" : String,
  "SubType" : String,
  "Type" : String
}

```

### YAML

```yaml

  Id: String
  Name: String
  SubType: String
  Type: String

```

## Properties

`Id`

A unique identifier for the input column.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of this column in the underlying data source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubType`

The sub data type of the column. Sub types are only available for decimal columns that are part of a SPICE dataset.

_Required_: No

_Type_: String

_Allowed values_: `FLOAT | FIXED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The data type of the column.

**Note:** `SEMISTRUCT` represents Athena's map, row, and struct data types. It is supported when using the new data preparation experience.

_Required_: Yes

_Type_: String

_Allowed values_: `STRING | INTEGER | DECIMAL | DATETIME | BIT | BOOLEAN | JSON | SEMISTRUCT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngestionWaitPolicy

IntegerDatasetParameter

All content copied from https://docs.aws.amazon.com/.
