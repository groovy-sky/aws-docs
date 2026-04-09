# GetMultiRegionAccessPoint

###### Note

This operation is not supported by directory buckets.

Returns configuration information about the specified Multi-Region Access Point.

This action will always be routed to the US West (Oregon) Region. For more information
about the restrictions around working with Multi-Region Access Points, see [Multi-Region Access Point\
restrictions and limitations](../userguide/multiregionaccesspointrestrictions.md) in the _Amazon S3 User Guide_.

The following actions are related to `GetMultiRegionAccessPoint`:

- [CreateMultiRegionAccessPoint](api-control-createmultiregionaccesspoint.md)

- [DeleteMultiRegionAccessPoint](api-control-deletemultiregionaccesspoint.md)

- [DescribeMultiRegionAccessPointOperation](api-control-describemultiregionaccesspointoperation.md)

- [ListMultiRegionAccessPoints](api-control-listmultiregionaccesspoints.md)

## Request Syntax

```nohighlight

GET /v20180820/mrap/instances/name+ HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetMultiRegionAccessPoint_RequestSyntax)**

The name of the Multi-Region Access Point whose configuration information you want to receive. The name of
the Multi-Region Access Point is different from the alias. For more information about the distinction between
the name and the alias of an Multi-Region Access Point, see [Rules for naming Amazon S3 Multi-Region Access Points](../userguide/creatingmultiregionaccesspoints.md#multi-region-access-point-naming) in the
_Amazon S3 User Guide_.

Length Constraints: Maximum length of 50.

Pattern: `^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$`

Required: Yes

**[x-amz-account-id](#API_control_GetMultiRegionAccessPoint_RequestSyntax)**

The AWS account ID for the owner of the Multi-Region Access Point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetMultiRegionAccessPointResult>
   <AccessPoint>
      <Alias>string</Alias>
      <CreatedAt>timestamp</CreatedAt>
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
            <Region>string</Region>
         </Region>
      </Regions>
      <Status>string</Status>
   </AccessPoint>
</GetMultiRegionAccessPointResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetMultiRegionAccessPointResult](#API_control_GetMultiRegionAccessPoint_ResponseSyntax)**

Root level tag for the GetMultiRegionAccessPointResult parameters.

Required: Yes

**[AccessPoint](#API_control_GetMultiRegionAccessPoint_ResponseSyntax)**

A container element containing the details of the requested Multi-Region Access Point.

Type: [MultiRegionAccessPointReport](api-control-multiregionaccesspointreport.md) data type

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getmultiregionaccesspoint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getmultiregionaccesspoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetJobTagging

GetMultiRegionAccessPointPolicy

All content copied from https://docs.aws.amazon.com/.
