# ModifyReportDefinition

Allows you to programmatically update your report preferences.

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
   "ReportName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ReportDefinition](#API_cur_ModifyReportDefinition_RequestSyntax)**

The definition of AWS Cost and Usage Report. You can specify the report name,
time unit, report format, compression format, S3 bucket, additional artifacts, and schema
elements in the definition.

Type: [ReportDefinition](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_cur_ReportDefinition.html) object

Required: Yes

**[ReportName](#API_cur_ModifyReportDefinition_RequestSyntax)**

The name of the report that you want to create. The name must be unique,
is case sensitive, and can't include spaces.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[0-9A-Za-z!\-_.*\'()]+`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalErrorException**

An error on the server occurred during the processing of your request. Try again later.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 500

**ValidationException**

The input fails to satisfy the constraints specified by an AWS service.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/cur-2017-01-06/ModifyReportDefinition)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/cur-2017-01-06/ModifyReportDefinition)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListTagsForResource

PutReportDefinition
