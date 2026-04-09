# CreateUsagePlanKey

Creates a usage plan key for adding an existing API key to a usage plan.

## Request Syntax

```nohighlight

POST /usageplans/usageplanId/keys HTTP/1.1
Content-type: application/json

{
   "keyId": "string",
   "keyType": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[usageplanId](#API_CreateUsagePlanKey_RequestSyntax)**

The Id of the UsagePlan resource representing the usage plan containing the to-be-created UsagePlanKey resource representing a plan customer.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[keyId](#API_CreateUsagePlanKey_RequestSyntax)**

The identifier of a UsagePlanKey resource for a plan customer.

Type: String

Required: Yes

**[keyType](#API_CreateUsagePlanKey_RequestSyntax)**

The type of a UsagePlanKey resource for a plan customer.

Type: String

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "id": "string",
   "name": "string",
   "type": "string",
   "value": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[id](#API_CreateUsagePlanKey_ResponseSyntax)**

The Id of a usage plan key.

Type: String

**[name](#API_CreateUsagePlanKey_ResponseSyntax)**

The name of a usage plan key.

Type: String

**[type](#API_CreateUsagePlanKey_ResponseSyntax)**

The type of a usage plan key. Currently, the valid key type is `API_KEY`.

Type: String

**[value](#API_CreateUsagePlanKey_ResponseSyntax)**

The value of a usage plan key.

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

### Create a usage plan key

This example illustrates one usage of CreateUsagePlanKey.

#### Sample Request

```

POST /usageplans/n371pt/keys HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160805T181755Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160805/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sigv4_hash}

{
    "keyId": "q5ugs7qjjh",
    "keyType": "API_KEY"
}
```

#### Sample Response

```

{
  "_links": {
    "self": {
      "href": "/usageplans/n371pt/keys/q5ugs7qjjh"
    }
  },
  "id": "q5ugs7qjjh",
  "name": " importedKey_2",
  "type": "API_KEY"
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/createusageplankey.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/createusageplankey.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateUsagePlan

CreateVpcLink

All content copied from https://docs.aws.amazon.com/.
