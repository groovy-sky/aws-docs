# GetDocumentationParts

Gets documentation parts.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/documentation/parts?limit=limit&locationStatus=locationStatus&name=nameQuery&path=path&position=position&type=type HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetDocumentationParts_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[locationStatus](#API_GetDocumentationParts_RequestSyntax)**

The status of the API documentation parts to retrieve. Valid values are `DOCUMENTED` for retrieving DocumentationPart resources with content and `UNDOCUMENTED` for DocumentationPart resources without content.

Valid Values: `DOCUMENTED | UNDOCUMENTED`

**[nameQuery](#API_GetDocumentationParts_RequestSyntax)**

The name of API entities of the to-be-retrieved documentation parts.

**[path](#API_GetDocumentationParts_RequestSyntax)**

The path of API entities of the to-be-retrieved documentation parts.

**[position](#API_GetDocumentationParts_RequestSyntax)**

The current pagination position in the paged result set.

**[restapi\_id](#API_GetDocumentationParts_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

**[type](#API_GetDocumentationParts_RequestSyntax)**

The type of API entities of the to-be-retrieved documentation parts.

Valid Values: `API | AUTHORIZER | MODEL | RESOURCE | METHOD | PATH_PARAMETER | QUERY_PARAMETER | REQUEST_HEADER | REQUEST_BODY | RESPONSE | RESPONSE_HEADER | RESPONSE_BODY`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "item": [
      {
         "id": "string",
         "location": {
            "method": "string",
            "name": "string",
            "path": "string",
            "statusCode": "string",
            "type": "string"
         },
         "properties": "string"
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetDocumentationParts_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [DocumentationPart](api-documentationpart.md) objects

**[position](#API_GetDocumentationParts_ResponseSyntax)**

The current pagination position in the paged result set.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getdocumentationparts.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getdocumentationparts.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetDocumentationPart

GetDocumentationVersion

All content copied from https://docs.aws.amazon.com/.
