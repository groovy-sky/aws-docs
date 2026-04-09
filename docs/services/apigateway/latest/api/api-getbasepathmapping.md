# GetBasePathMapping

Describe a BasePathMapping resource.

## Request Syntax

```nohighlight

GET /domainnames/domain_name/basepathmappings/base_path?domainNameId=domainNameId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[base\_path](#API_GetBasePathMapping_RequestSyntax)**

The base path name that callers of the API must provide as part of the URL after the domain name. This value must be unique for all of the mappings across a single API. Specify '(none)' if you do not want callers to specify any base path name after the domain name.

Required: Yes

**[domain\_name](#API_GetBasePathMapping_RequestSyntax)**

The domain name of the BasePathMapping resource to be described.

Required: Yes

**[domainNameId](#API_GetBasePathMapping_RequestSyntax)**

The identifier for the domain name resource. Supported only for private custom domain names.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "basePath": "string",
   "restApiId": "string",
   "stage": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[basePath](#API_GetBasePathMapping_ResponseSyntax)**

The base path name that callers of the API must provide as part of the URL after the domain name.

Type: String

**[restApiId](#API_GetBasePathMapping_ResponseSyntax)**

The string identifier of the associated RestApi.

Type: String

**[stage](#API_GetBasePathMapping_ResponseSyntax)**

The name of the associated stage.

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getbasepathmapping.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getbasepathmapping.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAuthorizers

GetBasePathMappings

All content copied from https://docs.aws.amazon.com/.
