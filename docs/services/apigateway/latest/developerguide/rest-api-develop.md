# Develop REST APIs in API Gateway

In Amazon API Gateway, you build a REST API as a collection of programmable entities known as
API Gateway [resources](../api/api-resource.md). For example, you use a [RestApi](../api/api-restapi.md) resource to represent
an API that can contain a collection of [Resource](../api/api-resource.md) entities.

Each `Resource` entity can have one or more [Method](../api/api-method.md)
resources. A `Method` is an incoming request submitted by the client and can contain the following
request parameters: a path parameter, a header, or a query string parameter. In addition, depending on the HTTP
method, the request can contain a body. Your method defines how the client accesses the exposed
`Resource`. To integrate the `Method` with a backend endpoint, also known as the integration
endpoint, you create an [Integration](../api/api-integration.md) resource. This forwards
the incoming request to a specified integration endpoint URI. If necessary, you can transform request parameters or
the request body to meet the backend requirements, or you can create a proxy integration, where API Gateway sends the
entire request in a standardized format to the integration endpoint URI and then directly
sends the response to the client.

For responses, you can create
a [MethodResponse](../api/api-methodresponse.md) resource to represent a
response received by the client and you create an [IntegrationResponse](../api/api-integrationresponse.md) resource to represent the response that is returned by the
backend. Use an integration response to transform the backend response data before returning the data
to the client or to pass the backend response as-is to the client.

## Example resource for a REST API

The following diagram shows how API Gateway implements this request/response model for an HTTP
proxy and an HTTP non-proxy integration for the `GET /pets` resource. The client sends the
`x-version:beta` header to API Gateway and API Gateway sends the `204` status code to the client.

In the non-proxy integration, API Gateway performs data transformations to meet the backend requirements, by
modifying the integration request and integration response. In a non-proxy integration, you can access the body in
the method request but you transform it in the integration request. When the integration endpoint returns a
response with a body, you access and transform it in the integration response. You can't modify the body in the method response.

In the proxy integration, the integration endpoint modifies the request and response. API Gateway doesn't modify the
integration request or integration response, and sends the incoming request to the backend as-is.

Regardless of the integration type, the client sent a request to API Gateway and API Gateway responded
synchronously.

Non-proxy integration

![Diagram of API Gateway non-proxy integration](https://docs.aws.amazon.com/images/apigateway/latest/developerguide/images/develop-non-proxy.png)

Proxy integration

![Diagram of API Gateway proxy integration](https://docs.aws.amazon.com/images/apigateway/latest/developerguide/images/develop-proxy.png)

The following example execution logs show what API Gateway would log in the previous example. For clarity, some values and initial logs have been removed:

Non-proxy integration

```nohighlight

Wed Feb 12 23:56:44 UTC 2025 : Starting execution for request: abcd-1234-5678
Wed Feb 12 23:56:44 UTC 2025 : HTTP Method: GET, Resource Path: /pets
Wed Feb 12 23:56:44 UTC 2025 : Method request path: {}
Wed Feb 12 23:56:44 UTC 2025 : Method request query string: {}
Wed Feb 12 23:56:44 UTC 2025 : Method request headers: {x-version=beta}
Wed Feb 12 23:56:44 UTC 2025 : Method request body before transformations:
Wed Feb 12 23:56:44 UTC 2025 : Endpoint request URI: http://petstore-demo-endpoint.execute-api.com/petstore/pets
Wed Feb 12 23:56:44 UTC 2025 : Endpoint request headers: {app-version=beta}
Wed Feb 12 23:56:44 UTC 2025 : Endpoint request body after transformations:
Wed Feb 12 23:56:44 UTC 2025 : Sending request to http://petstore-demo-endpoint.execute-api.com/petstore/pets
Wed Feb 12 23:56:45 UTC 2025 : Received response. Status: 200, Integration latency: 123 ms
Wed Feb 12 23:56:45 UTC 2025 : Endpoint response headers: {Date=Wed, 12 Feb 2025 23:56:45 GMT}
Wed Feb 12 23:56:45 UTC 2025 : Endpoint response body before transformations:
Wed Feb 12 23:56:45 UTC 2025 : Method response body after transformations: (null)
Wed Feb 12 23:56:45 UTC 2025 : Method response headers: {X-Amzn-Trace-Id=Root=1-abcd-12345}
Wed Feb 12 23:56:45 UTC 2025 : Successfully completed execution
Wed Feb 12 23:56:45 UTC 2025 : Method completed with status: 204
```

Proxy integration

```nohighlight

Wed Feb 12 23:59:42 UTC 2025 : Starting execution for request: abcd-1234-5678
Wed Feb 12 23:59:42 UTC 2025 : HTTP Method: GET, Resource Path: /pets
Wed Feb 12 23:59:42 UTC 2025 : Method request path: {}
Wed Feb 12 23:59:42 UTC 2025 : Method request query string: {}
Wed Feb 12 23:59:42 UTC 2025 : Method request headers: {x-version=beta}
Wed Feb 12 23:59:42 UTC 2025 : Method request body before transformations:
Wed Feb 12 23:59:42 UTC 2025 : Endpoint request URI: http://petstore-demo-endpoint.execute-api.com/petstore/pets
Wed Feb 12 23:59:42 UTC 2025 : Endpoint request headers: { x-version=beta}
Wed Feb 12 23:59:42 UTC 2025 : Endpoint request body after transformations:
Wed Feb 12 23:59:42 UTC 2025 : Sending request to http://petstore-demo-endpoint.execute-api.com/petstore/pets
Wed Feb 12 23:59:43 UTC 2025 : Received response. Status: 204, Integration latency: 123 ms
Wed Feb 12 23:59:43 UTC 2025 : Endpoint response headers: {Date=Wed, 12 Feb 2025 23:59:43 GMT}
Wed Feb 12 23:59:43 UTC 2025 : Endpoint response body before transformations:
Wed Feb 12 23:59:43 UTC 2025 : Method response body after transformations:
Wed Feb 12 23:59:43 UTC 2025 : Method response headers: {Date=Wed, 12 Feb 2025 23:59:43 GMT}
Wed Feb 12 23:59:43 UTC 2025 : Successfully completed execution
Wed Feb 12 23:59:43 UTC 2025 : Method completed with status: 204
```

To import a similar API and test it in the AWS Management Console, see the [example API](api-gateway-create-api-from-example.md).

## Additional REST API features for development

API Gateway supports additional features for the development of your REST API. For example, to help your customers
understand your API, you can provide documentation for the API. To enable this, add a
[DocumentationPart](https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationPart.html) resource for a supported API
entity.

To control how clients call an API, use [IAM\
permissions](permissions.md), a [Lambda\
authorizer](apigateway-use-lambda-authorizer.md), or an [Amazon Cognito user\
pool](apigateway-integrate-with-cognito.md). To meter the use of your API, set up [usage plans](api-gateway-api-usage-plans.md) to throttle API requests. You
can enable these when creating or updating your API.

The following diagram shows the features available for REST API development and where in the request/response
model these features are configured.

![Diagram of API Gateway features](https://docs.aws.amazon.com/images/apigateway/latest/developerguide/images/develop-features.png)

For an introduction on how to create an API, see
[Tutorial: Create a REST API with a Lambda proxy integration](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-create-api-as-simple-proxy-for-lambda.html). To learn more information about the capabilities of API Gateway that you might use while
developing a REST API, see the following topics. These topics contain conceptual information and procedures that you
can perform using the API Gateway console, the API Gateway REST API, the AWS CLI, or one of the AWS SDKs.

###### Topics

- [API endpoint types for REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-endpoint-types.html)

- [Security policies for REST APIs in API Gateway](apigateway-security-policies.md)

- [IP address types for REST APIs in API Gateway](api-gateway-ip-address-type.md)

- [Methods for REST APIs in API Gateway](how-to-method-settings.md)

- [Control and manage access to REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-control-access-to-api.html)

- [Integrations for REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-integration-settings.html)

- [Request validation for REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html)

- [Data transformations for REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/rest-api-data-transformations.html)

- [Gateway responses for REST APIs in API Gateway](api-gateway-gatewayresponse-definition.md)

- [CORS for REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-cors.html)

- [Binary media types for REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-payload-encodings.html)

- [Invoke REST APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-call-api.html)

- [Develop REST APIs using OpenAPI in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

API Gateway REST APIs

API Gateway endpoint types
