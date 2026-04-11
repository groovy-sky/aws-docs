This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping AmazonManagedKafkaEventSourceConfig

Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConsumerGroupId" : String,
  "SchemaRegistryConfig" : SchemaRegistryConfig
}

```

### YAML

```yaml

  ConsumerGroupId: String
  SchemaRegistryConfig:
    SchemaRegistryConfig

```

## Properties

`ConsumerGroupId`

The identifier for the Kafka consumer group to join. The consumer group ID must be unique among all your Kafka event sources.
After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value. For more information, see
[Customizable consumer group ID](../../../lambda/latest/dg/with-msk.md#services-msk-consumer-group-id).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9-\/*:_+=.@-]*`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaRegistryConfig`

Specific configuration settings for a Kafka schema registry.

_Required_: No

_Type_: [SchemaRegistryConfig](aws-properties-lambda-eventsourcemapping-schemaregistryconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lambda::EventSourceMapping

DestinationConfig

All content copied from https://docs.aws.amazon.com/.
