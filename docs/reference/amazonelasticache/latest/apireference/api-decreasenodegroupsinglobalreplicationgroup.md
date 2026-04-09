# DecreaseNodeGroupsInGlobalReplicationGroup

Decreases the number of node groups in a Global datastore

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ApplyImmediately**

Indicates that the shard reconfiguration process begins immediately. At present, the
only permitted value for this parameter is true.

Type: Boolean

Required: Yes

**GlobalReplicationGroupId**

The name of the Global datastore

Type: String

Required: Yes

**NodeGroupCount**

The number of node groups (shards) that results from the modification of the shard
configuration

Type: Integer

Required: Yes

**GlobalNodeGroupsToRemove.GlobalNodeGroupId.N**

If the value of NodeGroupCount is less than the current number of node groups
(shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
GlobalNodeGroupsToRemove is a list of NodeGroupIds to remove from the cluster.
ElastiCache will attempt to remove all node groups listed by
GlobalNodeGroupsToRemove from the cluster.

Type: Array of strings

Required: No

**GlobalNodeGroupsToRetain.GlobalNodeGroupId.N**

If the value of NodeGroupCount is less than the current number of node groups
(shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.
GlobalNodeGroupsToRetain is a list of NodeGroupIds to retain from the cluster.
ElastiCache will attempt to retain all node groups listed by
GlobalNodeGroupsToRetain from the cluster.

Type: Array of strings

Required: No

## Response Elements

The following element is returned by the service.

**GlobalReplicationGroup**

Consists of a primary cluster that accepts writes and an associated secondary cluster
that resides in a different Amazon region. The secondary cluster accepts only reads. The
primary cluster automatically replicates updates to the secondary cluster.

- The **GlobalReplicationGroupIdSuffix** represents
the name of the Global datastore, which is what you use to associate a secondary
cluster.

Type: [GlobalReplicationGroup](api-globalreplicationgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**GlobalReplicationGroupNotFoundFault**

The Global datastore does not exist

HTTP Status Code: 404

**InvalidGlobalReplicationGroupState**

The Global datastore is not available or in primary-only state.

HTTP Status Code: 400

**InvalidParameterCombination**

Two or more incompatible parameters were specified.

**message**

Two or more parameters that must not be used together were used together.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/decreasenodegroupsinglobalreplicationgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateUserGroup

DecreaseReplicaCount

All content copied from https://docs.aws.amazon.com/.
