# Route

Represents a route for an API.

## URI

`/v2/apis/apiId/routes/routeId`

## HTTP methods

### GET

**Operation ID:** `GetRoute`

Gets a `Route`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`routeId`StringTrue

The route ID.

ResponsesStatus codeResponse modelDescription`200``Route`

Success

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### DELETE

**Operation ID:** `DeleteRoute`

Deletes a `Route`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`routeId`StringTrue

The route ID.

ResponsesStatus codeResponse modelDescription`204`None

The request has succeeded, and there is no additional content to send in the response payload body.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### PATCH

**Operation ID:** `UpdateRoute`

Updates a `Route`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`routeId`StringTrue

The route ID.

ResponsesStatus codeResponse modelDescription`200``Route`

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
  "routeKey": "string",
  "authorizationType": enum,
  "authorizerId": "string",
  "authorizationScopes": [
    "string"
  ],
  "apiKeyRequired": boolean,
  "requestParameters": {
  },
  "requestModels": {
  },
  "modelSelectionExpression": "string",
  "target": "string",
  "operationName": "string",
  "routeResponseSelectionExpression": "string"
}
```

### Response bodies

```json

{
  "routeId": "string",
  "routeKey": "string",
  "authorizationType": enum,
  "authorizerId": "string",
  "authorizationScopes": [
    "string"
  ],
  "apiKeyRequired": boolean,
  "requestParameters": {
  },
  "requestModels": {
  },
  "modelSelectionExpression": "string",
  "target": "string",
  "operationName": "string",
  "routeResponseSelectionExpression": "string",
  "apiGatewayManaged": boolean
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

### AuthorizationType

The authorization type. For WebSocket APIs, valid values are `NONE` for
open access, `AWS_IAM` for using AWS IAM permissions, and
`CUSTOM` for using a Lambda authorizer. For HTTP APIs, valid values are `NONE` for open access, `JWT` for using JSON Web Tokens, `AWS_IAM` for using AWS IAM
permissions, and `CUSTOM` for using a Lambda authorizer.

- `NONE`

- `AWS_IAM`

- `CUSTOM`

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

### ParameterConstraints

Validation constraints imposed on parameters of a request (path, query string, headers).

PropertyTypeRequiredDescription`required`

boolean

False

Whether or not the parameter is required.

### Route

Represents a route.

PropertyTypeRequiredDescription`apiGatewayManaged`

boolean

False

Specifies whether a route is managed by API Gateway. If you created an API using
quick create, the `$default` route is managed by API Gateway. You can't modify the `$default` route key.

`apiKeyRequired`

boolean

False

Specifies whether an API key is required for this route. Supported only for WebSocket APIs.

`authorizationScopes`

Array of type string

False

A list of authorization scopes configured on a route. The scopes are used with a `JWT` authorizer to authorize the method invocation. The authorization works by matching the route scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any route scope matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the route scope is configured, the client must provide an access token instead of an identity token for authorization purposes.

`authorizationType`

[AuthorizationType](#apis-apiid-routes-routeid-model-authorizationtype)

False

The authorization type for the route. For WebSocket APIs, valid values are
`NONE` for open access, `AWS_IAM` for using AWS IAM
permissions, and `CUSTOM` for using a Lambda
authorizer. For HTTP APIs, valid values are `NONE` for open access, `JWT` for using JSON Web Tokens, `AWS_IAM` for using AWS IAM
permissions, and `CUSTOM` for using a Lambda authorizer.

`authorizerId`

string

False

The identifier of the `Authorizer` resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.

`modelSelectionExpression`

string

False

The model selection expression for the route. Supported only for WebSocket APIs.

`operationName`

string

False

The operation name for the route.

`requestModels`

[RouteModels](#apis-apiid-routes-routeid-model-routemodels)

False

The request models for the route. Supported only for WebSocket APIs.

`requestParameters`

[RouteParameters](#apis-apiid-routes-routeid-model-routeparameters)

False

The request parameters for the route. Supported only for WebSocket APIs.

`routeId`

string

False

The route ID.

`routeKey`

string

True

The route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.

`routeResponseSelectionExpression`

string

False

The route response selection expression for the route. Supported only for WebSocket APIs.

`target`

string

False

The target for the route.

### RouteModels

The route models.

PropertyTypeRequiredDescription

`*`

string

False

### RouteParameters

The route parameters.

PropertyTypeRequiredDescription

`*`

object

False

### UpdateRouteInput

Represents the input parameters for an `UpdateRoute` request.

PropertyTypeRequiredDescription`apiKeyRequired`

boolean

False

Specifies whether an API key is required for the route. Supported only for WebSocket APIs.

`authorizationScopes`

Array of type string

False

The authorization scopes supported by this route.

`authorizationType`

[AuthorizationType](#apis-apiid-routes-routeid-model-authorizationtype)

False

The authorization type for the route. For WebSocket APIs, valid values are
`NONE` for open access, `AWS_IAM` for using AWS IAM
permissions, and `CUSTOM` for using a Lambda
authorizer. For HTTP APIs, valid values are `NONE` for open access, `JWT` for using JSON Web Tokens, `AWS_IAM` for using AWS IAM
permissions, and `CUSTOM` for using a Lambda authorizer.

`authorizerId`

string

False

The identifier of the `Authorizer` resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.

`modelSelectionExpression`

string

False

The model selection expression for the route. Supported only for WebSocket APIs.

`operationName`

string

False

The operation name for the route.

`requestModels`

[RouteModels](#apis-apiid-routes-routeid-model-routemodels)

False

The request models for the route. Supported only for WebSocket APIs.

`requestParameters`

[RouteParameters](#apis-apiid-routes-routeid-model-routeparameters)

False

The request parameters for the route. Supported only for WebSocket APIs.

`routeKey`

string

False

The route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.

`routeResponseSelectionExpression`

string

False

The route response selection expression for the route. Supported only for WebSocket APIs.

`target`

string

False

The target for the route.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetRoute

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/getroute.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/getroute.md)

### DeleteRoute

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/deleteroute.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/deleteroute.md)

### UpdateRoute

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/updateroute.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/updateroute.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Publish

RouteRequestParameter

All content copied from https://docs.aws.amazon.com/.
