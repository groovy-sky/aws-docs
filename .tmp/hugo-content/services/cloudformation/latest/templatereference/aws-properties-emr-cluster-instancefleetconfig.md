This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster InstanceFleetConfig

Use `InstanceFleetConfig` to define instance fleets for an EMR cluster. A cluster can not use both instance fleets and instance groups. For more information, see [Configure Instance Fleets](../../../emr/latest/managementguide/emr-instance-group-configuration.md) in the _Amazon EMR Management Guide_.

###### Note

The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceTypeConfigs" : [ InstanceTypeConfig, ... ],
  "LaunchSpecifications" : InstanceFleetProvisioningSpecifications,
  "Name" : String,
  "ResizeSpecifications" : InstanceFleetResizingSpecifications,
  "TargetOnDemandCapacity" : Integer,
  "TargetSpotCapacity" : Integer
}

```

### YAML

```yaml

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

`InstanceTypeConfigs`

The instance type configurations that define the Amazon EC2 instances in the
instance fleet.

_Required_: No

_Type_: Array of [InstanceTypeConfig](aws-properties-emr-cluster-instancetypeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchSpecifications`

The launch specification for the instance fleet.

_Required_: No

_Type_: [InstanceFleetProvisioningSpecifications](aws-properties-emr-cluster-instancefleetprovisioningspecifications.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The friendly name of the instance fleet.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResizeSpecifications`

The resize specification for the instance fleet.

_Required_: No

_Type_: [InstanceFleetResizingSpecifications](aws-properties-emr-cluster-instancefleetresizingspecifications.md)

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HadoopJarStepConfig

InstanceFleetProvisioningSpecifications

All content copied from https://docs.aws.amazon.com/.
