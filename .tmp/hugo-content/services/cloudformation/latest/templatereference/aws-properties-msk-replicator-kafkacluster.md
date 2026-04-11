This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator KafkaCluster

Information about Kafka Cluster to be used as source / target for replication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmazonMskCluster" : AmazonMskCluster,
  "VpcConfig" : KafkaClusterClientVpcConfig
}

```

### YAML

```yaml

  AmazonMskCluster:
    AmazonMskCluster
  VpcConfig:
    KafkaClusterClientVpcConfig

```

## Properties

`AmazonMskCluster`

Details of an Amazon MSK Cluster.

_Required_: Yes

_Type_: [AmazonMskCluster](aws-properties-msk-replicator-amazonmskcluster.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConfig`

Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.

_Required_: Yes

_Type_: [KafkaClusterClientVpcConfig](aws-properties-msk-replicator-kafkaclusterclientvpcconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConsumerGroupReplication

KafkaClusterClientVpcConfig

All content copied from https://docs.aws.amazon.com/.
