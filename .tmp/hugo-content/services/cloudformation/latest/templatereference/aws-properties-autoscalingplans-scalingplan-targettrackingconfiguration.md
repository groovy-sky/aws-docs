This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScalingPlans::ScalingPlan TargetTrackingConfiguration

`TargetTrackingConfiguration` is a subproperty of [ScalingInstruction](../userguide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.md) that specifies a target tracking configuration for a scalable resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomizedScalingMetricSpecification" : CustomizedScalingMetricSpecification,
  "DisableScaleIn" : Boolean,
  "EstimatedInstanceWarmup" : Integer,
  "PredefinedScalingMetricSpecification" : PredefinedScalingMetricSpecification,
  "ScaleInCooldown" : Integer,
  "ScaleOutCooldown" : Integer,
  "TargetValue" : Number
}

```

### YAML

```yaml

  CustomizedScalingMetricSpecification:
    CustomizedScalingMetricSpecification
  DisableScaleIn: Boolean
  EstimatedInstanceWarmup: Integer
  PredefinedScalingMetricSpecification:
    PredefinedScalingMetricSpecification
  ScaleInCooldown: Integer
  ScaleOutCooldown: Integer
  TargetValue: Number

```

## Properties

`CustomizedScalingMetricSpecification`

A customized metric. You can specify either a predefined metric or a customized metric.

_Required_: No

_Type_: [CustomizedScalingMetricSpecification](aws-properties-autoscalingplans-scalingplan-customizedscalingmetricspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableScaleIn`

Indicates whether scale in by the target tracking scaling policy is disabled. If the
value is `true`, scale in is disabled and the target tracking scaling policy
doesn't remove capacity from the scalable resource. Otherwise, scale in is enabled and the
target tracking scaling policy can remove capacity from the scalable resource.

The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EstimatedInstanceWarmup`

The estimated time, in seconds, until a newly launched instance can contribute to the
CloudWatch metrics. This value is used only if the resource is an Auto Scaling group.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedScalingMetricSpecification`

A predefined metric. You can specify either a predefined metric or a customized
metric.

_Required_: No

_Type_: [PredefinedScalingMetricSpecification](aws-properties-autoscalingplans-scalingplan-predefinedscalingmetricspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleInCooldown`

The amount of time, in seconds, after a scale-in activity completes before another scale
in activity can start. This value is not used if the scalable resource is an Auto Scaling
group.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleOutCooldown`

The amount of time, in seconds, after a scale-out activity completes before another
scale-out activity can start. This value is not used if the scalable resource is an Auto
Scaling group.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

The target value for the metric. Although this property accepts numbers of type Double,
it won't accept values that are either too small or too large. Values must be in the range
of -2^360 to 2^360.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Scaling Plans User Guide](../../../autoscaling/plans/userguide/what-is-a-scaling-plan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagFilter

Next

All content copied from https://docs.aws.amazon.com/.
