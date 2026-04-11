This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy PredictiveScalingMetricSpecification

This structure specifies the metrics and target utilization settings for a predictive
scaling policy.

You must specify either a metric pair, or a load metric and a scaling metric
individually. Specifying a metric pair instead of individual metrics provides a simpler
way to configure metrics for a scaling policy. You choose the metric pair, and the
policy automatically knows the correct sum and average statistics to use for the load
metric and the scaling metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomizedCapacityMetricSpecification" : PredictiveScalingCustomizedCapacityMetric,
  "CustomizedLoadMetricSpecification" : PredictiveScalingCustomizedLoadMetric,
  "CustomizedScalingMetricSpecification" : PredictiveScalingCustomizedScalingMetric,
  "PredefinedLoadMetricSpecification" : PredictiveScalingPredefinedLoadMetric,
  "PredefinedMetricPairSpecification" : PredictiveScalingPredefinedMetricPair,
  "PredefinedScalingMetricSpecification" : PredictiveScalingPredefinedScalingMetric,
  "TargetValue" : Number
}

```

### YAML

```yaml

  CustomizedCapacityMetricSpecification:
    PredictiveScalingCustomizedCapacityMetric
  CustomizedLoadMetricSpecification:
    PredictiveScalingCustomizedLoadMetric
  CustomizedScalingMetricSpecification:
    PredictiveScalingCustomizedScalingMetric
  PredefinedLoadMetricSpecification:
    PredictiveScalingPredefinedLoadMetric
  PredefinedMetricPairSpecification:
    PredictiveScalingPredefinedMetricPair
  PredefinedScalingMetricSpecification:
    PredictiveScalingPredefinedScalingMetric
  TargetValue: Number

```

## Properties

`CustomizedCapacityMetricSpecification`

The customized capacity metric specification.

_Required_: No

_Type_: [PredictiveScalingCustomizedCapacityMetric](aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedcapacitymetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomizedLoadMetricSpecification`

The customized load metric specification.

_Required_: No

_Type_: [PredictiveScalingCustomizedLoadMetric](aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedloadmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomizedScalingMetricSpecification`

The customized scaling metric specification.

_Required_: No

_Type_: [PredictiveScalingCustomizedScalingMetric](aws-properties-applicationautoscaling-scalingpolicy-predictivescalingcustomizedscalingmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedLoadMetricSpecification`

The predefined load metric specification.

_Required_: No

_Type_: [PredictiveScalingPredefinedLoadMetric](aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedloadmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedMetricPairSpecification`

The predefined metric pair specification that determines the appropriate scaling metric and load metric to use.

_Required_: No

_Type_: [PredictiveScalingPredefinedMetricPair](aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedmetricpair.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedScalingMetricSpecification`

The predefined scaling metric specification.

_Required_: No

_Type_: [PredictiveScalingPredefinedScalingMetric](aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpredefinedscalingmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

Specifies the target utilization.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveScalingMetricDimension

PredictiveScalingMetricStat

All content copied from https://docs.aws.amazon.com/.
