This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationOutput DestinationSchema

Describes the data format when records are written to the destination in a SQL-based Kinesis Data Analytics application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecordFormatType" : String
}

```

### YAML

```yaml

  RecordFormatType: String

```

## Properties

`RecordFormatType`

Specifies the format of the records on the output stream.

_Required_: No

_Type_: String

_Allowed values_: `JSON | CSV`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DestinationSchema](../../../managed-flink/latest/apiv2/api-destinationschema.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KinesisAnalyticsV2::ApplicationOutput

KinesisFirehoseOutput

All content copied from https://docs.aws.amazon.com/.
