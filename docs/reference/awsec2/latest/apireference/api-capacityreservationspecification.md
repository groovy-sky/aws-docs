# CapacityReservationSpecification

Describes an instance's Capacity Reservation targeting option.

Use the `CapacityReservationPreference` parameter to configure the instance
to run as an On-Demand Instance, to run in any `open` Capacity Reservation
that has matching attributes, or to run only in a Capacity Reservation or Capacity
Reservation group. Use the `CapacityReservationTarget` parameter to
explicitly target a specific Capacity Reservation or a Capacity Reservation
group.

You can only specify `CapacityReservationPreference` and
`CapacityReservationTarget` if the
`CapacityReservationPreference` is
`capacity-reservations-only`.

## Contents

**CapacityReservationPreference**

Indicates the instance's Capacity Reservation preferences. Possible preferences
include:

- `capacity-reservations-only` \- The instance will only run in a
Capacity Reservation or Capacity Reservation group. If capacity isn't available,
the instance will fail to launch.

- `open` \- The instance can run in any `open` Capacity
Reservation that has matching attributes (instance type, platform, Availability
Zone, and tenancy). If capacity isn't available, the instance runs as an
On-Demand Instance.

- `none` \- The instance doesn't run in a Capacity Reservation even if
one is available. The instance runs as an On-Demand Instance.

Type: String

Valid Values: `capacity-reservations-only | open | none`

Required: No

**CapacityReservationTarget**

Information about the target Capacity Reservation or Capacity Reservation
group.

Type: [CapacityReservationTarget](api-capacityreservationtarget.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityreservationspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityreservationspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityreservationspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityReservationOptionsRequest

CapacityReservationSpecificationResponse
