# DescribeReplicationGroups

Returns information about a particular replication group. If no identifier is
specified, `DescribeReplicationGroups` returns information about all
replication groups.

###### Note

This operation is valid for Valkey or Redis OSS only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

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

Default: 100

Constraints: minimum 20; maximum 100.

Type: Integer

Required: No

**ReplicationGroupId**

The identifier for the replication group to be described. This parameter is not case
sensitive.

If you do not specify this parameter, information about all replication groups is
returned.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

Provides an identifier to allow retrieval of paginated results.

Type: String

**ReplicationGroups.ReplicationGroup.N**

A list of replication groups. Each item in the list contains detailed information
about one replication group.

Type: Array of [ReplicationGroup](api-replicationgroup.md) objects

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

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

## Examples

### DescribeReplicationGroups

This example illustrates one usage of DescribeReplicationGroups.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DescribeReplicationGroups
   &MaxRecords=100
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DescribeReplicationGroupsResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <DescribeReplicationGroupsResult>
      <ReplicationGroups>
         <ReplicationGroup>
            <SnapshottingClusterId>my-redis-primary</SnapshottingClusterId>
            <MemberClusters>
               <ClusterId>my-redis-primary</ClusterId>
            </MemberClusters>
            <NodeGroups>
               <NodeGroup>
                  <NodeGroupId>0001</NodeGroupId>
                  <PrimaryEndpoint>
                     <Port>6379</Port>
                     <Address>my-repgroup.q68zge.ng.0001.use1devo.elmo-dev.amazonaws.com</Address>
                  </PrimaryEndpoint>
                  <Status>available</Status>
                  <NodeGroupMembers>
                     <NodeGroupMember>
                        <CacheClusterId>my-redis-primary</CacheClusterId>
                        <ReadEndpoint>
                           <Port>6379</Port>
                           <Address>my-redis-primary.q68zge.0001.use1devo.elmo-dev.amazonaws.com</Address>
                        </ReadEndpoint>
                        <PreferredAvailabilityZone>us-west-2c</PreferredAvailabilityZone>
                        <CacheNodeId>0001</CacheNodeId>
                        <CurrentRole>primary</CurrentRole>
                     </NodeGroupMember>
                  </NodeGroupMembers>
               </NodeGroup>
            </NodeGroups>
            <ReplicationGroupId>my-repgroup</ReplicationGroupId>
            <Status>available</Status>
            <PendingModifiedValues />
            <Description>My replication group</Description>
         </ReplicationGroup>
      </ReplicationGroups>
   </DescribeReplicationGroupsResult>
   <ResponseMetadata>
      <RequestId>144745b0-b9d3-11e3-8a16-7978bb24ffdf</RequestId>
   </ResponseMetadata>
</DescribeReplicationGroupsResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describereplicationgroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describereplicationgroups.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeGlobalReplicationGroups

DescribeReservedCacheNodes
