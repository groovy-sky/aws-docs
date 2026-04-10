This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::Alarm MetricStat

This structure defines the metric to be returned, along with the statistics, period, and units.

`MetricStat` is a property of the
[MetricDataQuery](../userguide/aws-properties-cloudwatch-alarm-metricdataquery.md) property type.

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

The metric to return, including the metric name, namespace, and dimensions.

_Required_: Yes

_Type_: [Metric](aws-properties-cloudwatch-alarm-metric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The granularity, in seconds, of the returned data points. For metrics with regular
resolution, a period can be as short as one minute (60 seconds) and must be a multiple
of 60. For high-resolution metrics that are collected at intervals of less than one
minute, the period can be 1, 5, 10, 20, 30, 60, or any multiple of 60. High-resolution
metrics are those metrics stored by a `PutMetricData` call that includes a
`StorageResolution` of 1 second.

If the `StartTime` parameter specifies a time stamp that is greater than
3 hours ago, you must specify the period as follows or no data points in that time range
is returned:

- Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds
(1 minute).

- Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5
minutes).

- Start time greater than 63 days ago - Use a multiple of 3600 seconds (1
hour).

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stat`

The statistic to return. It can include any CloudWatch statistic or extended statistic.
For a list of valid values, see the table in [Statistics](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md#Statistic) in the _Amazon CloudWatch User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit to use for the returned data points.

Valid values are: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes,
Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count,
Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second,
Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.

_Required_: No

_Type_: String

_Allowed values_: `Seconds | Microseconds | Milliseconds | Bytes | Kilobytes | Megabytes | Gigabytes | Terabytes | Bits | Kilobits | Megabits | Gigabits | Terabits | Percent | Count | Bytes/Second | Kilobytes/Second | Megabytes/Second | Gigabytes/Second | Terabytes/Second | Bits/Second | Kilobits/Second | Megabits/Second | Gigabits/Second | Terabits/Second | Count/Second | None`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDataQuery

Tag

All content copied from https://docs.aws.amazon.com/.
