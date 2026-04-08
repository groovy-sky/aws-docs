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

Type: [ReportDefinition](api-cur-reportdefinition.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cur-2017-01-06/modifyreportdefinition.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cur-2017-01-06/modifyreportdefinition.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListTagsForResource

PutReportDefinition
