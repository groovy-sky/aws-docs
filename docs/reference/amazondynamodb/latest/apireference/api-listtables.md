# ListTables

Returns an array of table names associated with the current account and endpoint. The
output from `ListTables` is paginated, with each page returning a maximum of
100 table names.

## Request Syntax

```nohighlight

{
   "ExclusiveStartTableName": "string",
   "Limit": number
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ExclusiveStartTableName](#API_ListTables_RequestSyntax)**

The first table name that this operation will evaluate. Use the value that was
returned for `LastEvaluatedTableName` in a previous operation, so that you
can obtain the next page of results.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**[Limit](#API_ListTables_RequestSyntax)**

A maximum number of table names to return. If this parameter is not specified, the
limit is 100.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

## Response Syntax

```nohighlight

{
   "LastEvaluatedTableName": "string",
   "TableNames": [ "string" ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[LastEvaluatedTableName](#API_ListTables_ResponseSyntax)**

The name of the last table in the current page of results. Use this value as the
`ExclusiveStartTableName` in a new request to obtain the next page of
results, until all the table names are returned.

If you do not receive a `LastEvaluatedTableName` value in the response,
this means that there are no more table names to be retrieved.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

**[TableNames](#API_ListTables_ResponseSyntax)**

The names of the tables associated with the current account at the current endpoint.
The maximum size of this array is 100.

If `LastEvaluatedTableName` also appears in the output, you can use this
value as the `ExclusiveStartTableName` parameter in a subsequent
`ListTables` request and obtain the next page of results.

Type: Array of strings

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

## Examples

### List Tables

This example requests a list of tables, starting with a table named
`Forum` and ending after three table names have been
returned.

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
X-Amz-Target: DynamoDB_20120810.ListTables

{
    "ExclusiveStartTableName": "Forum",
    "Limit": 3
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
    "LastEvaluatedTableName": "Thread",
    "TableNames": ["Forum","Reply","Thread"]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/listtables.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/listtables.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListImports

ListTagsOfResource

All content copied from https://docs.aws.amazon.com/.
