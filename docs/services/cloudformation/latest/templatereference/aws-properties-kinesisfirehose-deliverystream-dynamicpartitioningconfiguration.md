This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DynamicPartitioningConfiguration

The `DynamicPartitioningConfiguration` property type specifies the
configuration of the dynamic partitioning mechanism that creates targeted data sets from
the streaming data by partitioning it based on partition keys.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "RetryOptions" : RetryOptions
}

```

### YAML

```yaml

  Enabled: Boolean
  RetryOptions:
    RetryOptions

```

## Properties

`Enabled`

Specifies whether dynamic partitioning is enabled for this Kinesis Data Firehose
delivery stream.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryOptions`

Specifies the retry behavior in case Kinesis Data Firehose is unable to deliver data
to an Amazon S3 prefix.

_Required_: No

_Type_: [RetryOptions](aws-properties-kinesisfirehose-deliverystream-retryoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentIdOptions

ElasticsearchBufferingHints

All content copied from https://docs.aws.amazon.com/.
