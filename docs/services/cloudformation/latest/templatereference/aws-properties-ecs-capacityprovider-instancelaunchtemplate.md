This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider InstanceLaunchTemplate

The launch template configuration for Amazon ECS Managed Instances. This defines how
Amazon ECS launches Amazon EC2 instances, including the instance profile for your tasks,
network and storage configuration, capacity options, and instance requirements for
flexible instance type selection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityOptionType" : String,
  "CapacityReservations" : CapacityReservationRequest,
  "Ec2InstanceProfileArn" : String,
  "FipsEnabled" : Boolean,
  "InstanceMetadataTagsPropagation" : Boolean,
  "InstanceRequirements" : InstanceRequirementsRequest,
  "LocalStorageConfiguration" : ManagedInstancesLocalStorageConfiguration,
  "Monitoring" : String,
  "NetworkConfiguration" : ManagedInstancesNetworkConfiguration,
  "StorageConfiguration" : ManagedInstancesStorageConfiguration
}

```

### YAML

```yaml

  CapacityOptionType: String
  CapacityReservations:
    CapacityReservationRequest
  Ec2InstanceProfileArn: String
  FipsEnabled: Boolean
  InstanceMetadataTagsPropagation: Boolean
  InstanceRequirements:
    InstanceRequirementsRequest
  LocalStorageConfiguration:
    ManagedInstancesLocalStorageConfiguration
  Monitoring: String
  NetworkConfiguration:
    ManagedInstancesNetworkConfiguration
  StorageConfiguration:
    ManagedInstancesStorageConfiguration

```

## Properties

`CapacityOptionType`

The capacity option type. This determines whether
Amazon ECS launches On-Demand, Spot or Capacity Reservation Instances for your managed instance
capacity provider.

Valid values are:

- `ON_DEMAND` \- Launches standard On-Demand Instances.
On-Demand Instances provide predictable pricing and availability.

- `SPOT` \- Launches Spot Instances that use spare Amazon EC2 capacity
at reduced cost. Spot Instances can be interrupted by Amazon EC2 with a two-minute
notification when the capacity is needed back.

- `RESERVED` \- Launches Instances using Amazon EC2 Capacity
Reservations. Capacity Reservations allow you to reserve compute capacity for Amazon
EC2 instances in a specific Availability Zone.

The default is On-Demand

For more information about Amazon EC2 capacity options, see [Instance purchasing\
options](../../../ec2/latest/userguide/instance-purchasing-options.md) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `ON_DEMAND | SPOT | RESERVED`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`CapacityReservations`

Capacity reservation specifications. You can specify:

- Capacity reservation preference

- Reservation resource group to be used for targeted capacity reservations

Amazon ECS will launch instances according to the specified criteria.

_Required_: No

_Type_: [CapacityReservationRequest](aws-properties-ecs-capacityprovider-capacityreservationrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2InstanceProfileArn`

The Amazon Resource Name (ARN) of the instance profile that Amazon ECS applies to
Amazon ECS Managed Instances. This instance profile must include the necessary
permissions for your tasks to access AWS services and
resources.

For more information, see [Amazon\
ECS instance profile for Managed Instances](../../../amazonecs/latest/developerguide/managed-instances-instance-profile.md) in the _Amazon ECS_
_Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FipsEnabled`

Determines whether to enable FIPS 140-2 validated cryptographic modules on EC2 instances launched by the capacity provider. If `true`, instances use FIPS-compliant cryptographic algorithms and modules for enhanced security compliance. If `false`, instances use standard cryptographic implementations.

If not specified, instances are launched with FIPS enabled in AWS GovCloud (US) regions and FIPS disabled in other regions.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceMetadataTagsPropagation`

Determines whether tags are propagated to the instance metadata service (IMDS) for
Amazon EC2 instances launched by the Managed Instances capacity provider. When enabled, all tags associated
with the instance are available through the instance metadata service. When disabled,
tags are not propagated to IMDS.

Disable this setting if your tags contain characters that are not compatible with IMDS,
such as `/`. IMDS requires tag keys to match the pattern
`[0-9a-zA-Z\-_+=,.@:]{1,255}`.

The default value is `true`.

For more information, see [Work with\
instance tags in instance metadata](../../../ec2/latest/userguide/using-tags.md#work-with-tags-in-IMDS) in the _Amazon EC2 User_
_Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceRequirements`

The instance requirements. You can specify:

- The instance types

- Instance requirements such as vCPU count, memory, network performance, and
accelerator specifications

Amazon ECS automatically selects the instances that match the specified
criteria.

_Required_: No

_Type_: [InstanceRequirementsRequest](aws-properties-ecs-capacityprovider-instancerequirementsrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalStorageConfiguration`

The local storage configuration for Amazon ECS Managed Instances. This defines how ECS
uses instance store volumes available on the container instance.

_Required_: No

_Type_: [ManagedInstancesLocalStorageConfiguration](aws-properties-ecs-capacityprovider-managedinstanceslocalstorageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Monitoring`

CloudWatch provides two categories of monitoring: basic monitoring and detailed
monitoring. By default, your managed instance is configured for basic monitoring. You
can optionally enable detailed monitoring to help you more quickly identify and act on
operational issues. You can enable or turn off detailed monitoring at launch or when the
managed instance is running or stopped. For more information, see [Detailed monitoring for Amazon ECS Managed Instances](../../../amazonecs/latest/developerguide/detailed-monitoring-managed-instances.md) in the Amazon ECS
Developer Guide.

_Required_: No

_Type_: String

_Allowed values_: `BASIC | DETAILED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

The network configuration for Amazon ECS Managed Instances. This specifies the subnets
and security groups that instances use for network connectivity.

_Required_: Yes

_Type_: [ManagedInstancesNetworkConfiguration](aws-properties-ecs-capacityprovider-managedinstancesnetworkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageConfiguration`

The storage configuration for Amazon ECS Managed Instances. This defines the data
volume properties for the instances.

_Required_: No

_Type_: [ManagedInstancesStorageConfiguration](aws-properties-ecs-capacityprovider-managedinstancesstorageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InfrastructureOptimization

InstanceRequirementsRequest

All content copied from https://docs.aws.amazon.com/.
