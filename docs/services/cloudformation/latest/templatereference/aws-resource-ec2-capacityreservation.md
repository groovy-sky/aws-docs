This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::CapacityReservation

Creates a new Capacity Reservation with the specified attributes. For more information,
see [Capacity\
Reservations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html) in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::CapacityReservation",
  "Properties" : {
      "AvailabilityZone" : String,
      "AvailabilityZoneId" : String,
      "EbsOptimized" : Boolean,
      "EndDate" : String,
      "EndDateType" : String,
      "EphemeralStorage" : Boolean,
      "InstanceCount" : Integer,
      "InstanceMatchCriteria" : String,
      "InstancePlatform" : String,
      "InstanceType" : String,
      "OutPostArn" : String,
      "PlacementGroupArn" : String,
      "TagSpecifications" : [ TagSpecification, ... ],
      "Tenancy" : String,
      "UnusedReservationBillingOwnerId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::CapacityReservation
Properties:
  AvailabilityZone: String
  AvailabilityZoneId: String
  EbsOptimized: Boolean
  EndDate: String
  EndDateType: String
  EphemeralStorage: Boolean
  InstanceCount: Integer
  InstanceMatchCriteria: String
  InstancePlatform: String
  InstanceType: String
  OutPostArn: String
  PlacementGroupArn: String
  TagSpecifications:
    - TagSpecification
  Tenancy: String
  UnusedReservationBillingOwnerId: String

```

## Properties

`AvailabilityZone`

The Availability Zone in which to create the Capacity Reservation.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityZoneId`

The ID of the Availability Zone in which the capacity is reserved.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EbsOptimized`

Indicates whether the Capacity Reservation supports EBS-optimized instances. This
optimization provides dedicated throughput to Amazon EBS and an optimized configuration
stack to provide optimal I/O performance. This optimization isn't available with all
instance types. Additional usage charges apply when using an EBS- optimized
instance.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndDate`

The date and time at which the Capacity Reservation expires. When a Capacity
Reservation expires, the reserved capacity is released and you can no longer launch
instances into it. The Capacity Reservation's state changes to `expired` when
it reaches its end date and time.

You must provide an `EndDate` value if `EndDateType` is
`limited`. Omit `EndDate` if `EndDateType` is
`unlimited`.

If the `EndDateType` is `limited`, the Capacity Reservation is
cancelled within an hour from the specified time. For example, if you specify 5/31/2019,
13:30:55, the Capacity Reservation is guaranteed to end between 13:30:55 and 14:30:55 on
5/31/2019.

If you are requesting a future-dated Capacity Reservation, you can't specify an end
date and time that is within the commitment duration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndDateType`

Indicates the way in which the Capacity Reservation ends. A Capacity Reservation can
have one of the following end types:

- `unlimited` \- The Capacity Reservation remains active until you
explicitly cancel it. Do not provide an `EndDate` if the
`EndDateType` is `unlimited`.

- `limited` \- The Capacity Reservation expires automatically at a
specified date and time. You must provide an `EndDate` value if the
`EndDateType` value is `limited`.

_Required_: No

_Type_: String

_Allowed values_: `unlimited | limited`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EphemeralStorage`

_Deprecated._

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceCount`

The number of instances for which to reserve capacity.

###### Note

You can request future-dated Capacity Reservations for an instance count with a
minimum of 32 vCPUs. For example, if you request a future-dated Capacity
Reservation for `m5.xlarge` instances, you must request at least 8
instances ( _8 \* m5.xlarge = 32 vCPUs_).

Valid range: 1 - 1000

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceMatchCriteria`

Indicates the type of instance launches that the Capacity Reservation accepts. The
options include:

- `open` \- The Capacity Reservation automatically matches all instances
that have matching attributes (instance type, platform, and Availability Zone).
Instances that have matching attributes run in the Capacity Reservation
automatically without specifying any additional parameters.

- `targeted` \- The Capacity Reservation only accepts instances that
have matching attributes (instance type, platform, and Availability Zone), and
explicitly target the Capacity Reservation. This ensures that only permitted
instances can use the reserved capacity.

###### Note

If you are requesting a future-dated Capacity Reservation, you must specify
`targeted`.

Default: `open`

_Required_: No

_Type_: String

_Allowed values_: `open | targeted`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstancePlatform`

The type of operating system for which to reserve capacity.

_Required_: Yes

_Type_: String

_Allowed values_: `Linux/UNIX | Red Hat Enterprise Linux | SUSE Linux | Windows | Windows with SQL Server | Windows with SQL Server Enterprise | Windows with SQL Server Standard | Windows with SQL Server Web | Linux with SQL Server Standard | Linux with SQL Server Web | Linux with SQL Server Enterprise | RHEL with SQL Server Standard | RHEL with SQL Server Enterprise | RHEL with SQL Server Web | RHEL with HA | RHEL with HA and SQL Server Standard | RHEL with HA and SQL Server Enterprise | Ubuntu Pro`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceType`

The instance type for which to reserve capacity.

###### Note

You can request future-dated Capacity Reservations for instance types in the C, M,
R, I, T, and G instance families only.

For more information, see [Instance types](../../../ec2/latest/userguide/instance-types.md) in the
_Amazon EC2 User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutPostArn`

###### Note

Not supported for future-dated Capacity Reservations.

The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity
Reservation.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws([a-z-]+)?:outposts:[a-z\d-]+:\d{12}:outpost/op-[a-f0-9]{17}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlacementGroupArn`

###### Note

Not supported for future-dated Capacity Reservations.

The Amazon Resource Name (ARN) of the cluster placement group in which to create the
Capacity Reservation. For more information, see [Capacity Reservations for cluster\
placement groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/cr-cpg.html) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws([a-z-]+)?:ec2:[a-z\d-]+:\d{12}:placement-group/^.{1,255}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagSpecifications`

The tags to apply to the Capacity Reservation during launch.

_Required_: No

_Type_: Array of [TagSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-capacityreservation-tagspecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tenancy`

Indicates the tenancy of the Capacity Reservation. A Capacity Reservation can have one
of the following tenancy settings:

- `default` \- The Capacity Reservation is created on hardware that is
shared with other AWS accounts.

- `dedicated` \- The Capacity Reservation is created on single-tenant
hardware that is dedicated to a single AWS account.

_Required_: No

_Type_: String

_Allowed values_: `default | dedicated`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UnusedReservationBillingOwnerId`

The ID of the AWS account to which to assign billing of the unused capacity of the Capacity
Reservation. A request will be sent to the specified account. That account must accept the request for the billing
to be assigned to their account. For more information, see [Billing assignment for shared Amazon EC2 Capacity Reservations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/assign-billing.html).

You can assign billing only for shared Capacity Reservations. To share a Capacity Reservation, you must add it
to a resource share. For more information, see [AWS::RAM::ResourceShare](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html).

_Required_: No

_Type_: String

_Pattern_: `[0-9]{12}`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ID, such as
`cr-1234ab5cd6789e0f1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AvailabilityZone`

Returns the Availability Zone in which the capacity is reserved. For example:
`us-east-1a`.

`AvailableInstanceCount`

Returns the remaining capacity, which indicates the number of instances that can be
launched in the Capacity Reservation. For example: `9`.

`CapacityAllocationSet`

Property description not available.

`CapacityReservationArn`

The Amazon Resource Name (ARN) of the Capacity Reservation.

`CapacityReservationFleetId`

The ID of the Capacity Reservation Fleet to which the Capacity Reservation belongs.
Only valid for Capacity Reservations that were created by a Capacity Reservation
Fleet.

`CreateDate`

The date and time the Capacity Reservation was created.

`DeliveryPreference`

The delivery method for a future-dated Capacity Reservation. `incremental`
indicates that the requested capacity is delivered in addition to any running instances
and reserved capacity that you have in your account at the requested date and
time.

`Id`

The ID of the Capacity Reservation.

`InstanceType`

Returns the type of instance for which the capacity is reserved. For example:
`m4.large`.

`OwnerId`

The ID of the AWS account that owns the Capacity Reservation.

`ReservationType`

The type of Capacity Reservation.

`StartDate`

The date and time the Capacity Reservation was started.

`State`

The current state of the Capacity Reservation. A Capacity Reservation can be in one of
the following states:

- `active` \- The capacity is available for use.

- `expired` \- The Capacity Reservation expired automatically at the date and time
specified in your reservation request. The reserved capacity is no longer available for your use.

- `cancelled` \- The Capacity Reservation was canceled. The reserved capacity is no
longer available for your use.

- `pending` \- The Capacity Reservation request was successful but the capacity
provisioning is still pending.

- `failed` \- The Capacity Reservation request has failed. A request can fail due to
request parameters that are not valid, capacity constraints, or instance limit constraints. You
can view a failed request for 60 minutes.

- `scheduled` \- ( _Future-dated Capacity Reservations_) The
future-dated Capacity Reservation request was approved and the Capacity Reservation is scheduled
for delivery on the requested start date.

- `payment-pending` \- ( _Capacity Blocks_) The upfront
payment has not been processed yet.

- `payment-failed` \- ( _Capacity Blocks_) The upfront
payment was not processed in the 12-hour time frame. Your Capacity Block was released.

- `assessing` \- ( _Future-dated Capacity Reservations_)
Amazon EC2 is assessing your request for a future-dated Capacity Reservation.

- `delayed` \- ( _Future-dated Capacity Reservations_) Amazon EC2
encountered a delay in provisioning the requested future-dated Capacity Reservation. Amazon EC2 is
unable to deliver the requested capacity by the requested start date and time.

- `unsupported` \- ( _Future-dated Capacity Reservations_) Amazon EC2
can't support the future-dated Capacity Reservation request due to capacity constraints. You can view
unsupported requests for 30 days. The Capacity Reservation will not be delivered.

`Tenancy`

Returns the tenancy of the Capacity Reservation. For example:
`dedicated`.

`TotalInstanceCount`

Returns the total number of instances for which the Capacity Reservation reserves
capacity. For example: `15`.

## See also

- [On-Demand\
Capacity Reservations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html) in the _Amazon EC2 User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CapacityAllocation
