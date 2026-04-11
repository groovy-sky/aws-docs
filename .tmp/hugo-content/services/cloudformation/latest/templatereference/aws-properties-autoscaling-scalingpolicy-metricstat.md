This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScalingPolicy MetricStat

`MetricStat` is a property of the [AWS::AutoScaling::ScalingPolicy MetricDataQuery](../userguide/aws-properties-autoscaling-scalingpolicy-metricdataquery.md) property type.

This structure defines the CloudWatch metric to return, along with the statistic and
unit.

For more information about the CloudWatch terminology below, see [Amazon CloudWatch concepts](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md) in the _Amazon CloudWatch User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Metric" : Metric,
  "Stat" : String,
  "Unit" : String
}

```

### YAML

```yaml

  Metric:
    Metric
  Stat: String
  Unit: String

```

## Properties

`Metric`

The CloudWatch metric to return, including the metric name, namespace, and dimensions. To
get the exact metric name, namespace, and dimensions, inspect the [Metric](../../../../reference/amazoncloudwatch/latest/apireference/api-metric.md) object that is returned by a call to [ListMetrics](../../../../reference/amazoncloudwatch/latest/apireference/api-listmetrics.md).

_Required_: Yes

_Type_: [Metric](aws-properties-autoscaling-scalingpolicy-metric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stat`

The statistic to return. It can include any CloudWatch statistic or extended statistic. For
a list of valid values, see the table in [Statistics](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md#Statistic) in the _Amazon CloudWatch User Guide_.

The most commonly used metrics for predictive scaling are `Average` and
`Sum`.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit to use for the returned data points. For a complete list of the units that
CloudWatch supports, see the [MetricDatum](../../../../reference/amazoncloudwatch/latest/apireference/api-metricdatum.md)
data type in the _Amazon CloudWatch API Reference_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDimension

PredefinedMetricSpecification

All content copied from https://docs.aws.amazon.com/.
