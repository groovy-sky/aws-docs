This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream ProcessingConfiguration

The `ProcessingConfiguration` property configures data processing for an
Amazon Kinesis Data Firehose delivery stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "Processors" : [ Processor, ... ]
}

```

### YAML

```yaml

  Enabled: Boolean
  Processors:
    - Processor

```

## Properties

`Enabled`

Indicates whether data processing is enabled (true) or disabled (false).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Processors`

The data processors.

_Required_: No

_Type_: Array of [Processor](aws-properties-kinesisfirehose-deliverystream-processor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PartitionSpec

Processor

All content copied from https://docs.aws.amazon.com/.
