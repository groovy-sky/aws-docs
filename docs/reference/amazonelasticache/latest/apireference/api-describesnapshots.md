# DescribeSnapshots

Returns information about cluster or replication group snapshots. By default,
`DescribeSnapshots` lists all of your snapshots; it can optionally
describe a single snapshot, or just the snapshots associated with a particular cache
cluster.

###### Note

This operation is valid for Valkey or Redis OSS only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheClusterId**

A user-supplied cluster identifier. If this parameter is specified, only snapshots
associated with that specific cluster are described.

Type: String

Required: No

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than
the specified `MaxRecords` value, a marker is included in the response so
that the remaining results can be retrieved.

Default: 50

Constraints: minimum 20; maximum 50.

Type: Integer

Required: No

**ReplicationGroupId**

A user-supplied replication group identifier. If this parameter is specified, only
snapshots associated with that specific replication group are described.

Type: String

Required: No

**ShowNodeGroupConfig**

A Boolean value which if true, the node group (shard) configuration is included in the
snapshot description.

Type: Boolean

Required: No

**SnapshotName**

A user-supplied name of the snapshot. If this parameter is specified, only this
snapshot are described.

Type: String

Required: No

**SnapshotSource**

If set to `system`, the output shows snapshots that were automatically
created by ElastiCache. If set to `user` the output shows snapshots that were
manually created. If omitted, the output shows both automatically and manually created
snapshots.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by `MaxRecords`.

Type: String

**Snapshots.Snapshot.N**

A list of snapshots. Each item in the list contains detailed information about one
snapshot.

Type: Array of [Snapshot](api-snapshot.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheClusterNotFound**

The requested cluster ID does not refer to an existing cluster.

HTTP Status Code: 404

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

**SnapshotNotFoundFault**

The requested snapshot name does not refer to an existing snapshot.

HTTP Status Code: 404

## Examples

### DescribeSnapshots

This example illustrates one usage of DescribeSnapshots.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DescribeSnapshots
   &MaxRecords=50
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DescribeSnapshotsResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <DescribeSnapshotsResult>
      <Snapshots>
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
            <SnapshotStatus>creating</SnapshotStatus>
            <NumCacheNodes>1</NumCacheNodes>
            <SnapshotWindow>07:30-08:30</SnapshotWindow>
         </Snapshot>
      </Snapshots>
   </DescribeSnapshotsResult>
   <ResponseMetadata>
      <RequestId>51b0b25e-b9cf-11e3-8a16-7978bb24ffdf</RequestId>
   </ResponseMetadata>
</DescribeSnapshotsResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describesnapshots.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describesnapshots.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeServiceUpdates

DescribeUpdateActions

All content copied from https://docs.aws.amazon.com/.
