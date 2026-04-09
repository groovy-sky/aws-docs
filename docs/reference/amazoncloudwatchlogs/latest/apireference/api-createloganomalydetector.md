# CreateLogAnomalyDetector

Creates an _anomaly detector_ that regularly scans one or more log
groups and look for patterns and anomalies in the logs.

An anomaly detector can help surface issues by automatically discovering anomalies in your
log event traffic. An anomaly detector uses machine learning algorithms to scan log events and
find _patterns_. A pattern is a shared text structure that recurs among
your log fields. Patterns provide a useful tool for analyzing large sets of logs because a
large number of log events can often be compressed into a few patterns.

The anomaly detector uses pattern recognition to find `anomalies`, which are
unusual log events. It uses the `evaluationFrequency` to compare current log events
and patterns with trained baselines.

Fields within a pattern are called _tokens_. Fields that vary within a
pattern, such as a request ID or timestamp, are referred to as _dynamic_
_tokens_ and represented by `<*>`.

The following is an example of a pattern:

`[INFO] Request time: <*> ms`

This pattern represents log events like `[INFO] Request time: 327 ms` and other
similar log events that differ only by the number, in this csse 327. When the pattern is
displayed, the different numbers are replaced by `<*>`

###### Note

Any parts of log events that are masked as sensitive data are not scanned for anomalies.
For more information about masking sensitive data, see [Help protect sensitive log\
data with masking](../../../../services/amazoncloudwatch/latest/logs/mask-sensitive-log-data.md).

## Request Syntax

```nohighlight

{
   "anomalyVisibilityTime": number,
   "detectorName": "string",
   "evaluationFrequency": "string",
   "filterPattern": "string",
   "kmsKeyId": "string",
   "logGroupArnList": [ "string" ],
   "tags": {
      "string" : "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[anomalyVisibilityTime](#API_CreateLogAnomalyDetector_RequestSyntax)**

The number of days to have visibility on an anomaly. After this time period has elapsed
for an anomaly, it will be automatically baselined and the anomaly detector will treat new
occurrences of a similar anomaly as normal. Therefore, if you do not correct the cause of an
anomaly during the time period specified in `anomalyVisibilityTime`, it will be
considered normal going forward and will not be detected as an anomaly.

Type: Long

Valid Range: Minimum value of 7. Maximum value of 90.

Required: No

**[detectorName](#API_CreateLogAnomalyDetector_RequestSyntax)**

A name for this anomaly detector.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[evaluationFrequency](#API_CreateLogAnomalyDetector_RequestSyntax)**

Specifies how often the anomaly detector is to run and look for anomalies. Set this value
according to the frequency that the log group receives new logs. For example, if the log group
receives new log events every 10 minutes, then 15 minutes might be a good setting for
`evaluationFrequency` .

Type: String

Valid Values: `ONE_MIN | FIVE_MIN | TEN_MIN | FIFTEEN_MIN | THIRTY_MIN | ONE_HOUR`

Required: No

**[filterPattern](#API_CreateLogAnomalyDetector_RequestSyntax)**

You can use this parameter to limit the anomaly detection model to examine only log events
that match the pattern you specify here. For more information, see [Filter and Pattern\
Syntax](../../../../services/amazoncloudwatch/latest/logs/filterandpatternsyntax.md).

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**[kmsKeyId](#API_CreateLogAnomalyDetector_RequestSyntax)**

Optionally assigns a AWS KMS key to secure this anomaly detector and its
findings. If a key is assigned, the anomalies found and the model used by this detector are
encrypted at rest with the key. If a key is assigned to an anomaly detector, a user must have
permissions for both this key and for the anomaly detector to retrieve information about the
anomalies that it finds.

Make sure the value provided is a valid AWS KMS key ARN. For more information
about using a AWS KMS key and to see the required IAM policy, see
[Use a AWS KMS key with an anomaly detector](../../../../services/amazoncloudwatch/latest/logs/logsanomalydetection-kms.md).

Type: String

Length Constraints: Maximum length of 256.

Pattern: `^arn:aws[a-z\-]*:kms:[-a-z0-9]*:[0-9]*:key/[-a-z0-9]*$`

Required: No

**[logGroupArnList](#API_CreateLogAnomalyDetector_RequestSyntax)**

An array containing the ARN of the log group that this anomaly detector will watch. You
can specify only one log group ARN.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[tags](#API_CreateLogAnomalyDetector_RequestSyntax)**

An optional list of key-value pairs to associate with the resource.

For more information about tagging, see [Tagging AWS resources](../../../../general/latest/gr/aws-tagging.md)

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: No

## Response Syntax

```nohighlight

{
   "anomalyDetectorArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[anomalyDetectorArn](#API_CreateLogAnomalyDetector_ResponseSyntax)**

The ARN of the log anomaly detector that you just created.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[\w#+=/:,.@-]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**LimitExceededException**

You have reached the maximum number of resources that can be created.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/createloganomalydetector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/createloganomalydetector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateImportTask

CreateLogGroup

All content copied from https://docs.aws.amazon.com/.
