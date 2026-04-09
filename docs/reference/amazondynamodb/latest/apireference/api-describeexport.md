# DescribeExport

Describes an existing table export.

## Request Syntax

```nohighlight

{
   "ExportArn": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ExportArn](#API_DescribeExport_RequestSyntax)**

The Amazon Resource Name (ARN) associated with the export.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: Yes

## Response Syntax

```nohighlight

{
   "ExportDescription": {
      "BilledSizeBytes": number,
      "ClientToken": "string",
      "EndTime": number,
      "ExportArn": "string",
      "ExportFormat": "string",
      "ExportManifest": "string",
      "ExportStatus": "string",
      "ExportTime": number,
      "ExportType": "string",
      "FailureCode": "string",
      "FailureMessage": "string",
      "IncrementalExportSpecification": {
         "ExportFromTime": number,
         "ExportToTime": number,
         "ExportViewType": "string"
      },
      "ItemCount": number,
      "S3Bucket": "string",
      "S3BucketOwner": "string",
      "S3Prefix": "string",
      "S3SseAlgorithm": "string",
      "S3SseKmsKeyId": "string",
      "StartTime": number,
      "TableArn": "string",
      "TableId": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ExportDescription](#API_DescribeExport_ResponseSyntax)**

Represents the properties of the export.

Type: [ExportDescription](api-exportdescription.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ExportNotFoundException**

The specified export was not found.

HTTP Status Code: 400

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**LimitExceededException**

There is no limit to the number of daily on-demand backups that can be taken.

For most purposes, up to 500 simultaneous table operations are allowed per account.
These operations include `CreateTable`, `UpdateTable`,
`DeleteTable`, `UpdateTimeToLive`,
`RestoreTableFromBackup`, and `RestoreTableToPointInTime`.

When you are creating a table with one or more secondary indexes, you can have up
to 250 such requests running at a time. However, if the table or index specifications
are complex, then DynamoDB might temporarily reduce the number of concurrent
operations.

When importing into DynamoDB, up to 50 simultaneous import table operations are
allowed per account.

There is a soft account quota of 2,500 tables.

GetRecords was called with a value of more than 1000 for the limit request
parameter.

More than 2 processes are reading from the same streams shard at the same time.
Exceeding this limit may result in request throttling.

**message**

Too many operations for a given subscriber.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/describeexport.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/describeexport.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeEndpoints

DescribeGlobalTable

All content copied from https://docs.aws.amazon.com/.
