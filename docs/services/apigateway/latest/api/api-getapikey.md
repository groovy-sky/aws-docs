# GetApiKey

Gets information about the current ApiKey resource.

## Request Syntax

```nohighlight

GET /apikeys/api_Key?includeValue=includeValue HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[api\_Key](#API_GetApiKey_RequestSyntax)**

The identifier of the ApiKey resource.

Required: Yes

**[includeValue](#API_GetApiKey_RequestSyntax)**

A boolean flag to specify whether ( `true`) or not ( `false`) the result contains the key value.

## Request Body

The request does not have a request body.

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

**[createdDate](#API_GetApiKey_ResponseSyntax)**

The timestamp when the API Key was created.

Type: Timestamp

**[customerId](#API_GetApiKey_ResponseSyntax)**

An AWS Marketplace customer identifier, when integrating with the AWS SaaS Marketplace.

Type: String

**[description](#API_GetApiKey_ResponseSyntax)**

The description of the API Key.

Type: String

**[enabled](#API_GetApiKey_ResponseSyntax)**

Specifies whether the API Key can be used by callers.

Type: Boolean

**[id](#API_GetApiKey_ResponseSyntax)**

The identifier of the API Key.

Type: String

**[lastUpdatedDate](#API_GetApiKey_ResponseSyntax)**

The timestamp when the API Key was last updated.

Type: Timestamp

**[name](#API_GetApiKey_ResponseSyntax)**

The name of the API Key.

Type: String

**[stageKeys](#API_GetApiKey_ResponseSyntax)**

A list of Stage resources that are associated with the ApiKey resource.

Type: Array of strings

**[tags](#API_GetApiKey_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

**[value](#API_GetApiKey_ResponseSyntax)**

The value of the API Key.

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

## Examples

### Retrieve an API Key

The following example request retrieves an API key.

The successful response returns
`200 OK` status code and a payload similar to the following:

#### Sample Request

```

GET /apikeys/hzYAVO9Sg98nsNh65VfX81M84O2kyXVy6K1xwHD76 HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20170223T221142Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}

```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-apikey-{rel}.html",
      "name": "apikey",
      "templated": true
    },
    "self": {
      "href": "/apikeys/hzYAVO9Sg98nsNh65VfX81M84O2kyXVy6K1xwHD76"
    },
    "apikey:delete": {
      "href": "/apikeys/hzYAVO9Sg98nsNh65VfX81M84O2kyXVy6K1xwHD76"
    },
    "apikey:update": {
      "href": "/apikeys/hzYAVO9Sg98nsNh65VfX81M84O2kyXVy6K1xwHD76"
    }
  },
  "createdDate": "2015-11-06T23:51:03Z",
  "enabled": true,
  "id": "hzYAVO9Sg98nsNh65VfX81M84O2kyXVy6K1xwHD76",
  "lastUpdatedDate": "2016-06-06T23:44:43Z",
  "name": "my_test_gateway_service",
  "stageKeys": [
    "h4ah70cvmb/beta",
    "fugvjdxtri/stage2"
  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getapikey.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getapikey.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAccount

GetApiKeys

All content copied from https://docs.aws.amazon.com/.
