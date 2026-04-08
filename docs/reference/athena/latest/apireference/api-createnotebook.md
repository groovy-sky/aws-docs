# CreateNotebook

Creates an empty `ipynb` file in the specified Apache Spark enabled
workgroup. Throws an error if a file in the workgroup with the same name already
exists.

## Request Syntax

```nohighlight

{
   "ClientRequestToken": "string",
   "Name": "string",
   "WorkGroup": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ClientRequestToken](#API_CreateNotebook_RequestSyntax)**

A unique case-sensitive string used to ensure the request to create the notebook is
idempotent (executes only once).

###### Important

This token is listed as not required because AWS SDKs (for example
the AWS SDK for Java) auto-generate the token for you. If you are not
using the AWS SDK or the AWS CLI, you must provide
this token or the action will fail.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 36.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: No

**[Name](#API_CreateNotebook_RequestSyntax)**

The name of the `ipynb` file to be created in the Spark workgroup, without
the `.ipynb` extension.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `(?!.*[/:\\])[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]+`

Required: Yes

**[WorkGroup](#API_CreateNotebook_RequestSyntax)**

The name of the Spark enabled workgroup in which the notebook will be created.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: Yes

## Response Syntax

```nohighlight

{
   "NotebookId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NotebookId](#API_CreateNotebook_ResponseSyntax)**

A unique identifier for the notebook.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 36.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

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

**TooManyRequestsException**

Indicates that the request was throttled.

**Reason**

The reason for the query throttling, for example, when it exceeds the concurrent query
limit.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/createnotebook.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/createnotebook.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/createnotebook.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/createnotebook.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/createnotebook.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/createnotebook.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/createnotebook.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/createnotebook.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/createnotebook.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/createnotebook.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateNamedQuery

CreatePreparedStatement
