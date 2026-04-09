# CalculationStatus

Contains information about the status of a notebook calculation.

## Contents

**CompletionDateTime**

The date and time the calculation completed processing.

Type: Timestamp

Required: No

**State**

The state of the calculation execution. A description of each state follows.

`CREATING` \- The calculation is in the process of being created.

`CREATED` \- The calculation has been created and is ready to run.

`QUEUED` \- The calculation has been queued for processing.

`RUNNING` \- The calculation is running.

`CANCELING` \- A request to cancel the calculation has been received and the
system is working to stop it.

`CANCELED` \- The calculation is no longer running as the result of a cancel
request.

`COMPLETED` \- The calculation has completed without error.

`FAILED` \- The calculation failed and is no longer running.

Type: String

Valid Values: `CREATING | CREATED | QUEUED | RUNNING | CANCELING | CANCELED | COMPLETED | FAILED`

Required: No

**StateChangeReason**

The reason for the calculation state change (for example, the calculation was canceled
because the session was terminated).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**SubmissionDateTime**

The date and time the calculation was submitted for processing.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/calculationstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/calculationstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/calculationstatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CalculationStatistics

CalculationSummary

All content copied from https://docs.aws.amazon.com/.
