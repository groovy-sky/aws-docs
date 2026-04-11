# GetAccessGrantsInstanceResourcePolicy

Returns the resource policy of the S3 Access Grants instance.

Permissions

You must have the `s3:GetAccessGrantsInstanceResourcePolicy` permission to use this operation.

## Request Syntax

```nohighlight

GET /v20180820/accessgrantsinstance/resourcepolicy HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[x-amz-account-id](#API_control_GetAccessGrantsInstanceResourcePolicy_RequestSyntax)**

The AWS account ID of the S3 Access Grants instance.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetAccessGrantsInstanceResourcePolicyResult>
   <Policy>string</Policy>
   <Organization>string</Organization>
   <CreatedAt>timestamp</CreatedAt>
</GetAccessGrantsInstanceResourcePolicyResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetAccessGrantsInstanceResourcePolicyResult](#API_control_GetAccessGrantsInstanceResourcePolicy_ResponseSyntax)**

Root level tag for the GetAccessGrantsInstanceResourcePolicyResult parameters.

Required: Yes

**[CreatedAt](#API_control_GetAccessGrantsInstanceResourcePolicy_ResponseSyntax)**

The date and time when you created the S3 Access Grants instance resource policy.

Type: Timestamp

**[Organization](#API_control_GetAccessGrantsInstanceResourcePolicy_ResponseSyntax)**

The Organization of the resource policy of the S3 Access Grants instance.

Type: String

Length Constraints: Minimum length of 12. Maximum length of 34.

Pattern: `^o-[a-z0-9]{10,32}$`

**[Policy](#API_control_GetAccessGrantsInstanceResourcePolicy_ResponseSyntax)**

The resource policy of the S3 Access Grants instance.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 350000.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getaccessgrantsinstanceresourcepolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAccessGrantsInstanceForPrefix

GetAccessGrantsLocation

All content copied from https://docs.aws.amazon.com/.
