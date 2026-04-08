# DescribeDBClusterParameterGroups

Returns a list of `DBClusterParameterGroup` descriptions. If a
`DBClusterParameterGroupName` parameter is specified,
the list will contain only the description of the specified DB cluster parameter group.

For more information on Amazon Aurora, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../../services/amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBClusterParameterGroupName**

The name of a specific DB cluster parameter group to return details for.

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
`DescribeDBClusterParameterGroups` request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more records exist than the specified `MaxRecords` value,
a pagination token called a marker is included in the response so you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**DBClusterParameterGroups.DBClusterParameterGroup.N**

A list of DB cluster parameter groups.

Type: Array of [DBClusterParameterGroup](api-dbclusterparametergroup.md) objects

**Marker**

An optional pagination token provided by a previous
`DescribeDBClusterParameterGroups` request.
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

This example illustrates one usage of DescribeDBClusterParameterGroups.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=DescribeDBClusterParameterGroups
   &MaxRecords=30
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20150318/us-east-1/rds/aws4_request
   &X-Amz-Date=20150318T184307Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=d9922fdf06b86b870c072b896745251ea8b52bad64bf90e30b0e46f1bb488cca

```

#### Sample Response

```

<DescribeDBClusterParameterGroupsResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBClusterParameterGroupsResult>
    <DBClusterParameterGroups>
      <DBClusterParameterGroup>
        <DBParameterGroupFamily>aurora5.6</DBParameterGroupFamily>
        <Description>Default cluster parameter group for aurora5.6</Description>
        <DBClusterParameterGroupName>default.aurora5.6</DBClusterParameterGroupName>
      </DBClusterParameterGroup>
      <DBClusterParameterGroup>
        <DBParameterGroupFamily>aurora5.6</DBParameterGroupFamily>
        <Description>Sample group</Description>
        <DBClusterParameterGroupName>samplegroup</DBClusterParameterGroupName>
      </DBClusterParameterGroup>
      <DBClusterParameterGroup>
        <DBParameterGroupFamily>aurora5.6</DBParameterGroupFamily>
        <Description>Custom group</Description>
        <DBClusterParameterGroupName>custom-group</DBClusterParameterGroupName>
      </DBClusterParameterGroup>
    </DBClusterParameterGroups>
  </DescribeDBClusterParameterGroupsResult>
  <ResponseMetadata>
    <RequestId>9e6503d0-cd9e-11e4-ccf9-7528e6a28483</RequestId>
  </ResponseMetadata>
</DescribeDBClusterParameterGroupsResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbclusterparametergroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbclusterparametergroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBClusterEndpoints

DescribeDBClusterParameters

All content copied from https://docs.aws.amazon.com/.
