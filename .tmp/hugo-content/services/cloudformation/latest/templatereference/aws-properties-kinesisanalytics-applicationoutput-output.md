This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationOutput Output

Describes application output configuration in which you identify an in-application
stream and a destination where you want the in-application stream data to be written.
The destination can be an Amazon Kinesis stream or an Amazon Kinesis Firehose delivery
stream.

For limits on how many destinations an application can write and other limitations,
see [Limits](../../../kinesisanalytics/latest/dev/limits.md).

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

Describes the data format when records are written to the destination. For more
information, see [Configuring Application\
Output](../../../kinesisanalytics/latest/dev/how-it-works-output.md).

_Required_: Yes

_Type_: [DestinationSchema](aws-properties-kinesisanalytics-applicationoutput-destinationschema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisFirehoseOutput`

Identifies an Amazon Kinesis Firehose delivery stream as the destination.

_Required_: No

_Type_: [KinesisFirehoseOutput](aws-properties-kinesisanalytics-applicationoutput-kinesisfirehoseoutput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamsOutput`

Identifies an Amazon Kinesis stream as the destination.

_Required_: No

_Type_: [KinesisStreamsOutput](aws-properties-kinesisanalytics-applicationoutput-kinesisstreamsoutput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaOutput`

Identifies an AWS Lambda function as the destination.

_Required_: No

_Type_: [LambdaOutput](aws-properties-kinesisanalytics-applicationoutput-lambdaoutput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the in-application stream.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaOutput

AWS::KinesisAnalytics::ApplicationReferenceDataSource

All content copied from https://docs.aws.amazon.com/.
