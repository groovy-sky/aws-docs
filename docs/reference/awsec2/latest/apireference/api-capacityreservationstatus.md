# CapacityReservationStatus

Describes the availability of capacity for a Capacity Reservation.

## Contents

**capacityReservationId**

The ID of the Capacity Reservation.

Type: String

Required: No

**totalAvailableCapacity**

The remaining capacity. Indicates the amount of resources that can be launched into the Capacity Reservation.

Type: Integer

Required: No

**totalCapacity**

The combined amount of `Available` and `Unavailable` capacity in the Capacity Reservation.

Type: Integer

Required: No

**totalUnavailableCapacity**

The used capacity. Indicates that the capacity is in use by resources that are running in the Capacity Reservation.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityreservationstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityreservationstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityreservationstatus.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityReservationSpecificationResponse

CapacityReservationTarget
