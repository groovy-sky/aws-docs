This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup CapacityReservationSpecification

Describes the Capacity Reservation preference and targeting options. If you specify `open` or `none` for `CapacityReservationPreference`, do not specify a `CapacityReservationTarget`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityReservationPreference" : String,
  "CapacityReservationTarget" : CapacityReservationTarget
}

```

### YAML

```yaml

  CapacityReservationPreference: String
  CapacityReservationTarget:
    CapacityReservationTarget

```

## Properties

`CapacityReservationPreference`

The capacity reservation preference. The following options are available:

- `capacity-reservations-only` \- Auto Scaling will only launch instances into a Capacity Reservation or Capacity Reservation resource group. If capacity isn't available, instances will fail to launch.

- `capacity-reservations-first` \- Auto Scaling will try to launch instances into a Capacity Reservation or Capacity Reservation resource group first. If capacity isn't available, instances will run in On-Demand capacity.

- `none` \- Auto Scaling will not launch instances into a Capacity Reservation. Instances will run in On-Demand capacity.

- `default` \- Auto Scaling uses the Capacity Reservation preference from your launch template or an open Capacity Reservation.

_Required_: Yes

_Type_: String

_Allowed values_: `capacity-reservations-only | capacity-reservations-first | none | default`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityReservationTarget`

Describes a target Capacity Reservation or Capacity Reservation resource group.

_Required_: No

_Type_: [CapacityReservationTarget](aws-properties-autoscaling-autoscalinggroup-capacityreservationtarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BaselinePerformanceFactorsRequest

CapacityReservationTarget

All content copied from https://docs.aws.amazon.com/.
