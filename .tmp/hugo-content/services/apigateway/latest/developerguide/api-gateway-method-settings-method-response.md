# Set up a method response in API Gateway

An API method response encapsulates the output of an API method request that the
client will receive. The output data includes an HTTP status code, some headers, and
possibly a body.

With non-proxy integrations, the specified response parameters and body can be mapped
from the associated integration response data or can be assigned certain static values
according to mappings. These mappings are specified in the integration response. The
mapping can be an identical transformation that passes the integration response through
as-is.

With a proxy integration, API Gateway passes the backend response through to the method
response automatically. There is no need for you to set up the API method response.
However, with the Lambda proxy integration, the Lambda function must return a result of
[this output\
format](set-up-lambda-proxy-integrations.md#api-gateway-simple-proxy-for-lambda-output-format) for API Gateway to successfully map the integration response to a method
response.

Programmatically, the method response setup amounts to creating a [MethodResponse](../api/api-methodresponse.md) resource of
API Gateway and setting the properties of [statusCode](../api/api-methodresponse.md#statusCode), [responseParameters](../api/api-methodresponse.md#responseParameters), and [responseModels](../api/api-methodresponse.md#responseModels).

When setting status codes for an API method, you should choose one as the default to
handle any integration response of an unanticipated status code. It is reasonable to set
`500` as the default because this amounts to casting otherwise unmapped
responses as a server-side error. For instructional reasons, the API Gateway console sets the
`200` response as the default. But you can reset it to the
`500` response.

To set up a method response, you must have created the method request.

## Set up method response status code

The status code of a method response defines a type of response. For example,
responses of 200, 400, and 500 indicate successful, client-side error and
server-side error responses, respectively.

To set up a method response status code, set the [`statusCode`](../api/api-methodresponse.md#statusCode) property to an HTTP status code. The following
[put-method-response](../../../cli/latest/reference/apigateway/put-method-response.md) command creates `200` method response.

```nohighlight

aws apigateway put-method-response \
    --rest-api-id vaz7da96z6 \
    --resource-id 6sxz2j \
    --http-method GET \
    --status-code 200
```

## Set up method response parameters

Method response parameters define which headers the client receives in response to
the associated method request. Response parameters also specify a target to which
API Gateway maps an integration response parameter, according to mappings prescribed in
the API method's integration response.

To set up the method response parameters, add to the [`responseParameters`](../api/api-methodresponse.md#responseParameters) map of `MethodResponse`
key-value pairs of the `"{parameter-name}":"{boolean}"` format. The
following
[put-method-response](../../../cli/latest/reference/apigateway/put-method-response.md) command sets the `my-header` header.

```nohighlight

aws apigateway put-method-response \
        --rest-api-id vaz7da96z6 \
        --resource-id 6sxz2j \
        --http-method GET \
        --status-code 200  \
        --response-parameters method.response.header.my-header=false
```

## Set up method response models

A method response model defines a format of the method response body.
Setting up a method response model is necessary when you generate a strongly typed SDK for the API. It
ensures that the output is cast into an appropriate class in Java or Objective-C. In other cases, setting a
model is optional.

Before
setting up the response model, you must first create the model in API Gateway. To do so,
you can call the `create-model` command. The following
[create-model](../../../cli/latest/reference/apigateway/create-model.md) command
creates `PetStorePet` model to describe the body of the response to the
`GET /pets/{petId}` method request.

```nohighlight

aws apigateway create-model \
    --rest-api-id vaz7da96z6 \
    --content-type application/json \
    --name PetStorePet \
    --schema '{ \
                  "$schema": "http://json-schema.org/draft-04/schema#", \
                  "title": "PetStorePet", \
                  "type": "object", \
                  "properties": { \
                    "id": { "type": "number" }, \
                    "type": { "type": "string" }, \
                    "price": { "type": "number" } \
                  } \
              }'
```

The result is created as an API Gateway [`Model`](../api/api-model.md) resource.

To set up the method response models to define the payload format, add the "application/json":"PetStorePet"
key-value pair to the [`requestModels`](../api/api-methodresponse.md#responseModels) map of [`MethodResponse`](../api/api-methodresponse.md) resource. The following [put-method-response](../../../cli/latest/reference/apigateway/put-method-response.md) command creates a method
response that uses a response model to define the payload format:

```nohighlight

aws apigateway put-method-response \
    --rest-api-id vaz7da96z6 \
    --resource-id 6sxz2j \
    --http-method GET \
    --status-code 200  \
    --response-parameters method.response.header.my-header=false \
    --response-models '{"application/json":"PetStorePet"}'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Set up method request

Set up method using the console

All content copied from https://docs.aws.amazon.com/.
