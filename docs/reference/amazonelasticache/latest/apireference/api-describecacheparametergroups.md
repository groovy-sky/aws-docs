# DescribeCacheParameterGroups

Returns a list of cache parameter group descriptions. If a cache parameter group name
is specified, the list contains only the descriptions for that group.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheParameterGroupName**

The name of a specific cache parameter group to return details for.

Type: String

Required: No

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

The following elements are returned by the service.

**CacheParameterGroups.CacheParameterGroup.N**

A list of cache parameter groups. Each element in the list contains detailed
information about one cache parameter group.

Type: Array of [CacheParameterGroup](api-cacheparametergroup.md) objects

**Marker**

Provides an identifier to allow retrieval of paginated results.

Type: String

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

### DescribeCacheParameterGroups

This example illustrates one usage of DescribeCacheParameterGroups.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DescribeCacheParameterGroups
   &MaxRecords=100
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DescribeCacheParameterGroupsResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
  <DescribeCacheParameterGroupsResult>
    <CacheParameterGroups>
      <CacheParameterGroup>
        <CacheParameterGroupName>default.memcached1.4</CacheParameterGroupName>
        <CacheParameterGroupFamily>memcached1.4</CacheParameterGroupFamily>
        <Description>Default parameter group for memcached1.4</Description>
      </CacheParameterGroup>
      <CacheParameterGroup>
        <CacheParameterGroupName>mycacheparametergroup</CacheParameterGroupName>
        <CacheParameterGroupFamily>memcached1.4</CacheParameterGroupFamily>
        <Description>My cache parameter group</Description>
      </CacheParameterGroup>
      <CacheParameterGroup>
        <CacheParameterGroupName>mycacheparametergroup1</CacheParameterGroupName>
        <CacheParameterGroupFamily>memcached1.4</CacheParameterGroupFamily>
        <Description>My first cache parameter group</Description>
      </CacheParameterGroup>
      <CacheParameterGroup>
        <CacheParameterGroupName>mycacheparametergroup3</CacheParameterGroupName>
        <CacheParameterGroupFamily>memcached1.4</CacheParameterGroupFamily>
        <Description>My first cache parameter group</Description>
      </CacheParameterGroup>
    </CacheParameterGroups>
  </DescribeCacheParameterGroupsResult>
  <ResponseMetadata>
    <RequestId>7193fbb8-b7fc-11e0-9b0b-a9261be2b354</RequestId>
  </ResponseMetadata>
</DescribeCacheParameterGroupsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describecacheparametergroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describecacheparametergroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeCacheEngineVersions

DescribeCacheParameters

All content copied from https://docs.aws.amazon.com/.
