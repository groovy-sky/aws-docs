# DescribeAlarmsForMetric

Retrieves the alarms for the specified metric. To filter the results, specify a
statistic, period, or unit.

This operation retrieves only standard alarms that are based on the specified
metric. It does not return alarms based on math expressions that use the specified
metric, or composite alarms that use the specified metric.

## Request Parameters

**Dimensions**

The dimensions associated with the metric. If the metric has any associated
dimensions, you must specify them in order for the call to succeed.

Type: Array of [Dimension](api-dimension.md) objects

Array Members: Maximum number of 30 items.

Required: No

**ExtendedStatistic**

The percentile statistic for the metric. Specify a value between p0.0 and
p100.

Type: String

Required: No

**MetricName**

The name of the metric.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**Namespace**

The namespace of the metric.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[^:].*`

Required: Yes

**Period**

The period, in seconds, over which the statistic is applied.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**Statistic**

The statistic for the metric, other than percentiles. For percentile statistics,
use `ExtendedStatistics`.

Type: String

Valid Values: `SampleCount | Average | Sum | Minimum | Maximum`

Required: No

**Unit**

The unit for the metric.

Type: String

Valid Values: `Seconds | Microseconds | Milliseconds | Bytes | Kilobytes | Megabytes | Gigabytes | Terabytes | Bits | Kilobits | Megabits | Gigabits | Terabits | Percent | Count | Bytes/Second | Kilobytes/Second | Megabytes/Second | Gigabytes/Second | Terabytes/Second | Bits/Second | Kilobits/Second | Megabits/Second | Gigabits/Second | Terabits/Second | Count/Second | None`

Required: No

## Response Elements

The following element is returned by the service.

**MetricAlarms**

The information for each alarm with the specified metric.

Type: Array of [MetricAlarm](api-metricalarm.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/describealarmsformetric.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/describealarmsformetric.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeAlarms

DescribeAnomalyDetectors

All content copied from https://docs.aws.amazon.com/.
