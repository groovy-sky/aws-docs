This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationReferenceDataSource CSVMappingParameters

Provides additional mapping information when the record format uses delimiters, such
as CSV. For example, the following sample records use CSV format, where the records use
the _'\\n'_ as the row delimiter and a comma (",") as the column
delimiter:

`"name1", "address1"`

`"name2", "address2"`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecordColumnDelimiter" : String,
  "RecordRowDelimiter" : String
}

```

### YAML

```yaml

  RecordColumnDelimiter: String
  RecordRowDelimiter: String

```

## Properties

`RecordColumnDelimiter`

Column delimiter. For example, in a CSV format, a comma (",") is the typical column
delimiter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordRowDelimiter`

Row delimiter. For example, in a CSV format, _'\\n'_ is the typical
row delimiter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KinesisAnalytics::ApplicationReferenceDataSource

JSONMappingParameters

All content copied from https://docs.aws.amazon.com/.
