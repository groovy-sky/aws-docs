# WorkflowExecutionInfo

Contains information about a workflow execution.

## Contents

**execution**

The workflow execution this information is about.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: Yes

**executionStatus**

The current status of the execution.

Type: String

Valid Values: `OPEN | CLOSED`

Required: Yes

**startTimestamp**

The time when the execution was started.

Type: Timestamp

Required: Yes

**workflowType**

The type of the workflow execution.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

**cancelRequested**

Set to true if a cancellation is requested for this workflow execution.

Type: Boolean

Required: No

**closeStatus**

If the execution status is closed then this specifies how the execution was closed:

- `COMPLETED` – the execution was successfully completed.

- `CANCELED` – the execution was canceled.Cancellation allows the implementation to gracefully clean
up before the execution is closed.

- `TERMINATED` – the execution was force terminated.

- `FAILED` – the execution failed to complete.

- `TIMED_OUT` – the execution did not complete in the alloted time and was automatically timed
out.

- `CONTINUED_AS_NEW` – the execution is logically continued. This means the current execution was
completed and a new execution was started to carry on the workflow.

Type: String

Valid Values: `COMPLETED | FAILED | CANCELED | TERMINATED | CONTINUED_AS_NEW | TIMED_OUT`

Required: No

**closeTimestamp**

The time when the workflow execution was closed. Set only if the execution status is CLOSED.

Type: Timestamp

Required: No

**parent**

If this workflow execution is a child of another execution then contains the workflow execution that started this execution.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: No

**tagList**

The list of tags associated with the workflow execution. Tags can be used to identify and list workflow executions of interest through the visibility APIs. A workflow execution can have a maximum of 5 tags.

Type: Array of strings

Array Members: Maximum number of 5 items.

Length Constraints: Minimum length of 0. Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutioninfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutioninfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutioninfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionFilter

WorkflowExecutionOpenCounts

All content copied from https://docs.aws.amazon.com/.
