This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator TopicReplication

Details about topic replication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopyAccessControlListsForTopics" : Boolean,
  "CopyTopicConfigurations" : Boolean,
  "DetectAndCopyNewTopics" : Boolean,
  "StartingPosition" : ReplicationStartingPosition,
  "TopicNameConfiguration" : ReplicationTopicNameConfiguration,
  "TopicsToExclude" : [ String, ... ],
  "TopicsToReplicate" : [ String, ... ]
}

```

### YAML

```yaml

  CopyAccessControlListsForTopics: Boolean
  CopyTopicConfigurations: Boolean
  DetectAndCopyNewTopics: Boolean
  StartingPosition:
    ReplicationStartingPosition
  TopicNameConfiguration:
    ReplicationTopicNameConfiguration
  TopicsToExclude:
    - String
  TopicsToReplicate:
    - String

```

## Properties

`CopyAccessControlListsForTopics`

Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTopicConfigurations`

Whether to periodically configure remote topics to match their corresponding upstream topics.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectAndCopyNewTopics`

Whether to periodically check for new topics and partitions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartingPosition`

Specifies the position in the topics to start replicating from.

_Required_: No

_Type_: [ReplicationStartingPosition](aws-properties-msk-replicator-replicationstartingposition.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicNameConfiguration`

Configuration for specifying replicated topic names will be the same as their corresponding upstream topics or prefixed with source cluster alias.

_Required_: No

_Type_: [ReplicationTopicNameConfiguration](aws-properties-msk-replicator-replicationtopicnameconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicsToExclude`

List of regular expression patterns indicating the topics that should not be replicated.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `249 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicsToReplicate`

List of regular expression patterns indicating the topics to copy.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `249 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::MSK::ServerlessCluster

All content copied from https://docs.aws.amazon.com/.
