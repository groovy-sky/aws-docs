# DeleteReplicationGroup

Deletes an existing replication group. By default, this operation deletes the entire
replication group, including the primary/primaries and all of the read replicas. If the
replication group has only one primary, you can optionally delete only the read
replicas, while retaining the primary by setting
`RetainPrimaryCluster=true`.

When you receive a successful response from this operation, Amazon ElastiCache
immediately begins deleting the selected resources; you cannot cancel or revert this
operation.

###### Note

- `CreateSnapshot` permission is required to create a final snapshot.
Without this permission, the API call will fail with an `Access Denied` exception.

- This operation is valid for Redis OSS only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ReplicationGroupId**

The identifier for the cluster to be deleted. This parameter is not case
sensitive.

Type: String

Required: Yes

**FinalSnapshotIdentifier**

The name of a final node group (shard) snapshot. ElastiCache creates the snapshot from
the primary node in the cluster, rather than one of the replicas; this is to ensure that
it captures the freshest data. After the final snapshot is taken, the replication group
is immediately deleted.

Type: String

Required: No

**RetainPrimaryCluster**

If set to `true`, all of the read replicas are deleted, but the primary
node is retained.

Type: Boolean

Required: No

## Response Elements

The following element is returned by the service.

**ReplicationGroup**

Contains all of the attributes of a specific Valkey or Redis OSS replication group.

Type: [ReplicationGroup](api-replicationgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

**SnapshotAlreadyExistsFault**

You already have a snapshot with the given name.

HTTP Status Code: 400

**SnapshotFeatureNotSupportedFault**

You attempted one of the following operations:

- Creating a snapshot of a Valkey or Redis OSS cluster running on a
`cache.t1.micro` cache node.

- Creating a snapshot of a cluster that is running Memcached rather than
Valkey or Redis OSS.

Neither of these are supported by ElastiCache.

HTTP Status Code: 400

**SnapshotQuotaExceededFault**

The request cannot be processed because it would exceed the maximum number of
snapshots.

HTTP Status Code: 400

## Examples

### DeleteReplicationGroup

This example illustrates one usage of DeleteReplicationGroup.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DeleteReplicationGroup &RetainPrimaryCluster=false
   &FinalSnapshotIdentifier=my-final-snapshot
   &ReplicationGroupId=my-repgroup
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DeleteReplicationGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <DeleteReplicationGroupResult>
      <ReplicationGroup>
         <SnapshottingClusterId>my-redis-primary</SnapshottingClusterId>
decrease-replica-count         <ReplicationGroupId>my-repgroup</ReplicationGroupId>
         <Status>deleting</Status>
         <PendingModifiedValues />
         <Description>My replication group</Description>
      </ReplicationGroup>
   </DeleteReplicationGroupResult>
   <ResponseMetadata>
      <RequestId>93eb37db-b9d7-11e3-8a16-7978bb24ffdf</RequestId>
   </ResponseMetadata>
</DeleteReplicationGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deletereplicationgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deletereplicationgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteGlobalReplicationGroup

DeleteServerlessCache

All content copied from https://docs.aws.amazon.com/.
