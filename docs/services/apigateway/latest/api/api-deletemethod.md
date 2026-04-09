# DeleteMethod

Deletes an existing Method resource.

## Request Syntax

```nohighlight

DELETE /restapis/restapi_id/resources/resource_id/methods/http_method HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[http\_method](#API_DeleteMethod_RequestSyntax)**

The HTTP verb of the Method resource.

Required: Yes

**[resource\_id](#API_DeleteMethod_RequestSyntax)**

The Resource identifier for the Method resource.

Required: Yes

**[restapi\_id](#API_DeleteMethod_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/deletemethod.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/deletemethod.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteIntegrationResponse

DeleteMethodResponse

All content copied from https://docs.aws.amazon.com/.
