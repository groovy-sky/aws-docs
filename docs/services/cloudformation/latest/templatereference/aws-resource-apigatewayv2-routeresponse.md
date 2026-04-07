This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::RouteResponse

The `AWS::ApiGatewayV2::RouteResponse` resource creates a route
response for a WebSocket API. For more information, see [Set up Route Responses for a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-route-response.html) in the
_API Gateway Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::RouteResponse",
  "Properties" : {
      "ApiId" : String,
      "ModelSelectionExpression" : String,
      "ResponseModels" : Json,
      "ResponseParameters" : ParameterConstraints,
      "RouteId" : String,
      "RouteResponseKey" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::RouteResponse
Properties:
  ApiId: String
  ModelSelectionExpression: String
  ResponseModels: Json
  ResponseParameters:
    ParameterConstraints
  RouteId: String
  RouteResponseKey: String

```

## Properties

`ApiId`

The API identifier.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelSelectionExpression`

The model selection expression for the route response. Supported only for WebSocket APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseModels`

The response models for the route response.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseParameters`

The route response parameters.

_Required_: No

_Type_: [ParameterConstraints](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-routeresponse-parameterconstraints.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouteId`

The route ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RouteResponseKey`

The route response key.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Route Response resource ID, such as
`abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RouteResponseId`

The route response ID.

## Examples

### Route response creation example

The following example creates a `RouteResponse` resource
for a WebSocket API called `MyApi` that already has an
`integration` called `MyIntegration` and a
`route` called `MyRoute`.

#### JSON

```json

{
    "MyRouteResponse": {
        "Type": "AWS::ApiGatewayV2::RouteResponse",
        "Properties": {
            "RouteId": {
                "Ref": "MyRoute"
            },
            "ApiId": {
                "Ref": "MyApi"
            },
            "RouteResponseKey": "$default"
        }
    }
}
```

#### YAML

```yaml

MyRouteResponse:
  Type: 'AWS::ApiGatewayV2::RouteResponse'
  Properties:
    RouteId: !Ref MyRoute
    ApiId: !Ref MyApi
    RouteResponseKey: $default

```

## See also

- [CreateRouteResponse](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-routes-routeid-routeresponses.html#CreateRouteResponse) in the _Amazon API_
_Gateway Version 2 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGatewayV2::Route

ParameterConstraints
