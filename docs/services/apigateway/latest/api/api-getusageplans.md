# GetUsagePlans

Gets all the usage plans of the caller's account.

## Request Syntax

```nohighlight

GET /usageplans?keyId=keyId&limit=limit&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[keyId](#API_GetUsagePlans_RequestSyntax)**

The identifier of the API key associated with the usage plans.

**[limit](#API_GetUsagePlans_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetUsagePlans_RequestSyntax)**

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
         "apiStages": [
            {
               "apiId": "string",
               "stage": "string",
               "throttle": {
                  "string" : {
                     "burstLimit": number,
                     "rateLimit": number
                  }
               }
            }
         ],
         "description": "string",
         "id": "string",
         "name": "string",
         "productCode": "string",
         "quota": {
            "limit": number,
            "offset": number,
            "period": "string"
         },
         "tags": {
            "string" : "string"
         },
         "throttle": {
            "burstLimit": number,
            "rateLimit": number
         }
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetUsagePlans_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [UsagePlan](api-usageplan.md) objects

**[position](#API_GetUsagePlans_ResponseSyntax)**

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

### Retrieve usage plans

The following example request gets the UsagePlans of the caller's account.

The successful response returns a `200 OK` status code and a payload similar to
the following:

#### Sample Request

```

GET /usageplans HTTP/1.1
Content-Type: application/json
Host: apigateway.ap-southeast-1.amazonaws.com
Content-Length: 254
X-Amz-Date: 20160801T204414Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160801/ap-southeast-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sigv4_hash}
```

#### Sample Response

```

{
  "_links": {
    "curies": [
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-usage-{rel}.html",
        "name": "usage",
        "templated": true
      },
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-usageplan-{rel}.html",
        "name": "usageplan",
        "templated": true
      },
      {
        "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-usageplankey-{rel}.html",
        "name": "usageplankey",
        "templated": true
      }
    ],
    "self": {
      "href": "/usageplans"
    },
    "item": {
      "href": "/usageplans/ywbqww"
    },
    "usageplan:by-id": {
      "href": "/usageplans/{usageplanId}",
      "templated": true
    },
    "usageplan:create": {
      "href": "/usageplans"
    }
  },
  "_embedded": {
    "item": {
      "_links": {
        "self": {
          "href": "/usageplans/ywbqww"
        },
        "usage:get": {
          "href": "/usageplans/ywbqww/usage?startDate=2016-07-02&endDate=2016-08-01"
        },
        "usageplan:delete": {
          "href": "/usageplans/ywbqww"
        },
        "usageplan:update": {
          "href": "/usageplans/ywbqww"
        },
        "usageplan:usageplankeys": {
          "href": "/usageplans/ywbqww/keys"
        },
        "usageplankey:create": {
          "href": "/usageplans/ywbqww/keys"
        }
      },
      "apiStages": {
        "stage": "testStage",
        "apiId": "xbvxlpijch"
      },
      "description": "Plan A",
      "id": "ywbqww",
      "name": "Plan_A",
      "quota": {
        "period": "MONTH",
        "offset": 0,
        "limit": 1000
      },
      "throttle": {
        "rateLimit": 100,
        "burstLimit": 200
      }
    }
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getusageplans.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getusageplans.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetUsagePlanKeys

GetVpcLink

All content copied from https://docs.aws.amazon.com/.
