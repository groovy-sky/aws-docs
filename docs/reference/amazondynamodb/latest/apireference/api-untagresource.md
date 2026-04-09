# UntagResource

Removes the association of tags from an Amazon DynamoDB resource. You can call
`UntagResource` up to five times per second, per account.

- `UntagResource` is an asynchronous operation. If you issue a [ListTagsOfResource](api-listtagsofresource.md) request immediately after an
`UntagResource` request, DynamoDB might return your
previous tag set, if there was one, or an empty tag set. This is because
`ListTagsOfResource` uses an eventually consistent query, and the
metadata for your tags or table might not be available at that moment. Wait for
a few seconds, and then try the `ListTagsOfResource` request
again.

- The application or removal of tags using `TagResource` and
`UntagResource` APIs is eventually consistent.
`ListTagsOfResource` API will only reflect the changes after a
few seconds.

For an overview on tagging DynamoDB resources, see [Tagging for DynamoDB](../../../../services/dynamodb/latest/developerguide/tagging.md)
in the _Amazon DynamoDB Developer Guide_.

## Request Syntax

```nohighlight

{
   "ResourceArn": "string",
   "TagKeys": [ "string" ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ResourceArn](#API_UntagResource_RequestSyntax)**

The DynamoDB resource that the tags will be removed from. This value is an Amazon
Resource Name (ARN).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: Yes

**[TagKeys](#API_UntagResource_RequestSyntax)**

A list of tag keys. Existing tags of the resource whose keys are members of this list
will be removed from the DynamoDB resource.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/untagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/untagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransactWriteItems

UpdateContinuousBackups

All content copied from https://docs.aws.amazon.com/.
