# CreateNamedQuery

Creates a named query in the specified workgroup. Requires that you have access to the
workgroup.

## Request Syntax

```nohighlight

{
   "ClientRequestToken": "string",
   "Database": "string",
   "Description": "string",
   "Name": "string",
   "QueryString": "string",
   "WorkGroup": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ClientRequestToken](#API_CreateNamedQuery_RequestSyntax)**

A unique case-sensitive string used to ensure the request to create the query is
idempotent (executes only once). If another `CreateNamedQuery` request is
received, the same response is returned and another query is not created. If a parameter
has changed, for example, the `QueryString`, an error is returned.

###### Important

This token is listed as not required because AWS SDKs (for example
the AWS SDK for Java) auto-generate the token for users. If you are
not using the AWS SDK or the AWS CLI, you must provide
this token or the action will fail.

Type: String

Length Constraints: Minimum length of 32. Maximum length of 128.

Required: No

**[Database](#API_CreateNamedQuery_RequestSyntax)**

The database to which the query belongs.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[Description](#API_CreateNamedQuery_RequestSyntax)**

The query description.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**[Name](#API_CreateNamedQuery_RequestSyntax)**

The query name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**[QueryString](#API_CreateNamedQuery_RequestSyntax)**

The contents of the query with all query statements.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 262144.

Required: Yes

**[WorkGroup](#API_CreateNamedQuery_RequestSyntax)**

The name of the workgroup in which the named query is being created.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: No

## Response Syntax

```nohighlight

{
   "NamedQueryId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NamedQueryId](#API_CreateNamedQuery_ResponseSyntax)**

The unique ID of the query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `\S+`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/createnamedquery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/createnamedquery.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDataCatalog

CreateNotebook

All content copied from https://docs.aws.amazon.com/.
