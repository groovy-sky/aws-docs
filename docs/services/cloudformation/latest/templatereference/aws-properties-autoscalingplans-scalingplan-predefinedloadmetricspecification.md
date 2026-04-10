This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScalingPlans::ScalingPlan PredefinedLoadMetricSpecification

`PredefinedLoadMetricSpecification` is a subproperty of
[ScalingInstruction](../userguide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.md)
that specifies a predefined load metric for predictive
scaling to use with a scaling plan.

After creating your scaling plan, you can use the AWS Auto Scaling console to
visualize forecasts for the specified metric. For more information, see [View\
scaling information for a resource](../../../autoscaling/plans/userguide/gs-create-scaling-plan.md#gs-view-resource) in the _Scaling Plans User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PredefinedLoadMetricType" : String,
  "ResourceLabel" : String
}

```

### YAML

```yaml

  PredefinedLoadMetricType: String
  ResourceLabel: String

```

## Properties

`PredefinedLoadMetricType`

The metric type.

_Required_: Yes

_Type_: String

_Allowed values_: `ASGTotalCPUUtilization | ASGTotalNetworkIn | ASGTotalNetworkOut | ALBTargetGroupRequestCount`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLabel`

Identifies the resource associated with the metric type. You can't specify a resource
label unless the metric type is `ALBTargetGroupRequestCount` and there is a
target group for an Application Load Balancer attached to the Auto Scaling group.

You create the resource label by appending the final portion of the load balancer ARN
and the final portion of the target group ARN into a single value, separated by a forward
slash (/). The format is
app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>,
where:

- app/<load-balancer-name>/<load-balancer-id> is the final portion of
the load balancer ARN

- targetgroup/<target-group-name>/<target-group-id> is the final portion
of the target group ARN.

This is an example:
app/EC2Co-EcsEl-1TKLTMITMM0EO/f37c06a68c1748aa/targetgroup/EC2Co-Defau-LDNM7Q3ZH1ZN/6d4ea56ca2d6a18d.

To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](../../../../reference/elasticloadbalancing/latest/apireference/api-describeloadbalancers.md) API operation. To find the ARN for the target group, use
the [DescribeTargetGroups](../../../../reference/elasticloadbalancing/latest/apireference/api-describetargetgroups.md) API operation.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `1023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Scaling Plans User Guide](../../../autoscaling/plans/userguide/what-is-a-scaling-plan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDimension

PredefinedScalingMetricSpecification

All content copied from https://docs.aws.amazon.com/.
