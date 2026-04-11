# AnomalyDetector

Contains information about one anomaly detector in the account.

## Contents

**anomalyDetectorArn**

The ARN of the anomaly detector.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `[\w#+=/:,.@-]*`

Required: No

**anomalyDetectorStatus**

Specifies the current status of the anomaly detector. To pause an anomaly detector, use
the `enabled` parameter in the [UpdateLogAnomalyDetector](api-updateloganomalydetector.md) operation.

Type: String

Valid Values: `INITIALIZING | TRAINING | ANALYZING | FAILED | DELETED | PAUSED`

Required: No

**anomalyVisibilityTime**

The number of days used as the life cycle of anomalies. After this time, anomalies are
automatically baselined and the anomaly detector model will treat new occurrences of similar
event as normal.

Type: Long

Valid Range: Minimum value of 7. Maximum value of 90.

Required: No

**creationTimeStamp**

The date and time when this anomaly detector was created.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**detectorName**

The name of the anomaly detector.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**evaluationFrequency**

Specifies how often the anomaly detector runs and look for anomalies.

Type: String

Valid Values: `ONE_MIN | FIVE_MIN | TEN_MIN | FIFTEEN_MIN | THIRTY_MIN | ONE_HOUR`

Required: No

**filterPattern**

A symbolic description of how CloudWatch Logs should interpret the data in each log
event. For example, a log event can contain timestamps, IP addresses, strings, and so on. You
use the filter pattern to specify what to look for in the log event message.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**kmsKeyId**

The ARN of the AWS KMS key assigned to this anomaly detector, if any.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**lastModifiedTimeStamp**

The date and time when this anomaly detector was most recently modified.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logGroupArnList**

A list of the ARNs of the log groups that this anomaly detector watches.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/anomalydetector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/anomalydetector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/anomalydetector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Anomaly

ConfigurationTemplate

All content copied from https://docs.aws.amazon.com/.
