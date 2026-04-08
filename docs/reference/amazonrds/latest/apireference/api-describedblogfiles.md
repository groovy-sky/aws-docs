# DescribeDBLogFiles

Returns a list of DB log files for the DB instance.

This command doesn't apply to RDS Custom.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBInstanceIdentifier**

The customer-assigned name of the DB instance that contains the log files you want to list.

Constraints:

- Must match the identifier of an existing DBInstance.

Type: String

Required: Yes

**FileLastWritten**

Filters the available log files for files written since the specified date, in POSIX timestamp format with milliseconds.

Type: Long

Required: No

**FilenameContains**

Filters the available log files for log file names that contain the specified string.

Type: String

Required: No

**FileSize**

Filters the available log files for files larger than the specified size.

Type: Long

Required: No

**Filters.Filter.N**

This parameter isn't currently supported.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

The pagination token provided in the previous request. If this parameter is specified the response includes only records beyond the marker, up to MaxRecords.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than the specified MaxRecords value, a pagination token called a marker is included in the response so you can retrieve the remaining results.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**DescribeDBLogFiles.DescribeDBLogFilesDetails.N**

The DB log files returned.

Type: Array of [DescribeDBLogFilesDetails](api-describedblogfilesdetails.md) objects

**Marker**

A pagination token that can be used in a later `DescribeDBLogFiles` request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**DBInstanceNotReady**

An attempt to download or examine log files didn't succeed because an Aurora Serverless v2 instance was paused.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of DescribeDBLogFiles.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=DescribeDBLogFiles
   &DBInstanceIdentifier=mysqldb
   &MaxRecords=100
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140421/us-east-1/rds/aws4_request
   &X-Amz-Date=20140421T225750Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=9020fd1bcd658614e058cd2eb8c58572cf6c11460b1e96380635ee428a52e8a1

```

#### Sample Response

```

<DescribeDBLogFilesResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBLogFilesResult>
    <DescribeDBLogFiles>
      <DescribeDBLogFilesDetails>
        <LastWritten>1398119101000</LastWritten>
        <LogFileName>error/mysql-error-running.log</LogFileName>
        <Size>1599</Size>
      </DescribeDBLogFilesDetails>
      <DescribeDBLogFilesDetails>
        <LastWritten>1398120900000</LastWritten>
        <LogFileName>error/mysql-error.log</LogFileName>
        <Size>0</Size>
      </DescribeDBLogFilesDetails>
    </DescribeDBLogFiles>
  </DescribeDBLogFilesResult>
  <ResponseMetadata>
    <RequestId>4c6ed648-b9f7-11d3-97bd-7999dd5a8f72</RequestId>
  </ResponseMetadata>
</DescribeDBLogFilesResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedblogfiles.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedblogfiles.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBInstances

DescribeDBMajorEngineVersions

All content copied from https://docs.aws.amazon.com/.
