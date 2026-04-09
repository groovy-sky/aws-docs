# DescribeEngineDefaultParameters

Returns the default engine and system parameter information for the specified cache
engine.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheParameterGroupFamily**

The name of the cache parameter group family.

Valid values are: `memcached1.4` \| `memcached1.5` \|
`memcached1.6` \| `redis2.6` \| `redis2.8` \|
`redis3.2` \| `redis4.0` \| `redis5.0` \|
`redis6.x` \| `redis6.2` \| `redis7`

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

## Response Elements

The following element is returned by the service.

**EngineDefaults**

Represents the output of a `DescribeEngineDefaultParameters`
operation.

Type: [EngineDefaults](api-enginedefaults.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

### DescribeEngineDefaultParameters

Some of the output has been omitted for brevity.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DescribeEngineDefaultParameters
   &CacheParameterGroupFamily=memcached1.4
   &MaxRecords=100
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DescribeEngineDefaultParametersResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
  <DescribeEngineDefaultParametersResult>
    <EngineDefaults>
      <CacheParameterGroupFamily>memcached1.4</CacheParameterGroupFamily>
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
        <Parameter>

 (...output omitted...)

        </Parameter>
      </Parameters>
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

 (...output omitted...)

      </CacheNodeTypeSpecificParameters>
    </EngineDefaults>
  </DescribeEngineDefaultParametersResult>
  <ResponseMetadata>
    <RequestId>061282fe-b7fd-11e0-9326-b7275b9d4a6c</RequestId>
  </ResponseMetadata>
</DescribeEngineDefaultParametersResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describeenginedefaultparameters.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describeenginedefaultparameters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeCacheSubnetGroups

DescribeEvents

All content copied from https://docs.aws.amazon.com/.
