# x-amazon-apigateway-gateway-responses object

Defines the gateway responses for an API as a string-to- [GatewayResponse](../api/api-gatewayresponse.md) map of key-value pairs. The extension applies to the root-level OpenAPI structure.

Property nameTypeDescription`responseType`[x-amazon-apigateway-gateway-responses.gatewayResponse](api-gateway-swagger-extensions-gateway-responses-gatewayresponse.md)

A `GatewayResponse` for the specified `responseType`.

## x-amazon-apigateway-gateway-responses example

The following API Gateway extension to OpenAPI example defines a [GatewayResponses](../api/api-getgatewayresponses.md) map that
contains two [GatewayResponse](../api/api-gatewayresponse.md) instances—one for the `DEFAULT_4XX`
type and another for the `INVALID_API_KEY` type.

```nohighlight

{
  "x-amazon-apigateway-gateway-responses": {
    "DEFAULT_4XX": {
      "responseParameters": {
        "gatewayresponse.header.Access-Control-Allow-Origin": "'domain.com'"
      },
      "responseTemplates": {
        "application/json": "{\"message\": test 4xx b }"
      }
    },
    "INVALID_API_KEY": {
      "statusCode": "429",
      "responseTemplates": {
        "application/json": "{\"message\": test forbidden }"
      }
    }
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

x-amazon-apigateway-endpoint-configuration

x-amazon-apigateway-gateway-responses.gatewayResponse

All content copied from https://docs.aws.amazon.com/.
