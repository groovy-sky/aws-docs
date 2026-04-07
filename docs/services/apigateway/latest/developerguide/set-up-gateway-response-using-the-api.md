# Set up a gateway response using the API Gateway REST API

Before customizing a gateway response using the API Gateway REST API, you must have already
created an API and have obtained its identifier. To retrieve the API identifier, you can
follow [restapi:gateway-responses](https://docs.aws.amazon.com/apigateway/latest/api/API_GetGatewayResponses.html) link relation and examine the result.

###### To customize a gateway response using the API Gateway REST API

1. To overwrite an entire [GatewayResponse](../api/api-gatewayresponse.md) instance, call the [gatewayresponse:put](https://docs.aws.amazon.com/apigateway/latest/api/API_PutGatewayResponse.html) action. Specify a desired [responseType](../api/api-gatewayresponse.md#responseType) in the URL path parameter, and supply in the request
    payload the [statusCode](../api/api-gatewayresponse.md#statusCode), [responseParameters](../api/api-gatewayresponse.md#responseParameters), and [responseTemplates](../api/api-gatewayresponse.md#responseTemplates) mappings.

2. To update part of a `GatewayResponse` instance, call the [gatewayresponse:update](https://docs.aws.amazon.com/apigateway/latest/api/API_UpdateGatewayResponse.html) action. Specify a desired
    `responseType` in the URL path parameter, and supply in the
    request payload the individual `GatewayResponse` properties you
    want—for example, the `responseParameters` or the
    `responseTemplates` mapping.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Set up a gateway response for a REST API using the API Gateway console

Set up gateway response customization in OpenAPI
