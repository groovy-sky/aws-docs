# CreateMultiRegionAccessPoint

###### Note

This operation is not supported by directory buckets.

Creates a Multi-Region Access Point and associates it with the specified buckets. For more information
about creating Multi-Region Access Points, see [Creating Multi-Region Access Points](../userguide/creatingmultiregionaccesspoints.md) in the _Amazon S3 User Guide_.

This action will always be routed to the US West (Oregon) Region. For more information
about the restrictions around working with Multi-Region Access Points, see [Multi-Region Access Point\
restrictions and limitations](../userguide/multiregionaccesspointrestrictions.md) in the _Amazon S3 User Guide_.

This request is asynchronous, meaning that you might receive a response before the
command has completed. When this request provides a response, it provides a token that you
can use to monitor the status of the request with
`DescribeMultiRegionAccessPointOperation`.

The following actions are related to `CreateMultiRegionAccessPoint`:

- [DeleteMultiRegionAccessPoint](api-control-deletemultiregionaccesspoint.md)

- [DescribeMultiRegionAccessPointOperation](api-control-describemultiregionaccesspointoperation.md)

- [GetMultiRegionAccessPoint](api-control-getmultiregionaccesspoint.md)

- [ListMultiRegionAccessPoints](api-control-listmultiregionaccesspoints.md)

## Request Syntax

```nohighlight

POST /v20180820/async-requests/mrap/create HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<CreateMultiRegionAccessPointRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <ClientToken>string</ClientToken>
   <Details>
      <Name>string</Name>
      <PublicAccessBlock>
         <BlockPublicAcls>boolean</BlockPublicAcls>
         <BlockPublicPolicy>boolean</BlockPublicPolicy>
         <IgnorePublicAcls>boolean</IgnorePublicAcls>
         <RestrictPublicBuckets>boolean</RestrictPublicBuckets>
      </PublicAccessBlock>
      <Regions>
         <Region>
            <Bucket>string</Bucket>
            <BucketAccountId>string</BucketAccountId>
         </Region>
      </Regions>
   </Details>
</CreateMultiRegionAccessPointRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[x-amz-account-id](#API_control_CreateMultiRegionAccessPoint_RequestSyntax)**

The AWS account ID for the owner of the Multi-Region Access Point. The owner of the Multi-Region Access Point also must own
the underlying buckets.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[CreateMultiRegionAccessPointRequest](#API_control_CreateMultiRegionAccessPoint_RequestSyntax)**

Root level tag for the CreateMultiRegionAccessPointRequest parameters.

Required: Yes

**[ClientToken](#API_control_CreateMultiRegionAccessPoint_RequestSyntax)**

An idempotency token used to identify the request and guarantee that requests are
unique.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `\S+`

Required: Yes

**[Details](#API_control_CreateMultiRegionAccessPoint_RequestSyntax)**

A container element containing details about the Multi-Region Access Point.

Type: [CreateMultiRegionAccessPointInput](api-control-createmultiregionaccesspointinput.md) data type

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<CreateMultiRegionAccessPointResult>
   <RequestTokenARN>string</RequestTokenARN>
</CreateMultiRegionAccessPointResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[CreateMultiRegionAccessPointResult](#API_control_CreateMultiRegionAccessPoint_ResponseSyntax)**

Root level tag for the CreateMultiRegionAccessPointResult parameters.

Required: Yes

**[RequestTokenARN](#API_control_CreateMultiRegionAccessPoint_ResponseSyntax)**

The request token associated with the request. You can use this token with [DescribeMultiRegionAccessPointOperation](api-control-describemultiregionaccesspointoperation.md) to determine the status of asynchronous
requests.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `arn:.+`

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/createmultiregionaccesspoint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/createmultiregionaccesspoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateJob

CreateStorageLensGroup

All content copied from https://docs.aws.amazon.com/.
