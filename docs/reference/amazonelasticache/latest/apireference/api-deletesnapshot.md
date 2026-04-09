# DeleteSnapshot

Deletes an existing snapshot. When you receive a successful response from this
operation, ElastiCache immediately begins deleting the snapshot; you cannot cancel or
revert this operation.

###### Note

This operation is valid for Valkey or Redis OSS only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SnapshotName**

The name of the snapshot to be deleted.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**Snapshot**

Represents a copy of an entire Valkey or Redis OSS cluster as of the time when the snapshot was
taken.

Type: [Snapshot](api-snapshot.md) object

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

**InvalidSnapshotState**

The current state of the snapshot does not allow the requested operation to
occur.

HTTP Status Code: 400

**SnapshotNotFoundFault**

The requested snapshot name does not refer to an existing snapshot.

HTTP Status Code: 404

## Examples

### DeleteSnapshot

This example illustrates one usage of DeleteSnapshot.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DeleteSnapshot
   &SnapshotName=my-manual-snapshot
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DeleteSnapshotResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <DeleteSnapshotResult>
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
               <SnapshotCreateTime>2015-02-02T18:54:12Z</SnapshotCreateTime>
               <CacheNodeCreateTime>2015-02-02T18:46:57.972Z</CacheNodeCreateTime>
               <CacheNodeId>0001</CacheNodeId>
               <CacheSize>3 MB</CacheSize>
            </NodeSnapshot>
         </NodeSnapshots>
         <SnapshotStatus>deleting</SnapshotStatus>
         <NumCacheNodes>1</NumCacheNodes>
         <SnapshotWindow>07:30-08:30</SnapshotWindow>
      </Snapshot>
   </DeleteSnapshotResult>
   <ResponseMetadata>
      <RequestId>694d7017-b9d2-11e3-8a16-7978bb24ffdf</RequestId>
   </ResponseMetadata>
</DeleteSnapshotResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deletesnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deletesnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteServerlessCacheSnapshot

DeleteUser

All content copied from https://docs.aws.amazon.com/.
