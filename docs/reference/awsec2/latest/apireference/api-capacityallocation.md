# CapacityAllocation

Information about instance capacity usage for a Capacity Reservation.

## Contents

**AllocationMetadataList.N**

Additional metadata associated with the capacity allocation. Each entry contains a key-value pair providing context
about the allocation.

Type: Array of [CapacityAllocationMetadataEntry](api-capacityallocationmetadataentry.md) objects

Required: No

**allocationType**

The usage type. `used` indicates that the instance capacity is in use by
instances that are running in the Capacity Reservation.

Type: String

Valid Values: `used`

Required: No

**count**

The amount of instance capacity associated with the usage. For example a value of
`4` indicates that instance capacity for 4 instances is currently in
use.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityallocation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityallocation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityallocation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CancelSpotFleetRequestsSuccessItem

CapacityAllocationMetadataEntry
