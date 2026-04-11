# UpdateTimeToLive

The `UpdateTimeToLive` method enables or disables Time to Live (TTL) for
the specified table. A successful `UpdateTimeToLive` call returns the current
`TimeToLiveSpecification`. It can take up to one hour for the change to
fully process. Any additional `UpdateTimeToLive` calls for the same table
during this one hour duration result in a `ValidationException`.

TTL compares the current time in epoch time format to the time stored in the TTL
attribute of an item. If the epoch time value stored in the attribute is less than the
current time, the item is marked as expired and subsequently deleted.

###### Note

The epoch time format is the number of seconds elapsed since 12:00:00 AM January
1, 1970 UTC.

DynamoDB deletes expired items on a best-effort basis to ensure availability of
throughput for other data operations.

###### Important

DynamoDB typically deletes expired items within two days of expiration. The exact
duration within which an item gets deleted after expiration is specific to the
nature of the workload. Items that have expired and not been deleted will still show
up in reads, queries, and scans.

As items are deleted, they are removed from any local secondary index and global
secondary index immediately in the same eventually consistent way as a standard delete
operation.

For more information, see [Time To Live](../../../../services/dynamodb/latest/developerguide/ttl.md) in the
Amazon DynamoDB Developer Guide.

## Request Syntax

```nohighlight

{
   "TableName": "string",
   "TimeToLiveSpecification": {
      "AttributeName": "string",
      "Enabled": boolean
   }
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[TableName](#API_UpdateTimeToLive_RequestSyntax)**

The name of the table to be configured. You can also provide the Amazon Resource Name (ARN) of the
table in this parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**[TimeToLiveSpecification](#API_UpdateTimeToLive_RequestSyntax)**

Represents the settings used to enable or disable Time to Live for the specified
table.

Type: [TimeToLiveSpecification](api-timetolivespecification.md) object

Required: Yes

## Response Syntax

```nohighlight

{
   "TimeToLiveSpecification": {
      "AttributeName": "string",
      "Enabled": boolean
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[TimeToLiveSpecification](#API_UpdateTimeToLive_ResponseSyntax)**

Represents the output of an `UpdateTimeToLive` operation.

Type: [TimeToLiveSpecification](api-timetolivespecification.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

**ResourceInUseException**

The operation conflicts with the resource's availability. For example:

- You attempted to recreate an existing table.

- You tried to delete a table currently in the `CREATING`
state.

- You tried to update a resource that was already being updated.

When appropriate, wait for the ongoing update to complete and attempt the request
again.

**message**

The resource which is being attempted to be changed is in use.

HTTP Status Code: 400

**ResourceNotFoundException**

The operation tried to access a nonexistent table or index. The resource might not
be specified correctly, or its status might not be `ACTIVE`.

**message**

The resource which is being requested does not exist.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/updatetimetolive.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/updatetimetolive.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateTableReplicaAutoScaling

DynamoDB Accelerator

All content copied from https://docs.aws.amazon.com/.
