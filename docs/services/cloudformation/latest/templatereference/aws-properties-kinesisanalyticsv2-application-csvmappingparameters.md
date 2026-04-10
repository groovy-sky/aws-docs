This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application CSVMappingParameters

For a SQL-based Kinesis Data Analytics application, provides additional mapping information when the record
format uses delimiters, such as CSV. For example, the following sample records use CSV format,
where the records use the _'\\n'_ as the row delimiter and a comma (",") as
the column delimiter:

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

The column delimiter. For example, in a CSV format, a comma (",") is the typical column
delimiter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordRowDelimiter`

The row delimiter. For example, in a CSV format, _'\\n'_ is the typical
row delimiter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CSVMappingParameters](../../../managed-flink/latest/apiv2/api-csvmappingparameters.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeContent

CustomArtifactConfiguration

All content copied from https://docs.aws.amazon.com/.
