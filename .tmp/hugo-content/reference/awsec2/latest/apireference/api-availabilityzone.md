# AvailabilityZone

Describes Availability Zones, Local Zones, and Wavelength Zones.

## Contents

**GeographySet.N**

The geography information for the Availability Zone or Local Zone. The geography is returned as a list.

Type: Array of [AvailabilityZoneGeography](api-availabilityzonegeography.md) objects

Required: No

**groupLongName**

The long name of the Availability Zone group, Local Zone group, or Wavelength Zone
group.

Type: String

Required: No

**groupName**

The name of the zone group. For example:

- Availability Zones - `us-east-1-zg-1`

- Local Zones - `us-west-2-lax-1`

- Wavelength Zones - `us-east-1-wl1-bos-wlz-1`

Type: String

Required: No

**MessageSet.N**

Any messages about the Availability Zone, Local Zone, or Wavelength Zone.

Type: Array of [AvailabilityZoneMessage](api-availabilityzonemessage.md) objects

Required: No

**networkBorderGroup**

The name of the network border group.

Type: String

Required: No

**optInStatus**

For Availability Zones, this parameter always has the value of
`opt-in-not-required`.

For Local Zones and Wavelength Zones, this parameter is the opt-in status. The possible
values are `opted-in` and `not-opted-in`.

Type: String

Valid Values: `opt-in-not-required | opted-in | not-opted-in`

Required: No

**parentZoneId**

The ID of the zone that handles some of the Local Zone or Wavelength Zone control plane
operations, such as API calls.

Type: String

Required: No

**parentZoneName**

The name of the zone that handles some of the Local Zone or Wavelength Zone control plane
operations, such as API calls.

Type: String

Required: No

**regionName**

The name of the Region.

Type: String

Required: No

**SubGeographySet.N**

The sub-geography information for the Availability Zone or Local Zone. The sub-geography is returned
as a list.

Type: Array of [AvailabilityZoneSubGeography](api-availabilityzonesubgeography.md) objects

Required: No

**zoneId**

The ID of the Availability Zone, Local Zone, or Wavelength Zone.

Type: String

Required: No

**zoneName**

The name of the Availability Zone, Local Zone, or Wavelength Zone.

Type: String

Required: No

**zoneState**

The state of the Availability Zone, Local Zone, or Wavelength Zone. The possible values are
`available`, `unavailable`, and `constrained`.

Type: String

Valid Values: `available | information | impaired | unavailable | constrained`

Required: No

**zoneType**

The type of zone.

Valid values: `availability-zone` \| `local-zone` \|
`wavelength-zone`

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/availabilityzone.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/availabilityzone.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/availabilityzone.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AuthorizationRule

AvailabilityZoneAddress
