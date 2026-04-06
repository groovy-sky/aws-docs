# PutMultiRegionAccessPointPolicy

###### Note

This operation is not supported by directory buckets.

Associates an access control policy with the specified Multi-Region Access Point. Each Multi-Region Access Point can have only
one policy, so a request made to this action replaces any existing policy that is
associated with the specified Multi-Region Access Point.

This action will always be routed to the US West (Oregon) Region. For more information
about the restrictions around working with Multi-Region Access Points, see [Multi-Region Access Point\
restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html) in the _Amazon S3 User Guide_.

The following actions are related to
`PutMultiRegionAccessPointPolicy`:

- [GetMultiRegionAccessPointPolicy](api-control-getmultiregionaccesspointpolicy.md)

- [GetMultiRegionAccessPointPolicyStatus](api-control-getmultiregionaccesspointpolicystatus.md)

## Request Syntax

```nohighlight

POST /v20180820/async-requests/mrap/put-policy HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<PutMultiRegionAccessPointPolicyRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <ClientToken>string</ClientToken>
   <Details>
      <Name>string</Name>
      <Policy>string</Policy>
   </Details>
</PutMultiRegionAccessPointPolicyRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[x-amz-account-id](#API_control_PutMultiRegionAccessPointPolicy_RequestSyntax)**

The AWS account ID for the owner of the Multi-Region Access Point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[PutMultiRegionAccessPointPolicyRequest](#API_control_PutMultiRegionAccessPointPolicy_RequestSyntax)**

Root level tag for the PutMultiRegionAccessPointPolicyRequest parameters.

Required: Yes

**[ClientToken](#API_control_PutMultiRegionAccessPointPolicy_RequestSyntax)**

An idempotency token used to identify the request and guarantee that requests are
unique.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `\S+`

Required: Yes

**[Details](#API_control_PutMultiRegionAccessPointPolicy_RequestSyntax)**

A container element containing the details of the policy for the Multi-Region Access Point.

Type: [PutMultiRegionAccessPointPolicyInput](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutMultiRegionAccessPointPolicyInput.html) data type

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<PutMultiRegionAccessPointPolicyResult>
   <RequestTokenARN>string</RequestTokenARN>
</PutMultiRegionAccessPointPolicyResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[PutMultiRegionAccessPointPolicyResult](#API_control_PutMultiRegionAccessPointPolicy_ResponseSyntax)**

Root level tag for the PutMultiRegionAccessPointPolicyResult parameters.

Required: Yes

**[RequestTokenARN](#API_control_PutMultiRegionAccessPointPolicy_ResponseSyntax)**

The request token associated with the request. You can use this token with [DescribeMultiRegionAccessPointOperation](api-control-describemultiregionaccesspointoperation.md) to determine the status of asynchronous
requests.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `arn:.+`

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/PutMultiRegionAccessPointPolicy)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PutJobTagging

PutPublicAccessBlock
