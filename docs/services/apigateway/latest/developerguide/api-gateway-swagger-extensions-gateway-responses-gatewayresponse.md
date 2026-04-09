# x-amazon-apigateway-gateway-responses.gatewayResponse object

Defines a gateway response of a given response type, including the status code, any
applicable response parameters, or response templates.

Property nameTypeDescription`responseParameters`[x-amazon-apigateway-gateway-responses.responseParameters](api-gateway-swagger-extensions-gateway-responses-responseparameters.md)

Specifies the [GatewayResponse](../api/api-gatewayresponse.md) parameters, namely the header parameters. The parameter values can take any incoming [request parameter](rest-api-parameter-mapping.md) value or a static custom value.

`responseTemplates`[x-amazon-apigateway-gateway-responses.responseTemplates](api-gateway-swagger-extensions-gateway-responses-responsetemplates.md)

Specifies the mapping templates of the gateway response. The templates are not processed by the VTL engine.

`statusCode``string`

An HTTP status code for the gateway response.

## x-amazon-apigateway-gateway-responses.gatewayResponse example

The following example of the API Gateway extension to OpenAPI defines a [GatewayResponse](../api/api-gatewayresponse.md) to
customize the `INVALID_API_KEY` response to return the status code of
`456`, the incoming request's `api-key` header value, and
a `"Bad api-key"` message.

```nohighlight

    "INVALID_API_KEY": {
      "statusCode": "456",
      "responseParameters": {
        "gatewayresponse.header.api-key": "method.request.header.api-key"
      },
      "responseTemplates": {
        "application/json": "{\"message\": \"Bad api-key\" }"
      }
    }

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

x-amazon-apigateway-gateway-responses

x-amazon-apigateway-gateway-responses.responseParameters

All content copied from https://docs.aws.amazon.com/.
