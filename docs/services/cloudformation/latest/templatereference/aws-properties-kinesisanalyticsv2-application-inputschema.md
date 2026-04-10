This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application InputSchema

For a SQL-based Kinesis Data Analytics application, describes the format of the
data in the streaming source, and how each data element maps to corresponding columns created
in the in-application stream.

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

A list of `RecordColumn` objects.

_Required_: Yes

_Type_: Array of [RecordColumn](aws-properties-kinesisanalyticsv2-application-recordcolumn.md)

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordEncoding`

Specifies the encoding of the records in the streaming source. For example, UTF-8.

_Required_: No

_Type_: String

_Allowed values_: `UTF-8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordFormat`

Specifies the format of the records on the streaming source.

_Required_: Yes

_Type_: [RecordFormat](aws-properties-kinesisanalyticsv2-application-recordformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SourceSchema](../../../managed-flink/latest/apiv2/api-sourceschema.md) in the _Amazon Kinesis Data Analytics API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputProcessingConfiguration

JSONMappingParameters

All content copied from https://docs.aws.amazon.com/.
