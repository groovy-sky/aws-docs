This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster CloudWatchAlarmDefinition

`CloudWatchAlarmDefinition` is a subproperty of the `ScalingTrigger ` property, which determines when to trigger an automatic scaling activity. Scaling activity begins when you satisfy the defined alarm conditions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "Dimensions" : [ MetricDimension, ... ],
  "EvaluationPeriods" : Integer,
  "MetricName" : String,
  "Namespace" : String,
  "Period" : Integer,
  "Statistic" : String,
  "Threshold" : Number,
  "Unit" : String
}

```

### YAML

```yaml

  ComparisonOperator: String
  Dimensions:
    - MetricDimension
  EvaluationPeriods: Integer
  MetricName: String
  Namespace: String
  Period: Integer
  Statistic: String
  Threshold: Number
  Unit: String

```

## Properties

`ComparisonOperator`

Determines how the metric specified by `MetricName` is compared to the value
specified by `Threshold`.

_Required_: Yes

_Type_: String

_Allowed values_: `GREATER_THAN_OR_EQUAL | GREATER_THAN | LESS_THAN | LESS_THAN_OR_EQUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dimensions`

A CloudWatch metric dimension.

_Required_: No

_Type_: Array of [MetricDimension](aws-properties-emr-cluster-metricdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationPeriods`

The number of periods, in five-minute increments, during which the alarm condition must
exist before the alarm triggers automatic scaling activity. The default value is
`1`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the CloudWatch metric that is watched to determine an alarm
condition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace for the CloudWatch metric. The default is
`AWS/ElasticMapReduce`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The period, in seconds, over which the statistic is applied. CloudWatch metrics for
Amazon EMR are emitted every five minutes (300 seconds), so if you specify a
CloudWatch metric, specify `300`.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistic`

The statistic to apply to the metric associated with the alarm. The default is
`AVERAGE`.

_Required_: No

_Type_: String

_Allowed values_: `SAMPLE_COUNT | AVERAGE | SUM | MINIMUM | MAXIMUM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The value against which the specified statistic is compared.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of measure associated with the CloudWatch metric being watched. The value
specified for `Unit` must correspond to the units specified in the CloudWatch
metric.

_Required_: No

_Type_: String

_Allowed values_: `NONE | SECONDS | MICRO_SECONDS | MILLI_SECONDS | BYTES | KILO_BYTES | MEGA_BYTES | GIGA_BYTES | TERA_BYTES | BITS | KILO_BITS | MEGA_BITS | GIGA_BITS | TERA_BITS | PERCENT | COUNT | BYTES_PER_SECOND | KILO_BYTES_PER_SECOND | MEGA_BYTES_PER_SECOND | GIGA_BYTES_PER_SECOND | TERA_BYTES_PER_SECOND | BITS_PER_SECOND | KILO_BITS_PER_SECOND | MEGA_BITS_PER_SECOND | GIGA_BITS_PER_SECOND | TERA_BITS_PER_SECOND | COUNT_PER_SECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BootstrapActionConfig

ComputeLimits

All content copied from https://docs.aws.amazon.com/.
