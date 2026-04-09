# Integration

Represents an API integration.

## URI

`/v2/apis/apiId/integrations/integrationId`

## HTTP methods

### GET

**Operation ID:** `GetIntegration`

Gets an `Integration`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`integrationId`StringTrue

The integration ID.

ResponsesStatus codeResponse modelDescription`200``Integration`

Success

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### DELETE

**Operation ID:** `DeleteIntegration`

Deletes an `Integration`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`integrationId`StringTrue

The integration ID.

ResponsesStatus codeResponse modelDescription`204`None

The request has succeeded, and there is no additional content to send in the response payload body.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### PATCH

**Operation ID:** `UpdateIntegration`

Updates an `Integration`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`integrationId`StringTrue

The integration ID.

ResponsesStatus codeResponse modelDescription`200``Integration`

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
  "description": "string",
  "integrationType": enum,
  "integrationSubtype": "string",
  "integrationMethod": "string",
  "integrationUri": "string",
  "tlsConfig": {
    "serverNameToVerify": "string"
  },
  "credentialsArn": "string",
  "connectionType": enum,
  "connectionId": "string",
  "requestParameters": {
  },
  "responseParameters": {
  },
  "requestTemplates": {
  },
  "templateSelectionExpression": "string",
  "passthroughBehavior": enum,
  "contentHandlingStrategy": enum,
  "timeoutInMillis": integer,
  "payloadFormatVersion": "string"
}
```

### Response bodies

```json

{
  "integrationId": "string",
  "description": "string",
  "integrationType": enum,
  "integrationSubtype": "string",
  "integrationMethod": "string",
  "integrationUri": "string",
  "tlsConfig": {
    "serverNameToVerify": "string"
  },
  "credentialsArn": "string",
  "connectionType": enum,
  "connectionId": "string",
  "requestParameters": {
  },
  "responseParameters": {
  },
  "requestTemplates": {
  },
  "templateSelectionExpression": "string",
  "passthroughBehavior": enum,
  "contentHandlingStrategy": enum,
  "timeoutInMillis": integer,
  "integrationResponseSelectionExpression": "string",
  "payloadFormatVersion": "string",
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

### ConnectionType

Represents a connection type.

- `INTERNET`

- `VPC_LINK`

### ContentHandlingStrategy

Specifies how to handle response payload content type conversions. Supported only for WebSocket APIs.

- `CONVERT_TO_BINARY`

- `CONVERT_TO_TEXT`

### Integration

Represents an integration.

PropertyTypeRequiredDescription`apiGatewayManaged`

boolean

False

Specifies whether an integration is managed by API Gateway. If you created an API using
using quick create, the resulting integration is managed by API Gateway. You can update a managed integration, but you can't delete it.

`connectionId`

string

False

The ID of the VPC link for a private integration. Supported only for HTTP APIs.

`connectionType`

[ConnectionType](#apis-apiid-integrations-integrationid-model-connectiontype)

False

The type of the network connection to the integration endpoint. Specify `INTERNET` for connections through the public routable internet or
`VPC_LINK` for private connections between API Gateway and resources in a VPC. The default value is `INTERNET`.

`contentHandlingStrategy`

[ContentHandlingStrategy](#apis-apiid-integrations-integrationid-model-contenthandlingstrategy)

False

Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`, with the following behaviors:

`CONVERT_TO_BINARY`: Converts a response payload from a Base64-encoded string to the corresponding binary blob.

`CONVERT_TO_TEXT`: Converts a response payload from a binary blob to a Base64-encoded string.

If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.

`credentialsArn`

string

False

Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::*:user/*`. To use resource-based permissions on supported AWS services, don't specify this parameter.

`description`

string

False

Represents the description of an integration.

`integrationId`

string

False

Represents the identifier of an integration.

`integrationMethod`

string

False

Specifies the integration's HTTP method type.

`integrationResponseSelectionExpression`

string

False

The integration response selection expression for the integration. Supported only for WebSocket APIs. See [Integration Response Selection Expressions](../../../apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.md#apigateway-websocket-api-integration-response-selection-expressions).

`integrationSubtype`

string

False

Supported only for HTTP API `AWS_PROXY` integrations. Specifies the AWS service action to invoke. To learn more, see [Integration subtype reference](../../../apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.md).

`integrationType`

[IntegrationType](#apis-apiid-integrations-integrationid-model-integrationtype)

False

The integration type of an integration. One of the following:

`AWS`: for integrating the route or method request with an AWS service action, including the Lambda function-invoking action. With the Lambda function-invoking action, this is referred to as the Lambda custom integration. With any other AWS service action, this is known as AWS integration. Supported only for WebSocket APIs.

`AWS_PROXY`: for integrating the route or method request with a Lambda function or other AWS service action. This integration is also referred to as a Lambda proxy integration.

`HTTP`: for integrating the route or method request with an HTTP endpoint. This integration is also referred to as the HTTP custom integration. Supported only for WebSocket APIs.

`HTTP_PROXY`: for integrating the route or method request with an HTTP endpoint, with the client request passed through as-is. This is also referred to as HTTP proxy integration.

`MOCK`: for integrating the route or method request with API Gateway as a "loopback" endpoint without invoking any backend. Supported only for WebSocket APIs.

`integrationUri`

string

False

For a Lambda integration, specify the URI of a Lambda function.

For an HTTP integration, specify a fully-qualified URL.

For an HTTP API private integration, specify the ARN of an Application Load
Balancer listener, Network Load Balancer listener, or AWS Cloud Map service. If
you specify the ARN of an AWS Cloud Map service, API Gateway uses
`DiscoverInstances` to identify resources. You can use query
parameters to target specific resources. To learn more, see [DiscoverInstances](../../../cloud-map/latest/api/api-discoverinstances.md). For private integrations, all resources must be
owned by the same AWS account.

`passthroughBehavior`

[PassthroughBehavior](#apis-apiid-integrations-integrationid-model-passthroughbehavior)

False

Specifies the pass-through behavior for incoming requests based on the `Content-Type` header in the request, and the available mapping templates specified as the `requestTemplates` property on the `Integration` resource. There are three valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, and `NEVER`. Supported only for WebSocket APIs.

`WHEN_NO_MATCH` passes the request body for unmapped content types through to the integration backend without transformation.

`NEVER` rejects unmapped content types with an `HTTP 415 Unsupported Media Type` response.

`WHEN_NO_TEMPLATES` allows pass-through when the integration has no content types mapped to templates. However, if there is at least one content type defined, unmapped content types will be rejected with the same `HTTP 415 Unsupported Media Type` response.

`payloadFormatVersion`

string

False

Specifies the format of the payload sent to an integration. Required for HTTP
APIs. For HTTP APIs, supported values for Lambda proxy integrations are
`1.0` and `2.0`. For all other integrations,
`1.0` is the only supported value. To learn more, see
[Working with AWS Lambda proxy integrations for HTTP APIs](../../../apigateway/latest/developerguide/http-api-develop-integrations-lambda.md).

`requestParameters`

[IntegrationParameters](#apis-apiid-integrations-integrationid-model-integrationparameters)

False

For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend. The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the backend. The method request parameter value must match the pattern of `method.request.{location}.{name}
          `, where `
            {location}
          ` is `querystring`, `path`, or `header`; and `
            {name}
          ` must be a valid and unique method request parameter name.

For HTTP API integrations with a specified `integrationSubtype`, request parameters are a key-value map specifying parameters that are passed to `AWS_PROXY` integrations. You can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Working with AWS service integrations for HTTP APIs](../../../apigateway/latest/developerguide/http-api-develop-integrations-aws-services.md).

For HTTP API integrations without a specified `integrationSubtype` request parameters are a key-value map specifying how to transform HTTP requests before sending them to backend integrations. The key should follow the pattern `<action>:<header|querystring|path>.<location>`. The action can be `append`, `overwrite` or ` remove`. For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

`requestTemplates`

[TemplateMap](#apis-apiid-integrations-integrationid-model-templatemap)

False

Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value. Supported only for WebSocket APIs.

`responseParameters`

[ResponseParameters](#apis-apiid-integrations-integrationid-model-responseparameters)

False

Supported only for HTTP APIs. You use response parameters to transform the HTTP response from a backend
integration before returning the response to clients. Specify a key-value map from a selection key to response
parameters. The selection key must be a valid HTTP status code within the range of 200-599. Response parameters are a key-value map. The key
must match the pattern `<action>:<header>.<location>` or
`overwrite.statuscode`. The action can be `append`, `overwrite` or
`remove`. The value can be a static value, or map to response data, stage variables, or context
variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

`templateSelectionExpression`

string

False

The template selection expression for the integration. Supported only for WebSocket APIs.

`timeoutInMillis`

integer

False

Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and
between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29
seconds for WebSocket APIs and 30 seconds for HTTP APIs.

`tlsConfig`

[TlsConfig](#apis-apiid-integrations-integrationid-model-tlsconfig)

False

The TLS configuration for a private integration. If you specify a TLS configuration, private integration traffic uses the HTTPS
protocol. Supported only for HTTP APIs.

### IntegrationParameters

For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend. The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the backend. The method request parameter value must match the pattern of `method.request.{location}.{name}
          `, where `
            {location}
          ` is `querystring`, `path`, or `header`; and `
            {name}
          ` must be a valid and unique method request parameter name.

For HTTP API integrations with a specified `integrationSubtype`, request parameters are a key-value map specifying parameters that are passed to `AWS_PROXY` integrations. You can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Working with AWS service integrations for HTTP APIs](../../../apigateway/latest/developerguide/http-api-develop-integrations-aws-services.md).

For HTTP API integrations without a specified `integrationSubtype` request parameters are a key-value map specifying how to transform HTTP requests before sending them to the backend. The key should follow the pattern <action>:<header\|querystring\|path>.<location> where action can be `append`, `overwrite` or ` remove`. For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

PropertyTypeRequiredDescription

`*`

string

False

### IntegrationType

Represents an API method integration type.

- `AWS`

- `HTTP`

- `MOCK`

- `HTTP_PROXY`

- `AWS_PROXY`

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

### PassthroughBehavior

Represents passthrough behavior for an integration response. Supported only for WebSocket APIs.

- `WHEN_NO_MATCH`

- `NEVER`

- `WHEN_NO_TEMPLATES`

### ResponseParameters

Supported only for HTTP APIs. You use response parameters to transform the HTTP response from a backend
integration before returning the response to clients.

PropertyTypeRequiredDescription

`*`

object

False

### TemplateMap

A mapping of identifier keys to templates. The value is an actual template script. The key is typically a `SelectionKey` which is chosen based on evaluating a selection expression.

PropertyTypeRequiredDescription

`*`

string

False

### TlsConfig

The TLS configuration for a private integration. If you specify a TLS configuration, private integration traffic uses the HTTPS
protocol. Supported only for HTTP APIs.

PropertyTypeRequiredDescription`serverNameToVerify`

string

False

If you specify a server name, API Gateway uses it to verify the hostname on
the integration's certificate. The server name is also included in the TLS
handshake to support Server Name Indication (SNI) or virtual hosting.

### TlsConfigInput

The TLS configuration for a private integration. If you specify a TLS configuration, private integration traffic uses the HTTPS
protocol. Supported only for HTTP APIs.

PropertyTypeRequiredDescription`serverNameToVerify`

string

False

If you specify a server name, API Gateway uses it to verify the hostname on
the integration's certificate. The server name is also included in the TLS
handshake to support Server Name Indication (SNI) or virtual hosting.

### UpdateIntegrationInput

Represents the input parameters for an `UpdateIntegration` request.

PropertyTypeRequiredDescription`connectionId`

string

False

The ID of the VPC link for a private integration. Supported only for HTTP APIs.

`connectionType`

[ConnectionType](#apis-apiid-integrations-integrationid-model-connectiontype)

False

The type of the network connection to the integration endpoint. Specify `INTERNET` for connections through the public routable internet or
`VPC_LINK` for private connections between API Gateway and resources in a VPC. The default value is `INTERNET`.

`contentHandlingStrategy`

[ContentHandlingStrategy](#apis-apiid-integrations-integrationid-model-contenthandlingstrategy)

False

Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`, with the following behaviors:

`CONVERT_TO_BINARY`: Converts a response payload from a Base64-encoded string to the corresponding binary blob.

`CONVERT_TO_TEXT`: Converts a response payload from a binary blob to a Base64-encoded string.

If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.

`credentialsArn`

string

False

Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::*:user/*`. To use resource-based permissions on supported AWS services, don't specify this parameter.

`description`

string

False

The description of the integration

`integrationMethod`

string

False

Specifies the integration's HTTP method type. For WebSocket APIs, if you use a Lambda integration, you must set the integration method to `POST`.

`integrationSubtype`

string

False

Supported only for HTTP API `AWS_PROXY` integrations. Specifies the AWS service action to invoke. To learn more, see [Integration subtype reference](../../../apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.md).

`integrationType`

[IntegrationType](#apis-apiid-integrations-integrationid-model-integrationtype)

False

The integration type of an integration. One of the following:

`AWS`: for integrating the route or method request with an AWS service action, including the Lambda function-invoking action. With the Lambda function-invoking action, this is referred to as the Lambda custom integration. With any other AWS service action, this is known as AWS integration. Supported only for WebSocket APIs.

`AWS_PROXY`: for integrating the route or method request with a Lambda function or other AWS service action. This integration is also referred to as a Lambda proxy integration.

`HTTP`: for integrating the route or method request with an HTTP endpoint. This integration is also referred to as the HTTP custom integration. Supported only for WebSocket APIs.

`HTTP_PROXY`: for integrating the route or method request with an
HTTP endpoint, with the
client request passed through as-is. This is also referred to as HTTP proxy
integration. For HTTP API private integrations, use an `HTTP_PROXY`
integration.

`MOCK`: for integrating the route or method request with API Gateway as a "loopback" endpoint without invoking any backend. Supported only for WebSocket APIs.

`integrationUri`

string

False

For a Lambda integration, specify the URI of a Lambda function.

For an HTTP integration, specify a fully-qualified URL.

For an HTTP API private integration, specify the ARN of an Application Load
Balancer listener, Network Load Balancer listener, or AWS Cloud Map service. If
you specify the ARN of an AWS Cloud Map service, API Gateway uses
`DiscoverInstances` to identify resources. You can use query
parameters to target specific resources. To learn more, see [DiscoverInstances](../../../cloud-map/latest/api/api-discoverinstances.md). For private integrations, all resources must be
owned by the same AWS account.

`passthroughBehavior`

[PassthroughBehavior](#apis-apiid-integrations-integrationid-model-passthroughbehavior)

False

Specifies the pass-through behavior for incoming requests based on the `Content-Type` header in the request, and the available mapping templates specified as the `requestTemplates` property on the `Integration` resource. There are three valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, and `NEVER`. Supported only for WebSocket APIs.

`WHEN_NO_MATCH` passes the request body for unmapped content types through to the integration backend without transformation.

`NEVER` rejects unmapped content types with an `HTTP 415 Unsupported Media Type` response.

`WHEN_NO_TEMPLATES` allows pass-through when the integration has no content types mapped to templates. However, if there is at least one content type defined, unmapped content types will be rejected with the same `HTTP 415 Unsupported Media Type` response.

`payloadFormatVersion`

string

False

Specifies the format of the payload sent to an integration. Required for HTTP
APIs. For HTTP APIs, supported values for Lambda proxy integrations are
`1.0` and `2.0`. For all other integrations,
`1.0` is the only supported value. To learn more, see
[Working with AWS Lambda proxy integrations for HTTP APIs](../../../apigateway/latest/developerguide/http-api-develop-integrations-lambda.md).

`requestParameters`

[IntegrationParameters](#apis-apiid-integrations-integrationid-model-integrationparameters)

False

For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend. The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the backend. The method request parameter value must match the pattern of `method.request.{location}.{name}
          `, where `
            {location}
          ` is `querystring`, `path`, or `header`; and `
            {name}
          ` must be a valid and unique method request parameter name.

For HTTP API integrations with a specified `integrationSubtype`, request parameters are a key-value map specifying parameters that are passed to `AWS_PROXY` integrations. You can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Working with AWS service integrations for HTTP APIs](../../../apigateway/latest/developerguide/http-api-develop-integrations-aws-services.md).

For HTTP API integrations without a specified `integrationSubtype` request parameters are a key-value map specifying how to transform HTTP requests before sending them to the backend. The key should follow the pattern <action>:<header\|querystring\|path>.<location> where action can be `append`, `overwrite` or ` remove`. For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

`requestTemplates`

[TemplateMap](#apis-apiid-integrations-integrationid-model-templatemap)

False

Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value. Supported only for WebSocket APIs.

`responseParameters`

[ResponseParameters](#apis-apiid-integrations-integrationid-model-responseparameters)

False

Supported only for HTTP APIs. You use response parameters to transform the HTTP response from a backend
integration before returning the response to clients. Specify a key-value map from a selection key to response
parameters. The selection key must be a valid HTTP status code within the range of 200-599. Response parameters are a key-value map. The key
must match the pattern `<action>:<header>.<location>` or
`overwrite.statuscode`. The action can be `append`, `overwrite` or
`remove`. The value can be a static value, or map to response data, stage variables, or context
variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

`templateSelectionExpression`

string

False

The template selection expression for the integration.

`timeoutInMillis`

integer

False

Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and
between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29
seconds for WebSocket APIs and 30 seconds for HTTP APIs.

`tlsConfig`

[TlsConfigInput](#apis-apiid-integrations-integrationid-model-tlsconfiginput)

False

The TLS configuration for a private integration. If you specify a TLS configuration, private integration traffic uses the HTTPS
protocol. Supported only for HTTP APIs.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetIntegration

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/getintegration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/getintegration.md)

### DeleteIntegration

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/deleteintegration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/deleteintegration.md)

### UpdateIntegration

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/updateintegration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/updateintegration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExportedAPI

IntegrationResponse

All content copied from https://docs.aws.amazon.com/.
