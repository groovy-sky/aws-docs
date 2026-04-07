# x-amazon-apigateway-gateway-responses.responseParameters object

Defines a string-to-string map of key-value pairs to generate gateway response parameters from the incoming request parameters or using literal strings. Supported only for REST APIs.

Property nameTypeDescription`gatewayresponse.param-position.param-name``string`

`param-position` can be
`header`, `path`, or
`querystring`. For more information, see [Parameter mapping for REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/rest-api-parameter-mapping.html).

## x-amazon-apigateway-gateway-responses.responseParameters example

The following OpenAPI extensions example shows a
[GatewayResponse](../api/api-gatewayresponse.md) response parameter mapping expression to enable CORS support for resources on the `*.example.domain` domains.

```nohighlight

      "responseParameters": {
        "gatewayresponse.header.Access-Control-Allow-Origin": '*.example.domain',
        "gatewayresponse.header.from-request-header" : method.request.header.Accept,
        "gatewayresponse.header.from-request-path" : method.request.path.petId,
        "gatewayresponse.header.from-request-query" : method.request.querystring.qname
      }
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

x-amazon-apigateway-gateway-responses.gatewayResponse

x-amazon-apigateway-gateway-responses.responseTemplates
