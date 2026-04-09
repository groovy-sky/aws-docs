# AssociateAccessGrantsIdentityCenter

Associate your S3 Access Grants instance with an AWS IAM Identity Center instance. Use this action if you want to create access grants for users or groups from your corporate identity directory. First, you must add your corporate identity directory to AWS IAM Identity Center. Then, you can associate this IAM Identity Center instance with your S3 Access Grants instance.

Permissions

You must have the `s3:AssociateAccessGrantsIdentityCenter` permission to use this operation.

Additional Permissions

You must also have the following permissions: `sso:CreateApplication`, `sso:PutApplicationGrant`, and `sso:PutApplicationAuthenticationMethod`.

## Request Syntax

```nohighlight

POST /v20180820/accessgrantsinstance/identitycenter HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<AssociateAccessGrantsIdentityCenterRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <IdentityCenterArn>string</IdentityCenterArn>
</AssociateAccessGrantsIdentityCenterRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[x-amz-account-id](#API_control_AssociateAccessGrantsIdentityCenter_RequestSyntax)**

The AWS account ID of the S3 Access Grants instance.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[AssociateAccessGrantsIdentityCenterRequest](#API_control_AssociateAccessGrantsIdentityCenter_RequestSyntax)**

Root level tag for the AssociateAccessGrantsIdentityCenterRequest parameters.

Required: Yes

**[IdentityCenterArn](#API_control_AssociateAccessGrantsIdentityCenter_RequestSyntax)**

The Amazon Resource Name (ARN) of the AWS IAM Identity Center instance that you are associating with your S3 Access Grants instance. An IAM Identity Center instance is your corporate identity directory that you added to the IAM Identity Center. You can use the [ListInstances](../../../../reference/singlesignon/latest/apireference/api-listinstances.md) API operation to retrieve a list of your Identity Center instances and their ARNs.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 1224.

Pattern: `arn:[^:]+:sso::(\d{12}){0,1}:instance/.*$`

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/associateaccessgrantsidentitycenter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3 Control

CreateAccessGrant

All content copied from https://docs.aws.amazon.com/.
