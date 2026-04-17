---
title: "DeleteCacheSubnetGroup"
---

# DeleteCacheSubnetGroup

Deletes a cache subnet group.

###### Note

You cannot delete a default cache subnet group or one that is associated with any
clusters.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheSubnetGroupName**

The name of the cache subnet group to delete.

Constraints: Must contain no more than 255 alphanumeric characters or hyphens.

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheSubnetGroupInUse**

The requested cache subnet group is currently in use.

HTTP Status Code: 400

**CacheSubnetGroupNotFoundFault**

The requested cache subnet group name does not refer to an existing cache subnet
group.

HTTP Status Code: 400

## Examples

### DeleteCacheSubnetGroup

This example illustrates one usage of DeleteCacheSubnetGroup.

#### Sample Request

```

https://elasticache.amazonaws.com/
   ?Action=DeleteCacheSubnetGroup
   &CacheSubnetGroupName=mysubnetgroup
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DeleteCacheSubnetGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
      <ResponseMetadata>
        <RequestId>5d013245-4172-11df-8520-e7e1e602a915</RequestId>
      </ResponseMetadata>
    </DeleteCacheSubnetGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/elasticache-2015-02-02/DeleteCacheSubnetGroup)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/DeleteCacheSubnetGroup)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteCacheSecurityGroup

DeleteGlobalReplicationGroup

All content copied from https://docs.aws.amazon.com/.
