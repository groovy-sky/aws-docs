# PutAccountSetting

Allows you to change the basic scan type version or registry policy scope.

## Request Syntax

```nohighlight

{
   "name": "string",
   "value": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[name](#API_PutAccountSetting_RequestSyntax)**

The name of the account setting, such as `BASIC_SCAN_TYPE_VERSION`,
`REGISTRY_POLICY_SCOPE`, or `BLOB_MOUNTING`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[value](#API_PutAccountSetting_RequestSyntax)**

Setting value that is specified. Valid value for basic scan type: `AWS_NATIVE`.
Valid values for registry policy scope: `V1` or `V2`.
Valid values for blob mounting: `ENABLED` or `DISABLED`.

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "name": "string",
   "value": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[name](#API_PutAccountSetting_ResponseSyntax)**

Retrieves the name of the account setting.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

**[value](#API_PutAccountSetting_ResponseSyntax)**

Retrieves the value of the specified account setting.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**LimitExceededException**

The operation did not succeed because it would have exceeded a service limit for your
account. For more information, see [Amazon ECR service quotas](../../../../services/amazonecr/latest/userguide/service-quotas.md) in
the Amazon Elastic Container Registry User Guide.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## Examples

In the following example or examples, the Authorization header contents
( `AUTHPARAMS`) must be replaced with an AWS Signature Version 4
signature. For more information about creating these signatures, see [Signature\
Version 4 Signing Process](../../../../general/latest/gr/signature-version-4.md) in the _AWS General_
_Reference_.

You only need to learn how to sign HTTP requests if you intend to manually
create them. When you use the [AWS Command Line Interface (AWS CLI)](http://aws.amazon.com/cli) or one of the [AWS SDKs](http://aws.amazon.com/tools) to make requests to AWS, these tools automatically sign the
requests for you with the access key that you specify when you configure the tools.
When you use these tools, you don't need to learn how to sign requests
yourself.

### Example

This example assigns the `BASIC_SCAN_TYPE_VERSION` to be used in
the registry. The accepted value is `AWS_NATIVE`.

#### Sample Request

```

POST / HTTP/1.1
Host: ecr.us-west-2.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.PutAccountSetting
Content-Type: application/x-amz-json-1.1
User-Agent: aws-cli/1.11.144 Python/3.6.1 Darwin/16.6.0 botocore/1.7.2
X-Amz-Date: 20170901T223937Z
Authorization: AUTHPARAMS
Content-Length: 48

{
   aws ecr put-account-setting --name BASIC_SCAN_TYPE_VERSION --value AWS_NATIVE,
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Server: Server
Date: Fri, 01 Sep 2017 19:42:18 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 340
Connection: keep-alive
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f

{
    "name": "BASIC_SCAN_TYPE_VERSION",
    "value": "AWS_NATIVE"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/putaccountsetting.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/putaccountsetting.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTagsForResource

PutImage

All content copied from https://docs.aws.amazon.com/.
