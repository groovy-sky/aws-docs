# Apis

Represents your APIs as a collection. A collection offers a paginated view of your APIs.

## URI

`/v2/apis`

## HTTP methods

### GET

**Operation ID:** `GetApis`

Gets a collection of `Api` resources.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The next page of elements from this collection. Not valid for the last element of the collection.

`maxResults`StringFalse

The maximum number of elements to be returned for this resource.

ResponsesStatus codeResponse modelDescription`200``Apis`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### POST

**Operation ID:** `CreateApi`

Creates an `Api` resource.

ResponsesStatus codeResponse modelDescription`201``Api`

The request has succeeded and has resulted in the creation of a resource.

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`409``ConflictException`

The resource already exists.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### PUT

**Operation ID:** `ImportApi`

Imports an API.

Query parametersNameTypeRequiredDescription`failOnWarnings`StringFalse

Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.

`basepath`StringFalse

Specifies how to interpret the base path of the API during import. Valid values are `ignore`, `prepend`, and
`split`. The default value is `ignore`. To learn more, see [Set the OpenAPI basePath\
Property](../../../apigateway/latest/developerguide/api-gateway-import-api-basepath.md). Supported only for HTTP APIs.

ResponsesStatus codeResponse modelDescription`201``Api`

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
  "name": "string",
  "description": "string",
  "version": "string",
  "protocolType": enum,
  "ipAddressType": enum,
  "routeSelectionExpression": "string",
  "apiKeySelectionExpression": "string",
  "disableSchemaValidation": boolean,
  "tags": {
  },
  "target": "string",
  "credentialsArn": "string",
  "routeKey": "string",
  "corsConfiguration": {
    "allowOrigins": [
      "string"
    ],
    "allowCredentials": boolean,
    "exposeHeaders": [
      "string"
    ],
    "maxAge": integer,
    "allowMethods": [
      "string"
    ],
    "allowHeaders": [
      "string"
    ]
  },
  "disableExecuteApiEndpoint": boolean
}
```

```json

{
  "body": "string"
}
```

### Response bodies

```json

{
  "items": [
    {
      "apiId": "string",
      "name": "string",
      "description": "string",
      "version": "string",
      "protocolType": enum,
      "ipAddressType": enum,
      "routeSelectionExpression": "string",
      "apiKeySelectionExpression": "string",
      "disableSchemaValidation": boolean,
      "warnings": [
        "string"
      ],
      "importInfo": [
        "string"
      ],
      "apiEndpoint": "string",
      "apiGatewayManaged": boolean,
      "createdDate": "string",
      "tags": {
      },
      "corsConfiguration": {
        "allowOrigins": [
          "string"
        ],
        "allowCredentials": boolean,
        "exposeHeaders": [
          "string"
        ],
        "maxAge": integer,
        "allowMethods": [
          "string"
        ],
        "allowHeaders": [
          "string"
        ]
      },
      "disableExecuteApiEndpoint": boolean
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "apiId": "string",
  "name": "string",
  "description": "string",
  "version": "string",
  "protocolType": enum,
  "ipAddressType": enum,
  "routeSelectionExpression": "string",
  "apiKeySelectionExpression": "string",
  "disableSchemaValidation": boolean,
  "warnings": [
    "string"
  ],
  "importInfo": [
    "string"
  ],
  "apiEndpoint": "string",
  "apiGatewayManaged": boolean,
  "createdDate": "string",
  "tags": {
  },
  "corsConfiguration": {
    "allowOrigins": [
      "string"
    ],
    "allowCredentials": boolean,
    "exposeHeaders": [
      "string"
    ],
    "maxAge": integer,
    "allowMethods": [
      "string"
    ],
    "allowHeaders": [
      "string"
    ]
  },
  "disableExecuteApiEndpoint": boolean
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

### Api

Represents an API.

PropertyTypeRequiredDescription`apiEndpoint`

string

Format: uri

False

The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com. The stage name is typically appended to this URI to form a complete path to a deployed API stage.

`apiGatewayManaged`

boolean

False

Specifies whether an API is managed by API Gateway. You can't update or delete a managed API by using API Gateway. A managed API can be deleted only through the tooling or service that created it.

`apiId`

string

False

The API ID.

`apiKeySelectionExpression`

string

False

An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](../../../apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.md#apigateway-websocket-api-apikey-selection-expressions).

`corsConfiguration`

[Cors](#apis-model-cors)

False

A CORS configuration. Supported only for HTTP APIs.

`createdDate`

string

Format: date-time

False

The timestamp when the API was created.

`description`

string

False

The description of the API.

`disableExecuteApiEndpoint`

boolean

False

Specifies whether clients can invoke your API by using the default
`execute-api` endpoint. By default, clients can invoke your API
with the default https://{api\_id}.execute-api.{region}.amazonaws.com endpoint.
To require that clients use a custom domain name to invoke your API, disable the
default endpoint.

`disableSchemaValidation`

boolean

False

Avoid validating models when creating a deployment. Supported only for WebSocket APIs.

`importInfo`

Array of type string

False

The validation information during API import. This may include particular
properties of your OpenAPI definition which are ignored during import. Supported only
for HTTP APIs.

`ipAddressType`

[IpAddressType](#apis-model-ipaddresstype)

False

The IP address types that can invoke the API.

`name`

string

True

The name of the API.

`protocolType`

[ProtocolType](#apis-model-protocoltype)

True

The API protocol.

`routeSelectionExpression`

string

True

The route selection expression for the API. For HTTP APIs, the `routeSelectionExpression` must be `${request.method} ${request.path}`. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.

`tags`

[Tags](#apis-model-tags)

False

A collection of tags associated with the API.

`version`

string

False

A version identifier for the API.

`warnings`

Array of type string

False

The warning messages reported when `failonwarnings` is turned on during API import.

### Apis

Represents a collection of APIs.

PropertyTypeRequiredDescription`items`

Array of type Api

False

The elements from this collection.

`nextToken`

string

False

The next page of elements from this collection. Not valid for the last element of the collection.

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

### Cors

Represents a CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](../../../apigateway/latest/developerguide/http-api-cors.md) for more information.

PropertyTypeRequiredDescription`allowCredentials`

boolean

False

Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.

`allowHeaders`

Array of type string

False

Represents a collection of allowed headers. Supported only for HTTP APIs.

`allowMethods`

Array of type string

False

Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.

`allowOrigins`

Array of type string

False

Represents a collection of allowed origins. Supported only for HTTP APIs.

`exposeHeaders`

Array of type string

False

Represents a collection of exposed headers. Supported only for HTTP APIs.

`maxAge`

integer

False

The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.

### CreateApiInput

Represents the input parameters for a `CreateApi` request.

PropertyTypeRequiredDescription`apiKeySelectionExpression`

string

False

An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](../../../apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.md#apigateway-websocket-api-apikey-selection-expressions).

`corsConfiguration`

[Cors](#apis-model-cors)

False

A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](../../../apigateway/latest/developerguide/http-api-cors.md) for more information.

`credentialsArn`

string

False

This property is part of quick create. It specifies the credentials required for
the integration, if any. For a Lambda integration, three options are available. To
specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name
(ARN). To require that the caller's identity be passed through from the request,
specify `arn:aws:iam::*:user/*`. To use resource-based permissions on
supported AWS services, specify `null`. Currently, this property is not used for HTTP
integrations. Supported only for HTTP APIs.

`description`

string

False

The description of the API.

`disableExecuteApiEndpoint`

boolean

False

Specifies whether clients can invoke your API by using the default
`execute-api` endpoint. By default, clients can invoke your API
with the default https://{api\_id}.execute-api.{region}.amazonaws.com endpoint.
To require that clients use a custom domain name to invoke your API, disable the
default endpoint.

`disableSchemaValidation`

boolean

False

Avoid validating models when creating a deployment. Supported only for WebSocket APIs.

`ipAddressType`

[IpAddressType](#apis-model-ipaddresstype)

False

The IP address types that can invoke the API. Use `ipv4` to allow only IPv4 addresses to invoke your API, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your API.

`name`

string

True

The name of the API.

`protocolType`

[ProtocolType](#apis-model-protocoltype)

True

The API protocol.

`routeKey`

string

False

This property is part of quick create. If you don't specify a
`routeKey`, a default route of `$default` is created. The
`$default` route acts as a catch-all for any request made to your API,
for a particular stage. The `$default` route key can't be modified. You
can add routes after creating the API, and you can update the route keys of
additional routes. Supported only for HTTP APIs.

`routeSelectionExpression`

string

False

The route selection expression for the API. For HTTP APIs, the `routeSelectionExpression` must be `${request.method} ${request.path}`. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.

`tags`

[Tags](#apis-model-tags)

False

The collection of tags. Each tag element is associated with a given resource.

`target`

string

False

This property is part of quick create. Quick create produces an API with an
integration, a default catch-all route, and a default stage which is configured to
automatically deploy changes. For HTTP integrations, specify a fully qualified URL.
For Lambda integrations, specify a function ARN. The type of the integration will be
HTTP\_PROXY or AWS\_PROXY, respectively. Supported only for HTTP APIs.

`version`

string

False

A version identifier for the API.

### ImportApiInput

Represents the input to `ImportAPI`. Supported only for HTTP APIs.

PropertyTypeRequiredDescription`body`

string

True

The OpenAPI definition. Supported only for HTTP APIs.

### IpAddressType

The IP address types that can invoke your API or domain name.

- `ipv4`

- `dualstack`

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

### ProtocolType

Represents a protocol type.

- `WEBSOCKET`

- `HTTP`

### Tags

Represents a collection of tags associated with the resource.

PropertyTypeRequiredDescription

`*`

string

False

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetApis

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/GetApis)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/GetApis)

### CreateApi

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/CreateApi)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/CreateApi)

### ImportApi

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/ImportApi)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/ImportApi)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ApiMappings

Authorizer
