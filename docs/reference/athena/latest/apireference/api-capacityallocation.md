# CapacityAllocation

Contains the submission time of a single allocation request for a capacity reservation
and the most recent status of the attempted allocation.

## Contents

**RequestTime**

The time when the capacity allocation was requested.

Type: Timestamp

Required: Yes

**Status**

The status of the capacity allocation.

Type: String

Valid Values: `PENDING | SUCCEEDED | FAILED`

Required: Yes

**RequestCompletionTime**

The time when the capacity allocation request was completed.

Type: Timestamp

Required: No

**StatusMessage**

The status message of the capacity allocation.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/capacityallocation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/capacityallocation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/capacityallocation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CalculationSummary

CapacityAssignment
