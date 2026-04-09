# ListImports

Lists completed imports within the past 90 days.

## Request Syntax

```nohighlight

{
   "NextToken": "string",
   "PageSize": number,
   "TableArn": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[NextToken](#API_ListImports_RequestSyntax)**

An optional string that, if supplied, must be copied from the output of a previous
call to `ListImports`. When provided in this manner, the API fetches the next
page of results.

Type: String

Length Constraints: Minimum length of 112. Maximum length of 1024.

Pattern: `([0-9a-f]{16})+`

Required: No

**[PageSize](#API_ListImports_RequestSyntax)**

The number of `ImportSummary ` objects returned in a single page.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 25.

Required: No

**[TableArn](#API_ListImports_RequestSyntax)**

The Amazon Resource Name (ARN) associated with the table that was imported to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## Response Syntax

```nohighlight

{
   "ImportSummaryList": [
      {
         "CloudWatchLogGroupArn": "string",
         "EndTime": number,
         "ImportArn": "string",
         "ImportStatus": "string",
         "InputFormat": "string",
         "S3BucketSource": {
            "S3Bucket": "string",
            "S3BucketOwner": "string",
            "S3KeyPrefix": "string"
         },
         "StartTime": number,
         "TableArn": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ImportSummaryList](#API_ListImports_ResponseSyntax)**

A list of `ImportSummary` objects.

Type: Array of [ImportSummary](api-importsummary.md) objects

**[NextToken](#API_ListImports_ResponseSyntax)**

If this value is returned, there are additional results to be displayed. To retrieve
them, call `ListImports` again, with `NextToken` set to this
value.

Type: String

Length Constraints: Minimum length of 112. Maximum length of 1024.

Pattern: `([0-9a-f]{16})+`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/listimports.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/listimports.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListGlobalTables

ListTables

All content copied from https://docs.aws.amazon.com/.
