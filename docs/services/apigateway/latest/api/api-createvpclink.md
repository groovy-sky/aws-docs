# CreateVpcLink

Creates a VPC link, under the caller's account in a selected region, in an asynchronous operation that typically takes 2-4 minutes to complete and become operational. The caller must have permissions to create and update VPC Endpoint services.

## Request Syntax

```nohighlight

POST /vpclinks HTTP/1.1
Content-type: application/json

{
   "description": "string",
   "name": "string",
   "tags": {
      "string" : "string"
   },
   "targetArns": [ "string" ]
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[description](#API_CreateVpcLink_RequestSyntax)**

The description of the VPC link.

Type: String

Required: No

**[name](#API_CreateVpcLink_RequestSyntax)**

The name used to label and identify the VPC link.

Type: String

Required: Yes

**[tags](#API_CreateVpcLink_RequestSyntax)**

The key-value map of strings. The valid character set is \[a-zA-Z+-=.\_:/\]. The tag key can be up to 128 characters and must not start with `aws:`. The tag value can be up to 256 characters.

Type: String to string map

Required: No

**[targetArns](#API_CreateVpcLink_RequestSyntax)**

The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.

Type: Array of strings

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 202
Content-type: application/json

{
   "description": "string",
   "id": "string",
   "name": "string",
   "status": "string",
   "statusMessage": "string",
   "tags": {
      "string" : "string"
   },
   "targetArns": [ "string" ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 202 response.

The following data is returned in JSON format by the service.

**[description](#API_CreateVpcLink_ResponseSyntax)**

The description of the VPC link.

Type: String

**[id](#API_CreateVpcLink_ResponseSyntax)**

The identifier of the VpcLink. It is used in an Integration to reference this VpcLink.

Type: String

**[name](#API_CreateVpcLink_ResponseSyntax)**

The name used to label and identify the VPC link.

Type: String

**[status](#API_CreateVpcLink_ResponseSyntax)**

The status of the VPC link. The valid values are `AVAILABLE`, `PENDING`, `DELETING`, or `FAILED`. Deploying an API will wait if the status is `PENDING` and will fail if the status is `DELETING`.

Type: String

Valid Values: `AVAILABLE | PENDING | DELETING | FAILED`

**[statusMessage](#API_CreateVpcLink_ResponseSyntax)**

A description about the VPC link status.

Type: String

**[tags](#API_CreateVpcLink_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

**[targetArns](#API_CreateVpcLink_ResponseSyntax)**

The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.

Type: Array of strings

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

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Create a VPC link

This example illustrates one usage of CreateVpcLink.

#### Sample Request

```

POST /vpclinks HTTP/1.1
Content-Type: application/json
Host: apigateway.us-eas-t.amazonaws.com
Content-Length: ...
X-Amz-Date: 20160801T235803Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160801/ap-southeast-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sigv4_hash}

{
    "name":"my-test-vpc-link",
    "targetArns": ["arn:aws:elasticloadbalancing:us-east-1:123456789012:loadbalancer/net/my-vpclink-test-nlb/1f8df693cd094a72"]
}
```

#### Sample Response

```

{
    "id": "gim7c3",
    "name": "my-test-vpc-link",
    "status": "PENDING",
    "targetArns": "arn:aws:elasticloadbalancing:us-east-1:123456789012:loadbalancer/net/my-vpclink-test-nlb/1f8df693cd094a72"
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigateway-2015-07-09/CreateVpcLink)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigateway-2015-07-09/CreateVpcLink)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateUsagePlanKey

DeleteApiKey
