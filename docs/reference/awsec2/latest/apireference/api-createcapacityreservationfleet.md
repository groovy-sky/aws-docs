# CreateCapacityReservationFleet

Creates a Capacity Reservation Fleet. For more information, see [Create a\
Capacity Reservation Fleet](../../../../services/ec2/latest/userguide/work-with-cr-fleets.md#create-crfleet) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllocationStrategy**

The strategy used by the Capacity Reservation Fleet to determine which of the
specified instance types to use. Currently, only the `prioritized` allocation
strategy is supported. For more information, see [Allocation\
strategy](../../../../services/ec2/latest/userguide/crfleet-concepts.md#allocation-strategy) in the _Amazon EC2 User Guide_.

Valid values: `prioritized`

Type: String

Required: No

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensure Idempotency](run-instance-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EndDate**

The date and time at which the Capacity Reservation Fleet expires. When the Capacity
Reservation Fleet expires, its state changes to `expired` and all of the
Capacity Reservations in the Fleet expire.

The Capacity Reservation Fleet expires within an hour after the specified time. For
example, if you specify `5/31/2019`, `13:30:55`, the Capacity
Reservation Fleet is guaranteed to expire between `13:30:55` and
`14:30:55` on `5/31/2019`.

Type: Timestamp

Required: No

**InstanceMatchCriteria**

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

**InstanceTypeSpecification.N**

Information about the instance types for which to reserve the capacity.

Type: Array of [ReservationFleetInstanceSpecification](api-reservationfleetinstancespecification.md) objects

Required: Yes

**TagSpecification.N**

The tags to assign to the Capacity Reservation Fleet. The tags are automatically
assigned to the Capacity Reservations in the Fleet.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**Tenancy**

Indicates the tenancy of the Capacity Reservation Fleet. All Capacity Reservations in
the Fleet inherit this tenancy. The Capacity Reservation Fleet can have one of the
following tenancy settings:

- `default` \- The Capacity Reservation Fleet is created on hardware
that is shared with other AWS accounts.

- `dedicated` \- The Capacity Reservations are created on single-tenant
hardware that is dedicated to a single AWS account.

Type: String

Valid Values: `default`

Required: No

**TotalTargetCapacity**

The total number of capacity units to be reserved by the Capacity Reservation Fleet.
This value, together with the instance type weights that you assign to each instance
type used by the Fleet determine the number of instances for which the Fleet reserves
capacity. Both values are based on units that make sense for your workload. For more
information, see [Total target\
capacity](../../../../services/ec2/latest/userguide/crfleet-concepts.md#target-capacity) in the _Amazon EC2 User Guide_.

Type: Integer

Required: Yes

## Response Elements

The following elements are returned by the service.

**allocationStrategy**

The allocation strategy used by the Capacity Reservation Fleet.

Type: String

**capacityReservationFleetId**

The ID of the Capacity Reservation Fleet.

Type: String

**createTime**

The date and time at which the Capacity Reservation Fleet was created.

Type: Timestamp

**endDate**

The date and time at which the Capacity Reservation Fleet expires.

Type: Timestamp

**fleetCapacityReservationSet**

Information about the individual Capacity Reservations in the Capacity Reservation
Fleet.

Type: Array of [FleetCapacityReservation](api-fleetcapacityreservation.md) objects

**instanceMatchCriteria**

The instance matching criteria for the Capacity Reservation Fleet.

Type: String

Valid Values: `open`

**requestId**

The ID of the request.

Type: String

**state**

The status of the Capacity Reservation Fleet.

Type: String

Valid Values: `submitted | modifying | active | partially_fulfilled | expiring | expired | cancelling | cancelled | failed`

**tagSet**

The tags assigned to the Capacity Reservation Fleet.

Type: Array of [Tag](api-tag.md) objects

**tenancy**

Indicates the tenancy of Capacity Reservation Fleet.

Type: String

Valid Values: `default`

**totalFulfilledCapacity**

The requested capacity units that have been successfully reserved.

Type: Double

**totalTargetCapacity**

The total number of capacity units for which the Capacity Reservation Fleet reserves
capacity.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateCapacityReservationFleet)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateCapacityReservationFleet)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createcapacityreservationfleet.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createcapacityreservationfleet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createcapacityreservationfleet.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createcapacityreservationfleet.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createcapacityreservationfleet.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createcapacityreservationfleet.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateCapacityReservationFleet)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createcapacityreservationfleet.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateCapacityReservationBySplitting

CreateCarrierGateway
