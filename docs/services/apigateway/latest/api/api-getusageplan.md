# GetUsagePlan

Gets a usage plan of a given plan identifier.

## Request Syntax

```nohighlight

GET /usageplans/usageplanId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[usageplanId](#API_GetUsagePlan_RequestSyntax)**

The identifier of the UsagePlan resource to be retrieved.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

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
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[apiStages](#API_GetUsagePlan_ResponseSyntax)**

The associated API stages of a usage plan.

Type: Array of [ApiStage](api-apistage.md) objects

**[description](#API_GetUsagePlan_ResponseSyntax)**

The description of a usage plan.

Type: String

**[id](#API_GetUsagePlan_ResponseSyntax)**

The identifier of a UsagePlan resource.

Type: String

**[name](#API_GetUsagePlan_ResponseSyntax)**

The name of a usage plan.

Type: String

**[productCode](#API_GetUsagePlan_ResponseSyntax)**

The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on the AWS Marketplace.

Type: String

**[quota](#API_GetUsagePlan_ResponseSyntax)**

The target maximum number of permitted requests per a given unit time interval.

Type: [QuotaSettings](api-quotasettings.md) object

**[tags](#API_GetUsagePlan_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

**[throttle](#API_GetUsagePlan_ResponseSyntax)**

A map containing method level throttling information for API stage in a usage plan.

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

## Examples

### Retrieve information about a usage plan

This example illustrates one usage of GetUsagePlan.

#### Sample Request

```

GET /usageplans/n371pt HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160805T012305Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160805/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sigvv4_hash}
Response

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
      "href": "/usageplans/n371pt"
    },
    "usage:get": {
      "href": "/usageplans/n371pt/usage?startDate=2016-07-06&endDate=2016-08-05"
    },
    "usageplan:delete": {
      "href": "/usageplans/n371pt"
    },
    "usageplan:update": {
      "href": "/usageplans/n371pt"
    },
    "usageplan:usageplankeys": {
      "href": "/usageplans/n371pt/keys"
    },
    "usageplankey:create": {
      "href": "/usageplans/n371pt/keys"
    }
  },
  "apiStages": {
    "stage": "testStage",
    "apiId": "kdwpu39c2k"
  },
  "id": "n371pt",
  "name": "Plan_D",
  "quota": {
    "period": "DAY",
    "offset": 0,
    "limit": 5
  },
  "throttle": {
    "rateLimit": 100,
    "burstLimit": 200
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getusageplan.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getusageplan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetUsage

GetUsagePlanKey

All content copied from https://docs.aws.amazon.com/.
