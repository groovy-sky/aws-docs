# CreateVpcConnector

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Create an AWS App Runner VPC connector resource. App Runner requires this resource when you want to associate your App Runner service to a custom Amazon Virtual Private Cloud
(Amazon VPC).

## Request Syntax

```nohighlight

{
   "SecurityGroups": [ "string" ],
   "Subnets": [ "string" ],
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ],
   "VpcConnectorName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[SecurityGroups](#API_CreateVpcConnector_RequestSyntax)**

A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the
default security group of the Amazon VPC. The default security group allows all outbound traffic.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**[Subnets](#API_CreateVpcConnector_RequestSyntax)**

A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single
Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.

###### Note

App Runner only supports subnets of IP address type _IPv4_ and _dual stack_ (IPv4 and IPv6).

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: Yes

**[Tags](#API_CreateVpcConnector_RequestSyntax)**

A list of metadata items that you can associate with your VPC connector resource. A tag is a key-value pair.

Type: Array of [Tag](api-tag.md) objects

Required: No

**[VpcConnectorName](#API_CreateVpcConnector_RequestSyntax)**

A name for the VPC connector.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 40.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,39}`

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

**[VpcConnector](#API_CreateVpcConnector_ResponseSyntax)**

A description of the App Runner VPC connector that's created by this request.

Type: [VpcConnector](api-vpcconnector.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**ServiceQuotaExceededException**

App Runner can't create this resource. You've reached your account quota for this resource type.

For App Runner per-resource quotas, see [AWS App Runner endpoints and quotas](../../../../general/latest/gr/apprunner.md) in the
_AWS General Reference_.

HTTP Status Code: 400

## Examples

### Create a VPC connector

This example illustrates how to create a VPC connector.

#### Sample Request

```json

$ aws apprunner create-vpc-connector --cli-input-json "`cat`"
{
  "VpcConnectorName": "my-vpc-connector",
  "Subnets": ["subnet-123", "subnet-456"],
  "SecurityGroups": ["sg-123", "sg-456"]
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

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/createvpcconnector.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/createvpcconnector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateService

CreateVpcIngressConnection

All content copied from https://docs.aws.amazon.com/.
