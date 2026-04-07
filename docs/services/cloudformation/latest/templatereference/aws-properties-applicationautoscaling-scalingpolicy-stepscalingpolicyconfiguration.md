This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy StepScalingPolicyConfiguration

`StepScalingPolicyConfiguration` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a step scaling
policy configuration for Application Auto Scaling.

For more information, see [Step scaling policies](../../../autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.md) in the _Application Auto Scaling User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdjustmentType" : String,
  "Cooldown" : Integer,
  "MetricAggregationType" : String,
  "MinAdjustmentMagnitude" : Integer,
  "StepAdjustments" : [ StepAdjustment, ... ]
}

```

### YAML

```yaml

  AdjustmentType: String
  Cooldown: Integer
  MetricAggregationType: String
  MinAdjustmentMagnitude: Integer
  StepAdjustments:
    - StepAdjustment

```

## Properties

`AdjustmentType`

Specifies whether the `ScalingAdjustment` value in the
`StepAdjustment` property is an absolute number or a percentage of the current
capacity.

_Required_: No

_Type_: String

_Allowed values_: `ChangeInCapacity | PercentChangeInCapacity | ExactCapacity`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cooldown`

The amount of time, in seconds, to wait for a previous scaling activity to take effect. If
not specified, the default value is 300. For more information, see [Cooldown period](https://docs.aws.amazon.com/autoscaling/application/userguide/step-scaling-policy-overview.html#step-scaling-cooldown) in the _Application Auto Scaling User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricAggregationType`

The aggregation type for the CloudWatch metrics. Valid values are `Minimum`,
`Maximum`, and `Average`. If the aggregation type is null, the
value is treated as `Average`.

_Required_: No

_Type_: String

_Allowed values_: `Average | Minimum | Maximum`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinAdjustmentMagnitude`

The minimum value to scale by when the adjustment type is
`PercentChangeInCapacity`. For example, suppose that you create a step
scaling policy to scale out an Amazon ECS service by 25 percent and you specify a
`MinAdjustmentMagnitude` of 2. If the service has 4 tasks and the scaling
policy is performed, 25 percent of 4 is 1. However, because you specified a
`MinAdjustmentMagnitude` of 2, Application Auto Scaling scales out the service by 2
tasks.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StepAdjustments`

A set of adjustments that enable you to scale based on the size of the alarm
breach.

At least one step adjustment is required if you are adding a new step scaling policy
configuration.

_Required_: No

_Type_: Array of [StepAdjustment](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationautoscaling-scalingpolicy-stepadjustment.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md)

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StepAdjustment

TargetTrackingMetric
