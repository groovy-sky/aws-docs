This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule KafkaAction

Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientProperties" : {Key: Value, ...},
  "DestinationArn" : String,
  "Headers" : [ KafkaActionHeader, ... ],
  "Key" : String,
  "Partition" : String,
  "Topic" : String
}

```

### YAML

```yaml

  ClientProperties:
    Key: Value
  DestinationArn: String
  Headers:
    - KafkaActionHeader
  Key: String
  Partition: String
  Topic: String

```

## Properties

`ClientProperties`

Properties of the Apache Kafka producer client.

_Required_: Yes

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationArn`

The ARN of Kafka action's VPC `TopicRuleDestination`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Headers`

The list of Kafka headers that you specify.

_Required_: No

_Type_: Array of [KafkaActionHeader](aws-properties-iot-topicrule-kafkaactionheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The Kafka message key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Partition`

The Kafka message partition.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Topic`

The Kafka topic for messages to be sent to the Kafka broker.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IotSiteWiseAction

KafkaActionHeader

All content copied from https://docs.aws.amazon.com/.
