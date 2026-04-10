This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationOutput Output

Describes a SQL-based Kinesis Data Analytics application's output configuration,
in which you identify an in-application stream and a destination where you want the
in-application stream data to be written. The destination can be a Kinesis data stream or a
Kinesis Data Firehose delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationSchema" : DestinationSchema,
  "KinesisFirehoseOutput" : KinesisFirehoseOutput,
  "KinesisStreamsOutput" : KinesisStreamsOutput,
  "LambdaOutput" : LambdaOutput,
  "Name" : String
}

```

### YAML

```yaml

  DestinationSchema:
    DestinationSchema
  KinesisFirehoseOutput:
    KinesisFirehoseOutput
  KinesisStreamsOutput:
    KinesisStreamsOutput
  LambdaOutput:
    LambdaOutput
  Name: String

```

## Properties

`DestinationSchema`

Describes the data format when records are written to the destination.

_Required_: Yes

_Type_: [DestinationSchema](aws-properties-kinesisanalyticsv2-applicationoutput-destinationschema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisFirehoseOutput`

Identifies a Kinesis Data Firehose delivery stream as the destination.

_Required_: No

_Type_: [KinesisFirehoseOutput](aws-properties-kinesisanalyticsv2-applicationoutput-kinesisfirehoseoutput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamsOutput`

Identifies a Kinesis data stream
as the destination.

_Required_: No

_Type_: [KinesisStreamsOutput](aws-properties-kinesisanalyticsv2-applicationoutput-kinesisstreamsoutput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaOutput`

Identifies an Amazon Lambda function as the destination.

_Required_: No

_Type_: [LambdaOutput](aws-properties-kinesisanalyticsv2-applicationoutput-lambdaoutput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the in-application stream.

_Required_: No

_Type_: String

_Pattern_: `[^-\s<>&]*`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Output](../../../managed-flink/latest/apiv2/api-output.md) in the
_Amazon Kinesis Data Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaOutput

AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource

All content copied from https://docs.aws.amazon.com/.
