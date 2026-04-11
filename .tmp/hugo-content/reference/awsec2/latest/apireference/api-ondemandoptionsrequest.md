# OnDemandOptionsRequest

Describes the configuration of On-Demand Instances in an EC2 Fleet.

## Contents

**AllocationStrategy**

The strategy that determines the order of the launch template overrides to use in
fulfilling On-Demand capacity.

`lowest-price` \- EC2 Fleet uses price to determine the order, launching the lowest
price first.

`prioritized` \- EC2 Fleet uses the priority that you assigned to each launch
template override, launching the highest priority first.

Default: `lowest-price`

Type: String

Valid Values: `lowest-price | prioritized`

Required: No

**CapacityReservationOptions**

The strategy for using unused Capacity Reservations for fulfilling On-Demand
capacity.

Supported only for fleets of type `instant`.

Type: [CapacityReservationOptionsRequest](api-capacityreservationoptionsrequest.md) object

Required: No

**MaxTotalPrice**

The maximum amount per hour for On-Demand Instances that you're willing to pay.

###### Note

If your fleet includes T instances that are configured as `unlimited`,
and if their average CPU usage exceeds the baseline utilization, you will incur a charge
for surplus credits. The `MaxTotalPrice` does not account for surplus
credits, and, if you use surplus credits, your final cost might be higher than what you
specified for `MaxTotalPrice`. For more information, see [Surplus credits can incur charges](../../../../services/ec2/latest/userguide/burstable-performance-instances-unlimited-mode-concepts.md#unlimited-mode-surplus-credits) in the _Amazon EC2 User Guide_.

Type: String

Required: No

**MinTargetCapacity**

The minimum target capacity for On-Demand Instances in the fleet. If this minimum capacity isn't
reached, no instances are launched.

Constraints: Maximum value of `1000`. Supported only for fleets of type
`instant`.

At least one of the following must be specified: `SingleAvailabilityZone` \|
`SingleInstanceType`

Type: Integer

Required: No

**SingleAvailabilityZone**

Indicates that the fleet launches all On-Demand Instances into a single Availability Zone.

Supported only for fleets of type `instant`.

Type: Boolean

Required: No

**SingleInstanceType**

Indicates that the fleet uses a single instance type to launch all On-Demand Instances in the
fleet.

Supported only for fleets of type `instant`.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ondemandoptionsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ondemandoptionsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ondemandoptionsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

OnDemandOptions

OperatorRequest
