This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan ExecutionBlockConfiguration

Execution block configurations for a workflow in a Region switch plan. An execution block represents a specific type of action to perform during a Region switch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArcRoutingControlConfig" : ArcRoutingControlConfiguration,
  "CustomActionLambdaConfig" : CustomActionLambdaConfiguration,
  "DocumentDbConfig" : DocumentDbConfiguration,
  "Ec2AsgCapacityIncreaseConfig" : Ec2AsgCapacityIncreaseConfiguration,
  "EcsCapacityIncreaseConfig" : EcsCapacityIncreaseConfiguration,
  "EksResourceScalingConfig" : EksResourceScalingConfiguration,
  "ExecutionApprovalConfig" : ExecutionApprovalConfiguration,
  "GlobalAuroraConfig" : GlobalAuroraConfiguration,
  "ParallelConfig" : ParallelExecutionBlockConfiguration,
  "RdsCreateCrossRegionReadReplicaConfig" : RdsCreateCrossRegionReplicaConfiguration,
  "RdsPromoteReadReplicaConfig" : RdsPromoteReadReplicaConfiguration,
  "RegionSwitchPlanConfig" : RegionSwitchPlanConfiguration,
  "Route53HealthCheckConfig" : Route53HealthCheckConfiguration
}

```

### YAML

```yaml

  ArcRoutingControlConfig:
    ArcRoutingControlConfiguration
  CustomActionLambdaConfig:
    CustomActionLambdaConfiguration
  DocumentDbConfig:
    DocumentDbConfiguration
  Ec2AsgCapacityIncreaseConfig:
    Ec2AsgCapacityIncreaseConfiguration
  EcsCapacityIncreaseConfig:
    EcsCapacityIncreaseConfiguration
  EksResourceScalingConfig:
    EksResourceScalingConfiguration
  ExecutionApprovalConfig:
    ExecutionApprovalConfiguration
  GlobalAuroraConfig:
    GlobalAuroraConfiguration
  ParallelConfig:
    ParallelExecutionBlockConfiguration
  RdsCreateCrossRegionReadReplicaConfig:
    RdsCreateCrossRegionReplicaConfiguration
  RdsPromoteReadReplicaConfig:
    RdsPromoteReadReplicaConfiguration
  RegionSwitchPlanConfig:
    RegionSwitchPlanConfiguration
  Route53HealthCheckConfig:
    Route53HealthCheckConfiguration

```

## Properties

`ArcRoutingControlConfig`

An ARC routing control execution block.

_Required_: No

_Type_: [ArcRoutingControlConfiguration](aws-properties-arcregionswitch-plan-arcroutingcontrolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomActionLambdaConfig`

An AWS Lambda execution block.

_Required_: No

_Type_: [CustomActionLambdaConfiguration](aws-properties-arcregionswitch-plan-customactionlambdaconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentDbConfig`

Property description not available.

_Required_: No

_Type_: [DocumentDbConfiguration](aws-properties-arcregionswitch-plan-documentdbconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2AsgCapacityIncreaseConfig`

An EC2 Auto Scaling group execution block.

_Required_: No

_Type_: [Ec2AsgCapacityIncreaseConfiguration](aws-properties-arcregionswitch-plan-ec2asgcapacityincreaseconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcsCapacityIncreaseConfig`

The capacity increase specified for the configuration.

_Required_: No

_Type_: [EcsCapacityIncreaseConfiguration](aws-properties-arcregionswitch-plan-ecscapacityincreaseconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EksResourceScalingConfig`

An AWS EKS resource scaling execution block.

_Required_: No

_Type_: [EksResourceScalingConfiguration](aws-properties-arcregionswitch-plan-eksresourcescalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionApprovalConfig`

A manual approval execution block.

_Required_: No

_Type_: [ExecutionApprovalConfiguration](aws-properties-arcregionswitch-plan-executionapprovalconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalAuroraConfig`

An Aurora Global Database execution block.

_Required_: No

_Type_: [GlobalAuroraConfiguration](aws-properties-arcregionswitch-plan-globalauroraconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParallelConfig`

A parallel configuration execution block.

_Required_: No

_Type_: [ParallelExecutionBlockConfiguration](aws-properties-arcregionswitch-plan-parallelexecutionblockconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RdsCreateCrossRegionReadReplicaConfig`

An Amazon RDS create cross-Region replica execution block.

_Required_: No

_Type_: [RdsCreateCrossRegionReplicaConfiguration](aws-properties-arcregionswitch-plan-rdscreatecrossregionreplicaconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RdsPromoteReadReplicaConfig`

An Amazon RDS promote read replica execution block.

_Required_: No

_Type_: [RdsPromoteReadReplicaConfiguration](aws-properties-arcregionswitch-plan-rdspromotereadreplicaconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionSwitchPlanConfig`

A Region switch plan execution block.

_Required_: No

_Type_: [RegionSwitchPlanConfiguration](aws-properties-arcregionswitch-plan-regionswitchplanconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Route53HealthCheckConfig`

The Amazon Route 53 health check configuration.

_Required_: No

_Type_: [Route53HealthCheckConfiguration](aws-properties-arcregionswitch-plan-route53healthcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExecutionApprovalConfiguration

GlobalAuroraConfiguration

All content copied from https://docs.aws.amazon.com/.
