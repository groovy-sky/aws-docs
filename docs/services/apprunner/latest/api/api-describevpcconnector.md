# DescribeVpcConnector

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Return a description of an AWS App Runner VPC connector resource.

## Request Syntax

```nohighlight

{
   "VpcConnectorArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[VpcConnectorArn](#API_DescribeVpcConnector_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner VPC connector that you want a description for.

The ARN must be a full VPC connector ARN.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

## Response Syntax

```nohighlight

{
   "VpcConnector": {
      "CreatedAt": number,
      "DeletedAt": number,
      "SecurityGroups": [ "string" ],
      "Status": "string",
      "Subnets": [ "string" ],
      "VpcConnectorArn": "string",
      "VpcConnectorName": "string",
      "VpcConnectorRevision": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[VpcConnector](#API_DescribeVpcConnector_ResponseSyntax)**

A description of the App Runner VPC connector that you specified in this request.

Type: [VpcConnector](api-vpcconnector.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource doesn't exist for the specified Amazon Resource Name (ARN) in your AWS account.

HTTP Status Code: 400

## Examples

### Describe a VPC connector

This example illustrates how to get a description of an App Runner VPC connector.

#### Sample Request

```json

$ aws apprunner describe-vpc-connector --cli-input-json "`cat`"
{
  "VpcConnectorArn": "arn:aws:apprunner:us-east-1:123456789012:vpcconnector/my-vpc-connector/1/3f2eb10e2c494674952026f646844e3d"
}
```

#### Sample Response

```json

{
  "VpcConnector": {
    "VpcConnectorArn": "arn:aws:apprunner:us-east-1:123456789012:vpcconnector/my-vpc-connector/1/3f2eb10e2c494674952026f646844e3d",
    "VpcConnectorName": "my-vpc-connector",
    "VpcConnectorRevision": 1,
    "Subnets": ["subnet-123", "subnet-456"],
    "SecurityGroups": ["sg-123", "sg-456"],
    "Status": "ACTIVE",
    "CreatedAt": "2021-08-18T23:36:45.374Z"
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/describevpcconnector.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/describevpcconnector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeService

DescribeVpcIngressConnection

All content copied from https://docs.aws.amazon.com/.
