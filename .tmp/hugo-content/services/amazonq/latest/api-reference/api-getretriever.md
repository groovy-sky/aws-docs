# GetRetriever

Gets information about an existing retriever used by an Amazon Q Business
application.

## Request Syntax

```nohighlight

GET /applications/applicationId/retrievers/retrieverId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_GetRetriever_RequestSyntax)**

The identifier of the Amazon Q Business application using the retriever.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[retrieverId](#API_GetRetriever_RequestSyntax)**

The identifier of the retriever.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "applicationId": "string",
   "configuration": { ... },
   "createdAt": number,
   "displayName": "string",
   "retrieverArn": "string",
   "retrieverId": "string",
   "roleArn": "string",
   "status": "string",
   "type": "string",
   "updatedAt": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[applicationId](#API_GetRetriever_ResponseSyntax)**

The identifier of the Amazon Q Business application using the retriever.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[configuration](#API_GetRetriever_ResponseSyntax)**

Provides information on how the retriever used for your Amazon Q Business application is
configured.

Type: [RetrieverConfiguration](api-retrieverconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

**[createdAt](#API_GetRetriever_ResponseSyntax)**

The Unix timestamp when the retriever was created.

Type: Timestamp

**[displayName](#API_GetRetriever_ResponseSyntax)**

The name of the retriever.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

**[retrieverArn](#API_GetRetriever_ResponseSyntax)**

The Amazon Resource Name (ARN) of the IAM role associated with the retriever.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

**[retrieverId](#API_GetRetriever_ResponseSyntax)**

The identifier of the retriever.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[roleArn](#API_GetRetriever_ResponseSyntax)**

The Amazon Resource Name (ARN) of the role with the permission to access the retriever
and required resources.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

**[status](#API_GetRetriever_ResponseSyntax)**

The status of the retriever.

Type: String

Valid Values: `CREATING | ACTIVE | FAILED`

**[type](#API_GetRetriever_ResponseSyntax)**

The type of the retriever.

Type: String

Valid Values: `NATIVE_INDEX | KENDRA_INDEX`

**[updatedAt](#API_GetRetriever_ResponseSyntax)**

The Unix timestamp when the retriever was last updated.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**ResourceNotFoundException**

The application or plugin resource you want to use doesn’t exist. Make sure you have
provided the correct resource and try again.

**message**

The message describing a `ResourceNotFoundException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/getretriever.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/getretriever.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetPolicy

GetUser

All content copied from https://docs.aws.amazon.com/.
