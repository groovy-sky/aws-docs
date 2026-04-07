# CapacityBlockStatus

Describes the availability of capacity for a Capacity Block.

## Contents

**capacityBlockId**

The ID of the Capacity Block.

Type: String

Required: No

**CapacityReservationStatusSet.N**

The availability of capacity for the Capacity Block reservations.

Type: Array of [CapacityReservationStatus](api-capacityreservationstatus.md) objects

Required: No

**interconnectStatus**

The status of the high-bandwidth accelerator interconnect. Possible states include:

- `ok` the accelerator interconnect is healthy.

- `impaired` \- accelerator interconnect communication is impaired.

- `insufficient-data` \- insufficient data to determine accelerator interconnect status.

Type: String

Valid Values: `ok | impaired | insufficient-data`

Required: No

**totalAvailableCapacity**

The remaining capacity. Indicates the number of resources that can be launched into the Capacity Block.

Type: Integer

Required: No

**totalCapacity**

The combined amount of `Available` and `Unavailable` capacity in the Capacity Block.

Type: Integer

Required: No

**totalUnavailableCapacity**

The unavailable capacity. Indicates the instance capacity that is unavailable for use
due to a system status check failure.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityblockstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityblockstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityblockstatus.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityBlockOffering

CapacityManagerCondition
