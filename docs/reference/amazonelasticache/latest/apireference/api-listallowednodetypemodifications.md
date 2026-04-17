---
title: "ListAllowedNodeTypeModifications"
---

# ListAllowedNodeTypeModifications

Lists all available node types that you can scale with your cluster's replication
group's current node type.

When you use the `ModifyCacheCluster` or
`ModifyReplicationGroup` operations to scale your cluster or replication
group, the value of the `CacheNodeType` parameter must be one of the node
types returned by this operation.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheClusterId**

The name of the cluster you want to scale up to a larger node instanced type.
ElastiCache uses the cluster id to identify the current node type of this cluster and
from that to create a list of node types you can scale up to.

###### Important

You must provide a value for either the `CacheClusterId` or the
`ReplicationGroupId`.

Type: String

Required: No

**ReplicationGroupId**

The name of the replication group want to scale up to a larger node type. ElastiCache
uses the replication group id to identify the current node type being used by this
replication group, and from that to create a list of node types you can scale up
to.

###### Important

You must provide a value for either the `CacheClusterId` or the
`ReplicationGroupId`.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**ScaleDownModifications.member.N**

A string list, each element of which specifies a cache node type which you can use to
scale your cluster or replication group. When scaling down a Valkey or Redis OSS cluster or
replication group using ModifyCacheCluster or ModifyReplicationGroup, use a value from
this list for the CacheNodeType parameter.

Type: Array of strings

**ScaleUpModifications.member.N**

A string list, each element of which specifies a cache node type which you can use to
scale your cluster or replication group.

When scaling up a Valkey or Redis OSS cluster or replication group using
`ModifyCacheCluster` or `ModifyReplicationGroup`, use a value
from this list for the `CacheNodeType` parameter.

Type: Array of strings

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

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

## Examples

### ListAllowedNodeTypeModifications for a Cluster

The following example request a list of node types you can use to scale
`myCluster` up.

#### Sample Request

```

https://elasticache.us-east-1.amazonaws.com/
   ?Action=ListAllowedNodeTypeModifications
   &CacheClusterId=mycachecluster
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Version=2015-02-02
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

### ListAllowedNodeTypeModifications for a Replication Group

The following example requests a list of node types you can use to scale
`myReplGroup` up.

#### Sample Request

```

https://elasticache.us-east-1.amazonaws.com/
   ?Action=ListAllowedNodeTypeModifications
   &ReplicationGroupId=myreplgroup
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Version=2015-02-02
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/ListAllowedNodeTypeModifications)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IncreaseReplicaCount

ListTagsForResource

All content copied from https://docs.aws.amazon.com/.
