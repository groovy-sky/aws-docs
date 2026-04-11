This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet CapacityReservationOptionsRequest

Describes the strategy for using unused Capacity Reservations for fulfilling On-Demand
capacity.

###### Note

This strategy can only be used if the EC2 Fleet is of type `instant`.

For more information about Capacity Reservations, see [On-Demand Capacity\
Reservations](../../../ec2/latest/userguide/ec2-capacity-reservations.md) in the _Amazon EC2 User Guide_. For examples of using
Capacity Reservations in an EC2 Fleet, see [EC2 Fleet example\
configurations](../../../ec2/latest/userguide/ec2-fleet-examples.md) in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UsageStrategy" : String
}

```

### YAML

```yaml

  UsageStrategy: String

```

## Properties

`UsageStrategy`

Indicates whether to use unused Capacity Reservations for fulfilling On-Demand capacity.

If you specify `use-capacity-reservations-first`, the fleet uses unused
Capacity Reservations to fulfill On-Demand capacity up to the target On-Demand capacity. If
multiple instance pools have unused Capacity Reservations, the On-Demand allocation
strategy ( `lowest-price` or `prioritized`) is applied. If the number
of unused Capacity Reservations is less than the On-Demand target capacity, the remaining
On-Demand target capacity is launched according to the On-Demand allocation strategy
( `lowest-price` or `prioritized`).

If you do not specify a value, the fleet fulfils the On-Demand capacity according to the
chosen On-Demand allocation strategy.

_Required_: No

_Type_: String

_Allowed values_: `use-capacity-reservations-first`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityRebalance

CpuPerformanceFactorRequest

All content copied from https://docs.aws.amazon.com/.
