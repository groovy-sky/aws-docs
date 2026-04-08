# DescribeDBMajorEngineVersions

Describes the properties of specific major versions of DB engines.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Engine**

The database engine to return major version details for.

Valid Values:

- `aurora-mysql`

- `aurora-postgresql`

- `custom-sqlserver-ee`

- `custom-sqlserver-se`

- `custom-sqlserver-web`

- `db2-ae`

- `db2-se`

- `mariadb`

- `mysql`

- `oracle-ee`

- `oracle-ee-cdb`

- `oracle-se2`

- `oracle-se2-cdb`

- `postgres`

- `sqlserver-ee`

- `sqlserver-se`

- `sqlserver-ex`

- `sqlserver-web`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Required: No

**MajorEngineVersion**

A specific database major engine version to return details for.

Example: `8.4`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Required: No

**Marker**

An optional pagination token provided by a previous request. If this parameter is
specified, the response includes only records beyond the marker, up to the value
specified by `MaxRecords`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 340.

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more than the `MaxRecords` value is available, a pagination token called a marker is
included in the response so you can retrieve the remaining results.

Default: 100

Type: Integer

Valid Range: Minimum value of 20. Maximum value of 100.

Required: No

## Response Elements

The following elements are returned by the service.

**DBMajorEngineVersions.DBMajorEngineVersion.N**

A list of `DBMajorEngineVersion` elements.

Type: Array of [DBMajorEngineVersion](api-dbmajorengineversion.md) objects

**Marker**

An optional pagination token provided by a previous request. If this parameter is
specified, the response includes only records beyond the marker, up to the value
specified by `MaxRecords`.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### Example

This example illustrates one usage of DescribeDBMajorEngineVersions.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=DescribeDBMajorEngineVersions
   &MaxRecords=100
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140421/us-west-2/rds/aws4_request
   &X-Amz-Date=20140421T194732Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=4772d17a4c43bcd209ff42a0778dd23e73f8434253effd7ac53b89ade3dad45f

```

#### Sample Response

```

<DescribeDBMajorEngineVersionsResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBMajorEngineVersionsResult>
    <DBMajorEngineVersions>
      <DBMajorEngineVersion>
        <Engine>mysql</Engine>
        <MajorEngineVersion>8.0</MajorEngineVersion>
      </DBMajorEngineVersion>
      <DBMajorEngineVersion>
        <Engine>mysql</Engine>
        <MajorEngineVersion>8.0</MajorEngineVersion>
        <SupportedEngineLifecycles>
          <LifecycleSupportName>open-source-rds-standard-support</LifecycleSupportName>
          <LifecycleSupportStartDate>2021-08-26T00:00:00+00:00</LifecycleSupportStartDate>
          <LifecycleSupportEndDate>2026-02-28T23:59:59.999000+00:00</LifecycleSupportEndDate>
        </SupportedEngineLifecycles>
      </DBMajorEngineVersion>
  </DescribeDBMajorEngineVersionsResult>
  <ResponseMetadata>
    <RequestId>b74d2635-b98c-11d3-fbc7-5c0aad74da7c</RequestId>
  </ResponseMetadata>
</DescribeDBMajorEngineVersionsResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbmajorengineversions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbmajorengineversions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBLogFiles

DescribeDBParameterGroups

All content copied from https://docs.aws.amazon.com/.
