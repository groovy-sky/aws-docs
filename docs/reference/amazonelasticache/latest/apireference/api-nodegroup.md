# NodeGroup

Represents a collection of cache nodes in a replication group. One node in the node
group is the read/write primary node. All the other nodes are read-only Replica
nodes.

## Contents

###### Note

In the following list, the required parameters are described first.

**NodeGroupId**

The identifier for the node group (shard). A Valkey or Redis OSS (cluster mode disabled) replication
group contains only 1 node group; therefore, the node group ID is 0001. A Valkey or Redis OSS (cluster mode enabled) replication group contains 1 to 90 node groups numbered 0001 to 0090.
Optionally, the user can provide the id for a node group.

Type: String

Required: No

**NodeGroupMembers.NodeGroupMember.N**

A list containing information about individual nodes within the node group
(shard).

Type: Array of [NodeGroupMember](api-nodegroupmember.md) objects

Required: No

**PrimaryEndpoint**

The endpoint of the primary node in this node group (shard).

Type: [Endpoint](api-endpoint.md) object

Required: No

**ReaderEndpoint**

The endpoint of the replica nodes in this node group (shard). This value is read-only.

Type: [Endpoint](api-endpoint.md) object

Required: No

**Slots**

The keyspace for this node group (shard).

Type: String

Required: No

**Status**

The current state of this replication group - `creating`,
`available`, `modifying`, `deleting`.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/nodegroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/nodegroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/nodegroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogDeliveryConfigurationRequest

NodeGroupConfiguration

All content copied from https://docs.aws.amazon.com/.
