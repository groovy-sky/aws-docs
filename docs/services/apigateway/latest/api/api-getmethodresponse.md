# GetMethodResponse

Describes a MethodResponse resource.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/resources/resource_id/methods/http_method/responses/status_code HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[http\_method](#API_GetMethodResponse_RequestSyntax)**

The HTTP verb of the Method resource.

Required: Yes

**[resource\_id](#API_GetMethodResponse_RequestSyntax)**

The Resource identifier for the MethodResponse resource.

Required: Yes

**[restapi\_id](#API_GetMethodResponse_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

**[status\_code](#API_GetMethodResponse_RequestSyntax)**

The status code for the MethodResponse resource.

Pattern: `[1-5]\d\d`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "responseModels": {
      "string" : "string"
   },
   "responseParameters": {
      "string" : boolean
   },
   "statusCode": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[responseModels](#API_GetMethodResponse_ResponseSyntax)**

Specifies the Model resources used for the response's content-type. Response models are represented as a key/value map, with a content-type as the key and a Model name as the value.

Type: String to string map

**[responseParameters](#API_GetMethodResponse_ResponseSyntax)**

A key-value map specifying required or optional response parameters that API Gateway can send back to the caller. A key defines a method response header and the value specifies whether the associated method response header is required or not. The expression of the key must match the pattern `method.response.header.{name}`, where `name` is a valid and unique header name. API Gateway passes certain integration response data to the method response headers specified here according to the mapping you prescribe in the API's IntegrationResponse. The integration response data that can be mapped include an integration response header expressed in `integration.response.header.{name}`, a static value enclosed within a pair of single quotes (e.g., `'application/json'`), or a JSON expression from the back-end response payload in the form of `integration.response.body.{JSON-expression}`, where `JSON-expression` is a valid JSON expression without the `$` prefix.)

Type: String to boolean map

**[statusCode](#API_GetMethodResponse_ResponseSyntax)**

The method response's status code.

Type: String

Pattern: `[1-5]\d\d`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

### Get a 200 OK response of a GET method

This example illustrates one usage of GetMethodResponse.

#### Sample Request

```

GET /restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200 HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
Content-Length: 117
X-Amz-Date: 20160613T215008Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160613/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
      "name": "methodresponse",
      "templated": true
    },
    "self": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200",
      "title": "200"
    },
    "methodresponse:delete": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200"
    },
    "methodresponse:update": {
      "href": "/restapis/uojnr9hd57/resources/0cjtch/methods/GET/responses/200"
    }
  },
  "responseModels": {
    "application/json": "Empty"
  },
  "responseParameters": {
    "method.response.header.operator": false,
    "method.response.header.operand_2": false,
    "method.response.header.operand_1": false
  },
  "statusCode": "200"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getmethodresponse.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getmethodresponse.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetMethod

GetModel

All content copied from https://docs.aws.amazon.com/.
