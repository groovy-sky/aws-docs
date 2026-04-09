# DeleteCacheParameterGroup

Deletes the specified cache parameter group. You cannot delete a cache parameter group
if it is associated with any cache clusters. You cannot delete the default cache
parameter groups in your account.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheParameterGroupName**

The name of the cache parameter group to delete.

###### Note

The specified cache security group must not be associated with any
clusters.

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheParameterGroupNotFound**

The requested cache parameter group name does not refer to an existing cache parameter
group.

HTTP Status Code: 404

**InvalidCacheParameterGroupState**

The current state of the cache parameter group does not allow the requested operation
to occur.

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

## Examples

### DeleteCacheParameterGroup

This example illustrates one usage of DeleteCacheParameterGroup.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DeleteCacheParameterGroup
   &CacheParameterGroupName=myparametergroup
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DeleteCacheParameterGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
  <ResponseMetadata>
    <RequestId>d0a417cb-575b-11e0-8869-cd22b4f9d96f</RequestId>
  </ResponseMetadata>
</DeleteCacheParameterGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deletecacheparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deletecacheparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteCacheCluster

DeleteCacheSecurityGroup

All content copied from https://docs.aws.amazon.com/.
