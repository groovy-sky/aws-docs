# GetApiKeys

Gets information about the current ApiKeys resource.

## Request Syntax

```nohighlight

GET /apikeys?customerId=customerId&includeValues=includeValues&limit=limit&name=nameQuery&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[customerId](#API_GetApiKeys_RequestSyntax)**

The identifier of a customer in AWS Marketplace or an external system, such as a developer portal.

**[includeValues](#API_GetApiKeys_RequestSyntax)**

A boolean flag to specify whether ( `true`) or not ( `false`) the result contains key values.

**[limit](#API_GetApiKeys_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[nameQuery](#API_GetApiKeys_RequestSyntax)**

The name of queried API keys.

**[position](#API_GetApiKeys_RequestSyntax)**

The current pagination position in the paged result set.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "item": [
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
   ],
   "position": "string",
   "warnings": [ "string" ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetApiKeys_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [ApiKey](api-apikey.md) objects

**[position](#API_GetApiKeys_ResponseSyntax)**

The current pagination position in the paged result set.

Type: String

**[warnings](#API_GetApiKeys_ResponseSyntax)**

A list of warning messages logged during the import of API keys when the `failOnWarnings` option is set to true.

Type: Array of strings

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

### Retrieve an API key

The following example request retrieves the available API keys in the caller's
AWS account.

A successful response returns the requested `ApiKey` resources that can be navigated to by following the linked item or examining the embedded item resource.

#### Sample Request

```

GET /apikeys HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160601T180431Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160601/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
  "_links": {
    "curies": [
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-apikey-{rel}.html",
        "name": "apikey",
        "templated": true
      }
    ],
    "self": {
      "href": "/apikeys"
    },
    "apikey:by-key": {
      "href": "/apikeys/{api_Key}",
      "templated": true
    },
    "apikey:create": {
      "href": "/apikeys"
    },
    "apikey:delete": {
      "href": "/apikeys/{api_Key}",
      "templated": true
    },
    "item": {
      "href": "/apikeys/hzYAVO9Sg98nsNh45VfX81M84O2kyXVy6K1xwHD76"
    }
  },
  "_embedded": {
    "item": {
      "_links": {
        "self": {
          "href": "/apikeys/hzYAVO9Sg98nsNh45VfX81M84O2kyXVy6K1xwHD76"
        },
        "apikey:delete": {
          "href": "/apikeys/hzYAVO9Sg98nsNh45VfX81M84O2kyXVy6K1xwHD76"
        },
        "apikey:update": {
          "href": "/apikeys/hzYAVO9Sg98nsNh45VfX81M84O2kyXVy6K1xwHD76"
        }
      },
      "createdDate": "2015-11-06T23:51:03Z",
      "enabled": true,
      "id": "hzYAVO9Sg98nsNh45VfX81M84O2kyXVy6K1xwHD76",
      "lastUpdatedDate": "2016-01-26T20:05:38Z",
      "name": "my_test_gateway_service",
      "stageKeys": "h4ah70cvmb/beta"
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getapikeys.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getapikeys.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetApiKey

GetAuthorizer

All content copied from https://docs.aws.amazon.com/.
