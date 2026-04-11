# GetAuthorizers

Describe an existing Authorizers resource.

## Request Syntax

```nohighlight

GET /restapis/restapi_id/authorizers?limit=limit&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetAuthorizers_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetAuthorizers_RequestSyntax)**

The current pagination position in the paged result set.

**[restapi\_id](#API_GetAuthorizers_RequestSyntax)**

The string identifier of the associated RestApi.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "item": [
      {
         "authorizerCredentials": "string",
         "authorizerResultTtlInSeconds": number,
         "authorizerUri": "string",
         "authType": "string",
         "id": "string",
         "identitySource": "string",
         "identityValidationExpression": "string",
         "name": "string",
         "providerARNs": [ "string" ],
         "type": "string"
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetAuthorizers_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [Authorizer](api-authorizer.md) objects

**[position](#API_GetAuthorizers_ResponseSyntax)**

The current pagination position in the paged result set.

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

### Get the collection of custom authorizers defined for an API

This example illustrates one usage of GetAuthorizers.

#### Sample Request

```

GET /restapis/86l3267lf6/authorizers HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
Content-Length: 60
X-Amz-Date: 20170223T175134Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
  "_links": {
    "curies": {
      "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-authorizer-{rel}.html",
      "name": "authorizer",
      "templated": true
    },
    "self": {
      "href": "/restapis/86l3267lf6/authorizers"
    },
    "authorizer:by-id": {
      "href": "/restapis/86l3267lf6/authorizers/{authorizer_id}",
      "templated": true
    },
    "authorizer:create": {
      "href": "/restapis/86l3267lf6/authorizers"
    },
    "item": {
      "href": "/restapis/86l3267lf6/authorizers/bs9803"
    }
  },
  "_embedded": {
    "item": {
      "_links": {
        "self": {
          "href": "/restapis/86l3267lf6/authorizers/bs9803"
        },
        "authorizer:delete": {
          "href": "/restapis/86l3267lf6/authorizers/bs9803"
        },
        "authorizer:update": {
          "href": "/restapis/86l3267lf6/authorizers/bs9803"
        }
      },
      "authType": "custom",
      "authorizerResultTtlInSeconds": 300,
      "authorizerUri": "arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:123456789012:function:myCustomAuthorizer/invocations",
      "id": "bs9803",
      "identitySource": "method.request.header.Authorization",
      "name": "myCustomAuth",
      "type": "TOKEN"
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getauthorizers.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getauthorizers.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAuthorizer

GetBasePathMapping

All content copied from https://docs.aws.amazon.com/.
