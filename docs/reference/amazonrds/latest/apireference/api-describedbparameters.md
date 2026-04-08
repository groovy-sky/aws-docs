# DescribeDBParameters

Returns the detailed parameter list for a particular DB parameter group.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBParameterGroupName**

The name of a specific DB parameter group to return details for.

Constraints:

- If supplied, must match the name of an existing DBParameterGroup.

Type: String

Required: Yes

**Filters.Filter.N**

A filter that specifies one or more DB parameters to describe.

The only supported filter is `parameter-name`. The results list only includes information about the DB parameters with these names.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeDBParameters` request.
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

**Source**

The parameter types to return.

Default: All parameter types returned

Valid Values: `user | system | engine-default`

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

**Parameters.Parameter.N**

A list of `Parameter` values.

Type: Array of [Parameter](api-parameter.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBParameterGroupNotFound**

`DBParameterGroupName` doesn't refer to an
existing DB parameter group.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DescribeDBParameters.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=DescribeDBParameters
   &DBParameterGroupName=oracle-logs
   &MaxRecords=100
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140421/us-east-1/rds/aws4_request
   &X-Amz-Date=20140421T231357Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=ac9b18d6ae7cab4bf45ed2caa99cd8438101b293c0a84e80c3bab77f7369cc7

```

#### Sample Response

```

<DescribeDBParametersResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBParametersResult>
    <Marker>bGlzdGVuZXJfbmV0d29ya3M=</Marker>
    <Parameters>
      <Parameter>
        <DataType>integer</DataType>
        <Source>engine-default</Source>
        <IsModifiable>true</IsModifiable>
        <Description>number of AQ Time Managers to start</Description>
        <ApplyType>dynamic</ApplyType>
        <AllowedValues>0-40</AllowedValues>
        <ParameterName>aq_tm_processes</ParameterName>
      </Parameter>
      <Parameter>
        <ParameterValue>300</ParameterValue>
        <DataType>integer</DataType>
        <Source>system</Source>
        <IsModifiable>false</IsModifiable>
        <Description>Maximum number of seconds of redos the standby could lose</Description>
        <ApplyType>dynamic</ApplyType>
        <ParameterName>archive_lag_target</ParameterName>
      </Parameter>
      <Parameter>
        <ParameterValue>/rdsdbdata/admin/{dbName}/adump</ParameterValue>
        <DataType>string</DataType>
        <Source>system</Source>
        <IsModifiable>false</IsModifiable>
        <Description>Directory in which auditing files are to reside</Description>
        <ApplyType>dynamic</ApplyType>
        <ParameterName>audit_file_dest</ParameterName>
      </Parameter>
      <Parameter>
        <DataType>boolean</DataType>
        <Source>engine-default</Source>
        <IsModifiable>false</IsModifiable>
        <Description>enable sys auditing</Description>
        <ApplyType>static</ApplyType>
        <AllowedValues>TRUE,FALSE</AllowedValues>
        <ParameterName>audit_sys_operations</ParameterName>
      </Parameter>
      <Parameter>
        <DataType>string</DataType>
        <Source>engine-default</Source>
        <IsModifiable>false</IsModifiable>
        <Description>Syslog facility and level</Description>
        <ApplyType>static</ApplyType>
        <ParameterName>audit_syslog_level</ParameterName>
      </Parameter>
      <Parameter>
        <DataType>string</DataType>
        <Source>engine-default</Source>
        <IsModifiable>true</IsModifiable>
        <Description>enable system auditing</Description>
        <ApplyType>static</ApplyType>
        <AllowedValues>DB,OS,NONE,TRUE,FALSE,DB_EXTENDED,XML</AllowedValues>
        <ParameterName>audit_trail</ParameterName>
      </Parameter>
    </Parameters>
  </DescribeDBParametersResult>
  <ResponseMetadata>
    <RequestId>8c40488f-b9ff-11d3-a15e-7ac49293f4fa</RequestId>
  </ResponseMetadata>
</DescribeDBParametersResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbparameters.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbparameters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBParameterGroups

DescribeDBProxies

All content copied from https://docs.aws.amazon.com/.
