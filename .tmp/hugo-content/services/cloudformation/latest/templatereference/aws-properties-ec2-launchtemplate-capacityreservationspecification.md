This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate CapacityReservationSpecification

Specifies an instance's Capacity Reservation targeting option. You can specify only one
option at a time.

`CapacityReservationSpecification` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

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

Indicates the instance's Capacity Reservation preferences. Possible preferences
include:

- `capacity-reservations-only` \- The instance will only run in a
Capacity Reservation or Capacity Reservation group. If capacity isn't available,
the instance will fail to launch.

- `open` \- The instance can run in any `open` Capacity
Reservation that has matching attributes (instance type, platform, Availability
Zone, tenancy).

- `none` \- The instance avoids running in a Capacity Reservation even
if one is available. The instance runs in On-Demand capacity.

_Required_: No

_Type_: String

_Allowed values_: `capacity-reservations-only | open | none`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityReservationTarget`

Information about the target Capacity Reservation or Capacity Reservation
group.

_Required_: No

_Type_: [CapacityReservationTarget](aws-properties-ec2-launchtemplate-capacityreservationtarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlockDeviceMapping

CapacityReservationTarget

All content copied from https://docs.aws.amazon.com/.
