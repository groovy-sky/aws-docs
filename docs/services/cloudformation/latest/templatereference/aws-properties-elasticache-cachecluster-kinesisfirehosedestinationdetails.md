This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::CacheCluster KinesisFirehoseDestinationDetails

The configuration details of the Kinesis Data Firehose destination. Note that this field is marked as required
but only if Kinesis Data Firehose was chosen as the destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryStream" : String
}

```

### YAML

```yaml

  DeliveryStream: String

```

## Properties

`DeliveryStream`

The name of the Kinesis Data Firehose delivery stream.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationDetails

LogDeliveryConfigurationRequest

All content copied from https://docs.aws.amazon.com/.
