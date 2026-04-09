# DissociateAccessGrantsIdentityCenter

Dissociates the AWS IAM Identity Center instance from the S3 Access Grants instance.

Permissions

You must have the `s3:DissociateAccessGrantsIdentityCenter` permission to use this operation.

Additional Permissions

You must have the `sso:DeleteApplication` permission to use this operation.

## Request Syntax

```nohighlight

DELETE /v20180820/accessgrantsinstance/identitycenter HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[x-amz-account-id](#API_control_DissociateAccessGrantsIdentityCenter_RequestSyntax)**

The AWS account ID of the S3 Access Grants instance.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/dissociateaccessgrantsidentitycenter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeMultiRegionAccessPointOperation

GetAccessGrant

All content copied from https://docs.aws.amazon.com/.
