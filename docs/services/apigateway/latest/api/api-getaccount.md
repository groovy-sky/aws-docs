# GetAccount

Gets information about the current Account resource.

## Request Syntax

```

GET /account HTTP/1.1

```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "apiKeyVersion": "string",
   "cloudwatchRoleArn": "string",
   "features": [ "string" ],
   "throttleSettings": {
      "burstLimit": number,
      "rateLimit": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[apiKeyVersion](#API_GetAccount_ResponseSyntax)**

The version of the API keys used for the account.

Type: String

**[cloudwatchRoleArn](#API_GetAccount_ResponseSyntax)**

The ARN of an Amazon CloudWatch role for the current Account.

Type: String

**[features](#API_GetAccount_ResponseSyntax)**

A list of features supported for the account. When usage plans are enabled, the features list will include an entry of `"UsagePlans"`.

Type: Array of strings

**[throttleSettings](#API_GetAccount_ResponseSyntax)**

Specifies the API request limits configured for the current Account.

Type: [ThrottleSettings](api-throttlesettings.md) object

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getaccount.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getaccount.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GenerateClientCertificate

GetApiKey

All content copied from https://docs.aws.amazon.com/.
