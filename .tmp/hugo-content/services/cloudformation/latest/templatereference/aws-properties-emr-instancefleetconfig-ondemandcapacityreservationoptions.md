This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceFleetConfig OnDemandCapacityReservationOptions

Describes the strategy for using unused Capacity Reservations for fulfilling On-Demand
capacity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityReservationPreference" : String,
  "CapacityReservationResourceGroupArn" : String,
  "UsageStrategy" : String
}

```

### YAML

```yaml

  CapacityReservationPreference: String
  CapacityReservationResourceGroupArn: String
  UsageStrategy: String

```

## Properties

`CapacityReservationPreference`

Indicates the instance's Capacity Reservation preferences. Possible preferences
include:

- `open` \- The instance can run in any open Capacity Reservation that has
matching attributes (instance type, platform, Availability Zone).

- `none` \- The instance avoids running in a Capacity Reservation even if
one is available. The instance runs as an On-Demand Instance.

_Required_: No

_Type_: String

_Allowed values_: `open | none`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityReservationResourceGroupArn`

The ARN of the Capacity Reservation resource group in which to run the instance.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsageStrategy`

Indicates whether to use unused Capacity Reservations for fulfilling On-Demand
capacity.

If you specify `use-capacity-reservations-first`, the fleet uses unused
Capacity Reservations to fulfill On-Demand capacity up to the target On-Demand capacity. If
multiple instance pools have unused Capacity Reservations, the On-Demand allocation
strategy ( `lowest-price`) is applied. If the number of unused Capacity
Reservations is less than the On-Demand target capacity, the remaining On-Demand target
capacity is launched according to the On-Demand allocation strategy
( `lowest-price`).

If you do not specify a value, the fleet fulfills the On-Demand capacity according to
the chosen On-Demand allocation strategy.

_Required_: No

_Type_: String

_Allowed values_: `use-capacity-reservations-first`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceTypeConfig

OnDemandProvisioningSpecification

All content copied from https://docs.aws.amazon.com/.
