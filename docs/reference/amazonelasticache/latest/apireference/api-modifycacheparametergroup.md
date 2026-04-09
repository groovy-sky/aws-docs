# ModifyCacheParameterGroup

Modifies the parameters of a cache parameter group. You can modify up to 20 parameters
in a single request by submitting a list parameter name and value pairs.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheParameterGroupName**

The name of the cache parameter group to modify.

Type: String

Required: Yes

**ParameterNameValues.ParameterNameValue.N**

An array of parameter names and values for the parameter update. You must supply at
least one parameter name and value; subsequent arguments are optional. A maximum of 20
parameters may be modified per request.

Type: Array of [ParameterNameValue](api-parameternamevalue.md) objects

Required: Yes

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

### ModifyCacheParameterGroup

This example illustrates one usage of ModifyCacheParameterGroup.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=ModifyCacheParameterGroup
   &CacheParameterGroupName=mycacheparametergroup
   &ParameterNameValues.ParameterNameValue.1.ParameterName=chunk_size_growth_factor
   &ParameterNameValues.ParameterNameValue.1.ParameterValue=1.02
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<ModifyCacheParameterGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
  <ModifyCacheParameterGroupResult>
    <CacheParameterGroupName>mycacheparametergroup</CacheParameterGroupName>
  </ModifyCacheParameterGroupResult>
  <ResponseMetadata>
    <RequestId>fcedeef2-b7ff-11e0-9326-b7275b9d4a6c</RequestId>
  </ResponseMetadata>
</ModifyCacheParameterGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/modifycacheparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/modifycacheparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyCacheCluster

ModifyCacheSubnetGroup

All content copied from https://docs.aws.amazon.com/.
