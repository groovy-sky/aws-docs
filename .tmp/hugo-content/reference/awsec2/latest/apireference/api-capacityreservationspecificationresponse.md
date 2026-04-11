# CapacityReservationSpecificationResponse

Describes the instance's Capacity Reservation targeting preferences. The action returns the
`capacityReservationPreference` response element if the instance is
configured to run in On-Demand capacity, or if it is configured in run in any
`open` Capacity Reservation that has matching attributes (instance type, platform,
Availability Zone). The action returns the `capacityReservationTarget`
response element if the instance explicily targets a specific Capacity Reservation or Capacity Reservation group.

## Contents

**capacityReservationPreference**

Describes the instance's Capacity Reservation preferences. Possible preferences include:

- `open` \- The instance can run in any `open` Capacity Reservation that
has matching attributes (instance type, platform, Availability Zone).

- `none` \- The instance avoids running in a Capacity Reservation even if one is
available. The instance runs in On-Demand capacity.

Type: String

Valid Values: `capacity-reservations-only | open | none`

Required: No

**capacityReservationTarget**

Information about the targeted Capacity Reservation or Capacity Reservation group.

Type: [CapacityReservationTargetResponse](api-capacityreservationtargetresponse.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityreservationspecificationresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityreservationspecificationresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityreservationspecificationresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityReservationSpecification

CapacityReservationStatus
