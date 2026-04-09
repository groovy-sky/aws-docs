# ListTagsForResource

Lists all tags currently on a named resource.

A tag is a key-value pair where the key and value are case-sensitive. You can use
tags to categorize and track all your ElastiCache resources, with the exception of
global replication group. When you add or remove tags on replication groups, those
actions will be replicated to all nodes in the replication group. For more information,
see [Resource-level permissions](../../../../services/amazonelasticache/latest/dg/iam-resourcelevelpermissions.md).

If the cluster is not in the _available_ state,
`ListTagsForResource` returns an error.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ResourceName**

The Amazon Resource Name (ARN) of the resource for which you want the list of tags,
for example `arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster` or
`arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot`.

For more information about ARNs, see [Amazon Resource Names (ARNs)\
and AWS Service Namespaces](../../../../general/latest/gr/aws-arns-and-namespaces.md).

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**TagList.Tag.N**

A list of tags as key-value pairs.

Type: Array of [Tag](api-tag.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheClusterNotFound**

The requested cluster ID does not refer to an existing cluster.

HTTP Status Code: 404

**CacheParameterGroupNotFound**

The requested cache parameter group name does not refer to an existing cache parameter
group.

HTTP Status Code: 404

**CacheSecurityGroupNotFound**

The requested cache security group name does not refer to an existing cache security
group.

HTTP Status Code: 404

**CacheSubnetGroupNotFoundFault**

The requested cache subnet group name does not refer to an existing cache subnet
group.

HTTP Status Code: 400

**InvalidARN**

The requested Amazon Resource Name (ARN) does not refer to an existing
resource.

HTTP Status Code: 400

**InvalidReplicationGroupState**

The requested replication group is not in the `available` state.

HTTP Status Code: 400

**InvalidServerlessCacheSnapshotStateFault**

The state of the serverless cache snapshot was not received. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 400

**InvalidServerlessCacheStateFault**

The account for these credentials is not currently active.

HTTP Status Code: 400

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

**ReservedCacheNodeNotFound**

The requested reserved cache node was not found.

HTTP Status Code: 404

**ServerlessCacheNotFoundFault**

The serverless cache was not found or does not exist.

HTTP Status Code: 404

**ServerlessCacheSnapshotNotFoundFault**

This serverless cache snapshot could not be found or does not exist. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 404

**SnapshotNotFoundFault**

The requested snapshot name does not refer to an existing snapshot.

HTTP Status Code: 404

**UserGroupNotFound**

The user group was not found or does not exist

HTTP Status Code: 404

**UserNotFound**

The user does not exist or could not be found.

HTTP Status Code: 404

## Examples

### ListTagsForResource

This example illustrates one usage of ListTagsForResource.

#### Sample Request

```

https://elasticache.us-east-1.amazonaws.com/
   ?Action=ListTagsForResource
   &ResourceName=arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Version=2015-02-02
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/listtagsforresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/listtagsforresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAllowedNodeTypeModifications

ModifyCacheCluster

All content copied from https://docs.aws.amazon.com/.
