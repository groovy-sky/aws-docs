# JobDetails

Contains details about the policy generation request.

## Contents

**jobId**

The `JobId` that is returned by the `StartPolicyGeneration`
operation. The `JobId` can be used with `GetGeneratedPolicy` to
retrieve the generated policies or used with `CancelPolicyGeneration` to cancel
the policy generation request.

Type: String

Required: Yes

**startedOn**

A timestamp of when the job was started.

Type: Timestamp

Required: Yes

**status**

The status of the job request.

Type: String

Valid Values: `IN_PROGRESS | SUCCEEDED | FAILED | CANCELED`

Required: Yes

**completedOn**

A timestamp of when the job was completed.

Type: Timestamp

Required: No

**jobError**

The job error for the policy generation request.

Type: [JobError](api-joberror.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/jobdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/jobdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/jobdetails.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InternetConfiguration

JobError
