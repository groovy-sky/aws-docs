# PutReportDefinition

Creates a new report using the description that you provide.

## Request Syntax

```nohighlight

{
   "ReportDefinition": {
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
   },
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ReportDefinition](#API_cur_PutReportDefinition_RequestSyntax)**

Represents the output of the PutReportDefinition operation. The content consists of the detailed
metadata and data file information.

Type: [ReportDefinition](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_cur_ReportDefinition.html) object

Required: Yes

**[Tags](#API_cur_PutReportDefinition_RequestSyntax)**

The tags to be assigned to the report definition resource.

Type: Array of [Tag](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_cur_Tag.html) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: No

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DuplicateReportNameException**

A report with the specified name already exists in the account. Specify a different report name.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 400

**InternalErrorException**

An error on the server occurred during the processing of your request. Try again later.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 500

**ReportLimitReachedException**

This account already has five reports defined. To define a new report, you must delete an existing report.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified report ( `ReportName`) in the request doesn't exist.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 400

**ValidationException**

The input fails to satisfy the constraints specified by an AWS service.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 400

## Examples

### The following is a sample request of the PutReportDefinition operation.

This example illustrates one usage of PutReportDefinition.

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
X-Amz-Target: AWSOrigamiServiceGateway.PutReportDefinition
{
        "ReportDefinition": {
            "ReportName": "ExampleReport",
            "TimeUnit": "DAILY",
            "Format": "textORcsv",
            "Compression": "ZIP",
            "AdditionalSchemaElements": [
                "RESOURCES"
            ],
            "S3Bucket": "example-s3-bucket",
            "S3Prefix": "exampleprefix",
            "S3Region": "us-east-1",
            "AdditionalArtifacts": [
                "REDSHIFT",
                "QUICKSIGHT"
             },
        "Tags": [
          {
            "Key": "key-1",
            "Value": "value-1"
          }
        ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/cur-2017-01-06/PutReportDefinition)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/cur-2017-01-06/PutReportDefinition)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModifyReportDefinition

TagResource
