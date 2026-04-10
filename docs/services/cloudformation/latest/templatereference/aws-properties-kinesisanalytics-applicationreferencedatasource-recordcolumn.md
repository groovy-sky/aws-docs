This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationReferenceDataSource RecordColumn

Describes the mapping of each data element in the streaming source to the
corresponding column in the in-application stream.

Also used to describe the format of the reference data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mapping" : String,
  "Name" : String,
  "SqlType" : String
}

```

### YAML

```yaml

  Mapping: String
  Name: String
  SqlType: String

```

## Properties

`Mapping`

Reference to the data element in the streaming input or the reference data source.
This element is required if the [RecordFormatType](../../../kinesisanalytics/latest/dev/api-recordformat.md#analytics-Type-RecordFormat-RecordFormatTypel) is `JSON`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the column created in the in-application input stream or reference
table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlType`

Type of column created in the in-application input stream or reference table.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MappingParameters

RecordFormat

All content copied from https://docs.aws.amazon.com/.
