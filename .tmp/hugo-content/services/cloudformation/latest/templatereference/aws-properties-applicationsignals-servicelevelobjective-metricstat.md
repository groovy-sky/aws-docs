This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective MetricStat

This structure defines the metric to be used as the service level indicator, along with the statistics, period, and unit.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Metric" : Metric,
  "Period" : Integer,
  "Stat" : String,
  "Unit" : String
}

```

### YAML

```yaml

  Metric:
    Metric
  Period: Integer
  Stat: String
  Unit: String

```

## Properties

`Metric`

The metric to use as the service level indicator, including the metric name, namespace, and dimensions.

_Required_: Yes

_Type_: [Metric](aws-properties-applicationsignals-servicelevelobjective-metric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The granularity, in seconds, to be used for the metric. For metrics with regular resolution, a period can
be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected
at intervals of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of 60. High-resolution metrics
are those metrics stored by a `PutMetricData` call that includes a `StorageResolution` of 1 second.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stat`

The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic. For more information about statistics,
see [CloudWatch statistics definitions](../../../amazoncloudwatch/latest/monitoring/statistics-definitions.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

If you omit `Unit` then all data that was collected with any unit is returned, along with the corresponding units that were specified
when the data was reported to CloudWatch. If you specify a unit, the operation returns only data that was collected with that unit specified.
If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.

_Required_: No

_Type_: String

_Allowed values_: `Microseconds | Milliseconds | Seconds | Bytes | Kilobytes | Megabytes | Gigabytes | Terabytes | Bits | Kilobits | Megabits | Gigabits | Terabits | Percent | Count | Bytes/Second | Kilobytes/Second | Megabytes/Second | Gigabytes/Second | Terabytes/Second | Bits/Second | Kilobits/Second | Megabits/Second | Gigabits/Second | Terabits/Second | Count/Second | None`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDataQuery

MonitoredRequestCountMetric

All content copied from https://docs.aws.amazon.com/.
