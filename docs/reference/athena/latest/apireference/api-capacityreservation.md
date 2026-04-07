# CapacityReservation

A reservation for a specified number of data processing units (DPUs). When a
reservation is initially created, it has no DPUs. Athena allocates DPUs
until the allocated amount equals the requested amount.

## Contents

**AllocatedDpus**

The number of data processing units currently allocated.

Type: Integer

Valid Range: Minimum value of 0.

Required: Yes

**CreationTime**

The time in UTC epoch millis when the capacity reservation was created.

Type: Timestamp

Required: Yes

**Name**

The name of the capacity reservation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9._-]+`

Required: Yes

**Status**

The status of the capacity reservation.

Type: String

Valid Values: `PENDING | ACTIVE | CANCELLING | CANCELLED | FAILED | UPDATE_PENDING`

Required: Yes

**TargetDpus**

The number of data processing units requested.

Type: Integer

Valid Range: Minimum value of 24.

Required: Yes

**LastAllocation**

Contains the submission time of a single allocation request for a capacity reservation
and the most recent status of the attempted allocation.

Type: [CapacityAllocation](https://docs.aws.amazon.com/athena/latest/APIReference/API_CapacityAllocation.html) object

Required: No

**LastSuccessfulAllocationTime**

The time of the most recent capacity allocation that succeeded.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/CapacityReservation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/CapacityReservation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/CapacityReservation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CapacityAssignmentConfiguration

Classification
