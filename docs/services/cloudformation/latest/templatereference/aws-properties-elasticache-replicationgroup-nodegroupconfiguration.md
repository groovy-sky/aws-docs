This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::ReplicationGroup NodeGroupConfiguration

`NodeGroupConfiguration` is a property of the `AWS::ElastiCache::ReplicationGroup`
resource that configures an Amazon ElastiCache (ElastiCache) Valkey or Redis OSS cluster node group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NodeGroupId" : String,
  "PrimaryAvailabilityZone" : String,
  "ReplicaAvailabilityZones" : [ String, ... ],
  "ReplicaCount" : Integer,
  "Slots" : String
}

```

### YAML

```yaml

  NodeGroupId: String
  PrimaryAvailabilityZone: String
  ReplicaAvailabilityZones:
    - String
  ReplicaCount: Integer
  Slots: String

```

## Properties

`NodeGroupId`

Either the ElastiCache supplied 4-digit id or a user supplied id for the
node group these configuration values apply to.

_Required_: No

_Type_: String

_Pattern_: `\d+`

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryAvailabilityZone`

The Availability Zone where the primary node of this node group (shard) is
launched.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicaAvailabilityZones`

A list of Availability Zones to be used for the read replicas. The number of
Availability Zones in this list must match the value of `ReplicaCount` or
`ReplicasPerNodeGroup` if not specified.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicaCount`

The number of read replica nodes in this node group (shard).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slots`

A string of comma-separated values where the first set of values are the slot numbers (zero based), and the
second set of values are the keyspaces for each slot. The following example specifies three slots (numbered 0, 1, and
2): ` 0,1,2,0-4999,5000-9999,10000-16,383`.

If you don't specify a value, ElastiCache allocates keys equally among each slot.

When you use an `UseOnlineResharding` update policy to update the number of node groups without
interruption, ElastiCache evenly distributes the keyspaces between the specified number of slots. This cannot be
updated later. Therefore, after updating the number of node groups in this way, you should remove the value specified
for the `Slots` property of each `NodeGroupConfiguration` from the stack template, as it no
longer reflects the actual values in each node group. For more information, see [UseOnlineResharding Policy](../userguide/aws-attribute-updatepolicy.md#cfn-attributes-updatepolicy-useonlineresharding).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogDeliveryConfigurationRequest

Tag

All content copied from https://docs.aws.amazon.com/.
