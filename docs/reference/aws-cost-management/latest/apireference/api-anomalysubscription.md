# AnomalySubscription

An `AnomalySubscription` resource (also referred to as an alert
subscription) sends notifications about specific anomalies that meet an alerting
criteria defined by you.

You can specify the frequency of the alerts and the subscribers to notify.

Anomaly subscriptions can be associated with one or more [`AnomalyMonitor`](api-anomalymonitor.md) resources, and they only send
notifications about anomalies detected by those associated monitors. You can also
configure a threshold to further control which anomalies are included in the
notifications.

Anomalies that don’t exceed the chosen threshold and therefore don’t trigger
notifications from an anomaly subscription will still be available on the console and
from the [`GetAnomalies`](api-getanomalies.md) API.

## Contents

**Frequency**

The frequency that anomaly notifications are sent. Notifications are sent either over
email (for DAILY and WEEKLY frequencies) or SNS (for IMMEDIATE frequency). For more
information, see [Creating an Amazon SNS topic for\
anomaly notifications](../../../../services/cost-management/latest/userguide/ad-sns.md).

Type: String

Valid Values: `DAILY | IMMEDIATE | WEEKLY`

Required: Yes

**MonitorArnList**

A list of cost anomaly monitors.

Type: Array of strings

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+`

Required: Yes

**Subscribers**

A list of subscribers to notify.

Type: Array of [Subscriber](api-subscriber.md) objects

Required: Yes

**SubscriptionName**

The name for the subscription.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: Yes

**AccountId**

Your unique account identifier.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: No

**SubscriptionArn**

The `AnomalySubscription` Amazon Resource Name (ARN).

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: No

**Threshold**

(deprecated)

An absolute dollar value that must be exceeded by the anomaly's total impact (see
[Impact](api-impact.md) for
more details) for an anomaly notification to be generated.

This field has been deprecated. To specify a threshold, use ThresholdExpression.
Continued use of Threshold will be treated as shorthand syntax for a
ThresholdExpression.

One of Threshold or ThresholdExpression is required for this resource. You cannot
specify both.

Type: Double

Valid Range: Minimum value of 0.0.

Required: No

**ThresholdExpression**

An [Expression](api-expression.md)
object used to specify the anomalies that you want to generate alerts for. This supports
dimensions and nested expressions. The supported dimensions are
`ANOMALY_TOTAL_IMPACT_ABSOLUTE` and
`ANOMALY_TOTAL_IMPACT_PERCENTAGE`, corresponding to an anomaly’s
TotalImpact and TotalImpactPercentage, respectively (see [Impact](api-impact.md) for
more details). The supported nested expression types are `AND` and
`OR`. The match option `GREATER_THAN_OR_EQUAL` is required.
Values must be numbers between 0 and 10,000,000,000 in string format.

One of Threshold or ThresholdExpression is required for this resource. You cannot
specify both.

The following are examples of valid ThresholdExpressions:

- Absolute threshold: `{ "Dimensions": { "Key":
                          "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL"
                          ], "Values": [ "100" ] } }`

- Percentage threshold: `{ "Dimensions": { "Key":
                          "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL"
                          ], "Values": [ "100" ] } }`

- `AND` two thresholds together: `{ "And": [ { "Dimensions": {
                          "Key": "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [
                          "GREATER_THAN_OR_EQUAL" ], "Values": [ "100" ] } }, { "Dimensions": { "Key":
                          "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL"
                          ], "Values": [ "100" ] } } ] }`

- `OR` two thresholds together: `{ "Or": [ { "Dimensions": {
                          "Key": "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [
                          "GREATER_THAN_OR_EQUAL" ], "Values": [ "100" ] } }, { "Dimensions": { "Key":
                          "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL"
                          ], "Values": [ "100" ] } } ] }`

Type: [Expression](api-expression.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/anomalysubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/anomalysubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/anomalysubscription.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AnomalyScore

CommitmentPurchaseAnalysisConfiguration
