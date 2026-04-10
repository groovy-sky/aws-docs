This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector FirehoseLogDelivery

The settings for delivering logs to Amazon Kinesis Data Firehose.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryStream" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  DeliveryStream: String
  Enabled: Boolean

```

## Properties

`DeliveryStream`

The name of the Kinesis Data Firehose delivery stream that is the destination for log
delivery.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Enabled`

Specifies whether connector logs get delivered to Amazon Kinesis Data Firehose.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomPlugin

KafkaCluster

All content copied from https://docs.aws.amazon.com/.
