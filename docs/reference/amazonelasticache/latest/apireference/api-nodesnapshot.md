# NodeSnapshot

Represents an individual cache node in a snapshot of a cluster.

## Contents

###### Note

In the following list, the required parameters are described first.

**CacheClusterId**

A unique identifier for the source cluster.

Type: String

Required: No

**CacheNodeCreateTime**

The date and time when the cache node was created in the source cluster.

Type: Timestamp

Required: No

**CacheNodeId**

The cache node identifier for the node in the source cluster.

Type: String

Required: No

**CacheSize**

The size of the cache on the source cache node.

Type: String

Required: No

**NodeGroupConfiguration**

The configuration for the source node group (shard).

Type: [NodeGroupConfiguration](api-nodegroupconfiguration.md) object

Required: No

**NodeGroupId**

A unique identifier for the source node group (shard).

Type: String

Required: No

**SnapshotCreateTime**

The date and time when the source node's metadata and cache data set was obtained for
the snapshot.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/nodesnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/nodesnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/nodesnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeGroupUpdateStatus

NotificationConfiguration

All content copied from https://docs.aws.amazon.com/.
