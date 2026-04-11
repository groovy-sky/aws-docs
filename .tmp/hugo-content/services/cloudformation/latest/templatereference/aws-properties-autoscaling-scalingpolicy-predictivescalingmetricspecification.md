This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScalingPolicy PredictiveScalingMetricSpecification

A structure that specifies a metric specification for the
`MetricSpecifications` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingConfiguration](../userguide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.md) property
type.

You must specify either a metric pair, or a load metric and a scaling metric individually.
Specifying a metric pair instead of individual metrics provides a simpler way to configure
metrics for a scaling policy. You choose the metric pair, and the policy automatically knows
the correct sum and average statistics to use for the load metric and the scaling
metric.

Example

- You create a predictive scaling policy and specify `ALBRequestCount` as the
value for the metric pair and `1000.0` as the target value. For this type of
metric, you must provide the metric dimension for the corresponding target group, so you
also provide a resource label for the Application Load Balancer target group that is
attached to your Auto Scaling group.

- The number of requests the target group receives per minute provides the load metric,
and the request count averaged between the members of the target group provides the
scaling metric. In CloudWatch, this refers to the `RequestCount` and
`RequestCountPerTarget` metrics, respectively.

- For optimal use of predictive scaling, you adhere to the best practice of using a
dynamic scaling policy to automatically scale between the minimum capacity and maximum
capacity in response to real-time changes in resource utilization.

- Amazon EC2 Auto Scaling consumes data points for the load metric over the last 14 days
and creates an hourly load forecast for predictive scaling. (A minimum of 24 hours of data
is required.)

- After creating the load forecast, Amazon EC2 Auto Scaling determines when to reduce or
increase the capacity of your Auto Scaling group in each hour of the forecast period so
that the average number of requests received by each instance is as close to 1000 requests
per minute as possible at all times.

For information about using custom metrics with predictive scaling, see [Advanced predictive scaling policy configurations using custom metrics](../../../autoscaling/ec2/userguide/predictive-scaling-customized-metric-specification.md) in the
_Amazon EC2 Auto Scaling User Guide_.

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

_Type_: [PredictiveScalingCustomizedCapacityMetric](aws-properties-autoscaling-scalingpolicy-predictivescalingcustomizedcapacitymetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomizedLoadMetricSpecification`

The customized load metric specification.

_Required_: No

_Type_: [PredictiveScalingCustomizedLoadMetric](aws-properties-autoscaling-scalingpolicy-predictivescalingcustomizedloadmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomizedScalingMetricSpecification`

The customized scaling metric specification.

_Required_: No

_Type_: [PredictiveScalingCustomizedScalingMetric](aws-properties-autoscaling-scalingpolicy-predictivescalingcustomizedscalingmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedLoadMetricSpecification`

The predefined load metric specification.

_Required_: No

_Type_: [PredictiveScalingPredefinedLoadMetric](aws-properties-autoscaling-scalingpolicy-predictivescalingpredefinedloadmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedMetricPairSpecification`

The predefined metric pair specification from which Amazon EC2 Auto Scaling determines the
appropriate scaling metric and load metric to use.

_Required_: No

_Type_: [PredictiveScalingPredefinedMetricPair](aws-properties-autoscaling-scalingpolicy-predictivescalingpredefinedmetricpair.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredefinedScalingMetricSpecification`

The predefined scaling metric specification.

_Required_: No

_Type_: [PredictiveScalingPredefinedScalingMetric](aws-properties-autoscaling-scalingpolicy-predictivescalingpredefinedscalingmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

Specifies the target utilization.

###### Note

Some metrics are based on a count instead of a percentage, such as the request
count for an Application Load Balancer or the number of messages in an SQS queue. If the scaling policy
specifies one of these metrics, specify the target utilization as the optimal
average request or message count per instance during any one-minute interval.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveScalingCustomizedScalingMetric

PredictiveScalingPredefinedLoadMetric

All content copied from https://docs.aws.amazon.com/.
