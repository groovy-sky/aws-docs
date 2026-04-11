This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator

Creates the replicator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MSK::Replicator",
  "Properties" : {
      "Description" : String,
      "KafkaClusters" : [ KafkaCluster, ... ],
      "ReplicationInfoList" : [ ReplicationInfo, ... ],
      "ReplicatorName" : String,
      "ServiceExecutionRoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MSK::Replicator
Properties:
  Description: String
  KafkaClusters:
    - KafkaCluster
  ReplicationInfoList:
    - ReplicationInfo
  ReplicatorName: String
  ServiceExecutionRoleArn: String
  Tags:
    - Tag

```

## Properties

`Description`

A summary description of the replicator.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KafkaClusters`

Kafka Clusters to use in setting up sources / targets for replication.

_Required_: Yes

_Type_: Array of [KafkaCluster](aws-properties-msk-replicator-kafkacluster.md)

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicationInfoList`

A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.

_Required_: Yes

_Type_: Array of [ReplicationInfo](aws-properties-msk-replicator-replicationinfo.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicatorName`

The name of the replicator. Alpha-numeric characters with '-' are allowed.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z][0-9A-Za-z-]{0,}$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceExecutionRoleArn`

The ARN of the IAM role used by the replicator to access resources in the customer's account (e.g source and target clusters)

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):iam:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

List of tags to attach to created Replicator.

_Required_: No

_Type_: Array of [Tag](aws-properties-msk-replicator-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CurrentVersion`

The current version number of the replicator.

`ReplicatorArn`

Amazon Resource Name (ARN) for the created replicator.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LatestRevision

AmazonMskCluster

All content copied from https://docs.aws.amazon.com/.
