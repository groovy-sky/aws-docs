# CapacityReservationFleet

Information about a Capacity Reservation Fleet.

## Contents

**allocationStrategy**

The strategy used by the Capacity Reservation Fleet to determine which of the
specified instance types to use. For more information, see For more information, see
[Allocation\
strategy](../../../../services/ec2/latest/userguide/crfleet-concepts.md#allocation-strategy) in the _Amazon EC2 User Guide_.

Type: String

Required: No

**capacityReservationFleetArn**

The ARN of the Capacity Reservation Fleet.

Type: String

Required: No

**capacityReservationFleetId**

The ID of the Capacity Reservation Fleet.

Type: String

Required: No

**createTime**

The date and time at which the Capacity Reservation Fleet was created.

Type: Timestamp

Required: No

**endDate**

The date and time at which the Capacity Reservation Fleet expires.

Type: Timestamp

Required: No

**instanceMatchCriteria**

Indicates the type of instance launches that the Capacity Reservation Fleet accepts.
All Capacity Reservations in the Fleet inherit this instance matching criteria.

Currently, Capacity Reservation Fleets support `open` instance matching
criteria only. This means that instances that have matching attributes (instance type,
platform, and Availability Zone) run in the Capacity Reservations automatically.
Instances do not need to explicitly target a Capacity Reservation Fleet to use its
reserved capacity.

Type: String

Valid Values: `open`

Required: No

**InstanceTypeSpecificationSet.N**

Information about the instance types for which to reserve the capacity.

Type: Array of [FleetCapacityReservation](api-fleetcapacityreservation.md) objects

Required: No

**state**

The state of the Capacity Reservation Fleet. Possible states include:

- `submitted` \- The Capacity Reservation Fleet request has been
submitted and Amazon Elastic Compute Cloud is preparing to create the Capacity
Reservations.

- `modifying` \- The Capacity Reservation Fleet is being modified. The
Fleet remains in this state until the modification is complete.

- `active` \- The Capacity Reservation Fleet has fulfilled its total
target capacity and it is attempting to maintain this capacity. The Fleet
remains in this state until it is modified or deleted.

- `partially_fulfilled` \- The Capacity Reservation Fleet has
partially fulfilled its total target capacity. There is insufficient Amazon EC2 to fulfill the total target capacity. The Fleet is attempting to
asynchronously fulfill its total target capacity.

- `expiring` \- The Capacity Reservation Fleet has reach its end date
and it is in the process of expiring. One or more of its Capacity reservations
might still be active.

- `expired` \- The Capacity Reservation Fleet has reach its end date.
The Fleet and its Capacity Reservations are expired. The Fleet can't create new
Capacity Reservations.

- `cancelling` \- The Capacity Reservation Fleet is in the process of
being cancelled. One or more of its Capacity reservations might still be
active.

- `cancelled` \- The Capacity Reservation Fleet has been manually
cancelled. The Fleet and its Capacity Reservations are cancelled and the Fleet
can't create new Capacity Reservations.

- `failed` \- The Capacity Reservation Fleet failed to reserve
capacity for the specified instance types.

Type: String

Valid Values: `submitted | modifying | active | partially_fulfilled | expiring | expired | cancelling | cancelled | failed`

Required: No

**TagSet.N**

The tags assigned to the Capacity Reservation Fleet.

Type: Array of [Tag](api-tag.md) objects

Required: No

**tenancy**

The tenancy of the Capacity Reservation Fleet. Tenancies include:

- `default` \- The Capacity Reservation Fleet is created on hardware
that is shared with other AWS accounts.

- `dedicated` \- The Capacity Reservation Fleet is created on
single-tenant hardware that is dedicated to a single AWS account.

Type: String

Valid Values: `default`

Required: No

**totalFulfilledCapacity**

The capacity units that have been fulfilled.

Type: Double

Required: No

**totalTargetCapacity**

The total number of capacity units for which the Capacity Reservation Fleet reserves
capacity. For more information, see [Total target\
capacity](../../../../services/ec2/latest/userguide/crfleet-concepts.md#target-capacity) in the _Amazon EC2 User Guide_.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityreservationfleet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityreservationfleet.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityreservationfleet.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityReservationCommitmentInfo

CapacityReservationFleetCancellationState
