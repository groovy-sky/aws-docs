This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy TargetTrackingScalingPolicyConfiguration

`TargetTrackingScalingPolicyConfiguration` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a target
tracking scaling policy configuration for Application Auto Scaling. Use a target tracking
scaling policy to adjust the capacity of the specified scalable target in response to actual
workloads, so that resource utilization remains at or near the target utilization value.

For more information, see [Target\
tracking scaling policies](../../../autoscaling/application/userguide/application-auto-scaling-target-tracking.md) in the _Application Auto Scaling User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomizedMetricSpecification" : CustomizedMetricSpecification,
  "DisableScaleIn" : Boolean,
  "PredefinedMetricSpecification" : PredefinedMetricSpecification,
  "ScaleInCooldown" : Integer,
  "ScaleOutCooldown" : Integer,
  "TargetValue" : Number
}

```

### YAML

```yaml

  CustomizedMetricSpecification:
    CustomizedMetricSpecification
  DisableScaleIn: Boolean
  PredefinedMetricSpecification:
    PredefinedMetricSpecification
  ScaleInCooldown: Integer
  ScaleOutCooldown: Integer
  TargetValue: Number

```

## Properties

`CustomizedMetricSpecification`

A customized metric. You can specify either a predefined metric or a customized
metric.

_Required_: No

_Type_: [CustomizedMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableScaleIn`

Indicates whether scale in by the target tracking scaling policy is disabled. If the
value is `true`, scale in is disabled and the target tracking scaling policy
won't remove capacity from the scalable target. Otherwise, scale in is enabled and the
target tracking scaling policy can remove capacity from the scalable target. The default
value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedMetricSpecification`

A predefined metric. You can specify either a predefined metric or a customized
metric.

_Required_: No

_Type_: [PredefinedMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleInCooldown`

The amount of time, in seconds, after a scale-in activity completes before another
scale-in activity can start. For more information and for default values, see [Define cooldown periods](https://docs.aws.amazon.com/autoscaling/application/userguide/target-tracking-scaling-policy-overview.html#target-tracking-cooldown) in the _Application Auto Scaling User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleOutCooldown`

The amount of time, in seconds, to wait for a previous scale-out activity to take effect.
For more information and for default values, see [Define cooldown periods](https://docs.aws.amazon.com/autoscaling/application/userguide/target-tracking-scaling-policy-overview.html#target-tracking-cooldown) in the _Application Auto Scaling User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

The target value for the metric. Although this property accepts numbers of type Double, it
won't accept values that are either too small or too large. Values must be in the range of
 -2^360 to 2^360. The value must be a valid number based on the choice of metric. For example,
if the metric is CPU utilization, then the target value is a percent value that represents how
much of the CPU can be used before scaling out.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md)

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetTrackingMetricStat

Next
