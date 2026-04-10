This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::Application Input

When you configure the application input, you specify the streaming source, the
in-application stream name that is created, and the mapping between the two. For more
information, see [Configuring Application\
Input](../../../kinesisanalytics/latest/dev/how-it-works-input.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputParallelism" : InputParallelism,
  "InputProcessingConfiguration" : InputProcessingConfiguration,
  "InputSchema" : InputSchema,
  "KinesisFirehoseInput" : KinesisFirehoseInput,
  "KinesisStreamsInput" : KinesisStreamsInput,
  "NamePrefix" : String
}

```

### YAML

```yaml

  InputParallelism:
    InputParallelism
  InputProcessingConfiguration:
    InputProcessingConfiguration
  InputSchema:
    InputSchema
  KinesisFirehoseInput:
    KinesisFirehoseInput
  KinesisStreamsInput:
    KinesisStreamsInput
  NamePrefix: String

```

## Properties

`InputParallelism`

Describes the number of in-application streams to create.

Data from your source is routed to these in-application input streams.

See [Configuring Application Input](../../../kinesisanalytics/latest/dev/how-it-works-input.md).

_Required_: No

_Type_: [InputParallelism](aws-properties-kinesisanalytics-application-inputparallelism.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputProcessingConfiguration`

The [InputProcessingConfiguration](../userguide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.md) for the input. An input
processor transforms records as they are received from the stream, before the
application's SQL code executes. Currently, the only input processing configuration
available is [InputLambdaProcessor](../userguide/aws-properties-kinesisanalytics-application-inputlambdaprocessor.md).

_Required_: No

_Type_: [InputProcessingConfiguration](aws-properties-kinesisanalytics-application-inputprocessingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSchema`

Describes the format of the data in the streaming source, and how each data element
maps to corresponding columns in the in-application stream that is being created.

Also used to describe the format of the reference data source.

_Required_: Yes

_Type_: [InputSchema](aws-properties-kinesisanalytics-application-inputschema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisFirehoseInput`

If the streaming source is an Amazon Kinesis Firehose delivery stream, identifies the
delivery stream's ARN and an IAM role that enables Amazon Kinesis Analytics to access
the stream on your behalf.

Note: Either `KinesisStreamsInput` or `KinesisFirehoseInput` is
required.

_Required_: Conditional

_Type_: [KinesisFirehoseInput](aws-properties-kinesisanalytics-application-kinesisfirehoseinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamsInput`

If the streaming source is an Amazon Kinesis stream, identifies the stream's Amazon
Resource Name (ARN) and an IAM role that enables Amazon Kinesis Analytics to access the
stream on your behalf.

Note: Either `KinesisStreamsInput` or `KinesisFirehoseInput` is
required.

_Required_: Conditional

_Type_: [KinesisStreamsInput](aws-properties-kinesisanalytics-application-kinesisstreamsinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamePrefix`

Name prefix to use when creating an in-application stream. Suppose that you specify a
prefix "MyInApplicationStream." Amazon Kinesis Analytics then creates one or more (as
per the `InputParallelism` count you specified) in-application streams with
names "MyInApplicationStream\_001," "MyInApplicationStream\_002," and so on.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CSVMappingParameters

InputLambdaProcessor

All content copied from https://docs.aws.amazon.com/.
