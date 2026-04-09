# DescribeCacheParameters

Returns the detailed parameter list for a particular cache parameter group.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheParameterGroupName**

The name of a specific cache parameter group to return details for.

Type: String

Required: Yes

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

**Source**

The parameter types to return.

Valid values: `user` \| `system` \|
`engine-default`

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**CacheNodeTypeSpecificParameters.CacheNodeTypeSpecificParameter.N**

A list of parameters specific to a particular cache node type. Each element in the
list contains detailed information about one parameter.

Type: Array of [CacheNodeTypeSpecificParameter](api-cachenodetypespecificparameter.md) objects

**Marker**

Provides an identifier to allow retrieval of paginated results.

Type: String

**Parameters.Parameter.N**

A list of [Parameter](api-parameter.md) instances.

Type: Array of [Parameter](api-parameter.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheParameterGroupNotFound**

The requested cache parameter group name does not refer to an existing cache parameter
group.

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

## Examples

### DescribeCacheParameters

Some of the output has been omitted for brevity.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DescribeCacheParameters
   &CacheParameterGroupName=default.memcached1.4
   &MaxRecords=100
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DescribeCacheParametersResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
  <DescribeCacheParametersResult>
    <CacheNodeTypeSpecificParameters>
      <CacheNodeTypeSpecificParameter>
        <CacheNodeTypeSpecificValues>
          <CacheNodeTypeSpecificValue>
            <CacheNodeType>cache.c1.xlarge</CacheNodeType>
            <Value>6000</Value>
          </CacheNodeTypeSpecificValue>

 (...output omitted...)

        </CacheNodeTypeSpecificValues>
        <DataType>integer</DataType>
        <Source>system</Source>
        <IsModifiable>false</IsModifiable>
        <Description>The maximum configurable amount of memory to use to store items, in megabytes.</Description>
        <AllowedValues>1-100000</AllowedValues>
        <ParameterName>max_cache_memory</ParameterName>
        <MinimumEngineVersion>1.4.5</MinimumEngineVersion>
      </CacheNodeTypeSpecificParameter>
      <CacheNodeTypeSpecificParameter>

 (...output omitted...)

      </CacheNodeTypeSpecificParameter>
    </CacheNodeTypeSpecificParameters>
    <Parameters>
      <Parameter>
        <ParameterValue>1024</ParameterValue>
        <DataType>integer</DataType>
        <Source>system</Source>
        <IsModifiable>false</IsModifiable>
        <Description>The backlog queue limit.</Description>
        <AllowedValues>1-10000</AllowedValues>
        <ParameterName>backlog_queue_limit</ParameterName>
        <MinimumEngineVersion>1.4.5</MinimumEngineVersion>
      </Parameter>

 (...output omitted...)

    </Parameters>
  </DescribeCacheParametersResult>
  <ResponseMetadata>
    <RequestId>0c507368-b7fe-11e0-9326-b7275b9d4a6c</RequestId>
  </ResponseMetadata>
</DescribeCacheParametersResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describecacheparameters.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describecacheparameters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeCacheParameterGroups

DescribeCacheSecurityGroups

All content copied from https://docs.aws.amazon.com/.
