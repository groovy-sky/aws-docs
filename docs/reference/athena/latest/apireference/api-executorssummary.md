# ExecutorsSummary

Contains summary information about an executor.

## Contents

**ExecutorId**

The UUID of the executor.

Type: String

Length Constraints: Maximum length of 100000.

Pattern: `.*`

Required: Yes

**ExecutorSize**

The smallest unit of compute that a session can request from Athena. Size
is measured in data processing unit (DPU) values, a relative measure of processing
power.

Type: Long

Required: No

**ExecutorState**

The processing state of the executor. A description of each state follows.

`CREATING` \- The executor is being started, including acquiring
resources.

`CREATED` \- The executor has been started.

`REGISTERED` \- The executor has been registered.

`TERMINATING` \- The executor is in the process of shutting down.

`TERMINATED` \- The executor is no longer running.

`FAILED` \- Due to a failure, the executor is no longer running.

Type: String

Valid Values: `CREATING | CREATED | REGISTERED | TERMINATING | TERMINATED | FAILED`

Required: No

**ExecutorType**

The type of executor used for the application ( `COORDINATOR`,
`GATEWAY`, or `WORKER`).

Type: String

Valid Values: `COORDINATOR | GATEWAY | WORKER`

Required: No

**StartDateTime**

The date and time that the executor started.

Type: Long

Required: No

**TerminationDateTime**

The date and time that the executor was terminated.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/executorssummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/executorssummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/executorssummary.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EngineVersion

FilterDefinition
