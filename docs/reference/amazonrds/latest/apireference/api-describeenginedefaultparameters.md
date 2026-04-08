# DescribeEngineDefaultParameters

Returns the default engine and system parameter information for the specified database engine.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBParameterGroupFamily**

The name of the DB parameter group family.

Valid Values:

- `aurora-mysql5.7`

- `aurora-mysql8.0`

- `aurora-postgresql10`

- `aurora-postgresql11`

- `aurora-postgresql12`

- `aurora-postgresql13`

- `aurora-postgresql14`

- `custom-oracle-ee-19`

- `custom-oracle-ee-cdb-19`

- `db2-ae`

- `db2-se`

- `mariadb10.2`

- `mariadb10.3`

- `mariadb10.4`

- `mariadb10.5`

- `mariadb10.6`

- `mysql5.7`

- `mysql8.0`

- `oracle-ee-19`

- `oracle-ee-cdb-19`

- `oracle-ee-cdb-21`

- `oracle-se2-19`

- `oracle-se2-cdb-19`

- `oracle-se2-cdb-21`

- `postgres10`

- `postgres11`

- `postgres12`

- `postgres13`

- `postgres14`

- `sqlserver-ee-11.0`

- `sqlserver-ee-12.0`

- `sqlserver-ee-13.0`

- `sqlserver-ee-14.0`

- `sqlserver-ee-15.0`

- `sqlserver-ex-11.0`

- `sqlserver-ex-12.0`

- `sqlserver-ex-13.0`

- `sqlserver-ex-14.0`

- `sqlserver-ex-15.0`

- `sqlserver-se-11.0`

- `sqlserver-se-12.0`

- `sqlserver-se-13.0`

- `sqlserver-se-14.0`

- `sqlserver-se-15.0`

- `sqlserver-web-11.0`

- `sqlserver-web-12.0`

- `sqlserver-web-13.0`

- `sqlserver-web-14.0`

- `sqlserver-web-15.0`

Type: String

Required: Yes

**Filters.Filter.N**

A filter that specifies one or more parameters to describe.

The only supported filter is `parameter-name`. The results list only includes information about the parameters with these names.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeEngineDefaultParameters` request.
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

The following element is returned by the service.

**EngineDefaults**

Contains the result of a successful invocation of the `DescribeEngineDefaultParameters` action.

Type: [EngineDefaults](api-enginedefaults.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### Example

This example illustrates one usage of DescribeEngineDefaultParameters.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=DescribeEngineDefaultParameters
   &DBParameterGroupFamily=mysql5.1
   &MaxRecords=100
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140421/us-west-2/rds/aws4_request
   &X-Amz-Date=20140421T194732Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=747cc243a8a2385b0b06a9e2d145d08b905a39620b2782edd8382ea1712cf826

```

#### Sample Response

```

<DescribeEngineDefaultParametersResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeEngineDefaultParametersResult>
    <EngineDefaults>
      <DBParameterGroupFamily>mysql5.1</DBParameterGroupFamily>
      <Marker>bG9nX3FZXJpZ4Nfbm90X3VzaW5nX2luZGV4Z1M=</Marker>
      <Parameters>
        <Parameter>
          <DataType>boolean</DataType>
          <Source>engine-default</Source>
          <IsModifiable>false</IsModifiable>
          <Description>Controls whether user-defined functions that have only an xxx symbol for the main function can be loaded</Description>
          <ApplyType>static</ApplyType>
          <AllowedValues>0,1</AllowedValues>
          <ParameterName>allow-suspicious-udfs</ParameterName>
        </Parameter>
        <Parameter>
          <DataType>integer</DataType>
          <Source>engine-default</Source>
          <IsModifiable>true</IsModifiable>
          <Description>Intended for use with master-to-master replication, and can be used to control the operation of AUTO_INCREMENT columns</Description>
          <ApplyType>dynamic</ApplyType>
          <AllowedValues>1-65535</AllowedValues>
          <ParameterName>auto_increment_increment</ParameterName>
        </Parameter>
        <Parameter>
          <DataType>integer</DataType>
          <Source>engine-default</Source>
          <IsModifiable>true</IsModifiable>
          <Description>Determines the starting point for the AUTO_INCREMENT column value</Description>
          <ApplyType>dynamic</ApplyType>
          <AllowedValues>1-65535</AllowedValues>
          <ParameterName>auto_increment_offset</ParameterName>
        </Parameter>
      </Parameters>
    </EngineDefaults>
  </DescribeEngineDefaultParametersResult>
  <ResponseMetadata>
    <RequestId>b789ce01-b98c-11d3-a907-5a2c468b9cb0</RequestId>
  </ResponseMetadata>
</DescribeEngineDefaultParametersResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describeenginedefaultparameters.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describeenginedefaultparameters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeEngineDefaultClusterParameters

DescribeEventCategories

All content copied from https://docs.aws.amazon.com/.
