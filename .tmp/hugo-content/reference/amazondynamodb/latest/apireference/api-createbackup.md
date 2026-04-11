# CreateBackup

Creates a backup for an existing table.

Each time you create an on-demand backup, the entire table data is backed up. There
is no limit to the number of on-demand backups that can be taken.

When you create an on-demand backup, a time marker of the request is cataloged, and
the backup is created asynchronously, by applying all changes until the time of the
request to the last full table snapshot. Backup requests are processed instantaneously
and become available for restore within minutes.

You can call `CreateBackup` at a maximum rate of 50 times per
second.

All backups in DynamoDB work without consuming any provisioned throughput on the
table.

If you submit a backup request on 2018-12-14 at 14:25:00, the backup is guaranteed to
contain all data committed to the table up to 14:24:00, and data committed after
14:26:00 will not be. The backup might contain data modifications made between 14:24:00
and 14:26:00. On-demand backup does not support causal consistency.

Along with data, the following are also included on the backups:

- Global secondary indexes (GSIs)

- Local secondary indexes (LSIs)

- Streams

- Provisioned read and write capacity

## Request Syntax

```nohighlight

{
   "BackupName": "string",
   "TableName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[BackupName](#API_CreateBackup_RequestSyntax)**

Specified name for the backup.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**[TableName](#API_CreateBackup_RequestSyntax)**

The name of the table. You can also provide the Amazon Resource Name (ARN) of the table in this
parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

## Response Syntax

```nohighlight

{
   "BackupDetails": {
      "BackupArn": "string",
      "BackupCreationDateTime": number,
      "BackupExpiryDateTime": number,
      "BackupName": "string",
      "BackupSizeBytes": number,
      "BackupStatus": "string",
      "BackupType": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[BackupDetails](#API_CreateBackup_ResponseSyntax)**

Contains the details of the backup created for the table.

Type: [BackupDetails](api-backupdetails.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BackupInUseException**

There is another ongoing conflicting backup control plane operation on the table.
The backup is either being created, deleted or restored to a table.

HTTP Status Code: 400

**ContinuousBackupsUnavailableException**

Backups have not yet been enabled for this table.

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

**TableInUseException**

A target table with the specified name is either being created or deleted.

HTTP Status Code: 400

**TableNotFoundException**

A source table with the name `TableName` does not currently exist within
the subscriber's account or the subscriber is operating in the wrong AWS
Region.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/createbackup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/createbackup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchWriteItem

CreateGlobalTable

All content copied from https://docs.aws.amazon.com/.
