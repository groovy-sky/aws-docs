# ResetCacheParameterGroup

Modifies the parameters of a cache parameter group to the engine or system default
value. You can reset specific parameters by submitting a list of parameter names. To
reset the entire cache parameter group, specify the `ResetAllParameters` and
`CacheParameterGroupName` parameters.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheParameterGroupName**

The name of the cache parameter group to reset.

Type: String

Required: Yes

**ParameterNameValues.ParameterNameValue.N**

An array of parameter names to reset to their default values. If
`ResetAllParameters` is `true`, do not use
`ParameterNameValues`. If `ResetAllParameters` is
`false`, you must specify the name of at least one parameter to
reset.

Type: Array of [ParameterNameValue](api-parameternamevalue.md) objects

Required: No

**ResetAllParameters**

If `true`, all parameters in the cache parameter group are reset to their
default values. If `false`, only the parameters listed by
`ParameterNameValues` are reset to their default values.

Valid values: `true` \| `false`

Type: Boolean

Required: No

## Response Elements

The following element is returned by the service.

**CacheParameterGroupName**

The name of the cache parameter group.

Type: String

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

**InvalidGlobalReplicationGroupState**

The Global datastore is not available or in primary-only state.

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

### ResetCacheParameterGroup

This example illustrates one usage of ResetCacheParameterGroup.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=ResetCacheParameterGroup
   &ResetAllParameters=true
   &CacheParameterGroupName=mycacheparametergroup1
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<ResetCacheParameterGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <ResetCacheParameterGroupResult>
      <CacheParameterGroupName>mycacheparametergroup1</CacheParameterGroupName>
   </ResetCacheParameterGroupResult>
   <ResponseMetadata>
      <RequestId>cb7cc855-b9d2-11e3-8a16-7978bb24ffdf</RequestId>
   </ResponseMetadata>
</ResetCacheParameterGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/resetcacheparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/resetcacheparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemoveTagsFromResource

RevokeCacheSecurityGroupIngress

All content copied from https://docs.aws.amazon.com/.
