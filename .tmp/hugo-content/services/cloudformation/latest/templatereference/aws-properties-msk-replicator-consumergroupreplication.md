This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator ConsumerGroupReplication

Details about consumer group replication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConsumerGroupsToExclude" : [ String, ... ],
  "ConsumerGroupsToReplicate" : [ String, ... ],
  "DetectAndCopyNewConsumerGroups" : Boolean,
  "SynchroniseConsumerGroupOffsets" : Boolean
}

```

### YAML

```yaml

  ConsumerGroupsToExclude:
    - String
  ConsumerGroupsToReplicate:
    - String
  DetectAndCopyNewConsumerGroups: Boolean
  SynchroniseConsumerGroupOffsets: Boolean

```

## Properties

`ConsumerGroupsToExclude`

List of regular expression patterns indicating the consumer groups that should not be replicated.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConsumerGroupsToReplicate`

List of regular expression patterns indicating the consumer groups to copy.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `256 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectAndCopyNewConsumerGroups`

Enables synchronization of consumer groups to target cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SynchroniseConsumerGroupOffsets`

Enables synchronization of consumer group offsets to target cluster. The translated offsets will be written to topic \_\_consumer\_offsets.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmazonMskCluster

KafkaCluster

All content copied from https://docs.aws.amazon.com/.
