This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::GlobalReplicationGroup GlobalReplicationGroupMember

A member of a Global datastore. It contains the Replication Group Id, the Amazon
region and the role of the replication group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReplicationGroupId" : String,
  "ReplicationGroupRegion" : String,
  "Role" : String
}

```

### YAML

```yaml

  ReplicationGroupId: String
  ReplicationGroupRegion: String
  Role: String

```

## Properties

`ReplicationGroupId`

The replication group id of the Global datastore member.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationGroupRegion`

The Amazon region of the Global datastore member.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

Indicates the role of the replication group, `PRIMARY` or `SECONDARY`.

_Required_: No

_Type_: String

_Allowed values_: `PRIMARY | SECONDARY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ElastiCache::GlobalReplicationGroup

RegionalConfiguration
