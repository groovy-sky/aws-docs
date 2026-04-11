# GetUsage

Gets the usage data of a usage plan in a specified time interval.

## Request Syntax

```nohighlight

GET /usageplans/usageplanId/usage?endDate=endDate&keyId=keyId&limit=limit&position=position&startDate=startDate HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[endDate](#API_GetUsage_RequestSyntax)**

The ending date (e.g., 2016-12-31) of the usage data.

Required: Yes

**[keyId](#API_GetUsage_RequestSyntax)**

The Id of the API key associated with the resultant usage data.

**[limit](#API_GetUsage_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetUsage_RequestSyntax)**

The current pagination position in the paged result set.

**[startDate](#API_GetUsage_RequestSyntax)**

The starting date (e.g., 2016-01-01) of the usage data.

Required: Yes

**[usageplanId](#API_GetUsage_RequestSyntax)**

The Id of the usage plan associated with the usage data.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "endDate": "string",
   "values": {
      "string" : [
         [ number ]
      ]
   },
   "position": "string",
   "startDate": "string",
   "usagePlanId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[endDate](#API_GetUsage_ResponseSyntax)**

The ending date of the usage data.

Type: String

**[values](#API_GetUsage_ResponseSyntax)**

The usage data, as daily logs of used and remaining quotas, over the specified time interval indexed over the API keys in a usage plan. For example, `{..., "values" : { "{api_key}" : [ [0, 100], [10, 90], [100, 10]]}`, where `{api_key}` stands for an API key ID and the daily log entry is of the format `[used quota, remaining quota]`.

Type: String to array of arrays of longs map

**[position](#API_GetUsage_ResponseSyntax)**

The current pagination position in the paged result set.

Type: String

**[startDate](#API_GetUsage_ResponseSyntax)**

The starting date of the usage data.

Type: String

**[usagePlanId](#API_GetUsage_ResponseSyntax)**

The plan Id associated with this usage data.

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

### Get usage data

This example illustrates one usage of GetUsage.

#### Sample Request

```

GET /usageplans/q2hrol/usage?startDate=2016-08-01&endDate=2016-08-04 HTTP/1.1
Content-Type: application/json
Host: apigateway.ap-southeast-1.amazonaws.com
Content-Length: 254
X-Amz-Date: 20160801T211516Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160801/ap-southeast-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sigv4_hash}

```

#### Sample Response

```

{
  "_links": {
    "self": {
      "href": "/usageplans/q2hrol/usage?startDate=2016-08-01&endDate=2016-08-04"
    },
    "first": {
      "href": "/usageplans/q2hrol/usage?startDate=2016-08-01&endDate=2016-08-04"
    }
  },
  "endDate": "2016-08-04",
  "startDate": "2016-08-01",
  "usagePlanId": "q2hrol",
  "values": {
    "px1KW6TIqK6L8PfqZYR3R3rrFWSS74BB5qBazOJH6": [
      [
        0,
        5000
      ],
      [
        0,
        5000
      ],
      [
        0,
        10
      ],
      [
        0,
        10
      ]
    ]
  }
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getusage.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getusage.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getusage.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getusage.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getusage.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getusage.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getusage.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getusage.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getusage.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getusage.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetTags

GetUsagePlan

All content copied from https://docs.aws.amazon.com/.
