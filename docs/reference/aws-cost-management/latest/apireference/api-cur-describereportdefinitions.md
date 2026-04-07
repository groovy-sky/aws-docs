# DescribeReportDefinitions

Lists the AWS Cost and Usage Report available to this account.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_cur_DescribeReportDefinitions_RequestSyntax)**

The maximum number of results that AWS returns for the operation.

Type: Integer

Valid Range: Fixed value of 5.

Required: No

**[NextToken](#API_cur_DescribeReportDefinitions_RequestSyntax)**

A generic string.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[A-Za-z0-9_\.\-=]*`

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "ReportDefinitions": [
      {
         "AdditionalArtifacts": [ "string" ],
         "AdditionalSchemaElements": [ "string" ],
         "BillingViewArn": "string",
         "Compression": "string",
         "Format": "string",
         "RefreshClosedReports": boolean,
         "ReportName": "string",
         "ReportStatus": {
            "lastDelivery": "string",
            "lastStatus": "string"
         },
         "ReportVersioning": "string",
         "S3Bucket": "string",
         "S3Prefix": "string",
         "S3Region": "string",
         "TimeUnit": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_cur_DescribeReportDefinitions_ResponseSyntax)**

A generic string.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[A-Za-z0-9_\.\-=]*`

**[ReportDefinitions](#API_cur_DescribeReportDefinitions_ResponseSyntax)**

An AWS Cost and Usage Report list owned by the account.

Type: Array of [ReportDefinition](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_cur_ReportDefinition.html) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalErrorException**

An error on the server occurred during the processing of your request. Try again later.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 500

## Examples

### The following is a sample request and response of the DescribeReportDefinitions operation.

This example illustrates one usage of DescribeReportDefinitions.

#### Sample Request

```

POST / HTTP/1.1
Host: api.cur.<region>.<domain>
x-amz-Date: <Date>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=contenttype;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid,Signature=<Signature>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: AWSOrigamiServiceGateway.DescribeReportDefinitions
{
        "MaxResults": 5
}

```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
        "ReportDefinitions": [
        {
            "AdditionalArtifacts": ["QUICKSIGHT"],
            "AdditionalSchemaElements": ["RESOURCES"],
            "Compression": "GZIP",
            "Format": "textORcsv",
            "ReportName": "ExampleReport",
            "S3Bucket": "example-s3-bucket",
            "S3Prefix": "exampleprefix",
            "S3Region": "us-east-1",
            "TimeUnit": "HOURLY"
        },
        {
            "AdditionalArtifacts": ["QUICKSIGHT"],
            "AdditionalSchemaElements": ["RESOURCES"],
            "Compression": "GZIP",
            "Format": "textORcsv",
            "ReportName": "ExampleReport2",
            "S3Bucket": "example-s3-bucket",
            "S3Prefix": "exampleprefix",
            "S3Region": "us-east-1",
            "TimeUnit": "HOURLY"
            "ReportStatus": {
                "lastDelivery": "20191102T054923Z",
                "lastStatus": "SUCCESS"
        }
      }
    ]
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/cur-2017-01-06/DescribeReportDefinitions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/cur-2017-01-06/DescribeReportDefinitions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteReportDefinition

ListTagsForResource
