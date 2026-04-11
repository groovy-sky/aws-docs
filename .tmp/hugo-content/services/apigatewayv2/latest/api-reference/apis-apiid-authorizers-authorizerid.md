# Authorizer

Represents an authorizer. `JWT` authorizers are supported only for HTTP APIs.

## URI

`/v2/apis/apiId/authorizers/authorizerId`

## HTTP methods

### GET

**Operation ID:** `GetAuthorizer`

Gets an `Authorizer`.

Path parametersNameTypeRequiredDescription`authorizerId`StringTrue

The authorizer identifier.

`apiId`StringTrue

The API identifier.

ResponsesStatus codeResponse modelDescription`200``Authorizer`

Success

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### DELETE

**Operation ID:** `DeleteAuthorizer`

Deletes an `Authorizer`.

Path parametersNameTypeRequiredDescription`authorizerId`StringTrue

The authorizer identifier.

`apiId`StringTrue

The API identifier.

ResponsesStatus codeResponse modelDescription`204`None

The request has succeeded, and there is no additional content to send in the response payload body.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### PATCH

**Operation ID:** `UpdateAuthorizer`

Updates an `Authorizer`.

Path parametersNameTypeRequiredDescription`authorizerId`StringTrue

The authorizer identifier.

`apiId`StringTrue

The API identifier.

ResponsesStatus codeResponse modelDescription`200``Authorizer`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`409``ConflictException`

The resource already exists.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

## Schemas

### Request bodies

```json

{
  "name": "string",
  "authorizerType": enum,
  "authorizerUri": "string",
  "authorizerCredentialsArn": "string",
  "identitySource": [
    "string"
  ],
  "identityValidationExpression": "string",
  "authorizerResultTtlInSeconds": integer,
  "jwtConfiguration": {
    "issuer": "string",
    "audience": [
      "string"
    ]
  },
  "authorizerPayloadFormatVersion": "string",
  "enableSimpleResponses": boolean
}
```

### Response bodies

```json

{
  "authorizerId": "string",
  "name": "string",
  "authorizerType": enum,
  "authorizerUri": "string",
  "authorizerCredentialsArn": "string",
  "identitySource": [
    "string"
  ],
  "identityValidationExpression": "string",
  "authorizerResultTtlInSeconds": integer,
  "jwtConfiguration": {
    "issuer": "string",
    "audience": [
      "string"
    ]
  },
  "authorizerPayloadFormatVersion": "string",
  "enableSimpleResponses": boolean
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "resourceType": "string"
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "limitType": "string"
}
```

## Properties

### Authorizer

Represents an authorizer.

PropertyTypeRequiredDescription`authorizerCredentialsArn`

string

False

Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, don't specify this parameter. Supported only for `REQUEST` authorizers.

`authorizerId`

string

False

The authorizer identifier.

`authorizerPayloadFormatVersion`

string

False

Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are `1.0` and `2.0`. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

`authorizerResultTtlInSeconds`

integer

False

The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization
caching is disabled. If it is greater than 0, API Gateway caches authorizer
responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.

`authorizerType`

[AuthorizerType](#apis-apiid-authorizers-authorizerid-model-authorizertype)

False

The authorizer type. Specify `REQUEST` for a Lambda function using incoming request parameters. Specify `JWT` to use JSON Web Tokens (supported only for HTTP APIs).

`authorizerUri`

string

False

The authorizer's Uniform Resource Identifier (URI). For `REQUEST` authorizers, this must be a well-formed Lambda function URI, for example, `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations`. In general, the URI has this form: `arn:aws:apigateway:{region}:lambda:path/{service_api}
               `, where ``{region} is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial `/`. For Lambda functions, this is usually of the form `/2015-03-31/functions/[FunctionARN]/invocations`. Supported only for `REQUEST` authorizers.

`enableSimpleResponses`

boolean

False

Specifies whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

`identitySource`

Array of type string

False

The identity source for which authorization is requested.

For a `REQUEST` authorizer, this is optional. The value is a set of
one or more mapping expressions of the specified request parameters. The
identity source can be headers, query string parameters, stage variables, and
context parameters. For example, if an Auth header and a Name query string
parameter are defined as identity sources, this value is
route.request.header.Auth, route.request.querystring.Name for WebSocket APIs.
For HTTP APIs, use selection expressions prefixed with `$`, for
example, `$request.header.Auth`,
`$request.querystring.Name`. These parameters are used to perform
runtime validation for Lambda-based authorizers by verifying all of the
identity-related request parameters are present in the request, not null, and
non-empty. Only when this is true does the authorizer invoke the authorizer
Lambda function. Otherwise, it returns a 401 Unauthorized response without
calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

For `JWT`, a single entry that specifies where to extract the JSON
Web Token (JWT) from inbound requests. Currently only header-based and query
parameter-based selections are supported, for example
`$request.header.Authorization`.

`identityValidationExpression`

string

False

The validation expression does not apply to the `REQUEST` authorizer.

`jwtConfiguration`

[JWTConfiguration](#apis-apiid-authorizers-authorizerid-model-jwtconfiguration)

False

Represents the configuration of a JWT authorizer. Required for the `JWT`
authorizer type. Supported only for HTTP APIs.

`name`

string

True

The name of the authorizer.

### AuthorizerType

The authorizer type. Specify `REQUEST` for a Lambda function using incoming request parameters. Specify `JWT` to use JSON Web Tokens (supported only for HTTP APIs).

- `REQUEST`

- `JWT`

### BadRequestException

The request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### ConflictException

The requested operation would cause a conflict with the current state of a service resource associated with the request. Resolve the conflict before retrying this request. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### JWTConfiguration

Represents the configuration of a JWT authorizer. Required for the `JWT`
authorizer type. Supported only for HTTP APIs.

PropertyTypeRequiredDescription`audience`

Array of type string

False

A list of the intended recipients of the JWT. A valid JWT must provide an
`aud` that matches at least one entry in this list. See [RFC 7519](https://tools.ietf.org/html/rfc7519).
Required for the `JWT` authorizer type. Supported only for HTTP APIs.

`issuer`

string

False

The base domain of the identity provider that issues JSON Web Tokens. For example,
an Amazon Cognito user pool has the following format:
`https://cognito-idp.{region}.amazonaws.com/{userPoolId}
               `.
Required for the `JWT` authorizer type. Supported only for HTTP APIs.

### LimitExceededException

A limit has been exceeded. See the accompanying error message for details.

PropertyTypeRequiredDescription`limitType`

string

False

The limit type.

`message`

string

False

Describes the error encountered.

### NotFoundException

The resource specified in the request was not found. See the `message` field for more information.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

`resourceType`

string

False

The resource type.

### UpdateAuthorizerInput

The input parameters for an `UpdateAuthorizer` request.

PropertyTypeRequiredDescription`authorizerCredentialsArn`

string

False

Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, don't specify this parameter.

`authorizerPayloadFormatVersion`

string

False

Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are `1.0` and `2.0`. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

`authorizerResultTtlInSeconds`

integer

False

The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization
caching is disabled. If it is greater than 0, API Gateway caches authorizer
responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.

`authorizerType`

[AuthorizerType](#apis-apiid-authorizers-authorizerid-model-authorizertype)

False

The authorizer type. Specify `REQUEST` for a Lambda function using incoming request parameters. Specify `JWT` to use JSON Web Tokens (supported only for HTTP APIs).

`authorizerUri`

string

False

The authorizer's Uniform Resource Identifier (URI). For `REQUEST` authorizers, this must be a well-formed Lambda function URI, for example, `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations`. In general, the URI has this form: `arn:aws:apigateway:{region}:lambda:path/{service_api}
               `, where ``{region} is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial `/`. For Lambda functions, this is usually of the form `/2015-03-31/functions/[FunctionARN]/invocations`. Supported only for `REQUEST` authorizers.

`enableSimpleResponses`

boolean

False

Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

`identitySource`

Array of type string

False

The identity source for which authorization is requested.

For a `REQUEST` authorizer, this is optional. The value is a set of
one or more mapping expressions of the specified request parameters. The
identity source can be headers, query string parameters, stage variables, and
context parameters. For example, if an Auth header and a Name query string
parameter are defined as identity sources, this value is
route.request.header.Auth, route.request.querystring.Name for WebSocket APIs.
For HTTP APIs, use selection expressions prefixed with `$`, for
example, `$request.header.Auth`,
`$request.querystring.Name`. These parameters are used to perform
runtime validation for Lambda-based authorizers by verifying all of the
identity-related request parameters are present in the request, not null, and
non-empty. Only when this is true does the authorizer invoke the authorizer
Lambda function. Otherwise, it returns a 401 Unauthorized response without
calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

For `JWT`, a single entry that specifies where to extract the JSON
Web Token (JWT) from inbound requests. Currently only header-based and query
parameter-based selections are supported, for example
`$request.header.Authorization`.

`identityValidationExpression`

string

False

This parameter is not used.

`jwtConfiguration`

[JWTConfiguration](#apis-apiid-authorizers-authorizerid-model-jwtconfiguration)

False

Represents the configuration of a JWT authorizer. Required for the
`JWT` authorizer type. Supported only for HTTP APIs.

`name`

string

False

The name of the authorizer.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetAuthorizer

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/getauthorizer.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/getauthorizer.md)

### DeleteAuthorizer

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/deleteauthorizer.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/deleteauthorizer.md)

### UpdateAuthorizer

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/updateauthorizer.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/updateauthorizer.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Apis

Authorizers

All content copied from https://docs.aws.amazon.com/.
