# PutAccessGrantsInstanceResourcePolicy

Updates the resource policy of the S3 Access Grants instance.

Permissions

You must have the `s3:PutAccessGrantsInstanceResourcePolicy` permission to use this operation.

## Request Syntax

```nohighlight

PUT /v20180820/accessgrantsinstance/resourcepolicy HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<PutAccessGrantsInstanceResourcePolicyRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <Policy>string</Policy>
   <Organization>string</Organization>
</PutAccessGrantsInstanceResourcePolicyRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[x-amz-account-id](#API_control_PutAccessGrantsInstanceResourcePolicy_RequestSyntax)**

The AWS account ID of the S3 Access Grants instance.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[PutAccessGrantsInstanceResourcePolicyRequest](#API_control_PutAccessGrantsInstanceResourcePolicy_RequestSyntax)**

Root level tag for the PutAccessGrantsInstanceResourcePolicyRequest parameters.

Required: Yes

**[Organization](#API_control_PutAccessGrantsInstanceResourcePolicy_RequestSyntax)**

The Organization of the resource policy of the S3 Access Grants instance.

Type: String

Length Constraints: Minimum length of 12. Maximum length of 34.

Pattern: `^o-[a-z0-9]{10,32}$`

Required: No

**[Policy](#API_control_PutAccessGrantsInstanceResourcePolicy_RequestSyntax)**

The resource policy of the S3 Access Grants instance that you are updating.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 350000.

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<PutAccessGrantsInstanceResourcePolicyResult>
   <Policy>string</Policy>
   <Organization>string</Organization>
   <CreatedAt>timestamp</CreatedAt>
</PutAccessGrantsInstanceResourcePolicyResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[PutAccessGrantsInstanceResourcePolicyResult](#API_control_PutAccessGrantsInstanceResourcePolicy_ResponseSyntax)**

Root level tag for the PutAccessGrantsInstanceResourcePolicyResult parameters.

Required: Yes

**[CreatedAt](#API_control_PutAccessGrantsInstanceResourcePolicy_ResponseSyntax)**

The date and time when you created the S3 Access Grants instance resource policy.

Type: Timestamp

**[Organization](#API_control_PutAccessGrantsInstanceResourcePolicy_ResponseSyntax)**

The Organization of the resource policy of the S3 Access Grants instance.

Type: String

Length Constraints: Minimum length of 12. Maximum length of 34.

Pattern: `^o-[a-z0-9]{10,32}$`

**[Policy](#API_control_PutAccessGrantsInstanceResourcePolicy_ResponseSyntax)**

The updated resource policy of the S3 Access Grants instance.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 350000.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/PutAccessGrantsInstanceResourcePolicy)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListTagsForResource

PutAccessPointConfigurationForObjectLambda
