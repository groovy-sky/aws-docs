This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application RecordFormat

For a SQL-based Kinesis Data Analytics application, describes the record format
and relevant mapping information that should be applied to schematize the records on the
stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MappingParameters" : MappingParameters,
  "RecordFormatType" : String
}

```

### YAML

```yaml

  MappingParameters:
    MappingParameters
  RecordFormatType: String

```

## Properties

`MappingParameters`

When you configure application input at the time of creating or updating an application,
provides additional mapping information specific to the record format (such as JSON, CSV, or
record fields delimited by some delimiter) on the streaming source.

_Required_: No

_Type_: [MappingParameters](aws-properties-kinesisanalyticsv2-application-mappingparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordFormatType`

The type of record format.

_Required_: Yes

_Type_: String

_Allowed values_: `CSV | JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [RecordFormat](../../../managed-flink/latest/apiv2/api-recordformat.md) in the _Amazon Kinesis Data Analytics API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordColumn

RunConfiguration

All content copied from https://docs.aws.amazon.com/.
