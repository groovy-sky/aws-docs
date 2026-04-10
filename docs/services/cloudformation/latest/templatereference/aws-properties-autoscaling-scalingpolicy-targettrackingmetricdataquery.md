This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScalingPolicy TargetTrackingMetricDataQuery

The metric data to return. Also defines whether this call is returning data for one
metric only, or whether it is performing a math expression on the values of returned
metric statistics to create a new time series. A time series is a series of data points,
each of which is associated with a timestamp.

You can use `TargetTrackingMetricDataQuery` structures with a
[PutScalingPolicy](../../../../reference/autoscaling/ec2/apireference/api-putscalingpolicy.md) operation when you specify a
[TargetTrackingConfiguration](../../../../reference/autoscaling/ec2/apireference/api-targettrackingconfiguration.md)
in the request.

You can call for a single metric or perform math expressions on multiple
metrics. Any expressions used in a metric specification must eventually return a single
time series.

For more information, see the [Create a target tracking scaling policy for Amazon EC2 Auto Scaling using metric math](../../../autoscaling/ec2/userguide/ec2-auto-scaling-target-tracking-metric-math.md) in
the _Amazon EC2 Auto Scaling User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "Id" : String,
  "Label" : String,
  "MetricStat" : TargetTrackingMetricStat,
  "Period" : Integer,
  "ReturnData" : Boolean
}

```

### YAML

```yaml

  Expression: String
  Id: String
  Label: String
  MetricStat:
    TargetTrackingMetricStat
  Period: Integer
  ReturnData: Boolean

```

## Properties

`Expression`

The math expression to perform on the returned data, if this object is performing a
math expression. This expression can use the `Id` of the other metrics to
refer to those metrics, and can also use the `Id` of other expressions to use
the result of those expressions.

Conditional: Within each `TargetTrackingMetricDataQuery` object, you must
specify either `Expression` or `MetricStat`, but not both.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `2047`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

A short name that identifies the object's results in the response. This name must be
unique among all `TargetTrackingMetricDataQuery` objects specified for a
single scaling policy. If you are performing math expressions on this set of data, this
name represents that data and can serve as a variable in the mathematical expression.
The valid characters are letters, numbers, and underscores. The first character must be
a lowercase letter.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

A human-readable label for this metric or expression. This is especially useful if
this is a math expression, so that you know what the value represents.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Maximum_: `2047`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricStat`

Information about the metric data to return.

Conditional: Within each `TargetTrackingMetricDataQuery` object, you must
specify either `Expression` or `MetricStat`, but not both.

_Required_: No

_Type_: [TargetTrackingMetricStat](aws-properties-autoscaling-scalingpolicy-targettrackingmetricstat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The period of the metric in seconds. The default value is 60. Accepted values are 10, 30, and 60. For high resolution metric, set the value to less than 60. For more information, see
[Create a target tracking policy using high-resolution metrics for faster response](../../../autoscaling/ec2/userguide/policy-creating-high-resolution-metrics.md).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReturnData`

Indicates whether to return the timestamps and raw data values of this metric.

If you use any math expressions, specify `true` for this value for only the
final math expression that the metric specification is based on. You must specify
`false` for `ReturnData` for all the other metrics and
expressions used in the metric specification.

If you are only retrieving metrics and not performing any math expressions, do not
specify anything for `ReturnData`. This sets it to its default
( `true`).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetTrackingConfiguration

TargetTrackingMetricStat

All content copied from https://docs.aws.amazon.com/.
