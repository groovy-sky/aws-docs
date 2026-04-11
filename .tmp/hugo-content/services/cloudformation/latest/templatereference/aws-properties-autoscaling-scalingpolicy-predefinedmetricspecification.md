This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScalingPolicy PredefinedMetricSpecification

Contains predefined metric specification information for a target tracking scaling policy
for Amazon EC2 Auto Scaling.

`PredefinedMetricSpecification` is a property of the [AWS::AutoScaling::ScalingPolicy TargetTrackingConfiguration](../userguide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.md) property type.

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

The metric type. The following predefined metrics are available:

- `ASGAverageCPUUtilization` \- Average CPU utilization of the Auto Scaling
group.

- `ASGAverageNetworkIn` \- Average number of bytes received on all
network interfaces by the Auto Scaling group.

- `ASGAverageNetworkOut` \- Average number of bytes sent out on all
network interfaces by the Auto Scaling group.

- `ALBRequestCountPerTarget` \- Average Application Load Balancer request count per target
for your Auto Scaling group.

_Required_: Yes

_Type_: String

_Allowed values_: `ASGAverageCPUUtilization | ASGAverageNetworkIn | ASGAverageNetworkOut | ALBRequestCountPerTarget`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLabel`

A label that uniquely identifies a specific Application Load Balancer target group
from which to determine the average request count served by your Auto Scaling group. You can't
specify a resource label unless the target group is attached to the Auto Scaling group.

You create the resource label by appending the final portion of the load balancer ARN
and the final portion of the target group ARN into a single value, separated by a forward
slash (/). The format of the resource label is:

`app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff`.

Where:

- app/<load-balancer-name>/<load-balancer-id> is the final portion of
the load balancer ARN

- targetgroup/<target-group-name>/<target-group-id> is the final portion
of the target group ARN.

To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](../../../../reference/elasticloadbalancing/latest/apireference/api-describeloadbalancers.md) API operation. To find the ARN for the target group, use
the [DescribeTargetGroups](../../../../reference/elasticloadbalancing/latest/apireference/api-describetargetgroups.md) API operation.

_Required_: Conditional

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `1023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricStat

PredictiveScalingConfiguration

All content copied from https://docs.aws.amazon.com/.
