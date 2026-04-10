This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet ReservedCapacityOptionsRequest

Defines EC2 Fleet preferences for utilizing reserved capacity when DefaultTargetCapacityType is set to `reserved-capacity`.

###### Note

This configuration can only be used if the EC2 Fleet is of type
`instant`.

When you specify `ReservedCapacityOptions`, you must also set
`DefaultTargetCapacityType` to `reserved-capacity` in the
`TargetCapacitySpecification`.

For more information about Interruptible Capacity Reservations, see [Launch\
instances into an Interruptible Capacity Reservation](../../../ec2/latest/userguide/ec2-fleet-launch-instances-interruptible-cr-walkthrough.md) in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReservationTypes" : [ String, ... ]
}

```

### YAML

```yaml

  ReservationTypes:
    - String

```

## Properties

`ReservationTypes`

The types of Capacity Reservations to use for fulfilling the EC2 Fleet request.

_Required_: No

_Type_: Array of String

_Allowed values_: `interruptible-capacity-reservation`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Placement

SpotOptionsRequest

All content copied from https://docs.aws.amazon.com/.
