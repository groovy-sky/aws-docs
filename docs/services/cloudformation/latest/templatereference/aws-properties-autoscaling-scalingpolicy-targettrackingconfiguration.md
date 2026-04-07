This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScalingPolicy TargetTrackingConfiguration

`TargetTrackingConfiguration` is a property of the [AWS::AutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html) resource that specifies a target tracking scaling
policy configuration for Amazon EC2 Auto Scaling.

For more information about scaling policies, see [Dynamic scaling](../../../autoscaling/ec2/userguide/as-scale-based-on-demand.md) in the
_Amazon EC2 Auto Scaling User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomizedMetricSpecification" : CustomizedMetricSpecification,
  "DisableScaleIn" : Boolean,
  "PredefinedMetricSpecification" : PredefinedMetricSpecification,
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
  TargetValue: Number

```

## Properties

`CustomizedMetricSpecification`

A customized metric. You must specify either a predefined metric or a customized
metric.

_Required_: Conditional

_Type_: [CustomizedMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableScaleIn`

Indicates whether scaling in by the target tracking scaling policy is disabled. If
scaling in is disabled, the target tracking scaling policy doesn't remove instances from
the Auto Scaling group. Otherwise, the target tracking scaling policy can remove instances from
the Auto Scaling group. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedMetricSpecification`

A predefined metric. You must specify either a predefined metric or a customized
metric.

_Required_: Conditional

_Type_: [PredefinedMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-autoscaling-scalingpolicy-predefinedmetricspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

The target value for the metric.

###### Note

Some metrics are based on a count instead of a percentage, such as the request
count for an Application Load Balancer or the number of messages in an SQS queue. If the scaling policy
specifies one of these metrics, specify the target utilization as the optimal
average request or message count per instance during any one-minute interval.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StepAdjustment

TargetTrackingMetricDataQuery
