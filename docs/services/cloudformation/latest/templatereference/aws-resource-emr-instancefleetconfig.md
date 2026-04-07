This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceFleetConfig

Use `InstanceFleetConfig` to define instance fleets for an EMR cluster. A cluster can not use both instance fleets and instance groups. For more information, see [Configure Instance Fleets](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the _Amazon EMR Management Guide_.

###### Note

The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.

###### Important

You can currently only add a task instance fleet to a cluster with this resource. If you use
this resource, CloudFormation waits for the cluster launch to complete before adding the
task instance fleet to the cluster. In order to add a task instance fleet to the cluster
as part of the cluster launch and minimize delays in provisioning task nodes, use the
`TaskInstanceFleets` subproperty for the [AWS::EMR::Cluster JobFlowInstancesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html) property instead. To use this
subproperty, see [AWS::EMR::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html) for examples.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMR::InstanceFleetConfig",
  "Properties" : {
      "ClusterId" : String,
      "InstanceFleetType" : String,
      "InstanceTypeConfigs" : [ InstanceTypeConfig, ... ],
      "LaunchSpecifications" : InstanceFleetProvisioningSpecifications,
      "Name" : String,
      "ResizeSpecifications" : InstanceFleetResizingSpecifications,
      "TargetOnDemandCapacity" : Integer,
      "TargetSpotCapacity" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EMR::InstanceFleetConfig
Properties:
  ClusterId: String
  InstanceFleetType: String
  InstanceTypeConfigs:
    - InstanceTypeConfig
  LaunchSpecifications:
    InstanceFleetProvisioningSpecifications
  Name: String
  ResizeSpecifications:
    InstanceFleetResizingSpecifications
  TargetOnDemandCapacity: Integer
  TargetSpotCapacity: Integer

```

## Properties

`ClusterId`

The unique identifier of the EMR cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceFleetType`

The node type that the instance fleet hosts.

_Allowed Values_: TASK

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceTypeConfigs`

`InstanceTypeConfigs` determine the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.

###### Note

The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.

_Required_: No

_Type_: Array of [InstanceTypeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-instancefleetconfig-instancetypeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchSpecifications`

The launch specification for the instance fleet.

_Required_: No

_Type_: [InstanceFleetProvisioningSpecifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-instancefleetconfig-instancefleetprovisioningspecifications.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The friendly name of the instance fleet.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResizeSpecifications`

The resize specification for the instance fleet.

_Required_: No

_Type_: [InstanceFleetResizingSpecifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-instancefleetconfig-instancefleetresizingspecifications.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetOnDemandCapacity`

The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision. When the instance fleet launches, Amazon EMR tries to provision On-Demand instances as specified by `InstanceTypeConfig`. Each instance configuration has a specified `WeightedCapacity`. When an On-Demand instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.

###### Note

If not specified or set to 0, only Spot instances are provisioned for the instance fleet using `TargetSpotCapacity`. At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetSpotCapacity`

The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision. When the instance fleet launches, Amazon EMR tries to provision Spot instances as specified by `InstanceTypeConfig`. Each instance configuration has a specified `WeightedCapacity`. When a Spot instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.

###### Note

If not specified or set to 0, only On-Demand instances are provisioned for the instance fleet. At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns returns the ID of the instance fleet.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VolumeSpecification

Configuration
