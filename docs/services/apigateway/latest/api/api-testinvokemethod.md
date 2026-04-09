# TestInvokeMethod

Simulate the invocation of a Method in your RestApi with headers, parameters, and an incoming request body.

## Request Syntax

```nohighlight

POST /restapis/restapi_id/resources/resource_id/methods/http_method HTTP/1.1
Content-type: application/json

{
   "body": "string",
   "clientCertificateId": "string",
   "headers": {
      "string" : "string"
   },
   "multiValueHeaders": {
      "string" : [ "string" ]
   },
   "pathWithQueryString": "string",
   "stageVariables": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[http\_method](#API_TestInvokeMethod_RequestSyntax)**

Specifies a test invoke method request's HTTP method.

Required: Yes

**[resource\_id](#API_TestInvokeMethod_RequestSyntax)**

Specifies a test invoke method request's resource ID.

Required: Yes

**[restapi\_id](#API_TestInvokeMethod_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[body](#API_TestInvokeMethod_RequestSyntax)**

The simulated request body of an incoming invocation request.

Type: String

Required: No

**[clientCertificateId](#API_TestInvokeMethod_RequestSyntax)**

A ClientCertificate identifier to use in the test invocation. API Gateway will use the certificate when making the HTTPS request to the defined back-end endpoint.

Type: String

Required: No

**[headers](#API_TestInvokeMethod_RequestSyntax)**

A key-value map of headers to simulate an incoming invocation request.

Type: String to string map

Required: No

**[multiValueHeaders](#API_TestInvokeMethod_RequestSyntax)**

The headers as a map from string to list of values to simulate an incoming invocation request.

Type: String to array of strings map

Required: No

**[pathWithQueryString](#API_TestInvokeMethod_RequestSyntax)**

The URI path, including query string, of the simulated invocation request. Use this to specify path parameters and query string parameters.

Type: String

Required: No

**[stageVariables](#API_TestInvokeMethod_RequestSyntax)**

A key-value map of stage variables to simulate an invocation on a deployed Stage.

Type: String to string map

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "body": "string",
   "headers": {
      "string" : "string"
   },
   "latency": number,
   "log": "string",
   "multiValueHeaders": {
      "string" : [ "string" ]
   },
   "status": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[body](#API_TestInvokeMethod_ResponseSyntax)**

The body of the HTTP response.

Type: String

**[headers](#API_TestInvokeMethod_ResponseSyntax)**

The headers of the HTTP response.

Type: String to string map

**[latency](#API_TestInvokeMethod_ResponseSyntax)**

The execution latency, in ms, of the test invoke request.

Type: Long

**[log](#API_TestInvokeMethod_ResponseSyntax)**

The API Gateway execution log for the test invoke request.

Type: String

**[multiValueHeaders](#API_TestInvokeMethod_ResponseSyntax)**

The headers of the HTTP response as a map from string to list of values.

Type: String to array of strings map

**[status](#API_TestInvokeMethod_ResponseSyntax)**

The HTTP status code.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/testinvokemethod.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/testinvokemethod.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TestInvokeAuthorizer

UntagResource

All content copied from https://docs.aws.amazon.com/.
