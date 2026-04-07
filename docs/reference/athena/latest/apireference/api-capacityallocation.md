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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/CapacityAllocation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/CapacityAllocation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/CapacityAllocation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CalculationSummary

CapacityAssignment
