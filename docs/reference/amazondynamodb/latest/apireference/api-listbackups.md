# ListBackups

List DynamoDB backups that are associated with an AWS account and
weren't made with AWS Backup. To list these backups for a given table,
specify `TableName`. `ListBackups` returns a paginated list of
results with at most 1 MB worth of items in a page. You can also specify a maximum
number of entries to be returned in a page.

In the request, start time is inclusive, but end time is exclusive. Note that these
boundaries are for the time at which the original backup was requested.

You can call `ListBackups` a maximum of five times per second.

If you want to retrieve the complete list of backups made with AWS
Backup, use the [AWS Backup\
list API.](../../../../services/aws-backup/latest/devguide/api-listbackupjobs.md)

## Request Syntax

```nohighlight

{
   "BackupType": "string",
   "ExclusiveStartBackupArn": "string",
   "Limit": number,
   "TableName": "string",
   "TimeRangeLowerBound": number,
   "TimeRangeUpperBound": number
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[BackupType](#API_ListBackups_RequestSyntax)**

The backups from the table specified by `BackupType` are listed.

Where `BackupType` can be:

- `USER` \- On-demand backup created by you. (The default setting if no
other backup types are specified.)

- `SYSTEM` \- On-demand backup automatically created by DynamoDB.

- `ALL` \- All types of on-demand backups (USER and SYSTEM).

Type: String

Valid Values: `USER | SYSTEM | AWS_BACKUP | ALL`

Required: No

**[ExclusiveStartBackupArn](#API_ListBackups_RequestSyntax)**

`LastEvaluatedBackupArn` is the Amazon Resource Name (ARN) of the backup last
evaluated when the current page of results was returned, inclusive of the current page
of results. This value may be specified as the `ExclusiveStartBackupArn` of a
new `ListBackups` operation in order to fetch the next page of results.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**[Limit](#API_ListBackups_RequestSyntax)**

Maximum number of backups to return at once.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[TableName](#API_ListBackups_RequestSyntax)**

Lists the backups from the table specified in `TableName`. You can also
provide the Amazon Resource Name (ARN) of the table in this parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**[TimeRangeLowerBound](#API_ListBackups_RequestSyntax)**

Only backups created after this time are listed. `TimeRangeLowerBound` is
inclusive.

Type: Timestamp

Required: No

**[TimeRangeUpperBound](#API_ListBackups_RequestSyntax)**

Only backups created before this time are listed. `TimeRangeUpperBound` is
exclusive.

Type: Timestamp

Required: No

## Response Syntax

```nohighlight

{
   "BackupSummaries": [
      {
         "BackupArn": "string",
         "BackupCreationDateTime": number,
         "BackupExpiryDateTime": number,
         "BackupName": "string",
         "BackupSizeBytes": number,
         "BackupStatus": "string",
         "BackupType": "string",
         "TableArn": "string",
         "TableId": "string",
         "TableName": "string"
      }
   ],
   "LastEvaluatedBackupArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[BackupSummaries](#API_ListBackups_ResponseSyntax)**

List of `BackupSummary` objects.

Type: Array of [BackupSummary](api-backupsummary.md) objects

**[LastEvaluatedBackupArn](#API_ListBackups_ResponseSyntax)**

The ARN of the backup last evaluated when the current page of results was returned,
inclusive of the current page of results. This value may be specified as the
`ExclusiveStartBackupArn` of a new `ListBackups` operation in
order to fetch the next page of results.

If `LastEvaluatedBackupArn` is empty, then the last page of results has
been processed and there are no more results to be retrieved.

If `LastEvaluatedBackupArn` is not empty, this may or may not indicate
that there is more data to be returned. All results are guaranteed to have been returned
if and only if no value for `LastEvaluatedBackupArn` is returned.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/listbackups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/listbackups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportTable

ListContributorInsights

All content copied from https://docs.aws.amazon.com/.
