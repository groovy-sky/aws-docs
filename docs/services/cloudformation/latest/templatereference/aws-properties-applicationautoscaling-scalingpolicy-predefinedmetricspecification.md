This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy PredefinedMetricSpecification

Contains predefined metric specification information for a target tracking scaling policy
for Application Auto Scaling.

`PredefinedMetricSpecification` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy TargetTrackingScalingPolicyConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html)
property type.

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

The metric type. The `ALBRequestCountPerTarget` metric type applies only to
Spot fleet requests and ECS services.

_Required_: Yes

_Type_: String

_Allowed values_: `DynamoDBReadCapacityUtilization | DynamoDBWriteCapacityUtilization | ALBRequestCountPerTarget | RDSReaderAverageCPUUtilization | RDSReaderAverageDatabaseConnections | EC2SpotFleetRequestAverageCPUUtilization | EC2SpotFleetRequestAverageNetworkIn | EC2SpotFleetRequestAverageNetworkOut | SageMakerVariantInvocationsPerInstance | ECSServiceAverageCPUUtilization | ECSServiceAverageMemoryUtilization | AppStreamAverageCapacityUtilization | ComprehendInferenceUtilization | LambdaProvisionedConcurrencyUtilization | CassandraReadCapacityUtilization | CassandraWriteCapacityUtilization | KafkaBrokerStorageUtilization | ElastiCacheEngineCPUUtilization | ElastiCacheDatabaseMemoryUsagePercentage | ElastiCachePrimaryEngineCPUUtilization | ElastiCacheReplicaEngineCPUUtilization | ElastiCacheDatabaseMemoryUsageCountedForEvictPercentage | NeptuneReaderAverageCPUUtilization | SageMakerVariantProvisionedConcurrencyUtilization | ElastiCacheDatabaseCapacityUsageCountedForEvictPercentage | SageMakerInferenceComponentInvocationsPerCopy | WorkSpacesAverageUserSessionsCapacityUtilization | SageMakerInferenceComponentConcurrentRequestsPerCopyHighResolution | SageMakerVariantConcurrentRequestsPerModelHighResolution`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLabel`

Identifies the resource associated with the metric type. You can't specify a resource
label unless the metric type is `ALBRequestCountPerTarget` and there is a target
group attached to the Spot Fleet or ECS service.

You create the resource label by appending the final portion of the load balancer ARN
and the final portion of the target group ARN into a single value, separated by a forward
slash (/). The format of the resource label is:

`app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff`.

Where:

- app/<load-balancer-name>/<load-balancer-id> is the final portion of
the load balancer ARN

- targetgroup/<target-group-name>/<target-group-id> is the final portion
of the target group ARN.

To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use
the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `1023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md)

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetricDimension

PredictiveScalingCustomizedCapacityMetric
