This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator ReplicationInfo

Specifies configuration for replication between a source and target Kafka cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConsumerGroupReplication" : ConsumerGroupReplication,
  "SourceKafkaClusterArn" : String,
  "TargetCompressionType" : String,
  "TargetKafkaClusterArn" : String,
  "TopicReplication" : TopicReplication
}

```

### YAML

```yaml

  ConsumerGroupReplication:
    ConsumerGroupReplication
  SourceKafkaClusterArn: String
  TargetCompressionType: String
  TargetKafkaClusterArn: String
  TopicReplication:
    TopicReplication

```

## Properties

`ConsumerGroupReplication`

Configuration relating to consumer group replication.

_Required_: Yes

_Type_: [ConsumerGroupReplication](aws-properties-msk-replicator-consumergroupreplication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceKafkaClusterArn`

The ARN of the source Kafka cluster.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):kafka:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetCompressionType`

The compression type to use when producing records to target cluster.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | GZIP | SNAPPY | LZ4 | ZSTD`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetKafkaClusterArn`

The ARN of the target Kafka cluster.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):kafka:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicReplication`

Configuration relating to topic replication.

_Required_: Yes

_Type_: [TopicReplication](aws-properties-msk-replicator-topicreplication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KafkaClusterClientVpcConfig

ReplicationStartingPosition

All content copied from https://docs.aws.amazon.com/.
