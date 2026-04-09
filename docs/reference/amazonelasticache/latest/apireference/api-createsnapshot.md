# CreateSnapshot

Creates a copy of an entire cluster or replication group at a specific moment in
time.

###### Note

This operation is valid for Valkey or Redis OSS only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SnapshotName**

A name for the snapshot being created. This value is stored as a lowercase string.

Type: String

Required: Yes

**CacheClusterId**

The identifier of an existing cluster. The snapshot is created from this
cluster.

Type: String

Required: No

**KmsKeyId**

The ID of the KMS key used to encrypt the snapshot.

Type: String

Required: No

**ReplicationGroupId**

The identifier of an existing replication group. The snapshot is created from this
replication group.

Type: String

Required: No

**Tags.Tag.N**

A list of tags to be added to this resource. A tag is a key-value pair. A tag key must
be accompanied by a tag value, although null is accepted.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**Snapshot**

Represents a copy of an entire Valkey or Redis OSS cluster as of the time when the snapshot was
taken.

Type: [Snapshot](api-snapshot.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheClusterNotFound**

The requested cluster ID does not refer to an existing cluster.

HTTP Status Code: 404

**InvalidCacheClusterState**

The requested cluster is not in the `available` state.

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

**TagQuotaPerResourceExceeded**

The request cannot be processed because it would cause the resource to have more than
the allowed number of tags. The maximum number of tags permitted on a resource is
50.

HTTP Status Code: 400

## Examples

### CreateSnapshot

This example illustrates one usage of CreateSnapshot.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=CreateSnapshot
   &CacheClusterId=my-redis-primary
   &SnapshotName=my-manual-snapshot
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<CreateSnapshotResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <CreateSnapshotResult>
      <Snapshot>
         <CacheClusterId>my-redis-primary</CacheClusterId>
         <Port>6379</Port>
         <CacheNodeType>cache.m1.small</CacheNodeType>
         <CacheParameterGroupName>default.redis2.8</CacheParameterGroupName>
         <Engine>redis</Engine>
         <PreferredAvailabilityZone>us-west-2c</PreferredAvailabilityZone>
         <CacheClusterCreateTime>2015-02-02T18:46:57.972Z</CacheClusterCreateTime>
         <EngineVersion>2.8.6</EngineVersion>
         <SnapshotSource>manual</SnapshotSource>
         <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
         <PreferredMaintenanceWindow>wed:09:00-wed:10:00</PreferredMaintenanceWindow>
         <SnapshotName>my-manual-snapshot</SnapshotName>
         <SnapshotRetentionLimit>5</SnapshotRetentionLimit>
         <NodeSnapshots>
            <NodeSnapshot>
               <CacheNodeCreateTime>2015-02-02T18:46:57.972Z</CacheNodeCreateTime>
               <CacheNodeId>0001</CacheNodeId>
               <CacheSize />
            </NodeSnapshot>
         </NodeSnapshots>
         <SnapshotStatus>creating</SnapshotStatus>
         <NumCacheNodes>1</NumCacheNodes>
         <SnapshotWindow>07:30-08:30</SnapshotWindow>
      </Snapshot>
   </CreateSnapshotResult>
   <ResponseMetadata>
      <RequestId>faf5a232-b9ce-11e3-8a16-7978bb24ffdf</RequestId>
   </ResponseMetadata>
</CreateSnapshotResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/createsnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/createsnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateServerlessCacheSnapshot

CreateUser

All content copied from https://docs.aws.amazon.com/.
