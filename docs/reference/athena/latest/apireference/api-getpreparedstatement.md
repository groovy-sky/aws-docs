# GetPreparedStatement

Retrieves the prepared statement with the specified name from the specified
workgroup.

## Request Syntax

```nohighlight

{
   "StatementName": "string",
   "WorkGroup": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[StatementName](#API_GetPreparedStatement_RequestSyntax)**

The name of the prepared statement to retrieve.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[a-zA-Z_][a-zA-Z0-9_@:]{1,256}`

Required: Yes

**[WorkGroup](#API_GetPreparedStatement_RequestSyntax)**

The workgroup to which the statement to be retrieved belongs.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: Yes

## Response Syntax

```nohighlight

{
   "PreparedStatement": {
      "Description": "string",
      "LastModifiedTime": number,
      "QueryStatement": "string",
      "StatementName": "string",
      "WorkGroupName": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[PreparedStatement](#API_GetPreparedStatement_ResponseSyntax)**

The name of the prepared statement that was retrieved.

Type: [PreparedStatement](api-preparedstatement.md) object

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

**ResourceNotFoundException**

A resource, such as a workgroup, was not found.

**ResourceName**

The name of the Amazon resource.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/getpreparedstatement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/getpreparedstatement.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetNotebookMetadata

GetQueryExecution

All content copied from https://docs.aws.amazon.com/.
