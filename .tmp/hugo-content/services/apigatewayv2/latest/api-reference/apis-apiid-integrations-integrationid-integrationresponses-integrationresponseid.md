# IntegrationResponse

Represents an integration response. Supported only for WebSocket APIs.

## URI

`/v2/apis/apiId/integrations/integrationId/integrationresponses/integrationResponseId`

## HTTP methods

### GET

**Operation ID:** `GetIntegrationResponse`

Gets an `IntegrationResponses`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`integrationResponseId`StringTrue

The integration response ID.

`integrationId`StringTrue

The integration ID.

ResponsesStatus codeResponse modelDescription`200``IntegrationResponse`

Success

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### DELETE

**Operation ID:** `DeleteIntegrationResponse`

Deletes an `IntegrationResponses`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`integrationResponseId`StringTrue

The integration response ID.

`integrationId`StringTrue

The integration ID.

ResponsesStatus codeResponse modelDescription`204`None

The request has succeeded, and there is no additional content to send in the response payload body.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### PATCH

**Operation ID:** `UpdateIntegrationResponse`

Updates an `IntegrationResponses`.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

`integrationResponseId`StringTrue

The integration response ID.

`integrationId`StringTrue

The integration ID.

ResponsesStatus codeResponse modelDescription`200``IntegrationResponse`

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
  "integrationResponseKey": "string",
  "responseParameters": {
  },
  "responseTemplates": {
  },
  "templateSelectionExpression": "string",
  "contentHandlingStrategy": enum
}
```

### Response bodies

```json

{
  "integrationResponseId": "string",
  "integrationResponseKey": "string",
  "responseParameters": {
  },
  "responseTemplates": {
  },
  "templateSelectionExpression": "string",
  "contentHandlingStrategy": enum
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

### ContentHandlingStrategy

Specifies how to handle response payload content type conversions. Supported only for WebSocket APIs.

- `CONVERT_TO_BINARY`

- `CONVERT_TO_TEXT`

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

### IntegrationResponse

Represents an integration response.

PropertyTypeRequiredDescription`contentHandlingStrategy`

[ContentHandlingStrategy](#apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid-model-contenthandlingstrategy)

False

Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`, with the following behaviors:

`CONVERT_TO_BINARY`: Converts a response payload from a Base64-encoded string to the corresponding binary blob.

`CONVERT_TO_TEXT`: Converts a response payload from a binary blob to a Base64-encoded string.

If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.

`integrationResponseId`

string

False

The integration response ID.

`integrationResponseKey`

string

True

The integration response key.

`responseParameters`

[IntegrationParameters](#apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid-model-integrationparameters)

False

A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of method.response.header.{name}, where name is a valid and unique header name. The mapped non-static value must match the pattern of integration.response.header.{name} or integration.response.body.{JSON-expression}, where name is a valid and unique response header name and JSON-expression is a valid JSON expression without the $ prefix.

`responseTemplates`

[TemplateMap](#apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid-model-templatemap)

False

The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.

`templateSelectionExpression`

string

False

The template selection expressions for the integration response.

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

### TemplateMap

A mapping of identifier keys to templates. The value is an actual template script. The key is typically a `SelectionKey` which is chosen based on evaluating a selection expression.

PropertyTypeRequiredDescription

`*`

string

False

### UpdateIntegrationResponseInput

Represents the input parameters for an `UpdateIntegrationResponse` request.

PropertyTypeRequiredDescription`contentHandlingStrategy`

[ContentHandlingStrategy](#apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid-model-contenthandlingstrategy)

False

Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`, with the following behaviors:

`CONVERT_TO_BINARY`: Converts a response payload from a Base64-encoded string to the corresponding binary blob.

`CONVERT_TO_TEXT`: Converts a response payload from a binary blob to a Base64-encoded string.

If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.

`integrationResponseKey`

string

False

The integration response key.

`responseParameters`

[IntegrationParameters](#apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid-model-integrationparameters)

False

A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of `method.response.header.{name}
               `, where name is a valid and unique header name. The mapped non-static value must match the pattern of `integration.response.header.{name}
               ` or `integration.response.body.{JSON-expression}
               `, where `
                  {name}
               ` is a valid and unique response header name and `
                  {JSON-expression}
               ` is a valid JSON expression without the `$` prefix.

`responseTemplates`

[TemplateMap](#apis-apiid-integrations-integrationid-integrationresponses-integrationresponseid-model-templatemap)

False

The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.

`templateSelectionExpression`

string

False

The template selection expression for the integration response. Supported only for WebSocket APIs.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetIntegrationResponse

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/getintegrationresponse.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/getintegrationresponse.md)

### DeleteIntegrationResponse

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/deleteintegrationresponse.md)

### UpdateIntegrationResponse

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/updateintegrationresponse.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/updateintegrationresponse.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Integration

IntegrationResponses

All content copied from https://docs.aws.amazon.com/.
