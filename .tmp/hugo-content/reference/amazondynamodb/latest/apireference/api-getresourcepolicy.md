# GetResourcePolicy

Returns the resource-based policy document attached to the resource, which can be a
table or stream, in JSON format.

`GetResourcePolicy` follows an [_eventually consistent_](../../../../services/dynamodb/latest/developerguide/howitworks-readconsistency.md) model. The following list
describes the outcomes when you issue the `GetResourcePolicy` request
immediately after issuing another request:

- If you issue a `GetResourcePolicy` request immediately after a
`PutResourcePolicy` request, DynamoDB might return a
`PolicyNotFoundException`.

- If you issue a `GetResourcePolicy` request immediately after a
`DeleteResourcePolicy` request, DynamoDB might return
the policy that was present before the deletion request.

- If you issue a `GetResourcePolicy` request immediately after a
`CreateTable` request, which includes a resource-based policy,
DynamoDB might return a `ResourceNotFoundException` or
a `PolicyNotFoundException`.

Because `GetResourcePolicy` uses an _eventually_
_consistent_ query, the metadata for your policy or table might not be
available at that moment. Wait for a few seconds, and then retry the
`GetResourcePolicy` request.

After a `GetResourcePolicy` request returns a policy created using the
`PutResourcePolicy` request, the policy will be applied in the
authorization of requests to the resource. Because this process is eventually
consistent, it will take some time to apply the policy to all requests to a resource.
Policies that you attach while creating a table using the `CreateTable`
request will always be applied to all requests for that table.

## Request Syntax

```nohighlight

{
   "ResourceArn": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ResourceArn](#API_GetResourcePolicy_RequestSyntax)**

The Amazon Resource Name (ARN) of the DynamoDB resource to which the policy is attached. The
resources you can specify include tables and streams.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: Yes

## Response Syntax

```nohighlight

{
   "Policy": "string",
   "RevisionId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Policy](#API_GetResourcePolicy_ResponseSyntax)**

The resource-based policy document attached to the resource, which can be a table or
stream, in JSON format.

Type: String

**[RevisionId](#API_GetResourcePolicy_ResponseSyntax)**

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

**PolicyNotFoundException**

The operation tried to access a nonexistent resource-based policy.

If you specified an `ExpectedRevisionId`, it's possible that a policy is
present for the resource but its revision ID didn't match the expected value.

HTTP Status Code: 400

**ResourceNotFoundException**

The operation tried to access a nonexistent table or index. The resource might not
be specified correctly, or its status might not be `ACTIVE`.

**message**

The resource which is being requested does not exist.

HTTP Status Code: 400

## Examples

### Get the resource-based policy of a table

The following example retrieves the resource-based policy of a table named
`Thread`.

#### Sample Request

```

GET / HTTP/1.1
Host: dynamodb.<region>.<domain>;
Accept-Encoding: identity
Content-Length: <PayloadSizeBytes>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.0
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=<Headers>, Signature=<Signature>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDB_20120810.GetResourcePolicy
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
    "Policy": "{\"Version\":\"2012-10-17\",\"Statement\":{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"arn:aws:iam::111122223333:root\",\"arn:aws:iam::444455556666:root\"]},\"Action\":[\"dynamodb:GetItem\"],\"Resource\":\"arn:aws:dynamodb:us-west-2:123456789012:table/Thread\"}}",
    "RevisionId": "1683717331354"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/getresourcepolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/getresourcepolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetItem

ImportTable

All content copied from https://docs.aws.amazon.com/.
