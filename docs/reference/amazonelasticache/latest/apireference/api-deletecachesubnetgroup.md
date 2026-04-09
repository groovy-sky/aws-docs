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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deletecachesubnetgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deletecachesubnetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteCacheSecurityGroup

DeleteGlobalReplicationGroup

All content copied from https://docs.aws.amazon.com/.
