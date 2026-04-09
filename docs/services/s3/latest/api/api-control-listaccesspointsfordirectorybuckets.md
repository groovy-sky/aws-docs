# ListAccessPointsForDirectoryBuckets

Returns a list of the access points that are owned by the AWS account and that are associated with the specified directory bucket.

To list access points for general purpose buckets, see [ListAccesspoints](api-control-listaccesspoints.md).

To use this operation, you must have the permission to perform the
`s3express:ListAccessPointsForDirectoryBuckets` action.

For information about REST API errors, see [REST error responses](errorresponses.md#RESTErrorResponses).

## Request Syntax

```nohighlight

GET /v20180820/accesspointfordirectory?directoryBucket=DirectoryBucket&maxResults=MaxResults&nextToken=NextToken HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[directoryBucket](#API_control_ListAccessPointsForDirectoryBuckets_RequestSyntax)**

The name of the directory bucket associated with the access points you want to list.

Length Constraints: Minimum length of 3. Maximum length of 255.

**[maxResults](#API_control_ListAccessPointsForDirectoryBuckets_RequestSyntax)**

The maximum number of access points that you would like returned in the `ListAccessPointsForDirectoryBuckets` response. If the directory bucket is associated with more than this number of access points, the results include the pagination token `NextToken`. Make another call using the `NextToken` to retrieve more results.

Valid Range: Minimum value of 0. Maximum value of 1000.

**[nextToken](#API_control_ListAccessPointsForDirectoryBuckets_RequestSyntax)**

If `NextToken` is returned, there are more access points available than requested in the `maxResults` value. The value of `NextToken` is a
unique pagination token for each page. Make the call again using the returned token to
retrieve the next page. Keep all other arguments unchanged. Each pagination token expires
after 24 hours.

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[x-amz-account-id](#API_control_ListAccessPointsForDirectoryBuckets_RequestSyntax)**

The AWS account ID that owns the access points.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListAccessPointsForDirectoryBucketsResult>
   <AccessPointList>
      <AccessPoint>
         <AccessPointArn>string</AccessPointArn>
         <Alias>string</Alias>
         <Bucket>string</Bucket>
         <BucketAccountId>string</BucketAccountId>
         <DataSourceId>string</DataSourceId>
         <DataSourceType>string</DataSourceType>
         <Name>string</Name>
         <NetworkOrigin>string</NetworkOrigin>
         <VpcConfiguration>
            <VpcId>string</VpcId>
         </VpcConfiguration>
      </AccessPoint>
   </AccessPointList>
   <NextToken>string</NextToken>
</ListAccessPointsForDirectoryBucketsResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListAccessPointsForDirectoryBucketsResult](#API_control_ListAccessPointsForDirectoryBuckets_ResponseSyntax)**

Root level tag for the ListAccessPointsForDirectoryBucketsResult parameters.

Required: Yes

**[AccessPointList](#API_control_ListAccessPointsForDirectoryBuckets_ResponseSyntax)**

Contains identification and configuration information for one or more access points associated with the directory bucket.

Type: Array of [AccessPoint](api-control-accesspoint.md) data types

**[NextToken](#API_control_ListAccessPointsForDirectoryBuckets_ResponseSyntax)**

If `NextToken` is returned, there are more access points available than requested in the `maxResults` value. The value of `NextToken` is a
unique pagination token for each page. Make the call again using the returned token to
retrieve the next page. Keep all other arguments unchanged. Each pagination token expires
after 24 hours.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/listaccesspointsfordirectorybuckets.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAccessPoints

ListAccessPointsForObjectLambda

All content copied from https://docs.aws.amazon.com/.
