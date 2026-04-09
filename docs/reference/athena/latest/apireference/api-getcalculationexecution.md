# GetCalculationExecution

Describes a previously submitted calculation execution.

## Request Syntax

```nohighlight

{
   "CalculationExecutionId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CalculationExecutionId](#API_GetCalculationExecution_RequestSyntax)**

The calculation execution UUID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 36.

Required: Yes

## Response Syntax

```nohighlight

{
   "CalculationExecutionId": "string",
   "Description": "string",
   "Result": {
      "ResultS3Uri": "string",
      "ResultType": "string",
      "StdErrorS3Uri": "string",
      "StdOutS3Uri": "string"
   },
   "SessionId": "string",
   "Statistics": {
      "DpuExecutionInMillis": number,
      "Progress": "string"
   },
   "Status": {
      "CompletionDateTime": number,
      "State": "string",
      "StateChangeReason": "string",
      "SubmissionDateTime": number
   },
   "WorkingDirectory": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[CalculationExecutionId](#API_GetCalculationExecution_ResponseSyntax)**

The calculation execution UUID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 36.

**[Description](#API_GetCalculationExecution_ResponseSyntax)**

The description of the calculation execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**[Result](#API_GetCalculationExecution_ResponseSyntax)**

Contains result information. This field is populated only if the calculation is
completed.

Type: [CalculationResult](api-calculationresult.md) object

**[SessionId](#API_GetCalculationExecution_ResponseSyntax)**

The session ID that the calculation ran in.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

**[Statistics](#API_GetCalculationExecution_ResponseSyntax)**

Contains information about the data processing unit (DPU) execution time and progress.
This field is populated only when statistics are available.

Type: [CalculationStatistics](api-calculationstatistics.md) object

**[Status](#API_GetCalculationExecution_ResponseSyntax)**

Contains information about the status of the calculation.

Type: [CalculationStatus](api-calculationstatus.md) object

**[WorkingDirectory](#API_GetCalculationExecution_ResponseSyntax)**

The Amazon S3 location in which calculation results are stored.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `^(https|s3|S3)://([^/]+)/?(.*)$`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource, such as a workgroup, was not found.

**ResourceName**

The name of the Amazon resource.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/getcalculationexecution.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/getcalculationexecution.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExportNotebook

GetCalculationExecutionCode

All content copied from https://docs.aws.amazon.com/.
