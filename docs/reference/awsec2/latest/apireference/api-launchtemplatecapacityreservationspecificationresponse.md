# LaunchTemplateCapacityReservationSpecificationResponse

Information about the Capacity Reservation targeting option.

## Contents

**capacityReservationPreference**

Indicates the instance's Capacity Reservation preferences. Possible preferences
include:

- `open` \- The instance can run in any `open` Capacity
Reservation that has matching attributes (instance type, platform, Availability
Zone).

- `none` \- The instance avoids running in a Capacity Reservation even
if one is available. The instance runs in On-Demand capacity.

Type: String

Valid Values: `capacity-reservations-only | open | none`

Required: No

**capacityReservationTarget**

Information about the target Capacity Reservation or Capacity Reservation
group.

Type: [CapacityReservationTargetResponse](api-capacityreservationtargetresponse.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplatecapacityreservationspecificationresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplatecapacityreservationspecificationresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplatecapacityreservationspecificationresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateCapacityReservationSpecificationRequest

LaunchTemplateConfig
