# PutResourcePolicy

Attaches a resource-based policy document to the resource, which can be a table or
stream. When you attach a resource-based policy using this API, the policy application
is [_eventually consistent_](../../../../services/dynamodb/latest/developerguide/howitworks-readconsistency.md).

`PutResourcePolicy` is an idempotent operation; running it multiple times
on the same resource using the same policy document will return the same revision ID. If
you specify an `ExpectedRevisionId` that doesn't match the current policy's
`RevisionId`, the `PolicyNotFoundException` will be
returned.

###### Note

`PutResourcePolicy` is an asynchronous operation. If you issue a
`GetResourcePolicy` request immediately after a
`PutResourcePolicy` request, DynamoDB might return your
previous policy, if there was one, or return the
`PolicyNotFoundException`. This is because
`GetResourcePolicy` uses an eventually consistent query, and the
metadata for your policy or table might not be available at that moment. Wait for a
few seconds, and then try the `GetResourcePolicy` request again.

## Request Syntax

```nohighlight

{
   "ConfirmRemoveSelfResourceAccess": boolean,
   "ExpectedRevisionId": "string",
   "Policy": "string",
   "ResourceArn": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[Policy](#API_PutResourcePolicy_RequestSyntax)**

An AWS resource-based policy document in JSON format.

- The maximum size supported for a resource-based policy document is 20 KB.
DynamoDB counts whitespaces when calculating the size of a policy
against this limit.

- Within a resource-based policy, if the action for a DynamoDB
service-linked role (SLR) to replicate data for a global table is denied, adding
or deleting a replica will fail with an error.

For a full list of all considerations that apply while attaching a resource-based
policy, see [Resource-based\
policy considerations](../../../../services/dynamodb/latest/developerguide/rbac-considerations.md).

Type: String

Required: Yes

**[ResourceArn](#API_PutResourcePolicy_RequestSyntax)**

The Amazon Resource Name (ARN) of the DynamoDB resource to which the policy will be attached.
The resources you can specify include tables and streams.

You can control index permissions using the base table's policy. To specify the same permission level for your table and its indexes, you can provide both the table and index Amazon Resource Name (ARN)s in the `Resource` field of a given `Statement` in your policy document. Alternatively, to specify different permissions for your table, indexes, or both, you can define multiple `Statement` fields in your policy document.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: Yes

**[ConfirmRemoveSelfResourceAccess](#API_PutResourcePolicy_RequestSyntax)**

Set this parameter to `true` to confirm that you want to remove your
permissions to change the policy of this resource in the future.

Type: Boolean

Required: No

**[ExpectedRevisionId](#API_PutResourcePolicy_RequestSyntax)**

A string value that you can use to conditionally update your policy. You can provide
the revision ID of your existing policy to make mutating requests against that
policy.

###### Note

When you provide an expected revision ID, if the revision ID of the existing
policy on the resource doesn't match or if there's no policy attached to the
resource, your request will be rejected with a
`PolicyNotFoundException`.

To conditionally attach a policy when no policy exists for the resource, specify
`NO_POLICY` for the revision ID.

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

**[RevisionId](#API_PutResourcePolicy_ResponseSyntax)**

A unique string that represents the revision ID of the policy. If you're comparing revision IDs, make sure to always use string comparison logic.

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

### Attach a resource-based policy to a table

The following example attaches a resource-based policy to a table named
`Thread`.

To view more examples of resource-based policies, see [Resource-based\
policy examples](../../../../services/dynamodb/latest/developerguide/rbac-examples.md).

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
X-Amz-Target: DynamoDB_20120810.PutResourcePolicy
{
    "ResourceArn": "arn:aws:dynamodb:us-west-2:123456789012:table/Thread",
    "Policy": "{\"Version\":\"2012-10-17\",\"Statement\":{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"arn:aws:iam::111122223333:root\",\"arn:aws:iam::444455556666:root\"]},\"Action\":[\"dynamodb:GetItem\"],\"Resource\":\"arn:aws:dynamodb:us-west-2:123456789012:table/Thread\"}}"
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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/putresourcepolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/putresourcepolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutItem

Query

All content copied from https://docs.aws.amazon.com/.
