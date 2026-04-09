# Set up a gateway response using the API Gateway REST API

Before customizing a gateway response using the API Gateway REST API, you must have already
created an API and have obtained its identifier. To retrieve the API identifier, you can
follow [restapi:gateway-responses](../api/api-getgatewayresponses.md) link relation and examine the result.

###### To customize a gateway response using the API Gateway REST API

1. To overwrite an entire [GatewayResponse](../api/api-gatewayresponse.md) instance, call the [gatewayresponse:put](../api/api-putgatewayresponse.md) action. Specify a desired [responseType](../api/api-gatewayresponse.md#responseType) in the URL path parameter, and supply in the request
    payload the [statusCode](../api/api-gatewayresponse.md#statusCode), [responseParameters](../api/api-gatewayresponse.md#responseParameters), and [responseTemplates](../api/api-gatewayresponse.md#responseTemplates) mappings.

2. To update part of a `GatewayResponse` instance, call the [gatewayresponse:update](../api/api-updategatewayresponse.md) action. Specify a desired
    `responseType` in the URL path parameter, and supply in the
    request payload the individual `GatewayResponse` properties you
    want—for example, the `responseParameters` or the
    `responseTemplates` mapping.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Set up a gateway response for a REST API using the API Gateway console

Set up gateway response customization in OpenAPI

All content copied from https://docs.aws.amazon.com/.
