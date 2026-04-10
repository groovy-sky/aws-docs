This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeSourceSelfManagedKafkaParameters

The parameters for using a self-managed Apache Kafka stream as a source.

A _self managed_ cluster refers to any Apache Kafka cluster not hosted by AWS.
This includes both clusters you manage yourself, as well as those hosted by a third-party
provider, such as [Confluent\
Cloud](https://www.confluent.io/), [CloudKarafka](https://www.cloudkarafka.com/), or [Redpanda](https://redpanda.com/). For more information, see [Apache Kafka streams as a source](../../../eventbridge/latest/userguide/eb-pipes-kafka.md) in the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalBootstrapServers" : [ String, ... ],
  "BatchSize" : Integer,
  "ConsumerGroupID" : String,
  "Credentials" : SelfManagedKafkaAccessConfigurationCredentials,
  "MaximumBatchingWindowInSeconds" : Integer,
  "ServerRootCaCertificate" : String,
  "StartingPosition" : String,
  "TopicName" : String,
  "Vpc" : SelfManagedKafkaAccessConfigurationVpc
}

```

### YAML

```yaml

  AdditionalBootstrapServers:
    - String
  BatchSize: Integer
  ConsumerGroupID: String
  Credentials:
    SelfManagedKafkaAccessConfigurationCredentials
  MaximumBatchingWindowInSeconds: Integer
  ServerRootCaCertificate: String
  StartingPosition: String
  TopicName: String
  Vpc:
    SelfManagedKafkaAccessConfigurationVpc

```

## Properties

`AdditionalBootstrapServers`

An array of server URLs.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `300 | 2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BatchSize`

The maximum number of records to include in each batch.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConsumerGroupID`

The name of the destination queue to consume.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-\/*:_+=.@-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Credentials`

The credentials needed to access the resource.

_Required_: No

_Type_: [SelfManagedKafkaAccessConfigurationCredentials](aws-properties-pipes-pipe-selfmanagedkafkaaccessconfigurationcredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBatchingWindowInSeconds`

The maximum length of a time to wait for events.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerRootCaCertificate`

The ARN of the Secrets Manager secret used for certification.

_Required_: No

_Type_: String

_Pattern_: `^(^arn:aws([a-z]|\-)*:secretsmanager:([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}):(\d{12}):secret:.+)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartingPosition`

The position in a stream from which to start reading.

_Required_: No

_Type_: String

_Allowed values_: `TRIM_HORIZON | LATEST`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicName`

The name of the topic that the pipe will read from.

_Required_: Yes

_Type_: String

_Pattern_: `^[^.]([a-zA-Z0-9\-_.]+)$`

_Minimum_: `1`

_Maximum_: `249`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Vpc`

This structure specifies the VPC subnets and security groups for the stream, and whether a public IP address is to be used.

_Required_: No

_Type_: [SelfManagedKafkaAccessConfigurationVpc](aws-properties-pipes-pipe-selfmanagedkafkaaccessconfigurationvpc.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeSourceRabbitMQBrokerParameters

PipeSourceSqsQueueParameters

All content copied from https://docs.aws.amazon.com/.
