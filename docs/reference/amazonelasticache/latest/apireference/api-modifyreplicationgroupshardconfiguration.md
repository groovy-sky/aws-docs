# ModifyReplicationGroupShardConfiguration

Modifies a replication group's shards (node groups) by allowing you to add shards,
remove shards, or rebalance the keyspaces among existing shards.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ApplyImmediately**

Indicates that the shard reconfiguration process begins immediately. At present, the
only permitted value for this parameter is `true`.

Value: true

Type: Boolean

Required: Yes

**NodeGroupCount**

The number of node groups (shards) that results from the modification of the shard
configuration.

Type: Integer

Required: Yes

**ReplicationGroupId**

The name of the Valkey or Redis OSS (cluster mode enabled) cluster (replication group) on which the
shards are to be configured.

Type: String

Required: Yes

**NodeGroupsToRemove.NodeGroupToRemove.N**

If the value of `NodeGroupCount` is less than the current number of node
groups (shards), then either `NodeGroupsToRemove` or
`NodeGroupsToRetain` is required. `NodeGroupsToRemove` is a
list of `NodeGroupId` s to remove from the cluster.

ElastiCache will attempt to remove all node groups listed by
`NodeGroupsToRemove` from the cluster.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 4.

Pattern: `\d+`

Required: No

**NodeGroupsToRetain.NodeGroupToRetain.N**

If the value of `NodeGroupCount` is less than the current number of node
groups (shards), then either `NodeGroupsToRemove` or
`NodeGroupsToRetain` is required. `NodeGroupsToRetain` is a
list of `NodeGroupId` s to retain in the cluster.

ElastiCache will attempt to remove all node groups except those listed by
`NodeGroupsToRetain` from the cluster.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 4.

Pattern: `\d+`

Required: No

**ReshardingConfiguration.ReshardingConfiguration.N**

Specifies the preferred availability zones for each node group in the cluster. If the
value of `NodeGroupCount` is greater than the current number of node groups
(shards), you can use this parameter to specify the preferred availability zones of the
cluster's shards. If you omit this parameter ElastiCache selects availability zones for
you.

You can specify this parameter only if the value of `NodeGroupCount` is
greater than the current number of node groups (shards).

Type: Array of [ReshardingConfiguration](api-reshardingconfiguration.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**ReplicationGroup**

Contains all of the attributes of a specific Valkey or Redis OSS replication group.

Type: [ReplicationGroup](api-replicationgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InsufficientCacheClusterCapacity**

The requested cache node type is not available in the specified Availability Zone. For
more information, see [InsufficientCacheClusterCapacity](../../../../services/amazonelasticache/latest/dg/errormessages-errormessages-insufficient-cache-cluster-capacity.md) in the ElastiCache User Guide.

HTTP Status Code: 400

**InvalidCacheClusterState**

The requested cluster is not in the `available` state.

HTTP Status Code: 400

**InvalidKMSKeyFault**

The KMS key supplied is not valid.

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

**InvalidReplicationGroupState**

The requested replication group is not in the `available` state.

HTTP Status Code: 400

**InvalidVPCNetworkStateFault**

The VPC network is in an invalid state.

HTTP Status Code: 400

**NodeGroupsPerReplicationGroupQuotaExceeded**

The request cannot be processed because it would exceed the maximum allowed number of
node groups (shards) in a single replication group. The default maximum is 90

HTTP Status Code: 400

**NodeQuotaForCustomerExceeded**

The request cannot be processed because it would exceed the allowed number of cache
nodes per customer.

HTTP Status Code: 400

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

## Examples

### Add shards

The following example adds shards to the replication group
`my-cluster` so that at the completion of the call there are 4
shards. The availability zones for the nodes in the shards are specified by the
parameter
`ReshardingConfiguration.ReshardingConfiguration.N.PreferredAvailabilityZones.AvailabilityZone.N`.
If there are already 4 shards in this replication group, the call fails.

```

https://elasticache.us-east-2.amazonaws.com/
    ?Action=ModifyReplicationGroupShardConfiguration
    &ApplyImmediately=true
    &NodeGroupCount=4
    &ReplicationGroupId=my-cluster
    &ReshardingConfiguration.ReshardingConfiguration.1.PreferredAvailabilityZones.AvailabilityZone.1=us-east-2a
    &ReshardingConfiguration.ReshardingConfiguration.1.PreferredAvailabilityZones.AvailabilityZone.2=us-east-2c
    &ReshardingConfiguration.ReshardingConfiguration.2.PreferredAvailabilityZones.AvailabilityZone.1=us-east-2b
    &ReshardingConfiguration.ReshardingConfiguration.2.PreferredAvailabilityZones.AvailabilityZone.2=us-east-2a
    &ReshardingConfiguration.ReshardingConfiguration.3.PreferredAvailabilityZones.AvailabilityZone.1=us-east-2c
    &ReshardingConfiguration.ReshardingConfiguration.3.PreferredAvailabilityZones.AvailabilityZone.2=us-east-2d
    &ReshardingConfiguration.ReshardingConfiguration.4.PreferredAvailabilityZones.AvailabilityZone.1=us-east-2d
    &ReshardingConfiguration.ReshardingConfiguration.4.PreferredAvailabilityZones.AvailabilityZone.2=us-east-2c
    &Version=2015-02-02
    &SignatureVersion=4
    &SignatureMethod=HmacSHA256
    &Timestamp=20171002T192317Z
    &X-Amz-Credential=<credential>
```

### Remove shards

The following example removes two shards from the replication group
`my-cluster`. leaving 2 shards. When removing shards, the
parameter `NodeGroupsToRemove.NodeGroupToRemove` is required.

You cannot remove the last shard leaving zero shards.

```

https://elasticache.us-east-2.amazonaws.com/
    ?Action=ModifyReplicationGroupShardConfiguration
    &ApplyImmediately=true
    &NodeGroupCount=2
    &ReplicationGroupId=my-cluster
    &NodeGroupsToRemove.NodeGroupToRemove.1=0002
    &NodeGroupsToRemove.NodeGroupToRemove.2=0003
    &Version=2015-02-02
    &SignatureVersion=4
    &SignatureMethod=HmacSHA256
    &Timestamp=20171002T192317Z
    &X-Amz-Credential=<credential>
```

### Rebalance shards

The following rebalances the keyspaces among the existing shards in the
replication group `my-cluster`. The value specified by
`NodeGroupCount` must be the existing number of shards. If the
keyspaces are already balanced the call fails.

```

https://elasticache.us-east-2.amazonaws.com/
    ?Action=ModifyReplicationGroupShardConfiguration
    &ApplyImmediately=true
    &NodeGroupCount=4
    &ReplicationGroupId=my-cluster
    &Version=2015-02-02
    &SignatureVersion=4
    &SignatureMethod=HmacSHA256
    &Timestamp=20171002T192317Z
    &X-Amz-Credential=<credential>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/modifyreplicationgroupshardconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyReplicationGroup

ModifyServerlessCache
