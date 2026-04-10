This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::Alarm MetricDataQuery

The `MetricDataQuery` property type specifies the metric data to return, and whether this call is
just retrieving a batch set of data for one metric, or is performing a math expression on metric data.

Any expression used must return a single time series. For more information, see [Metric Math Syntax and Functions](../../../amazoncloudwatch/latest/monitoring/using-metric-math.md#metric-math-syntax) in the _Amazon CloudWatch User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "Expression" : String,
  "Id" : String,
  "Label" : String,
  "MetricStat" : MetricStat,
  "Period" : Integer,
  "ReturnData" : Boolean
}

```

### YAML

```yaml

  AccountId: String
  Expression: String
  Id: String
  Label: String
  MetricStat:
    MetricStat
  Period: Integer
  ReturnData: Boolean

```

## Properties

`AccountId`

The ID of the account where the metrics are located, if this is a cross-account alarm.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The math expression to be performed on the returned data, if this object is performing a math expression. This expression
can use the `Id` of the other metrics to refer to those metrics, and can also use the `Id` of other
expressions to use the result of those expressions. For more information about metric math expressions, see
[Metric Math Syntax and Functions](../../../amazoncloudwatch/latest/monitoring/using-metric-math.md#metric-math-syntax) in the
_Amazon CloudWatch User Guide_.

Within each MetricDataQuery object, you must specify either
`Expression` or `MetricStat` but not both.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

A short name used to tie this object to the results in the response. This name must be
unique within a single call to `GetMetricData`. If you are performing math
expressions on this set of data, this name represents that data and can serve as a
variable in the mathematical expression. The valid characters are letters, numbers, and
underscore. The first character must be a lowercase letter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know
what the value represents. If the metric or expression is shown in a CloudWatch dashboard widget, the label is shown. If `Label` is omitted, CloudWatch
generates a default.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricStat`

The metric to be returned, along with statistics, period, and units. Use this
parameter only if this object is retrieving a metric and not performing a math
expression on returned data.

Within one MetricDataQuery object, you must specify either `Expression` or
`MetricStat` but not both.

_Required_: No

_Type_: [MetricStat](aws-properties-cloudwatch-alarm-metricstat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The granularity, in seconds, of the returned data points. For metrics with regular
resolution, a period can be as short as one minute (60 seconds) and must be a multiple
of 60. For high-resolution metrics that are collected at intervals of less than one
minute, the period can be 1, 5, 10, 20, 30, 60, or any multiple of 60. High-resolution
metrics are those metrics stored by a `PutMetricData` operation that includes
a `StorageResolution of 1 second`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReturnData`

This option indicates whether to return the
timestamps and raw data values of this metric.

When you create an alarm based on a metric math expression, specify `True` for
this value for only the one math expression that the alarm is based on. You must specify
`False` for `ReturnData` for all the other metrics and expressions
used in the alarm.

This field is required.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Metric

MetricStat

All content copied from https://docs.aws.amazon.com/.
