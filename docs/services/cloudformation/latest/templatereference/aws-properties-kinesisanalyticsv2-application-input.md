This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application Input

When you configure the application input for a SQL-based Kinesis Data Analytics application, you specify the streaming source, the in-application stream
name that is created,
and the mapping between the two.

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

_Required_: No

_Type_: [InputParallelism](aws-properties-kinesisanalyticsv2-application-inputparallelism.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputProcessingConfiguration`

The [InputProcessingConfiguration](../../../managed-flink/latest/apiv2/api-inputprocessingconfiguration.md) for the input. An input processor transforms
records as they are received from the stream, before the application's SQL code
executes. Currently, the only input processing configuration available is [InputLambdaProcessor](../../../managed-flink/latest/apiv2/api-inputlambdaprocessor.md).

_Required_: No

_Type_: [InputProcessingConfiguration](aws-properties-kinesisanalyticsv2-application-inputprocessingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSchema`

Describes the format of the data in the streaming source, and how each data element maps
to corresponding columns in the in-application stream that is being created.

Also used to describe the format of the reference data source.

_Required_: Yes

_Type_: [InputSchema](aws-properties-kinesisanalyticsv2-application-inputschema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisFirehoseInput`

If the streaming source is an Amazon Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.

_Required_: No

_Type_: [KinesisFirehoseInput](aws-properties-kinesisanalyticsv2-application-kinesisfirehoseinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamsInput`

If the streaming source is an Amazon Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).

_Required_: No

_Type_: [KinesisStreamsInput](aws-properties-kinesisanalyticsv2-application-kinesisstreamsinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamePrefix`

The name prefix to use when creating an in-application stream. Suppose that you specify a
prefix " `MyInApplicationStream`." Kinesis Data Analytics then creates one or more
(as per the `InputParallelism` count you specified) in-application streams with the
names " `MyInApplicationStream_001`," " `MyInApplicationStream_002`," and
so on.

_Required_: Yes

_Type_: String

_Pattern_: `^[^-\s<>&]*$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Input](../../../managed-flink/latest/apiv2/api-input.md) in the
_Amazon Kinesis Data Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlueDataCatalogConfiguration

InputLambdaProcessor

All content copied from https://docs.aws.amazon.com/.
