# GetExport

Exports a deployed version of a RestApi in a specified format.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/stages/stage_name/exports/export_type?parameters HTTP/1.1
Accept: accepts

```

## URI Request Parameters

The request uses the following URI parameters.

**[accepts](#API_GetExport_RequestSyntax)**

The content-type of the export, for example `application/json`. Currently `application/json` and `application/yaml` are supported for `exportType` of `oas30` and `swagger`. This should be specified in the `Accept` header for direct API requests.

**[export\_type](#API_GetExport_RequestSyntax)**

The type of export. Acceptable values are 'oas30' for OpenAPI 3.0.x and 'swagger' for Swagger/OpenAPI 2.0.

Required: Yes

**[parameters](#API_GetExport_RequestSyntax)**

A key-value map of query string parameters that specify properties of the export, depending on the requested `exportType`. For `exportType` `oas30` and `swagger`, any combination of the following parameters are supported: `extensions='integrations'` or `extensions='apigateway'` will export the API with x-amazon-apigateway-integration extensions. `extensions='authorizers'` will export the API with x-amazon-apigateway-authorizer extensions. `postman` will export the API with Postman extensions, allowing for import to the Postman tool

**[restapi\_id](#API_GetExport_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

**[stage\_name](#API_GetExport_RequestSyntax)**

The name of the Stage that will be exported.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-Type: contentType
Content-Disposition: contentDisposition

body
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[contentDisposition](#API_GetExport_ResponseSyntax)**

The content-disposition header value in the HTTP response.

**[contentType](#API_GetExport_ResponseSyntax)**

The content-type header value in the HTTP response. This will correspond to a valid 'accept' type in the request.

The response returns the following as the HTTP body.

**[body](#API_GetExport_ResponseSyntax)**

The binary blob response to GetExport, which contains the export.

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getexport.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getexport.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getexport.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getexport.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getexport.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getexport.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getexport.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getexport.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getexport.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getexport.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetDomainNames

GetGatewayResponse

All content copied from https://docs.aws.amazon.com/.
