# CreateDocumentationVersion

Creates a documentation version

## Request Syntax

```nohighlight

POST /restapis/restapi_id/documentation/versions HTTP/1.1
Content-type: application/json

{
   "description": "string",
   "documentationVersion": "string",
   "stageName": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[restapi\_id](#API_CreateDocumentationVersion_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[description](#API_CreateDocumentationVersion_RequestSyntax)**

A description about the new documentation snapshot.

Type: String

Required: No

**[documentationVersion](#API_CreateDocumentationVersion_RequestSyntax)**

The version identifier of the new snapshot.

Type: String

Required: Yes

**[stageName](#API_CreateDocumentationVersion_RequestSyntax)**

The stage name to be associated with the new documentation snapshot.

Type: String

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "createdDate": number,
   "description": "string",
   "version": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[createdDate](#API_CreateDocumentationVersion_ResponseSyntax)**

The date when the API documentation snapshot is created.

Type: Timestamp

**[description](#API_CreateDocumentationVersion_ResponseSyntax)**

The description of the API documentation snapshot.

Type: String

**[version](#API_CreateDocumentationVersion_ResponseSyntax)**

The version identifier of the API documentation snapshot.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

**LimitExceededException**

The request exceeded the rate limit. Retry after the specified time period.

HTTP Status Code: 429

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/createdocumentationversion.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/createdocumentationversion.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDocumentationPart

CreateDomainName

All content copied from https://docs.aws.amazon.com/.
