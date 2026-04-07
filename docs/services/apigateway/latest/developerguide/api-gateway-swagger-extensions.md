# OpenAPI extensions for API Gateway

The API Gateway extensions support the AWS-specific authorization and API Gateway-specific API
integrations for REST APIs and HTTP APIs. In this section, we describe the
API Gateway extensions to the OpenAPI specification.

###### Tip

To understand how the API Gateway extensions are used in an application, you can use the
API Gateway console to create a REST API or HTTP API and export it to an
OpenAPI definition file. For more information on how to export an API, see [Export a REST API from API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-export-api.html) and [Export HTTP APIs from API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-export.html).

###### Topics

- [x-amazon-apigateway-any-method object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-any-method.html)

- [x-amazon-apigateway-cors object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-cors-configuration.html)

- [x-amazon-apigateway-api-key-source property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-api-key-source.html)

- [x-amazon-apigateway-auth object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-auth.html)

- [x-amazon-apigateway-authorizer object](api-gateway-swagger-extensions-authorizer.md)

- [x-amazon-apigateway-authtype property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-authtype.html)

- [x-amazon-apigateway-binary-media-types property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-binary-media-types.html)

- [x-amazon-apigateway-documentation object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-documentation.html)

- [x-amazon-apigateway-endpoint-access-mode](openapi-extensions-endpoint-access-mode.md)

- [x-amazon-apigateway-endpoint-configuration object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-endpoint-configuration.html)

- [x-amazon-apigateway-gateway-responses object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-gateway-responses.html)

- [x-amazon-apigateway-gateway-responses.gatewayResponse object](api-gateway-swagger-extensions-gateway-responses-gatewayresponse.md)

- [x-amazon-apigateway-gateway-responses.responseParameters object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-gateway-responses.responseParameters.html)

- [x-amazon-apigateway-gateway-responses.responseTemplates object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-gateway-responses.responseTemplates.html)

- [x-amazon-apigateway-importexport-version](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-extensions-importexport-version.html)

- [x-amazon-apigateway-integration object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-integration.html)

- [x-amazon-apigateway-integrations object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-extensions-integrations.html)

- [x-amazon-apigateway-integration.requestTemplates object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-integration-requestTemplates.html)

- [x-amazon-apigateway-integration.requestParameters object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-integration-requestParameters.html)

- [x-amazon-apigateway-integration.responses object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-integration-responses.html)

- [x-amazon-apigateway-integration.response object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-integration-response.html)

- [x-amazon-apigateway-integration.responseTemplates object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-integration-responseTemplates.html)

- [x-amazon-apigateway-integration.responseParameters object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-integration-responseParameters.html)

- [x-amazon-apigateway-integration.tlsConfig object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-extensions-integration-tls-config.html)

- [x-amazon-apigateway-minimum-compression-size](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-openapi-minimum-compression-size.html)

- [x-amazon-apigateway-policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/openapi-extensions-policy.html)

- [x-amazon-apigateway-request-validator property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-request-validator.html)

- [x-amazon-apigateway-request-validators object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-request-validators.html)

- [x-amazon-apigateway-request-validators.requestValidator object](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-request-validators.requestValidator.html)

- [x-amazon-apigateway-security-policy](openapi-extensions-security-policy.md)

- [x-amazon-apigateway-tag-value property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-openapi-extensions-x-amazon-apigateway-tag-value.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

API Gateway ARNs

x-amazon-apigateway-any-method
