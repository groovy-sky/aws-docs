This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet OnDemandOptionsRequest

Specifies the allocation strategy of On-Demand Instances in an EC2 Fleet.

`OnDemandOptionsRequest` is a property of the [AWS::EC2::EC2Fleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String,
  "CapacityReservationOptions" : CapacityReservationOptionsRequest,
  "MaxTotalPrice" : String,
  "MinTargetCapacity" : Integer,
  "SingleAvailabilityZone" : Boolean,
  "SingleInstanceType" : Boolean
}

```

### YAML

```yaml

  AllocationStrategy: String
  CapacityReservationOptions:
    CapacityReservationOptionsRequest
  MaxTotalPrice: String
  MinTargetCapacity: Integer
  SingleAvailabilityZone: Boolean
  SingleInstanceType: Boolean

```

## Properties

`AllocationStrategy`

The strategy that determines the order of the launch template overrides to use in
fulfilling On-Demand capacity.

`lowest-price` \- EC2 Fleet uses price to determine the order, launching the lowest
price first.

`prioritized` \- EC2 Fleet uses the priority that you assigned to each launch
template override, launching the highest priority first.

Default: `lowest-price`

_Required_: No

_Type_: String

_Allowed values_: `lowest-price | prioritized`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CapacityReservationOptions`

The strategy for using unused Capacity Reservations for fulfilling On-Demand
capacity.

Supported only for fleets of type `instant`.

_Required_: No

_Type_: [CapacityReservationOptionsRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ec2fleet-capacityreservationoptionsrequest.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxTotalPrice`

The maximum amount per hour for On-Demand Instances that you're willing to pay.

###### Note

If your fleet includes T instances that are configured as `unlimited`,
and if their average CPU usage exceeds the baseline utilization, you will incur a charge
for surplus credits. The `MaxTotalPrice` does not account for surplus
credits, and, if you use surplus credits, your final cost might be higher than what you
specified for `MaxTotalPrice`. For more information, see [Surplus credits can incur charges](../../../ec2/latest/userguide/burstable-performance-instances-unlimited-mode-concepts.md#unlimited-mode-surplus-credits) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinTargetCapacity`

The minimum target capacity for On-Demand Instances in the fleet. If this minimum capacity isn't
reached, no instances are launched.

Constraints: Maximum value of `1000`. Supported only for fleets of type
`instant`.

At least one of the following must be specified: `SingleAvailabilityZone` \|
`SingleInstanceType`

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SingleAvailabilityZone`

Indicates that the fleet launches all On-Demand Instances into a single Availability Zone.

Supported only for fleets of type `instant`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SingleInstanceType`

Indicates that the fleet uses a single instance type to launch all On-Demand Instances in the
fleet.

Supported only for fleets of type `instant`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [OnDemandOptionsRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_OnDemandOptionsRequest.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkInterfaceCountRequest

PerformanceFactorReferenceRequest
