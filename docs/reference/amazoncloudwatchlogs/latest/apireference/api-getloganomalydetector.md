# GetLogAnomalyDetector

Retrieves information about the log anomaly detector that you specify. The AWS KMS key ARN detected is valid.

## Request Syntax

```nohighlight

{
   "anomalyDetectorArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[anomalyDetectorArn](#API_GetLogAnomalyDetector_RequestSyntax)**

The ARN of the anomaly detector to retrieve information about. You can find the ARNs of
log anomaly detectors in your account by using the [ListLogAnomalyDetectors](api-listloganomalydetectors.md) operation.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

## Response Syntax

```nohighlight

{
   "anomalyDetectorStatus": "string",
   "anomalyVisibilityTime": number,
   "creationTimeStamp": number,
   "detectorName": "string",
   "evaluationFrequency": "string",
   "filterPattern": "string",
   "kmsKeyId": "string",
   "lastModifiedTimeStamp": number,
   "logGroupArnList": [ "string" ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[anomalyDetectorStatus](#API_GetLogAnomalyDetector_ResponseSyntax)**

Specifies whether the anomaly detector is currently active. To change its status, use the
`enabled` parameter in the [UpdateLogAnomalyDetector](api-updateloganomalydetector.md) operation.

Type: String

Valid Values: `INITIALIZING | TRAINING | ANALYZING | FAILED | DELETED | PAUSED`

**[anomalyVisibilityTime](#API_GetLogAnomalyDetector_ResponseSyntax)**

The number of days used as the life cycle of anomalies. After this time, anomalies are
automatically baselined and the anomaly detector model will treat new occurrences of similar
event as normal.

Type: Long

Valid Range: Minimum value of 7. Maximum value of 90.

**[creationTimeStamp](#API_GetLogAnomalyDetector_ResponseSyntax)**

The date and time when this anomaly detector was created.

Type: Long

Valid Range: Minimum value of 0.

**[detectorName](#API_GetLogAnomalyDetector_ResponseSyntax)**

The name of the log anomaly detector

Type: String

Length Constraints: Minimum length of 1.

**[evaluationFrequency](#API_GetLogAnomalyDetector_ResponseSyntax)**

Specifies how often the anomaly detector runs and look for anomalies. Set this value
according to the frequency that the log group receives new logs. For example, if the log group
receives new log events every 10 minutes, then setting `evaluationFrequency` to
`FIFTEEN_MIN` might be appropriate.

Type: String

Valid Values: `ONE_MIN | FIVE_MIN | TEN_MIN | FIFTEEN_MIN | THIRTY_MIN | ONE_HOUR`

**[filterPattern](#API_GetLogAnomalyDetector_ResponseSyntax)**

A symbolic description of how CloudWatch Logs should interpret the data in each log
event. For example, a log event can contain timestamps, IP addresses, strings, and so on. You
use the filter pattern to specify what to look for in the log event message.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

**[kmsKeyId](#API_GetLogAnomalyDetector_ResponseSyntax)**

The ARN of the AWS KMS key assigned to this anomaly detector, if any.

Type: String

Length Constraints: Maximum length of 256.

**[lastModifiedTimeStamp](#API_GetLogAnomalyDetector_ResponseSyntax)**

The date and time when this anomaly detector was most recently modified.

Type: Long

Valid Range: Minimum value of 0.

**[logGroupArnList](#API_GetLogAnomalyDetector_ResponseSyntax)**

An array of structures, where each structure contains the ARN of a log group associated
with this anomaly detector.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getloganomalydetector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getloganomalydetector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetIntegration

GetLogEvents

All content copied from https://docs.aws.amazon.com/.
