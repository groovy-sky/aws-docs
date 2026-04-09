# UpdateApiKey

Changes information about an ApiKey resource.

## Request Syntax

```nohighlight

PATCH /apikeys/api_Key HTTP/1.1
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

**[api\_Key](#API_UpdateApiKey_RequestSyntax)**

The identifier of the ApiKey resource to be updated.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateApiKey_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

Type: Array of [PatchOperation](api-patchoperation.md) objects

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "createdDate": number,
   "customerId": "string",
   "description": "string",
   "enabled": boolean,
   "id": "string",
   "lastUpdatedDate": number,
   "name": "string",
   "stageKeys": [ "string" ],
   "tags": {
      "string" : "string"
   },
   "value": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[createdDate](#API_UpdateApiKey_ResponseSyntax)**

The timestamp when the API Key was created.

Type: Timestamp

**[customerId](#API_UpdateApiKey_ResponseSyntax)**

An AWS Marketplace customer identifier, when integrating with the AWS SaaS Marketplace.

Type: String

**[description](#API_UpdateApiKey_ResponseSyntax)**

The description of the API Key.

Type: String

**[enabled](#API_UpdateApiKey_ResponseSyntax)**

Specifies whether the API Key can be used by callers.

Type: Boolean

**[id](#API_UpdateApiKey_ResponseSyntax)**

The identifier of the API Key.

Type: String

**[lastUpdatedDate](#API_UpdateApiKey_ResponseSyntax)**

The timestamp when the API Key was last updated.

Type: Timestamp

**[name](#API_UpdateApiKey_ResponseSyntax)**

The name of the API Key.

Type: String

**[stageKeys](#API_UpdateApiKey_ResponseSyntax)**

A list of Stage resources that are associated with the ApiKey resource.

Type: Array of strings

**[tags](#API_UpdateApiKey_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

**[value](#API_UpdateApiKey_ResponseSyntax)**

The value of the API Key.

Type: String

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/updateapikey.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/updateapikey.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateAccount

UpdateAuthorizer

All content copied from https://docs.aws.amazon.com/.
