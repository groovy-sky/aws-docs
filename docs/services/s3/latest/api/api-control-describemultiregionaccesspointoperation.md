# DescribeMultiRegionAccessPointOperation

###### Note

This operation is not supported by directory buckets.

Retrieves the status of an asynchronous request to manage a Multi-Region Access Point. For more information
about managing Multi-Region Access Points and how asynchronous requests work, see [Using Multi-Region Access Points](../userguide/mrapoperations.md) in the _Amazon S3 User Guide_.

The following actions are related to `GetMultiRegionAccessPoint`:

- [CreateMultiRegionAccessPoint](api-control-createmultiregionaccesspoint.md)

- [DeleteMultiRegionAccessPoint](api-control-deletemultiregionaccesspoint.md)

- [GetMultiRegionAccessPoint](api-control-getmultiregionaccesspoint.md)

- [ListMultiRegionAccessPoints](api-control-listmultiregionaccesspoints.md)

## Request Syntax

```nohighlight

GET /v20180820/async-requests/mrap/request_token+ HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[request\_token](#API_control_DescribeMultiRegionAccessPointOperation_RequestSyntax)**

The request token associated with the request you want to know about. This request token
is returned as part of the response when you make an asynchronous request. You provide this
token to query about the status of the asynchronous action.

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `arn:.+`

Required: Yes

**[x-amz-account-id](#API_control_DescribeMultiRegionAccessPointOperation_RequestSyntax)**

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
<DescribeMultiRegionAccessPointOperationResult>
   <AsyncOperation>
      <CreationTime>timestamp</CreationTime>
      <Operation>string</Operation>
      <RequestParameters>
         <CreateMultiRegionAccessPointRequest>
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
         </CreateMultiRegionAccessPointRequest>
         <DeleteMultiRegionAccessPointRequest>
            <Name>string</Name>
         </DeleteMultiRegionAccessPointRequest>
         <PutMultiRegionAccessPointPolicyRequest>
            <Name>string</Name>
            <Policy>string</Policy>
         </PutMultiRegionAccessPointPolicyRequest>
      </RequestParameters>
      <RequestStatus>string</RequestStatus>
      <RequestTokenARN>string</RequestTokenARN>
      <ResponseDetails>
         <ErrorDetails>
            <Code>string</Code>
            <Message>string</Message>
            <RequestId>string</RequestId>
            <Resource>string</Resource>
         </ErrorDetails>
         <MultiRegionAccessPointDetails>
            <Regions>
               <Region>
                  <Name>string</Name>
                  <RequestStatus>string</RequestStatus>
               </Region>
            </Regions>
         </MultiRegionAccessPointDetails>
      </ResponseDetails>
   </AsyncOperation>
</DescribeMultiRegionAccessPointOperationResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[DescribeMultiRegionAccessPointOperationResult](#API_control_DescribeMultiRegionAccessPointOperation_ResponseSyntax)**

Root level tag for the DescribeMultiRegionAccessPointOperationResult parameters.

Required: Yes

**[AsyncOperation](#API_control_DescribeMultiRegionAccessPointOperation_ResponseSyntax)**

A container element containing the details of the asynchronous operation.

Type: [AsyncOperation](api-control-asyncoperation.md) data type

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/describemultiregionaccesspointoperation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeJob

DissociateAccessGrantsIdentityCenter

All content copied from https://docs.aws.amazon.com/.
