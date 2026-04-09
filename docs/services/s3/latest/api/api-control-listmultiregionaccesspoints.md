# ListMultiRegionAccessPoints

###### Note

This operation is not supported by directory buckets.

Returns a list of the Multi-Region Access Points currently associated with the specified AWS account.
Each call can return up to 100 Multi-Region Access Points, the maximum number of Multi-Region Access Points that can be
associated with a single account.

This action will always be routed to the US West (Oregon) Region. For more information
about the restrictions around working with Multi-Region Access Points, see [Multi-Region Access Point\
restrictions and limitations](../userguide/multiregionaccesspointrestrictions.md) in the _Amazon S3 User Guide_.

The following actions are related to `ListMultiRegionAccessPoint`:

- [CreateMultiRegionAccessPoint](api-control-createmultiregionaccesspoint.md)

- [DeleteMultiRegionAccessPoint](api-control-deletemultiregionaccesspoint.md)

- [DescribeMultiRegionAccessPointOperation](api-control-describemultiregionaccesspointoperation.md)

- [GetMultiRegionAccessPoint](api-control-getmultiregionaccesspoint.md)

## Request Syntax

```nohighlight

GET /v20180820/mrap/instances?maxResults=MaxResults&nextToken=NextToken HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_control_ListMultiRegionAccessPoints_RequestSyntax)**

Not currently used. Do not use this parameter.

Valid Range: Minimum value of 0. Maximum value of 1000.

**[nextToken](#API_control_ListMultiRegionAccessPoints_RequestSyntax)**

Not currently used. Do not use this parameter.

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[x-amz-account-id](#API_control_ListMultiRegionAccessPoints_RequestSyntax)**

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
<ListMultiRegionAccessPointsResult>
   <AccessPoints>
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
   </AccessPoints>
   <NextToken>string</NextToken>
</ListMultiRegionAccessPointsResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListMultiRegionAccessPointsResult](#API_control_ListMultiRegionAccessPoints_ResponseSyntax)**

Root level tag for the ListMultiRegionAccessPointsResult parameters.

Required: Yes

**[AccessPoints](#API_control_ListMultiRegionAccessPoints_ResponseSyntax)**

The list of Multi-Region Access Points associated with the user.

Type: Array of [MultiRegionAccessPointReport](api-control-multiregionaccesspointreport.md) data types

**[NextToken](#API_control_ListMultiRegionAccessPoints_ResponseSyntax)**

If the specified bucket has more Multi-Region Access Points than can be returned in one call to this
action, this field contains a continuation token. You can use this token tin subsequent
calls to this action to retrieve additional Multi-Region Access Points.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/listmultiregionaccesspoints.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/listmultiregionaccesspoints.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListJobs

ListRegionalBuckets

All content copied from https://docs.aws.amazon.com/.
