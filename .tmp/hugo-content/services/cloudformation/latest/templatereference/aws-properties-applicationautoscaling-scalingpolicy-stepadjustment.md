This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy StepAdjustment

`StepAdjustment` specifies a step adjustment for the `StepAdjustments`
property of the [AWS::ApplicationAutoScaling::ScalingPolicy StepScalingPolicyConfiguration](../userguide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.md) property
type.

For the following examples, suppose that you have an alarm with a breach threshold of 50:

- To trigger a step adjustment when the metric is greater than or equal to 50 and less
than 60, specify a lower bound of 0 and an upper bound of 10.

- To trigger a step adjustment when the metric is greater than 40 and less than or equal
to 50, specify a lower bound of -10 and an upper bound of 0.

For more information, see [Step adjustments](../../../autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.md#as-scaling-steps) in the _Application Auto Scaling User_
_Guide_.

You can find a sample template snippet in the [Examples](../userguide/aws-resource-applicationautoscaling-scalingpolicy.md#aws-resource-applicationautoscaling-scalingpolicy--examples) section of the `AWS::ApplicationAutoScaling::ScalingPolicy`
documentation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricIntervalLowerBound" : Number,
  "MetricIntervalUpperBound" : Number,
  "ScalingAdjustment" : Integer
}

```

### YAML

```yaml

  MetricIntervalLowerBound: Number
  MetricIntervalUpperBound: Number
  ScalingAdjustment: Integer

```

## Properties

`MetricIntervalLowerBound`

The lower bound for the difference between the alarm threshold and the CloudWatch metric.
If the metric value is above the breach threshold, the lower bound is inclusive (the metric
must be greater than or equal to the threshold plus the lower bound). Otherwise, it is
exclusive (the metric must be greater than the threshold plus the lower bound). A null value
indicates negative infinity.

You must specify at least one upper or lower bound.

_Required_: Conditional

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricIntervalUpperBound`

The upper bound for the difference between the alarm threshold and the CloudWatch metric.
If the metric value is above the breach threshold, the upper bound is exclusive (the metric
must be less than the threshold plus the upper bound). Otherwise, it is inclusive (the metric
must be less than or equal to the threshold plus the upper bound). A null value indicates
positive infinity.

You must specify at least one upper or lower bound.

_Required_: Conditional

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingAdjustment`

The amount by which to scale. The adjustment is based on the value that you specified in
the `AdjustmentType` property (either an absolute number or a percentage). A
positive value adds to the current capacity and a negative number subtracts from the current
capacity.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md)

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveScalingPredefinedScalingMetric

StepScalingPolicyConfiguration

All content copied from https://docs.aws.amazon.com/.
