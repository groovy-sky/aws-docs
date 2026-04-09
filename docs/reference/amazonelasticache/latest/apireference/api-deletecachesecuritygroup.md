# DeleteCacheSecurityGroup

Deletes a cache security group.

###### Note

You cannot delete a cache security group if it is associated with any
clusters.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheSecurityGroupName**

The name of the cache security group to delete.

###### Note

You cannot delete the default security group.

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheSecurityGroupNotFound**

The requested cache security group name does not refer to an existing cache security
group.

HTTP Status Code: 404

**InvalidCacheSecurityGroupState**

The current state of the cache security group does not allow deletion.

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

### DeleteCacheSecurityGroup

This example illustrates one usage of DeleteCacheSecurityGroup.

#### Sample Request

```

   https://elasticache.us-west-2.amazonaws.com/
   ?Action=DeleteCacheSecurityGroup
   &CacheSecurityGroupName=mycachesecuritygroup3
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DeleteCacheSecurityGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
    <ResponseMetadata>
        <RequestId>c130cfb7-3650-11e0-ae57-f96cfe56749c</RequestId>
    </ResponseMetadata>
</DeleteCacheSecurityGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deletecachesecuritygroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deletecachesecuritygroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteCacheParameterGroup

DeleteCacheSubnetGroup

All content copied from https://docs.aws.amazon.com/.
