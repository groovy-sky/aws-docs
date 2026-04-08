# UpdateAnomaly

Use this operation to _suppress_ anomaly detection for a specified
anomaly or pattern. If you suppress an anomaly, CloudWatch Logs won't report new
occurrences of that anomaly and won't update that anomaly with new data. If you suppress a
pattern, CloudWatch Logs won't report any anomalies related to that pattern.

You must specify either `anomalyId` or `patternId`, but you can't
specify both parameters in the same operation.

If you have previously used this operation to suppress detection of a pattern or anomaly,
you can use it again to cause CloudWatch Logs to end the suppression. To do this, use this
operation and specify the anomaly or pattern to stop suppressing, and omit the
`suppressionType` and `suppressionPeriod` parameters.

## Request Syntax

```nohighlight

{
   "anomalyDetectorArn": "string",
   "anomalyId": "string",
   "baseline": boolean,
   "patternId": "string",
   "suppressionPeriod": {
      "suppressionUnit": "string",
      "value": number
   },
   "suppressionType": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[anomalyDetectorArn](#API_UpdateAnomaly_RequestSyntax)**

The ARN of the anomaly detector that this operation is to act on.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[anomalyId](#API_UpdateAnomaly_RequestSyntax)**

If you are suppressing or unsuppressing an anomaly, specify its unique ID here. You can
find anomaly IDs by using the [ListAnomalies](api-listanomalies.md)
operation.

Type: String

Length Constraints: Fixed length of 36.

Required: No

**[baseline](#API_UpdateAnomaly_RequestSyntax)**

Set this to `true` to prevent CloudWatch Logs from displaying this behavior
as an anomaly in the future. The behavior is then treated as baseline behavior. However, if
similar but more severe occurrences of this behavior occur in the future, those will still be
reported as anomalies.

The default is `false`

Type: Boolean

Required: No

**[patternId](#API_UpdateAnomaly_RequestSyntax)**

If you are suppressing or unsuppressing an pattern, specify its unique ID here. You can
find pattern IDs by using the [ListAnomalies](api-listanomalies.md)
operation.

Type: String

Length Constraints: Fixed length of 32.

Required: No

**[suppressionPeriod](#API_UpdateAnomaly_RequestSyntax)**

If you are temporarily suppressing an anomaly or pattern, use this structure to specify
how long the suppression is to last.

Type: [SuppressionPeriod](api-suppressionperiod.md) object

Required: No

**[suppressionType](#API_UpdateAnomaly_RequestSyntax)**

Use this to specify whether the suppression to be temporary or infinite. If you specify
`LIMITED`, you must also specify a `suppressionPeriod`. If you specify
`INFINITE`, any value for `suppressionPeriod` is ignored.

Type: String

Valid Values: `LIMITED | INFINITE`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/updateanomaly.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/updateanomaly.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UntagResource

UpdateDeliveryConfiguration
