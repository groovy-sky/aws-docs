# ListAccessPointsForObjectLambda

###### Note

This operation is not supported by directory buckets.

Returns some or all (up to 1,000) access points associated with the Object Lambda Access Point per call. If there
are more access points than what can be returned in one call, the response will include a
continuation token that you can use to list the additional access points.

The following actions are related to
`ListAccessPointsForObjectLambda`:

- [CreateAccessPointForObjectLambda](api-control-createaccesspointforobjectlambda.md)

- [DeleteAccessPointForObjectLambda](api-control-deleteaccesspointforobjectlambda.md)

- [GetAccessPointForObjectLambda](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointForObjectLambda.html)

## Request Syntax

```nohighlight

GET /v20180820/accesspointforobjectlambda?maxResults=MaxResults&nextToken=NextToken HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_control_ListAccessPointsForObjectLambda_RequestSyntax)**

The maximum number of access points that you want to include in the list. The response may
contain fewer access points but will never contain more. If there are more than this number of
access points, then the response will include a continuation token in the `NextToken`
field that you can use to retrieve the next page of access points.

Valid Range: Minimum value of 0. Maximum value of 1000.

**[nextToken](#API_control_ListAccessPointsForObjectLambda_RequestSyntax)**

If the list has more access points than can be returned in one call to this API, this field
contains a continuation token that you can provide in subsequent calls to this API to
retrieve additional access points.

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[x-amz-account-id](#API_control_ListAccessPointsForObjectLambda_RequestSyntax)**

The account ID for the account that owns the specified Object Lambda Access Point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListAccessPointsForObjectLambdaResult>
   <ObjectLambdaAccessPointList>
      <ObjectLambdaAccessPoint>
         <Alias>
            <Status>string</Status>
            <Value>string</Value>
         </Alias>
         <Name>string</Name>
         <ObjectLambdaAccessPointArn>string</ObjectLambdaAccessPointArn>
      </ObjectLambdaAccessPoint>
   </ObjectLambdaAccessPointList>
   <NextToken>string</NextToken>
</ListAccessPointsForObjectLambdaResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListAccessPointsForObjectLambdaResult](#API_control_ListAccessPointsForObjectLambda_ResponseSyntax)**

Root level tag for the ListAccessPointsForObjectLambdaResult parameters.

Required: Yes

**[NextToken](#API_control_ListAccessPointsForObjectLambda_ResponseSyntax)**

If the list has more access points than can be returned in one call to this API, this field
contains a continuation token that you can provide in subsequent calls to this API to
retrieve additional access points.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[ObjectLambdaAccessPointList](#API_control_ListAccessPointsForObjectLambda_ResponseSyntax)**

Returns list of Object Lambda Access Points.

Type: Array of [ObjectLambdaAccessPoint](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ObjectLambdaAccessPoint.html) data types

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3control-2018-08-20/ListAccessPointsForObjectLambda)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/ListAccessPointsForObjectLambda)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListAccessPointsForDirectoryBuckets

ListCallerAccessGrants
