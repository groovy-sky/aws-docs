# UpdateLogAnomalyDetector

Updates an existing log anomaly detector.

## Request Syntax

```nohighlight

{
   "anomalyDetectorArn": "string",
   "anomalyVisibilityTime": number,
   "enabled": boolean,
   "evaluationFrequency": "string",
   "filterPattern": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[anomalyDetectorArn](#API_UpdateLogAnomalyDetector_RequestSyntax)**

The ARN of the anomaly detector that you want to update.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[anomalyVisibilityTime](#API_UpdateLogAnomalyDetector_RequestSyntax)**

The number of days to use as the life cycle of anomalies. After this time, anomalies are
automatically baselined and the anomaly detector model will treat new occurrences of similar
event as normal. Therefore, if you do not correct the cause of an anomaly during this time, it
will be considered normal going forward and will not be detected.

Type: Long

Valid Range: Minimum value of 7. Maximum value of 90.

Required: No

**[enabled](#API_UpdateLogAnomalyDetector_RequestSyntax)**

Use this parameter to pause or restart the anomaly detector.

Type: Boolean

Required: Yes

**[evaluationFrequency](#API_UpdateLogAnomalyDetector_RequestSyntax)**

Specifies how often the anomaly detector runs and look for anomalies. Set this value
according to the frequency that the log group receives new logs. For example, if the log group
receives new log events every 10 minutes, then setting `evaluationFrequency` to
`FIFTEEN_MIN` might be appropriate.

Type: String

Valid Values: `ONE_MIN | FIVE_MIN | TEN_MIN | FIFTEEN_MIN | THIRTY_MIN | ONE_HOUR`

Required: No

**[filterPattern](#API_UpdateLogAnomalyDetector_RequestSyntax)**

A symbolic description of how CloudWatch Logs should interpret the data in each log
event. For example, a log event can contain timestamps, IP addresses, strings, and so on. You
use the filter pattern to specify what to look for in the log event message.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/updateloganomalydetector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/updateloganomalydetector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateDeliveryConfiguration

UpdateLookupTable

All content copied from https://docs.aws.amazon.com/.
