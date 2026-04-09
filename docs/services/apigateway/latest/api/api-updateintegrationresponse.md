# UpdateIntegrationResponse

Represents an update integration response.

## Request Syntax

```nohighlight

PATCH /restapis/restapi_id/resources/resource_id/methods/http_method/integration/responses/status_code HTTP/1.1
Content-type: application/json

{
   "patchOperations": [
      {
         "from": "string",
         "op": "string",
         "path": "string",
         "value": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[http\_method](#API_UpdateIntegrationResponse_RequestSyntax)**

Specifies an update integration response request's HTTP method.

Required: Yes

**[resource\_id](#API_UpdateIntegrationResponse_RequestSyntax)**

Specifies an update integration response request's resource identifier.

Required: Yes

**[restapi\_id](#API_UpdateIntegrationResponse_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

**[status\_code](#API_UpdateIntegrationResponse_RequestSyntax)**

Specifies an update integration response request's status code.

Pattern: `[1-5]\d\d`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateIntegrationResponse_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

Type: Array of [PatchOperation](api-patchoperation.md) objects

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "contentHandling": "string",
   "responseParameters": {
      "string" : "string"
   },
   "responseTemplates": {
      "string" : "string"
   },
   "selectionPattern": "string",
   "statusCode": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[contentHandling](#API_UpdateIntegrationResponse_ResponseSyntax)**

Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`, with the following behaviors:

If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.

Type: String

Valid Values: `CONVERT_TO_BINARY | CONVERT_TO_TEXT`

**[responseParameters](#API_UpdateIntegrationResponse_ResponseSyntax)**

A key-value map specifying response parameters that are passed to the method response from the back end.
The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of `method.response.header.{name}`, where `name` is a valid and unique header name. The mapped non-static value must match the pattern of `integration.response.header.{name}` or `integration.response.body.{JSON-expression}`, where `name` is a valid and unique response header name and `JSON-expression` is a valid JSON expression without the `$` prefix.

Type: String to string map

**[responseTemplates](#API_UpdateIntegrationResponse_ResponseSyntax)**

Specifies the templates used to transform the integration response body. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.

Type: String to string map

**[selectionPattern](#API_UpdateIntegrationResponse_ResponseSyntax)**

Specifies the regular expression (regex) pattern used to choose an integration response based on the response from the back end. For example, if the success response returns nothing and the error response returns some string, you could use the `.+` regex to match error response. However, make sure that the error response does not contain any newline ( `\n`) character in such cases. If the back end is an AWS Lambda function, the AWS Lambda function error header is matched. For all other HTTP and AWS back ends, the HTTP status code is matched.

Type: String

**[statusCode](#API_UpdateIntegrationResponse_ResponseSyntax)**

Specifies the status code that is used to map the integration response to an existing MethodResponse.

Type: String

Pattern: `[1-5]\d\d`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

**LimitExceededException**

The request exceeded the rate limit. Retry after the specified time period.

HTTP Status Code: 429

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Update a response template

This example illustrates one usage of UpdateIntegrationResponse.

#### Sample Request

```

PATCH /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200 HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160615T002050Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160615/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
  "patchOperations" : [ {
    "op" : "replace",
    "path" : "/responseTemplates/application~1json",
    "value": "{\n}"
  }]
}
```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-integration-response-{rel}.html",
      "name": "integrationresponse",
      "templated": true
    },
    "self": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200",
      "title": "200"
    },
    "integrationresponse:delete": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
    },
    "integrationresponse:update": {
      "href": "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/integration/responses/200"
    }
  },
  "responseTemplates": {
    "application/json": "{\n}"
  },
  "statusCode": "200"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/updateintegrationresponse.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/updateintegrationresponse.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateIntegration

UpdateMethod

All content copied from https://docs.aws.amazon.com/.
