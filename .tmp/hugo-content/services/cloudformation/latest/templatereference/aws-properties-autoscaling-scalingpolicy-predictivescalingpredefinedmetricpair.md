This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScalingPolicy PredictiveScalingPredefinedMetricPair

Contains metric pair information for the `PredefinedMetricPairSpecification`
property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingMetricSpecification](../userguide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.md) property
type.

For more information, see [Predictive\
scaling](../../../autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.md) in the _Amazon EC2 Auto Scaling User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PredefinedMetricType" : String,
  "ResourceLabel" : String
}

```

### YAML

```yaml

  PredefinedMetricType: String
  ResourceLabel: String

```

## Properties

`PredefinedMetricType`

Indicates which metrics to use. There are two different types of metrics for each
metric type: one is a load metric and one is a scaling metric. For example, if the
metric type is `ASGCPUUtilization`, the Auto Scaling group's total CPU metric is used
as the load metric, and the average CPU metric is used for the scaling metric.

_Required_: Yes

_Type_: String

_Allowed values_: `ASGCPUUtilization | ASGNetworkIn | ASGNetworkOut | ALBRequestCount`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLabel`

A label that uniquely identifies a specific Application Load Balancer target group from
which to determine the total and average request count served by your Auto Scaling group. You
can't specify a resource label unless the target group is attached to the Auto Scaling
group.

You create the resource label by appending the final portion of the load balancer ARN and
the final portion of the target group ARN into a single value, separated by a forward slash
(/). The format of the resource label is:

`app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff`.

Where:

- app/<load-balancer-name>/<load-balancer-id> is the final portion of the
load balancer ARN

- targetgroup/<target-group-name>/<target-group-id> is the final portion of
the target group ARN.

To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](../../../../reference/elasticloadbalancing/latest/apireference/api-describeloadbalancers.md) API operation. To find the ARN for the target group, use the
[DescribeTargetGroups](../../../../reference/elasticloadbalancing/latest/apireference/api-describetargetgroups.md) API operation.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `1023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveScalingPredefinedLoadMetric

PredictiveScalingPredefinedScalingMetric

All content copied from https://docs.aws.amazon.com/.
