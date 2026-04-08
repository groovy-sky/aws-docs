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

Type: [ReportDefinition](api-cur-reportdefinition.md) object

Required: Yes

**[Tags](#API_cur_PutReportDefinition_RequestSyntax)**

The tags to be assigned to the report definition resource.

Type: Array of [Tag](api-cur-tag.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cur-2017-01-06/putreportdefinition.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cur-2017-01-06/putreportdefinition.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyReportDefinition

TagResource
