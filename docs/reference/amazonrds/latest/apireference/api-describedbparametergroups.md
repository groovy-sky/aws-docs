# DescribeDBParameterGroups

Returns a list of `DBParameterGroup` descriptions. If a `DBParameterGroupName` is specified,
the list will contain only the description of the specified DB parameter group.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBParameterGroupName**

The name of a specific DB parameter group to return details for.

Constraints:

- If supplied, must match the name of an existing DBClusterParameterGroup.

Type: String

Required: No

**Filters.Filter.N**

This parameter isn't currently supported.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeDBParameterGroups` request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more records exist than the specified `MaxRecords` value,
a pagination token called a marker is included in the response so that
you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**DBParameterGroups.DBParameterGroup.N**

A list of `DBParameterGroup` instances.

Type: Array of [DBParameterGroup](api-dbparametergroup.md) objects

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBParameterGroupNotFound**

`DBParameterGroupName` doesn't refer to an
existing DB parameter group.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DescribeDBParameterGroups.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=DescribeDBParameterGroups
   &DBParameterGroupName=mysql-logs
   &MaxRecords=100
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140421/us-west-2/rds/aws4_request
   &X-Amz-Date=20140421T194732Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=e2753df1cb019f212057b51e8a2ac16dae05b344063355b195b560ef6e76661a

```

#### Sample Response

```

<DescribeDBParameterGroupsResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBParameterGroupsResult>
    <DBParameterGroups>
      <DBParameterGroup>
        <DBParameterGroupFamily>mysql5.1</DBParameterGroupFamily>
        <Description>Default parameter group for mysql5.1</Description>
        <DBParameterGroupName>default.mysql5.1</DBParameterGroupName>
      </DBParameterGroup>
      <DBParameterGroup>
        <DBParameterGroupFamily>mysql5.5</DBParameterGroupFamily>
        <Description>Default parameter group for mysql5.5</Description>
        <DBParameterGroupName>default.mysql5.5</DBParameterGroupName>
      </DBParameterGroup>
      <DBParameterGroup>
        <DBParameterGroupFamily>mysql5.6</DBParameterGroupFamily>
        <Description>Default parameter group for mysql5.6</Description>
        <DBParameterGroupName>default.mysql5.6</DBParameterGroupName>
      </DBParameterGroup>
    </DBParameterGroups>
  </DescribeDBParameterGroupsResult>
  <ResponseMetadata>
    <RequestId>b75d527a-b98c-11d3-f272-7cd6cce12cc5</RequestId>
  </ResponseMetadata>
</DescribeDBParameterGroupsResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbparametergroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbparametergroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBMajorEngineVersions

DescribeDBParameters

All content copied from https://docs.aws.amazon.com/.
