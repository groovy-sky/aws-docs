This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream Processor

The `Processor` property specifies a data processor for an Amazon Kinesis
Data Firehose delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameters" : [ ProcessorParameter, ... ],
  "Type" : String
}

```

### YAML

```yaml

  Parameters:
    - ProcessorParameter
  Type: String

```

## Properties

`Parameters`

The processor parameters.

_Required_: No

_Type_: Array of [ProcessorParameter](aws-properties-kinesisfirehose-deliverystream-processorparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of processor. Valid values: `Lambda`.

_Required_: Yes

_Type_: String

_Allowed values_: `RecordDeAggregation | Decompression | CloudWatchLogProcessing | Lambda | MetadataExtraction | AppendDelimiterToRecord`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProcessingConfiguration

ProcessorParameter

All content copied from https://docs.aws.amazon.com/.
