# GetDocumentationVersion

Gets a documentation version.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/documentation/versions/doc_version HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[doc\_version](#API_GetDocumentationVersion_RequestSyntax)**

The version identifier of the to-be-retrieved documentation snapshot.

Required: Yes

**[restapi\_id](#API_GetDocumentationVersion_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "createdDate": number,
   "description": "string",
   "version": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[createdDate](#API_GetDocumentationVersion_ResponseSyntax)**

The date when the API documentation snapshot is created.

Type: Timestamp

**[description](#API_GetDocumentationVersion_ResponseSyntax)**

The description of the API documentation snapshot.

Type: String

**[version](#API_GetDocumentationVersion_ResponseSyntax)**

The version identifier of the API documentation snapshot.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getdocumentationversion.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getdocumentationversion.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetDocumentationParts

GetDocumentationVersions

All content copied from https://docs.aws.amazon.com/.
