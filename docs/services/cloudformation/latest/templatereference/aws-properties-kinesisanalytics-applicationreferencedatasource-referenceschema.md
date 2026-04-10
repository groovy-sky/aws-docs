This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationReferenceDataSource ReferenceSchema

The ReferenceSchema property type specifies the format of the data in the reference source for a SQL-based Amazon Kinesis Data Analytics application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecordColumns" : [ RecordColumn, ... ],
  "RecordEncoding" : String,
  "RecordFormat" : RecordFormat
}

```

### YAML

```yaml

  RecordColumns:
    - RecordColumn
  RecordEncoding: String
  RecordFormat:
    RecordFormat

```

## Properties

`RecordColumns`

A list of RecordColumn objects.

_Required_: Yes

_Type_: Array of [RecordColumn](aws-properties-kinesisanalytics-applicationreferencedatasource-recordcolumn.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordEncoding`

Specifies the encoding of the records in the reference source. For example, UTF-8.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordFormat`

Specifies the format of the records on the reference source.

_Required_: Yes

_Type_: [RecordFormat](aws-properties-kinesisanalytics-applicationreferencedatasource-recordformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceDataSource

S3ReferenceDataSource

All content copied from https://docs.aws.amazon.com/.
