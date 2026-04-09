# BatchGetNamedQuery

Returns the details of a single named query or a list of up to 50 queries, which you
provide as an array of query ID strings. Requires you to have access to the workgroup in
which the queries were saved. Use [ListNamedQueriesInput](api-listnamedqueriesinput.md) to get the
list of named query IDs in the specified workgroup. If information could not be
retrieved for a submitted query ID, information about the query ID submitted is listed
under [UnprocessedNamedQueryId](api-unprocessednamedqueryid.md). Named queries differ from executed
queries. Use [BatchGetQueryExecutionInput](api-batchgetqueryexecutioninput.md) to get details about each
unique query execution, and [ListQueryExecutionsInput](api-listqueryexecutionsinput.md) to get a list of
query execution IDs.

## Request Syntax

```nohighlight

{
   "NamedQueryIds": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[NamedQueryIds](#API_BatchGetNamedQuery_RequestSyntax)**

An array of query IDs.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `\S+`

Required: Yes

## Response Syntax

```nohighlight

{
   "NamedQueries": [
      {
         "Database": "string",
         "Description": "string",
         "Name": "string",
         "NamedQueryId": "string",
         "QueryString": "string",
         "WorkGroup": "string"
      }
   ],
   "UnprocessedNamedQueryIds": [
      {
         "ErrorCode": "string",
         "ErrorMessage": "string",
         "NamedQueryId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NamedQueries](#API_BatchGetNamedQuery_ResponseSyntax)**

Information about the named query IDs submitted.

Type: Array of [NamedQuery](api-namedquery.md) objects

**[UnprocessedNamedQueryIds](#API_BatchGetNamedQuery_ResponseSyntax)**

Information about provided query IDs.

Type: Array of [UnprocessedNamedQueryId](api-unprocessednamedqueryid.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/batchgetnamedquery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/batchgetnamedquery.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

BatchGetPreparedStatement

All content copied from https://docs.aws.amazon.com/.
