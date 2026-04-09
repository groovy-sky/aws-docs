# TestInvokeAuthorizer

Simulate the execution of an Authorizer in your RestApi with headers, parameters, and an incoming request body.

## Request Syntax

```nohighlight

POST /restapis/restapi_id/authorizers/authorizer_id HTTP/1.1
Content-type: application/json

{
   "additionalContext": {
      "string" : "string"
   },
   "body": "string",
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

**[authorizer\_id](#API_TestInvokeAuthorizer_RequestSyntax)**

Specifies a test invoke authorizer request's Authorizer ID.

Required: Yes

**[restapi\_id](#API_TestInvokeAuthorizer_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[additionalContext](#API_TestInvokeAuthorizer_RequestSyntax)**

A key-value map of additional context variables.

Type: String to string map

Required: No

**[body](#API_TestInvokeAuthorizer_RequestSyntax)**

The simulated request body of an incoming invocation request.

Type: String

Required: No

**[headers](#API_TestInvokeAuthorizer_RequestSyntax)**

A key-value map of headers to simulate an incoming invocation request. This is where the incoming authorization token, or identity source, should be specified.

Type: String to string map

Required: No

**[multiValueHeaders](#API_TestInvokeAuthorizer_RequestSyntax)**

The headers as a map from string to list of values to simulate an incoming invocation request. This is where the incoming authorization token, or identity source, may be specified.

Type: String to array of strings map

Required: No

**[pathWithQueryString](#API_TestInvokeAuthorizer_RequestSyntax)**

The URI path, including query string, of the simulated invocation request. Use this to specify path parameters and query string parameters.

Type: String

Required: No

**[stageVariables](#API_TestInvokeAuthorizer_RequestSyntax)**

A key-value map of stage variables to simulate an invocation on a deployed Stage.

Type: String to string map

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "authorization": {
      "string" : [ "string" ]
   },
   "claims": {
      "string" : "string"
   },
   "clientStatus": number,
   "latency": number,
   "log": "string",
   "policy": "string",
   "principalId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[authorization](#API_TestInvokeAuthorizer_ResponseSyntax)**

The authorization response.

Type: String to array of strings map

**[claims](#API_TestInvokeAuthorizer_ResponseSyntax)**

The open identity claims, with any supported custom attributes, returned from the Cognito Your User Pool configured for the API.

Type: String to string map

**[clientStatus](#API_TestInvokeAuthorizer_ResponseSyntax)**

The HTTP status code that the client would have received. Value is 0 if the authorizer succeeded.

Type: Integer

**[latency](#API_TestInvokeAuthorizer_ResponseSyntax)**

The execution latency, in ms, of the test authorizer request.

Type: Long

**[log](#API_TestInvokeAuthorizer_ResponseSyntax)**

The API Gateway execution log for the test authorizer request.

Type: String

**[policy](#API_TestInvokeAuthorizer_ResponseSyntax)**

The JSON policy document returned by the Authorizer

Type: String

**[principalId](#API_TestInvokeAuthorizer_ResponseSyntax)**

The principal identity returned by the Authorizer

Type: String

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/testinvokeauthorizer.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/testinvokeauthorizer.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

TestInvokeMethod

All content copied from https://docs.aws.amazon.com/.
