# GetUsagePlanKeys

Gets all the usage plan keys representing the API keys added to a specified usage plan.

## Request Syntax

```nohighlight

GET /usageplans/usageplanId/keys?limit=limit&name=nameQuery&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetUsagePlanKeys_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[nameQuery](#API_GetUsagePlanKeys_RequestSyntax)**

A query parameter specifying the name of the to-be-returned usage plan keys.

**[position](#API_GetUsagePlanKeys_RequestSyntax)**

The current pagination position in the paged result set.

**[usageplanId](#API_GetUsagePlanKeys_RequestSyntax)**

The Id of the UsagePlan resource representing the usage plan containing the to-be-retrieved UsagePlanKey resource representing a plan customer.

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
         "id": "string",
         "name": "string",
         "type": "string",
         "value": "string"
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetUsagePlanKeys_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [UsagePlanKey](api-usageplankey.md) objects

**[position](#API_GetUsagePlanKeys_ResponseSyntax)**

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

### Retrieve usage plan keys

This example illustrates one usage of GetUsagePlanKeys.

#### Sample Request

```

GET /usageplans/n371pt/keys HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
Content-Length: 58
X-Amz-Date: 20160805T185459Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160805/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sigv4_hash}
Response
```

#### Sample Response

```

{
  "_links": {
    "self": {
      "href": "/usageplans/n371pt/keys"
    },
    "item": {
      "href": "/usageplans/n371pt/keys/4wj0d1lt91"
    }
  },
  "_embedded": {
    "item": {
      "_links": {
        "self": {
          "href": "/usageplans/n371pt/keys/4wj0d1lt91"
        }
      },
      "id": "4wj0d1lt91",
      "name": "MyApiKey",
      "type": "API_KEY"
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getusageplankeys.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getusageplankeys.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetUsagePlanKey

GetUsagePlans

All content copied from https://docs.aws.amazon.com/.
