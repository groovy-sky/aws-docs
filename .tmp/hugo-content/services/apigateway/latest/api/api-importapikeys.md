# ImportApiKeys

Import API keys from an external source, such as a CSV-formatted file.

## Request Syntax

```nohighlight

POST /apikeys?mode=import&failonwarnings=failOnWarnings&format=format HTTP/1.1

body
```

## URI Request Parameters

The request uses the following URI parameters.

**[failOnWarnings](#API_ImportApiKeys_RequestSyntax)**

A query parameter to indicate whether to rollback ApiKey importation ( `true`) or not ( `false`) when error is encountered.

**[format](#API_ImportApiKeys_RequestSyntax)**

A query parameter to specify the input format to imported API keys. Currently, only the `csv` format is supported.

Valid Values: `csv`

Required: Yes

## Request Body

The request accepts the following binary data.

**[body](#API_ImportApiKeys_RequestSyntax)**

The payload of the POST request to import API keys. For the payload format, see API Key File Format.

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "ids": [ "string" ],
   "warnings": [ "string" ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[ids](#API_ImportApiKeys_ResponseSyntax)**

A list of all the ApiKey identifiers.

Type: Array of strings

**[warnings](#API_ImportApiKeys_ResponseSyntax)**

A list of warning messages.

Type: Array of strings

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/importapikeys.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/importapikeys.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetVpcLinks

ImportDocumentationParts

All content copied from https://docs.aws.amazon.com/.
