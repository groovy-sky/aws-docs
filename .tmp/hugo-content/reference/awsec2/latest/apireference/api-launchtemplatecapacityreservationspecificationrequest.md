# LaunchTemplateCapacityReservationSpecificationRequest

Describes an instance's Capacity Reservation targeting option. You can specify only
one option at a time. Use the `CapacityReservationPreference` parameter to
configure the instance to run in On-Demand capacity or to run in any `open`
Capacity Reservation that has matching attributes (instance type, platform, Availability
Zone). Use the `CapacityReservationTarget` parameter to explicitly target a
specific Capacity Reservation or a Capacity Reservation group.

## Contents

**CapacityReservationPreference**

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplatecapacityreservationspecificationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplatecapacityreservationspecificationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplatecapacityreservationspecificationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateBlockDeviceMappingRequest

LaunchTemplateCapacityReservationSpecificationResponse
