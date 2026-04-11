This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream BufferingHints

The `BufferingHints` property type specifies how Amazon Kinesis Data
Firehose (Kinesis Data Firehose) buffers incoming data before delivering it to the
destination. The first buffer condition that is satisfied triggers Kinesis Data Firehose to
deliver the data.

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

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthenticationConfiguration

CatalogConfiguration

All content copied from https://docs.aws.amazon.com/.
