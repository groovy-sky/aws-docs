# ImportNotebook

Imports a single `ipynb` file to a Spark enabled workgroup. To import the
notebook, the request must specify a value for either `Payload` or
`NoteBookS3LocationUri`. If neither is specified or both are specified,
an `InvalidRequestException` occurs. The maximum file size that can be
imported is 10 megabytes. If an `ipynb` file with the same name already
exists in the workgroup, throws an error.

## Request Syntax

```nohighlight

{
   "ClientRequestToken": "string",
   "Name": "string",
   "NotebookS3LocationUri": "string",
   "Payload": "string",
   "Type": "string",
   "WorkGroup": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ClientRequestToken](#API_ImportNotebook_RequestSyntax)**

A unique case-sensitive string used to ensure the request to import the notebook is
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

**[Name](#API_ImportNotebook_RequestSyntax)**

The name of the notebook to import.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `(?!.*[/:\\])[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]+`

Required: Yes

**[NotebookS3LocationUri](#API_ImportNotebook_RequestSyntax)**

A URI that specifies the Amazon S3 location of a notebook file in
`ipynb` format.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `^(https|s3|S3)://([^/]+)/?(.*)$`

Required: No

**[Payload](#API_ImportNotebook_RequestSyntax)**

The notebook content to be imported. The payload must be in `ipynb`
format.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10485760.

Required: No

**[Type](#API_ImportNotebook_RequestSyntax)**

The notebook content type. Currently, the only valid type is
`IPYNB`.

Type: String

Valid Values: `IPYNB`

Required: Yes

**[WorkGroup](#API_ImportNotebook_RequestSyntax)**

The name of the Spark enabled workgroup to import the notebook to.

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

**[NotebookId](#API_ImportNotebook_ResponseSyntax)**

The ID assigned to the imported notebook.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/importnotebook.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/importnotebook.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/importnotebook.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/importnotebook.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/importnotebook.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/importnotebook.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/importnotebook.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/importnotebook.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/importnotebook.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/importnotebook.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetWorkGroup

ListApplicationDPUSizes
