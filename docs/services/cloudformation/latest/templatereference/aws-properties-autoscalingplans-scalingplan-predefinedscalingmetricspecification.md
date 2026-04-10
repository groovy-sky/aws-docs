This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScalingPlans::ScalingPlan PredefinedScalingMetricSpecification

`PredefinedScalingMetricSpecification` is a subproperty of [TargetTrackingConfiguration](../userguide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.md) that specifies a customized scaling metric for a
target tracking configuration to use with a scaling plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PredefinedScalingMetricType" : String,
  "ResourceLabel" : String
}

```

### YAML

```yaml

  PredefinedScalingMetricType: String
  ResourceLabel: String

```

## Properties

`PredefinedScalingMetricType`

The metric type. The `ALBRequestCountPerTarget` metric type applies only to
Auto Scaling groups, Spot Fleet requests, and ECS services.

_Required_: Yes

_Type_: String

_Allowed values_: `ASGAverageCPUUtilization | ASGAverageNetworkIn | ASGAverageNetworkOut | DynamoDBReadCapacityUtilization | DynamoDBWriteCapacityUtilization | ECSServiceAverageCPUUtilization | ECSServiceAverageMemoryUtilization | ALBRequestCountPerTarget | RDSReaderAverageCPUUtilization | RDSReaderAverageDatabaseConnections | EC2SpotFleetRequestAverageCPUUtilization | EC2SpotFleetRequestAverageNetworkIn | EC2SpotFleetRequestAverageNetworkOut`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLabel`

Identifies the resource associated with the metric type. You can't specify a resource
label unless the metric type is `ALBRequestCountPerTarget` and there is a target
group for an Application Load Balancer attached to the Auto Scaling group, Spot Fleet request, or
ECS service.

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

PredefinedLoadMetricSpecification

ScalingInstruction

All content copied from https://docs.aws.amazon.com/.
