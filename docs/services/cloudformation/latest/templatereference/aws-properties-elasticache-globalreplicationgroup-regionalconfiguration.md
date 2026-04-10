This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::GlobalReplicationGroup RegionalConfiguration

A list of the replication groups

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReplicationGroupId" : String,
  "ReplicationGroupRegion" : String,
  "ReshardingConfigurations" : [ ReshardingConfiguration, ... ]
}

```

### YAML

```yaml

  ReplicationGroupId: String
  ReplicationGroupRegion: String
  ReshardingConfigurations:
    - ReshardingConfiguration

```

## Properties

`ReplicationGroupId`

The name of the secondary cluster

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationGroupRegion`

The Amazon region where the cluster is stored

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReshardingConfigurations`

A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded
cluster.

_Required_: No

_Type_: Array of [ReshardingConfiguration](aws-properties-elasticache-globalreplicationgroup-reshardingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalReplicationGroupMember

ReshardingConfiguration

All content copied from https://docs.aws.amazon.com/.
