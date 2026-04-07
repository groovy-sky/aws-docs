# RouteResponses

Represents the collection of responses for a route. Supported only for WebSocket APIs.

## URI

`/v2/apis/apiId/routes/routeId/routeresponses`

## HTTP methods

### GET

**Operation ID:** `GetRouteResponses`

Gets the `RouteResponses` for a `Route`.

Path parametersNameTypeRequiredDescription`routeId`StringTrue

The route ID.

`apiId`StringTrue

The API identifier.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The next page of elements from this collection. Not valid for the last element of the collection.

`maxResults`StringFalse

The maximum number of elements to be returned for this resource.

ResponsesStatus codeResponse modelDescription`200``RouteResponses`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### POST

**Operation ID:** `CreateRouteResponse`

Creates a `RouteResponse` for a `Route`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`routeId`StringTrue

The route ID.

ResponsesStatus codeResponse modelDescription`201``RouteResponse`

The request has succeeded and has resulted in the creation of a resource.

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
  "routeResponseKey": "string",
  "responseParameters": {
  },
  "responseModels": {
  },
  "modelSelectionExpression": "string"
}
```

### Response bodies

```json

{
  "items": [
    {
      "routeResponseId": "string",
      "routeResponseKey": "string",
      "responseParameters": {
      },
      "responseModels": {
      },
      "modelSelectionExpression": "string"
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "routeResponseId": "string",
  "routeResponseKey": "string",
  "responseParameters": {
  },
  "responseModels": {
  },
  "modelSelectionExpression": "string"
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

### CreateRouteResponseInput

Represents the input parameters for an `CreateRouteResponse` request.

PropertyTypeRequiredDescription`modelSelectionExpression`

string

False

The model selection expression for the route response. Supported only for WebSocket APIs.

`responseModels`

[RouteModels](#apis-apiid-routes-routeid-routeresponses-model-routemodels)

False

The response models for the route response.

`responseParameters`

[RouteParameters](#apis-apiid-routes-routeid-routeresponses-model-routeparameters)

False

The route response parameters.

`routeResponseKey`

string

True

The route response key.

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

### RouteResponse

Represents a route response.

PropertyTypeRequiredDescription`modelSelectionExpression`

string

False

Represents the model selection expression of a route response. Supported only for WebSocket APIs.

`responseModels`

[RouteModels](#apis-apiid-routes-routeid-routeresponses-model-routemodels)

False

Represents the response models of a route response.

`responseParameters`

[RouteParameters](#apis-apiid-routes-routeid-routeresponses-model-routeparameters)

False

Represents the response parameters of a route response.

`routeResponseId`

string

False

Represents the identifier of a route response.

`routeResponseKey`

string

True

Represents the route response key of a route response.

### RouteResponses

Represents a collection of route responses.

PropertyTypeRequiredDescription`items`

Array of type RouteResponse

False

The elements from this collection.

`nextToken`

string

False

The next page of elements from this collection. Not valid for the last element of the collection.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetRouteResponses

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/GetRouteResponses)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/GetRouteResponses)

### CreateRouteResponse

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/CreateRouteResponse)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/CreateRouteResponse)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RouteResponse

Routes
