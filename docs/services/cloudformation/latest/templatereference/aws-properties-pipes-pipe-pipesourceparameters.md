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

_Type_: [PipeSourceActiveMQBrokerParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipesourceactivemqbrokerparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDBStreamParameters`

The parameters for using a DynamoDB stream as a source.

_Required_: No

_Type_: [PipeSourceDynamoDBStreamParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipesourcedynamodbstreamparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterCriteria`

The collection of event patterns used to filter events.

To remove a filter, specify a `FilterCriteria` object with an empty array of `Filter` objects.

For more information, see [Events and Event\
Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [FilterCriteria](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-filtercriteria.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisStreamParameters`

The parameters for using a Kinesis stream as a source.

_Required_: No

_Type_: [PipeSourceKinesisStreamParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipesourcekinesisstreamparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedStreamingKafkaParameters`

The parameters for using an MSK stream as a source.

_Required_: No

_Type_: [PipeSourceManagedStreamingKafkaParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipesourcemanagedstreamingkafkaparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RabbitMQBrokerParameters`

The parameters for using a Rabbit MQ broker as a source.

_Required_: No

_Type_: [PipeSourceRabbitMQBrokerParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipesourcerabbitmqbrokerparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfManagedKafkaParameters`

The parameters for using a self-managed Apache Kafka stream as a source.

A _self managed_ cluster refers to any Apache Kafka cluster not hosted by AWS.
This includes both clusters you manage yourself, as well as those hosted by a third-party
provider, such as [Confluent\
Cloud](https://www.confluent.io/), [CloudKarafka](https://www.cloudkarafka.com/), or [Redpanda](https://redpanda.com/). For more information, see [Apache Kafka streams as a source](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-kafka.html) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [PipeSourceSelfManagedKafkaParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipesourceselfmanagedkafkaparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqsQueueParameters`

The parameters for using a Amazon SQS stream as a source.

_Required_: No

_Type_: [PipeSourceSqsQueueParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pipes-pipe-pipesourcesqsqueueparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PipeSourceManagedStreamingKafkaParameters

PipeSourceRabbitMQBrokerParameters
