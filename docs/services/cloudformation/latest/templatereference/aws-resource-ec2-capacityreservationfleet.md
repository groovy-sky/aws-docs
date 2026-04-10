This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::CapacityReservationFleet

Creates a new Capacity Reservation Fleet with the specified attributes. For more
information, see [Capacity Reservation Fleets](../../../ec2/latest/userguide/cr-fleets.md) in the _Amazon EC2 User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::CapacityReservationFleet",
  "Properties" : {
      "AllocationStrategy" : String,
      "EndDate" : String,
      "InstanceMatchCriteria" : String,
      "InstanceTypeSpecifications" : [ InstanceTypeSpecification, ... ],
      "NoRemoveEndDate" : Boolean,
      "RemoveEndDate" : Boolean,
      "TagSpecifications" : [ TagSpecification, ... ],
      "Tenancy" : String,
      "TotalTargetCapacity" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EC2::CapacityReservationFleet
Properties:
  AllocationStrategy: String
  EndDate: String
  InstanceMatchCriteria: String
  InstanceTypeSpecifications:
    - InstanceTypeSpecification
  NoRemoveEndDate: Boolean
  RemoveEndDate: Boolean
  TagSpecifications:
    - TagSpecification
  Tenancy: String
  TotalTargetCapacity: Integer

```

## Properties

`AllocationStrategy`

The strategy used by the Capacity Reservation Fleet to determine which of the
specified instance types to use. Currently, only the `prioritized` allocation
strategy is supported. For more information, see [Allocation\
strategy](../../../ec2/latest/userguide/crfleet-concepts.md#allocation-strategy) in the _Amazon EC2 User Guide_.

Valid values: `prioritized`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndDate`

The date and time at which the Capacity Reservation Fleet expires. When the Capacity
Reservation Fleet expires, its state changes to `expired` and all of the
Capacity Reservations in the Fleet expire.

The Capacity Reservation Fleet expires within an hour after the specified time. For
example, if you specify `5/31/2019`, `13:30:55`, the Capacity
Reservation Fleet is guaranteed to expire between `13:30:55` and
`14:30:55` on `5/31/2019`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceMatchCriteria`

Indicates the type of instance launches that the Capacity Reservation Fleet accepts.
All Capacity Reservations in the Fleet inherit this instance matching criteria.

Currently, Capacity Reservation Fleets support `open` instance matching
criteria only. This means that instances that have matching attributes (instance type,
platform, and Availability Zone) run in the Capacity Reservations automatically.
Instances do not need to explicitly target a Capacity Reservation Fleet to use its
reserved capacity.

_Required_: No

_Type_: String

_Allowed values_: `open`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceTypeSpecifications`

Information about the instance types for which to reserve the capacity.

_Required_: No

_Type_: Array of [InstanceTypeSpecification](aws-properties-ec2-capacityreservationfleet-instancetypespecification.md)

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NoRemoveEndDate`

Used to add an end date to a Capacity Reservation Fleet that has no end date and
time. To add an end date to a Capacity Reservation Fleet, specify `true`
for this paramater and specify the end date and time (in UTC time format) for the
**EndDate** parameter.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveEndDate`

Used to remove an end date from a Capacity Reservation Fleet that is configured
to end automatically at a specific date and time. To remove the end date from a
Capacity Reservation Fleet, specify `true` for this paramater and omit
the **EndDate** parameter.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagSpecifications`

The tags to assign to the Capacity Reservation Fleet. The tags are automatically
assigned to the Capacity Reservations in the Fleet.

_Required_: No

_Type_: Array of [TagSpecification](aws-properties-ec2-capacityreservationfleet-tagspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tenancy`

Indicates the tenancy of the Capacity Reservation Fleet. All Capacity Reservations in
the Fleet inherit this tenancy. The Capacity Reservation Fleet can have one of the
following tenancy settings:

- `default` \- The Capacity Reservation Fleet is created on hardware
that is shared with other AWS accounts.

- `dedicated` \- The Capacity Reservations are created on single-tenant
hardware that is dedicated to a single AWS account.

_Required_: No

_Type_: String

_Allowed values_: `default`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TotalTargetCapacity`

The total number of capacity units to be reserved by the Capacity Reservation Fleet.
This value, together with the instance type weights that you assign to each instance
type used by the Fleet determine the number of instances for which the Fleet reserves
capacity. Both values are based on units that make sense for your workload. For more
information, see [Total target\
capacity](../../../ec2/latest/userguide/crfleet-concepts.md#target-capacity) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `25000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ID, such as
`crf-abcdef01234567890`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CapacityReservationFleetId`

The ID of the Capacity Reservation Fleet.

## Examples

### Creating a Capacity Reservation Fleet

The following example creates a Capacity Feservation Fleet that creates Capacity
Reservations for 2 `c4.large` or `c5.large` instances. The
Fleet prioritizes `c4.large` over `c5.large` and both instance
types have the same weight.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  EC2CapacityReservationFleetCanary:
    Type: 'AWS::EC2::CapacityReservationFleet'
    Properties:
      AllocationStrategy: "prioritized"
      InstanceTypeSpecifications:
        - InstanceType: "c4.large"
          InstancePlatform: "Linux/UNIX"
          AvailabilityZone: "us-east-1a"
          Weight: 1
          Priority: 1
        - InstanceType: "c5.large"
          InstancePlatform: "Linux/UNIX"
          AvailabilityZone: "us-east-1a"
          Weight: 1
          Priority: 2
      Tenancy: "default"
      TotalTargetCapacity: 2
      InstanceMatchCriteria: "open"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagSpecification

InstanceTypeSpecification

All content copied from https://docs.aws.amazon.com/.
