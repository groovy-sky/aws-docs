This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AnomalyDetector MetricDataQuery

This structure is used in both `GetMetricData` and
`PutMetricAlarm`. The supported use of this structure is different for
those two operations.

When used in `GetMetricData`, it indicates the metric data to return, and
whether this call is just retrieving a batch set of data for one metric, or is
performing a Metrics Insights query or a math expression. A single
`GetMetricData` call can include up to 500 `MetricDataQuery`
structures.

When used in `PutMetricAlarm`, it enables you to create an alarm based on a
metric math expression. Each `MetricDataQuery` in the array specifies either
a metric to retrieve, or a math expression to be performed on retrieved metrics. A
single `PutMetricAlarm` call can include up to 20
`MetricDataQuery` structures in the array. The 20 structures can include
as many as 10 structures that contain a `MetricStat` parameter to retrieve a
metric, and as many as 10 structures that contain the `Expression` parameter
to perform a math expression. Of those `Expression` structures, one must have
`true` as the value for `ReturnData`. The result of this
expression is the value the alarm watches.

Any expression used in a `PutMetricAlarm` operation must return a single
time series. For more information, see [Metric Math Syntax and Functions](../../../amazoncloudwatch/latest/monitoring/using-metric-math.md#metric-math-syntax) in the _Amazon CloudWatch User_
_Guide_.

Some of the parameters of this structure also have different uses whether you are
using this structure in a `GetMetricData` operation or a
`PutMetricAlarm` operation. These differences are explained in the
following parameter list.

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

The ID of the account where the metrics are located.

If you are performing a `GetMetricData` operation in a monitoring account, use this to specify
which account to retrieve this metric from.

If you are performing a `PutMetricAlarm` operation, use this to specify
which account contains the metric that the alarm is watching.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Expression`

This field can contain either a Metrics Insights query, or a metric math expression to
be performed on the returned data. For more information about Metrics Insights queries,
see [Metrics Insights query components and syntax](../../../amazoncloudwatch/latest/monitoring/cloudwatch-metrics-insights-querylanguage.md) in the _Amazon_
_CloudWatch User Guide_.

A math expression can use the `Id` of the other metrics or queries to refer
to those metrics, and can also use the `Id` of other expressions to use the
result of those expressions. For more information about metric math expressions, see
[Metric Math Syntax and Functions](../../../amazoncloudwatch/latest/monitoring/using-metric-math.md#metric-math-syntax) in the _Amazon CloudWatch User_
_Guide_.

Within each MetricDataQuery object, you must specify either `Expression` or
`MetricStat` but not both.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Label`

A human-readable label for this metric or expression. This is especially useful if
this is an expression, so that you know what the value represents. If the metric or
expression is shown in a CloudWatch dashboard widget, the label is shown. If Label is
omitted, CloudWatch generates a default.

You can put dynamic expressions into a label, so that it is more descriptive. For more
information, see [Using Dynamic\
Labels](../../../amazoncloudwatch/latest/monitoring/graph-dynamic-labels.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricStat`

The metric to be returned, along with statistics, period, and units. Use this
parameter only if this object is retrieving a metric and not performing a math
expression on returned data.

Within one MetricDataQuery object, you must specify either `Expression` or
`MetricStat` but not both.

_Required_: No

_Type_: [MetricStat](aws-properties-cloudwatch-anomalydetector-metricstat.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReturnData`

When used in `GetMetricData`, this option indicates whether to return the
timestamps and raw data values of this metric. If you are performing this call just to
do math expressions and do not also need the raw data returned, you can specify
`false`. If you omit this, the default of `true` is
used.

When used in `PutMetricAlarm`, specify `true` for the one
expression result to use as the alarm. For all other metrics and expressions in the same
`PutMetricAlarm` operation, specify `ReturnData` as
False.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricCharacteristics

MetricMathAnomalyDetector

All content copied from https://docs.aws.amazon.com/.
