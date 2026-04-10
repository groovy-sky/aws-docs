This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeSourceParameters

The parameters required to set up a source for your pipe.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActiveMQBrokerParameters" : PipeSourceActiveMQBrokerParameters,
  "DynamoDBStreamParameters" : PipeSourceDynamoDBStreamParameters,
  "FilterCriteria" : FilterCriteria,
  "KinesisStreamParameters" : PipeSourceKinesisStreamParameters,
  "ManagedStreamingKafkaParameters" : PipeSourceManagedStreamingKafkaParameters,
  "RabbitMQBrokerParameters" : PipeSourceRabbitMQBrokerParameters,
  "SelfManagedKafkaParameters" : PipeSourceSelfManagedKafkaParameters,
  "SqsQueueParameters" : PipeSourceSqsQueueParameters
}

```

### YAML

```yaml

  ActiveMQBrokerParameters:
    PipeSourceActiveMQBrokerParameters
  DynamoDBStreamParameters:
    PipeSourceDynamoDBStreamParameters
  FilterCriteria:
    FilterCriteria
  KinesisStreamParameters:
    PipeSourceKinesisStreamParameters
  ManagedStreamingKafkaParameters:
    PipeSourceManagedStreamingKafkaParameters
  RabbitMQBrokerParameters:
    PipeSourceRabbitMQBrokerParameters
  SelfManagedKafkaParameters:
    PipeSourceSelfManagedKafkaParameters
  SqsQueueParameters:
    PipeSourceSqsQueueParameters

```

## Properties

`ActiveMQBrokerParameters`

The parameters for using an Active MQ broker as a source.

_Required_: No

_Type_: [PipeSourceActiveMQBrokerParameters](aws-properties-pipes-pipe-pipesourceactivemqbrokerparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDBStreamParameters`

The parameters for using a DynamoDB stream as a source.

_Required_: No

_Type_: [PipeSourceDynamoDBStreamParameters](aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterCriteria`

The collection of event patterns used to filter events.

To remove a filter, specify a `FilterCriteria` object with an empty array of `Filter` objects.

For more information, see [Events and Event\
Patterns](../../../eventbridge/latest/userguide/eventbridge-and-event-patterns.md) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [FilterCriteria](aws-properties-pipes-pipe-filtercriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamParameters`

The parameters for using a Kinesis stream as a source.

_Required_: No

_Type_: [PipeSourceKinesisStreamParameters](aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedStreamingKafkaParameters`

The parameters for using an MSK stream as a source.

_Required_: No

_Type_: [PipeSourceManagedStreamingKafkaParameters](aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RabbitMQBrokerParameters`

The parameters for using a Rabbit MQ broker as a source.

_Required_: No

_Type_: [PipeSourceRabbitMQBrokerParameters](aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfManagedKafkaParameters`

The parameters for using a self-managed Apache Kafka stream as a source.

A _self managed_ cluster refers to any Apache Kafka cluster not hosted by AWS.
This includes both clusters you manage yourself, as well as those hosted by a third-party
provider, such as [Confluent\
Cloud](https://www.confluent.io/), [CloudKarafka](https://www.cloudkarafka.com/), or [Redpanda](https://redpanda.com/). For more information, see [Apache Kafka streams as a source](../../../eventbridge/latest/userguide/eb-pipes-kafka.md) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [PipeSourceSelfManagedKafkaParameters](aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqsQueueParameters`

The parameters for using a Amazon SQS stream as a source.

_Required_: No

_Type_: [PipeSourceSqsQueueParameters](aws-properties-pipes-pipe-pipesourcesqsqueueparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeSourceManagedStreamingKafkaParameters

PipeSourceRabbitMQBrokerParameters

All content copied from https://docs.aws.amazon.com/.
