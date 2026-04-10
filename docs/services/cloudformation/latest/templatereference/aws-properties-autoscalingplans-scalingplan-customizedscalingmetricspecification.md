This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScalingPlans::ScalingPlan CustomizedScalingMetricSpecification

`CustomizedScalingMetricSpecification` is a subproperty of [TargetTrackingConfiguration](../userguide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.md) that specifies a customized scaling metric for a
target tracking configuration to use with a scaling plan.

To create your customized scaling metric specification:

- Add values for each required property from CloudWatch. You can use an existing
metric, or a new metric that you create. To use your own metric, you must first
publish the metric to CloudWatch. For more information, see [Publish Custom\
Metrics](../../../amazoncloudwatch/latest/monitoring/publishingmetrics.md) in the _Amazon CloudWatch User Guide_.

- Choose a metric that changes proportionally with capacity. The value of the metric
should increase or decrease in inverse proportion to the number of capacity units.
That is, the value of the metric should decrease when capacity increases.

For information about terminology, available metrics, or how to publish new metrics, see
[Amazon CloudWatch\
Concepts](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md) in the _Amazon CloudWatch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ MetricDimension, ... ],
  "MetricName" : String,
  "Namespace" : String,
  "Statistic" : String,
  "Unit" : String
}

```

### YAML

```yaml

  Dimensions:
    - MetricDimension
  MetricName: String
  Namespace: String
  Statistic: String
  Unit: String

```

## Properties

`Dimensions`

The dimensions of the metric.

Conditional: If you published your metric with dimensions, you must specify the same
dimensions in your scaling policy.

_Required_: No

_Type_: Array of [MetricDimension](aws-properties-autoscalingplans-scalingplan-metricdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric. To get the exact metric name, namespace, and dimensions, inspect
the [Metrics](../../../../reference/amazoncloudwatch/latest/apireference/api-metric.md) object that is returned by a call to [ListMetrics](../../../../reference/amazoncloudwatch/latest/apireference/api-listmetrics.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistic`

The statistic of the metric.

_Required_: Yes

_Type_: String

_Allowed values_: `Average | Minimum | Maximum | SampleCount | Sum`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of the metric. For a complete list of the units that CloudWatch supports, see the
[MetricDatum](../../../../reference/amazoncloudwatch/latest/apireference/api-metricdatum.md) data
type in the _Amazon CloudWatch API Reference_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Scaling Plans User Guide](../../../autoscaling/plans/userguide/what-is-a-scaling-plan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomizedLoadMetricSpecification

MetricDimension

All content copied from https://docs.aws.amazon.com/.
