This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet SpotOptionsRequest

Specifies the configuration of Spot Instances for an EC2 Fleet.

`SpotOptionsRequest` is a property of the [AWS::EC2::EC2Fleet](../userguide/aws-resource-ec2-ec2fleet.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String,
  "InstanceInterruptionBehavior" : String,
  "InstancePoolsToUseCount" : Integer,
  "MaintenanceStrategies" : MaintenanceStrategies,
  "MaxTotalPrice" : String,
  "MinTargetCapacity" : Integer,
  "SingleAvailabilityZone" : Boolean,
  "SingleInstanceType" : Boolean
}

```

### YAML

```yaml

  AllocationStrategy: String
  InstanceInterruptionBehavior: String
  InstancePoolsToUseCount: Integer
  MaintenanceStrategies:
    MaintenanceStrategies
  MaxTotalPrice: String
  MinTargetCapacity: Integer
  SingleAvailabilityZone: Boolean
  SingleInstanceType: Boolean

```

## Properties

`AllocationStrategy`

Indicates how to allocate the target Spot Instance capacity across the Spot Instance
pools specified by the EC2 Fleet.

If the allocation strategy is `lowestPrice`, EC2 Fleet launches instances
from the Spot Instance pools with the lowest price. This is the default allocation
strategy.

If the allocation strategy is `diversified`, EC2 Fleet launches instances
from all the Spot Instance pools that you specify.

If the allocation strategy is `capacityOptimized`, EC2 Fleet launches
instances from Spot Instance pools that are optimally chosen based on the available Spot
Instance capacity.

_Allowed Values_: `lowestPrice` \| `diversified`
\| `capacityOptimized` \| `capacityOptimizedPrioritized`

_Required_: No

_Type_: String

_Allowed values_: `lowest-price | lowestPrice | diversified | capacityOptimized | capacity-optimized | capacityOptimizedPrioritized | capacity-optimized-prioritized | priceCapacityOptimized | price-capacity-optimized`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceInterruptionBehavior`

The behavior when a Spot Instance is interrupted.

Default: `terminate`

_Required_: No

_Type_: String

_Allowed values_: `hibernate | stop | terminate`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstancePoolsToUseCount`

The number of Spot pools across which to allocate your target Spot capacity. Supported
only when Spot `AllocationStrategy` is set to `lowest-price`. EC2 Fleet
selects the cheapest Spot pools and evenly allocates your target Spot capacity across the
number of Spot pools that you specify.

Note that EC2 Fleet attempts to draw Spot Instances from the number of pools that you specify on a
best effort basis. If a pool runs out of Spot capacity before fulfilling your target
capacity, EC2 Fleet will continue to fulfill your request by drawing from the next cheapest
pool. To ensure that your target capacity is met, you might receive Spot Instances from more than
the number of pools that you specified. Similarly, if most of the pools have no Spot
capacity, you might receive your full target capacity from fewer than the number of pools
that you specified.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaintenanceStrategies`

The strategies for managing your Spot Instances that are at an elevated risk of being
interrupted.

_Required_: No

_Type_: [MaintenanceStrategies](aws-properties-ec2-ec2fleet-maintenancestrategies.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxTotalPrice`

The maximum amount per hour for Spot Instances that you're willing to pay. We do not recommend
using this parameter because it can lead to increased interruptions. If you do not specify
this parameter, you will pay the current Spot price.

###### Important

If you specify a maximum price, your Spot Instances will be interrupted more frequently than if you do not specify this parameter.

###### Note

If your fleet includes T instances that are configured as `unlimited`, and
if their average CPU usage exceeds the baseline utilization, you will incur a charge for
surplus credits. The `MaxTotalPrice` does not account for surplus credits,
and, if you use surplus credits, your final cost might be higher than what you specified
for `MaxTotalPrice`. For more information, see [Surplus credits can incur charges](../../../ec2/latest/userguide/burstable-performance-instances-unlimited-mode-concepts.md#unlimited-mode-surplus-credits) in the
_Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinTargetCapacity`

The minimum target capacity for Spot Instances in the fleet. If this minimum capacity isn't
reached, no instances are launched.

Constraints: Maximum value of `1000`. Supported only for fleets of type
`instant`.

At least one of the following must be specified: `SingleAvailabilityZone` \|
`SingleInstanceType`

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SingleAvailabilityZone`

Indicates that the fleet launches all Spot Instances into a single Availability Zone.

Supported only for fleets of type `instant`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SingleInstanceType`

Indicates that the fleet uses a single instance type to launch all Spot Instances in the
fleet.

Supported only for fleets of type `instant`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [SpotOptionsRequest](../../../../reference/awsec2/latest/apireference/api-spotoptionsrequest.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReservedCapacityOptionsRequest

Tag

All content copied from https://docs.aws.amazon.com/.
