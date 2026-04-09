# UpdateUsage

Grants a temporary extension to the remaining quota of a usage plan associated with a specified API key.

## Request Syntax

```nohighlight

PATCH /usageplans/usageplanId/keys/keyId/usage HTTP/1.1
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

**[keyId](#API_UpdateUsage_RequestSyntax)**

The identifier of the API key associated with the usage plan in which a temporary extension is granted to the remaining quota.

Required: Yes

**[usageplanId](#API_UpdateUsage_RequestSyntax)**

The Id of the usage plan associated with the usage data.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateUsage_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

Type: Array of [PatchOperation](api-patchoperation.md) objects

Required: No

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

**[endDate](#API_UpdateUsage_ResponseSyntax)**

The ending date of the usage data.

Type: String

**[values](#API_UpdateUsage_ResponseSyntax)**

The usage data, as daily logs of used and remaining quotas, over the specified time interval indexed over the API keys in a usage plan. For example, `{..., "values" : { "{api_key}" : [ [0, 100], [10, 90], [100, 10]]}`, where `{api_key}` stands for an API key ID and the daily log entry is of the format `[used quota, remaining quota]`.

Type: String to array of arrays of longs map

**[position](#API_UpdateUsage_ResponseSyntax)**

The current pagination position in the paged result set.

Type: String

**[startDate](#API_UpdateUsage_ResponseSyntax)**

The starting date of the usage data.

Type: String

**[usagePlanId](#API_UpdateUsage_ResponseSyntax)**

The plan Id associated with this usage data.

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

## Examples

### Update usage

This example illustrates one usage of UpdateUsage.

#### Sample Request

```

PATCH /usageplans/ywbqww/keys/3JX4ISs7Ik23cDsgiaJdu6SiLffQpIsU7AyTMALs6/usage HTTP/1.1
Content-Type: application/json
Host: apigateway.ap-southeast-1.amazonaws.com
Content-Length: 114
X-Amz-Date: 20160801T235803Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160801/ap-southeast-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sigv4_hash}

{
    "patchOperation" : [ {
       "op": "replace",
       "path": "/remaining",
       "value": "10"
    } ]
}
```

#### Sample Response

```

{
  "_links": {
    "self": {
      "href": "/usageplans/ywbqww/keys/3JX4ISs7Ik23cDsgiaJdu6SiLffQpIsU7AyTMALs6/usage"
    }
  },
  "endDate": "2016-08-08",
  "startDate": "2016-08-08",
  "usagePlanId": "ywbqww",
  "values": {
    "3JX4ISs7Ik23cDsgiaJdu6SiLffQpIsU7AyTMALs6": [
      0,
      10
    ]
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/updateusage.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/updateusage.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateStage

UpdateUsagePlan

All content copied from https://docs.aws.amazon.com/.
