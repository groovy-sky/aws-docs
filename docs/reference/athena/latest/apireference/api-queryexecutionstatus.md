# QueryExecutionStatus

The completion date, current state, submission time, and state change reason (if
applicable) for the query execution.

## Contents

**AthenaError**

Provides information about an Athena query error.

Type: [AthenaError](api-athenaerror.md) object

Required: No

**CompletionDateTime**

The date and time that the query completed.

Type: Timestamp

Required: No

**State**

The state of query execution. `QUEUED` indicates that the query has been
submitted to the service, and Athena will execute the query as soon as
resources are available. `RUNNING` indicates that the query is in execution
phase. `SUCCEEDED` indicates that the query completed without errors.
`FAILED` indicates that the query experienced an error and did not
complete processing. `CANCELLED` indicates that a user input interrupted
query execution.

###### Note

For queries that experience certain transient errors, the state transitions from
`RUNNING` back to `QUEUED`. The `FAILED` state
is always terminal with no automatic retry.

Type: String

Valid Values: `QUEUED | RUNNING | SUCCEEDED | FAILED | CANCELLED`

Required: No

**StateChangeReason**

Further detail about the status of the query.

Type: String

Required: No

**SubmissionDateTime**

The date and time that the query was submitted.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/queryexecutionstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/queryexecutionstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/queryexecutionstatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryExecutionStatistics

QueryResultsS3AccessGrantsConfiguration

All content copied from https://docs.aws.amazon.com/.
