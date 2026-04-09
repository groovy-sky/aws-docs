# MetricTransformation

Indicates how to transform ingested log events to metric data in a CloudWatch
metric.

## Contents

**metricName**

The name of the CloudWatch metric.

Type: String

Length Constraints: Maximum length of 255.

Pattern: `[^:*$]*`

Required: Yes

**metricNamespace**

A custom namespace to contain your metric in CloudWatch. Use namespaces to group
together metrics that are similar. For more information, see [Namespaces](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md#Namespace).

Type: String

Length Constraints: Maximum length of 255.

Pattern: `[^:*$]*`

Required: Yes

**metricValue**

The value to publish to the CloudWatch metric when a filter pattern matches a log
event.

Type: String

Length Constraints: Maximum length of 100.

Required: Yes

**defaultValue**

(Optional) The value to emit when a filter pattern does not match a log event. This
value can be null.

Type: Double

Required: No

**dimensions**

The fields to use as dimensions for the metric. One metric filter can include as many as
three dimensions.

###### Important

Metrics extracted from log events are charged as custom metrics. To prevent unexpected
high charges, do not specify high-cardinality fields such as `IPAddress` or
`requestID` as dimensions. Each different value found for a dimension is
treated as a separate metric and accrues charges as a separate custom metric.

CloudWatch Logs disables a metric filter if it generates 1000 different name/value
pairs for your specified dimensions within a certain amount of time. This helps to prevent
accidental high charges.

You can also set up a billing alarm to alert you if your charges are higher than
expected. For more information, see [Creating a Billing Alarm to Monitor Your Estimated AWS Charges](../../../../services/amazoncloudwatch/latest/monitoring/monitor-estimated-charges-with-cloudwatch.md).

Type: String to string map

Key Length Constraints: Maximum length of 255.

Value Length Constraints: Maximum length of 255.

Required: No

**unit**

The unit to assign to the metric. If you omit this, the unit is set as
`None`.

Type: String

Valid Values: `Seconds | Microseconds | Milliseconds | Bytes | Kilobytes | Megabytes | Gigabytes | Terabytes | Bits | Kilobits | Megabits | Gigabits | Terabits | Percent | Count | Bytes/Second | Kilobytes/Second | Megabytes/Second | Gigabytes/Second | Terabytes/Second | Bits/Second | Kilobits/Second | Megabits/Second | Gigabits/Second | Terabits/Second | Count/Second | None`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/metrictransformation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/metrictransformation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/metrictransformation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricFilterMatchRecord

MoveKeyEntry

All content copied from https://docs.aws.amazon.com/.
