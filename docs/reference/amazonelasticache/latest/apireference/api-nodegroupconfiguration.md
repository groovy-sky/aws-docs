---
title: "NodeGroupConfiguration"
---

# NodeGroupConfiguration

Node group (shard) configuration options. Each node group (shard) configuration has
the following: `Slots`, `PrimaryAvailabilityZone`,
`ReplicaAvailabilityZones`, `ReplicaCount`.

## Contents

###### Note

In the following list, the required parameters are described first.

**NodeGroupId**

Either the ElastiCache supplied 4-digit id or a user supplied id for the
node group these configuration values apply to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 4.

Pattern: `\d+`

Required: No

**PrimaryAvailabilityZone**

The Availability Zone where the primary node of this node group (shard) is
launched.

Type: String

Required: No

**PrimaryOutpostArn**

The outpost ARN of the primary node.

Type: String

Required: No

**ReplicaAvailabilityZones.AvailabilityZone.N**

A list of Availability Zones to be used for the read replicas. The number of
Availability Zones in this list must match the value of `ReplicaCount` or
`ReplicasPerNodeGroup` if not specified.

Type: Array of strings

Required: No

**ReplicaCount**

The number of read replica nodes in this node group (shard).

Type: Integer

Required: No

**ReplicaOutpostArns.OutpostArn.N**

The outpost ARN of the node replicas.

Type: Array of strings

Required: No

**Slots**

A string that specifies the keyspace for a particular node group. Keyspaces range from
0 to 16,383. The string is in the format `startkey-endkey`.

Example: `"0-3999"`

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/NodeGroupConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/NodeGroupConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/NodeGroupConfiguration)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeGroup

NodeGroupMember

All content copied from https://docs.aws.amazon.com/.
