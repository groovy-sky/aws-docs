# UpdateAccount

Changes information about the current Account resource.

## Request Syntax

```nohighlight

PATCH /account HTTP/1.1
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

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[patchOperations](#API_UpdateAccount_RequestSyntax)**

For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).

Type: Array of [PatchOperation](https://docs.aws.amazon.com/apigateway/latest/api/API_PatchOperation.html) objects

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "apiKeyVersion": "string",
   "cloudwatchRoleArn": "string",
   "features": [ "string" ],
   "throttleSettings": {
      "burstLimit": number,
      "rateLimit": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[apiKeyVersion](#API_UpdateAccount_ResponseSyntax)**

The version of the API keys used for the account.

Type: String

**[cloudwatchRoleArn](#API_UpdateAccount_ResponseSyntax)**

The ARN of an Amazon CloudWatch role for the current Account.

Type: String

**[features](#API_UpdateAccount_ResponseSyntax)**

A list of features supported for the account. When usage plans are enabled, the features list will include an entry of `"UsagePlans"`.

Type: Array of strings

**[throttleSettings](#API_UpdateAccount_ResponseSyntax)**

Specifies the API request limits configured for the current Account.

Type: [ThrottleSettings](https://docs.aws.amazon.com/apigateway/latest/api/API_ThrottleSettings.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/apigateway/latest/api/CommonErrors.html).

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

### Update an Account's CloudWatch role

The following example updates an account's CloudWatch role ARN by replacing the existing one with a new CloudWatch role ARN.

If you specify an invalid role, i.e., a role that does not permit API Gateway to
invoke CloudWatch logs, you will receive an error response. If you attempt to replace the
throttleSettings, in part or as a whole, you will receive 400 Bad Request response with an
error message stating that /throttleSettings value cannot be changed this way, but the
/cloudwatchRoleArn value can.

#### Sample Request

```

PATCH /account HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160531T212738Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160531/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
    "patchOperations" : [{
        "op" : "replace",
        "path" : "/cloudwatchRoleArn",
        "value" : "arn:aws:iam::123456789012:role/apigAwsProxyRole"
    }]
}
```

#### Sample Response

```

{
    "_links": {
        "curies": {
            "href": "https://docs.aws.amazon.com/apigateway/latest/developerguide/account-apigateway-{rel}.html",
            "name": "account",
            "templated": true
        },
        "self": {
            "href": "/account"
        },
        "account:update": {
            "href": "/account"
        }
    },
    "cloudwatchRoleArn": "arn:aws:iam::123456789012:role/apigAwsProxyRole",
    "throttleSettings": {
        "rateLimit": 500,
        "burstLimit": 1000
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigateway-2015-07-09/UpdateAccount)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigateway-2015-07-09/UpdateAccount)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UntagResource

UpdateApiKey
