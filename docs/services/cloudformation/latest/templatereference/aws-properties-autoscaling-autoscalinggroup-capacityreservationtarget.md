This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup CapacityReservationTarget

The target for the Capacity Reservation. Specify Capacity Reservations IDs or Capacity Reservation resource group ARNs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityReservationIds" : [ String, ... ],
  "CapacityReservationResourceGroupArns" : [ String, ... ]
}

```

### YAML

```yaml

  CapacityReservationIds:
    - String
  CapacityReservationResourceGroupArns:
    - String

```

## Properties

`CapacityReservationIds`

The Capacity Reservation IDs to launch instances into.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityReservationResourceGroupArns`

The resource group ARNs of the Capacity Reservation to launch instances into.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityReservationSpecification

CpuPerformanceFactorRequest

All content copied from https://docs.aws.amazon.com/.
