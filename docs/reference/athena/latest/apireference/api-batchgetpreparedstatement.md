# BatchGetPreparedStatement

Returns the details of a single prepared statement or a list of up to 256 prepared
statements for the array of prepared statement names that you provide. Requires you to
have access to the workgroup to which the prepared statements belong. If a prepared
statement cannot be retrieved for the name specified, the statement is listed in
`UnprocessedPreparedStatementNames`.

## Request Syntax

```nohighlight

{
   "PreparedStatementNames": [ "string" ],
   "WorkGroup": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[PreparedStatementNames](#API_BatchGetPreparedStatement_RequestSyntax)**

A list of prepared statement names to return.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[a-zA-Z_][a-zA-Z0-9_@:]{1,256}`

Required: Yes

**[WorkGroup](#API_BatchGetPreparedStatement_RequestSyntax)**

The name of the workgroup to which the prepared statements belong.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: Yes

## Response Syntax

```nohighlight

{
   "PreparedStatements": [
      {
         "Description": "string",
         "LastModifiedTime": number,
         "QueryStatement": "string",
         "StatementName": "string",
         "WorkGroupName": "string"
      }
   ],
   "UnprocessedPreparedStatementNames": [
      {
         "ErrorCode": "string",
         "ErrorMessage": "string",
         "StatementName": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[PreparedStatements](#API_BatchGetPreparedStatement_ResponseSyntax)**

The list of prepared statements returned.

Type: Array of [PreparedStatement](api-preparedstatement.md) objects

**[UnprocessedPreparedStatementNames](#API_BatchGetPreparedStatement_ResponseSyntax)**

A list of one or more prepared statements that were requested but could not be
returned.

Type: Array of [UnprocessedPreparedStatementName](api-unprocessedpreparedstatementname.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/batchgetpreparedstatement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/batchgetpreparedstatement.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchGetNamedQuery

BatchGetQueryExecution

All content copied from https://docs.aws.amazon.com/.
