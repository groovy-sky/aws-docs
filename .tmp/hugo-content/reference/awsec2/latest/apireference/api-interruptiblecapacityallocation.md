# InterruptibleCapacityAllocation

Represents the allocation of capacity from a source reservation to an interruptible reservation, tracking current and target instance counts for allocation management.

## Contents

**instanceCount**

The current number of instances allocated to the interruptible reservation.

Type: Integer

Required: No

**interruptibleCapacityReservationId**

The ID of the interruptible Capacity Reservation created from the allocation.

Type: String

Required: No

**interruptionType**

The type of interruption policy applied to the interruptible reservation.

Type: String

Valid Values: `adhoc`

Required: No

**status**

The current status of the allocation (updating during reclamation, active when complete).

Type: String

Valid Values: `pending | active | updating | canceling | canceled | failed`

Required: No

**targetInstanceCount**

After your modify request, the requested number of instances allocated to interruptible reservation.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/interruptiblecapacityallocation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/interruptiblecapacityallocation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/interruptiblecapacityallocation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InternetGatewayAttachment

InterruptionInfo
