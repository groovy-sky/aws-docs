This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream ElasticsearchBufferingHints

The `ElasticsearchBufferingHints` property type specifies how Amazon
Kinesis Data Firehose (Kinesis Data Firehose) buffers incoming data while delivering it to
the destination. The first buffer condition that is satisfied triggers Kinesis Data
Firehose to deliver the data.

ElasticsearchBufferingHints is the property type for the `BufferingHints`
property of the [Amazon Kinesis Data Firehose DeliveryStream\
ElasticsearchDestinationConfiguration](../userguide/aws-properties-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IntervalInSeconds" : Integer,
  "SizeInMBs" : Integer
}

```

### YAML

```yaml

  IntervalInSeconds: Integer
  SizeInMBs: Integer

```

## Properties

`IntervalInSeconds`

The length of time, in seconds, that Kinesis Data Firehose buffers incoming data
before delivering it to the destination. For valid values, see the
`IntervalInSeconds` content for the [BufferingHints](../../../../reference/firehose/latest/apireference/api-bufferinghints.md) data
type in the _Amazon Kinesis Data Firehose API Reference_.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeInMBs`

The size of the buffer, in MBs, that Kinesis Data Firehose uses for incoming data
before delivering it to the destination. For valid values, see the `SizeInMBs`
content for the [BufferingHints](../../../../reference/firehose/latest/apireference/api-bufferinghints.md) data
type in the _Amazon Kinesis Data Firehose API Reference_.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamicPartitioningConfiguration

ElasticsearchDestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
