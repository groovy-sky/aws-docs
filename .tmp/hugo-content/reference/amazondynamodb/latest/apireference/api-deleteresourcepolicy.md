# DeleteResourcePolicy

Deletes the resource-based policy attached to the resource, which can be a table or
stream.

`DeleteResourcePolicy` is an idempotent operation; running it multiple
times on the same resource _doesn't_ result in an error response,
unless you specify an `ExpectedRevisionId`, which will then return a
`PolicyNotFoundException`.

###### Important

To make sure that you don't inadvertently lock yourself out of your own resources,
the root principal in your AWS account can perform
`DeleteResourcePolicy` requests, even if your resource-based policy
explicitly denies the root principal's access.

###### Note

`DeleteResourcePolicy` is an asynchronous operation. If you issue a
`GetResourcePolicy` request immediately after running the
`DeleteResourcePolicy` request, DynamoDB might still return
the deleted policy. This is because the policy for your resource might not have been
deleted yet. Wait for a few seconds, and then try the `GetResourcePolicy`
request again.

## Request Syntax

```nohighlight

{
   "ExpectedRevisionId": "string",
   "ResourceArn": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ResourceArn](#API_DeleteResourcePolicy_RequestSyntax)**

The Amazon Resource Name (ARN) of the DynamoDB resource from which the policy will be
removed. The resources you can specify include tables and streams. If you remove the
policy of a table, it will also remove the permissions for the table's indexes defined
in that policy document. This is because index permissions are defined in the table's
policy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: Yes

**[ExpectedRevisionId](#API_DeleteResourcePolicy_RequestSyntax)**

A string value that you can use to conditionally delete your policy. When you provide
an expected revision ID, if the revision ID of the existing policy on the resource
doesn't match or if there's no policy attached to the resource, the request will fail
and return a `PolicyNotFoundException`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

## Response Syntax

```nohighlight

{
   "RevisionId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[RevisionId](#API_DeleteResourcePolicy_ResponseSyntax)**

A unique string that represents the revision ID of the policy. If you're comparing revision IDs, make sure to always use string comparison logic.

This value will be empty if you make a request against a resource without a
policy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

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

**PolicyNotFoundException**

The operation tried to access a nonexistent resource-based policy.

If you specified an `ExpectedRevisionId`, it's possible that a policy is
present for the resource but its revision ID didn't match the expected value.

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

## Examples

### Delete the resource-based policy of a table

The following example deletes the resource-based policy of a table named
`Thread`.

#### Sample Request

```

POST / HTTP/1.1
Host: dynamodb.<region>.<domain>;
Accept-Encoding: identity
Content-Length: <PayloadSizeBytes>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.0
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=<Headers>, Signature=<Signature>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDB_20120810.DeleteResourcePolicy
{
    "ResourceArn": "arn:aws:dynamodb:us-west-2:123456789012:table/Thread"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
x-amz-crc32: <Checksum>
Content-Type: application/x-amz-json-1.0
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
    "RevisionId": "1683717331354"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/deleteresourcepolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/deleteresourcepolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteItem

DeleteTable

All content copied from https://docs.aws.amazon.com/.
