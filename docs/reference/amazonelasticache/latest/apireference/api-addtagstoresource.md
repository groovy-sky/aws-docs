# AddTagsToResource

A tag is a key-value pair where the key and value are case-sensitive. You can use tags
to categorize and track all your ElastiCache resources, with the exception of global
replication group. When you add or remove tags on replication groups, those actions will
be replicated to all nodes in the replication group. For more information, see [Resource-level permissions](../../../../services/amazonelasticache/latest/dg/iam-resourcelevelpermissions.md).

For example, you can use cost-allocation tags to your ElastiCache resources, Amazon
generates a cost allocation report as a comma-separated value (CSV) file with your usage
and costs aggregated by your tags. You can apply tags that represent business categories
(such as cost centers, application names, or owners) to organize your costs across
multiple services.

For more information, see [Using Cost Allocation Tags in\
Amazon ElastiCache](../../../../services/amazonelasticache/latest/dg/tagging.md) in the _ElastiCache User_
_Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ResourceName**

The Amazon Resource Name (ARN) of the resource to which the tags are to be added, for
example `arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster` or
`arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot`.
ElastiCache resources are _cluster_ and
_snapshot_.

For more information about ARNs, see [Amazon Resource Names (ARNs)\
and Amazon Service Namespaces](../../../../general/latest/gr/aws-arns-and-namespaces.md).

Type: String

Required: Yes

**Tags.Tag.N**

A list of tags to be added to this resource. A tag is a key-value pair. A tag key must
be accompanied by a tag value, although null is accepted.

Type: Array of [Tag](api-tag.md) objects

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

**TagQuotaPerResourceExceeded**

The request cannot be processed because it would cause the resource to have more than
the allowed number of tags. The maximum number of tags permitted on a resource is
50.

HTTP Status Code: 400

**UserGroupNotFound**

The user group was not found or does not exist

HTTP Status Code: 404

**UserNotFound**

The user does not exist or could not be found.

HTTP Status Code: 404

## Examples

### AddTagsToResource

This example illustrates one usage of AddTagsToResource.

#### Sample Request

```

https://elasticache.us-east-1.amazonaws.com/
   ?Action=AddTagsToResource
   &ResourceName=arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Tags.Tag.1.Key=Service
   &Tags.Tag.1.Value=elasticache
   &Tags.Tag.2.Key=Region
   &Tags.Tag.2.Value=us-west-2
   &Version=2015-02-02
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/addtagstoresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/addtagstoresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

AuthorizeCacheSecurityGroupIngress

All content copied from https://docs.aws.amazon.com/.
