This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::IntegrationResponse

The `AWS::ApiGatewayV2::IntegrationResponse` resource updates an
integration response for an WebSocket API. For more information, see [Set up WebSocket API Integration Responses in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-integration-responses.html) in the
_API Gateway Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::IntegrationResponse",
  "Properties" : {
      "ApiId" : String,
      "ContentHandlingStrategy" : String,
      "IntegrationId" : String,
      "IntegrationResponseKey" : String,
      "ResponseParameters" : Json,
      "ResponseTemplates" : Json,
      "TemplateSelectionExpression" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::IntegrationResponse
Properties:
  ApiId: String
  ContentHandlingStrategy: String
  IntegrationId: String
  IntegrationResponseKey: String
  ResponseParameters: Json
  ResponseTemplates: Json
  TemplateSelectionExpression: String

```

## Properties

`ApiId`

The API identifier.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContentHandlingStrategy`

Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`, with the following behaviors:

`CONVERT_TO_BINARY`: Converts a response payload from a Base64-encoded string to the corresponding binary blob.

`CONVERT_TO_TEXT`: Converts a response payload from a binary blob to a Base64-encoded string.

If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegrationId`

The integration ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IntegrationResponseKey`

The integration response key.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseParameters`

A key-value map specifying response parameters that are passed to the method
response from the backend. The key is a method response header parameter name and
the mapped value is an integration response header value, a static value enclosed
within a pair of single quotes, or a JSON expression from the integration response
body. The mapping key must match the pattern of
`method.response.header.{name}`, where name is a valid and unique header name. The mapped non-static value
must match the pattern of
`integration.response.header.{name}` or `integration.response.body.{JSON-expression}`, where `
                            {name}
                        ` is a valid and unique response header name and `
                            {JSON-expression}
                        ` is a valid JSON expression without the `$` prefix.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseTemplates`

The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateSelectionExpression`

The template selection expression for the integration response. Supported only for WebSocket APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the integration response resource ID, such as
`abcd123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IntegrationResponseId`

The integration response ID.

## Examples

### Integration response creation example

The following example creates an `IntegrationResponse`
resource for an API named `MyApi` that has an
`integration` resource called
`MyIntegration`.

#### JSON

```json

{
    "IntegrationResponse": {
        "Type": "AWS::ApiGatewayV2::IntegrationResponse",
        "Properties": {
            "IntegrationId": {
                "Ref": "MyIntegration"
            },
            "IntegrationResponseKey": "/400/",
            "ApiId": {
                "Ref": "MyApi"
            }
        }
    }
}
```

#### YAML

```yaml

IntegrationResponse:
  Type: 'AWS::ApiGatewayV2::IntegrationResponse'
  Properties:
    IntegrationId: !Ref MyIntegration
    IntegrationResponseKey: /400/
    ApiId: !Ref MyApi
```

## See also

- [CreateIntegrationResponse](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-integrations-integrationid-integrationresponses.html#CreateIntegrationResponse) in the _Amazon_
_API Gateway Version 2 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TlsConfig

AWS::ApiGatewayV2::Model
