This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::ExpressGatewayService

Creates an Express service that simplifies deploying containerized web applications on
Amazon ECS with managed AWS infrastructure. This operation provisions and configures
Application Load Balancers, target groups, security groups, and auto-scaling policies
automatically.

Specify a primary container configuration with your application image and basic
settings. Amazon ECS creates the necessary AWS resources for traffic distribution, health
monitoring, network access control, and capacity management.

Provide an execution role for task operations and an infrastructure role for managing
AWS resources on your behalf.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECS::ExpressGatewayService",
  "Properties" : {
      "Cluster" : String,
      "Cpu" : String,
      "ExecutionRoleArn" : String,
      "HealthCheckPath" : String,
      "InfrastructureRoleArn" : String,
      "Memory" : String,
      "NetworkConfiguration" : ExpressGatewayServiceNetworkConfiguration,
      "PrimaryContainer" : ExpressGatewayContainer,
      "ScalingTarget" : ExpressGatewayScalingTarget,
      "ServiceName" : String,
      "Tags" : [ Tag, ... ],
      "TaskRoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::ECS::ExpressGatewayService
Properties:
  Cluster: String
  Cpu: String
  ExecutionRoleArn: String
  HealthCheckPath: String
  InfrastructureRoleArn: String
  Memory: String
  NetworkConfiguration:
    ExpressGatewayServiceNetworkConfiguration
  PrimaryContainer:
    ExpressGatewayContainer
  ScalingTarget:
    ExpressGatewayScalingTarget
  ServiceName: String
  Tags:
    - Tag
  TaskRoleArn: String

```

## Properties

`Cluster`

The short name or full ARN of the cluster that hosts the Express service.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Cpu`

The CPU allocation for tasks in this service revision.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The ARN of the task execution role for the service revision.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckPath`

The health check path for this service revision.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InfrastructureRoleArn`

The ARN of the infrastructure role that manages AWS resources for the Express
service.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Memory`

The memory allocation for tasks in this service revision.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

The network configuration for tasks in this service revision.

_Required_: No

_Type_: [ExpressGatewayServiceNetworkConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryContainer`

The primary container configuration for this service revision.

_Required_: Yes

_Type_: [ExpressGatewayContainer](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-expressgatewayservice-expressgatewaycontainer.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingTarget`

The auto-scaling configuration for this service revision.

_Required_: No

_Type_: [ExpressGatewayScalingTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-expressgatewayservice-expressgatewayscalingtarget.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

The name of the Express service.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The metadata applied to the Express service.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-expressgatewayservice-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TaskRoleArn`

The ARN of the task role for the service revision.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ActiveConfigurations`

The list of active service configurations for the Express service.

`CreatedAt`

The Unix timestamp for when the Express service was created.

`ECSManagedResourceArns.AutoScaling.ApplicationAutoScalingPolicies`

The list of Auto Scaling policy ARNs associated with the express service.

`ECSManagedResourceArns.AutoScaling.ScalableTarget`

The Auto Scaling Scalable Target ARN associated with the express service.

`ECSManagedResourceArns.IngressPath.CertificateArn`

The Certificate ARN associated with the express service.

`ECSManagedResourceArns.IngressPath.ListenerArn`

The ARN of the Load Balancer listener associated with the express service.

`ECSManagedResourceArns.IngressPath.ListenerRuleArn`

The ARN of the Load Balancer listener rule associated with the express service.

`ECSManagedResourceArns.IngressPath.LoadBalancerArn`

The ARN of the Load Balancer associated with the express service.

`ECSManagedResourceArns.IngressPath.LoadBalancerSecurityGroups`

The list of Load Balancer Security Group ARNs associated with the express service.

`ECSManagedResourceArns.IngressPath.TargetGroupArns`

The list of Target Group ARNs associated with the express service.

`ECSManagedResourceArns.LogGroups`

The list of Log Group ARNs associated with the express service.

`ECSManagedResourceArns.MetricAlarms`

The list of Metric Alarm ARNs associated with the express service.

`ECSManagedResourceArns.ServiceSecurityGroups`

The list of Security Group ARNs associated with the express service.

`Endpoint`

The Endpoint of the express service.

`ServiceArn`

The ARN that identifies the Express service.

`UpdatedAt`

The Unix timestamp for when the Express service was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Volume

AutoScalingArns
