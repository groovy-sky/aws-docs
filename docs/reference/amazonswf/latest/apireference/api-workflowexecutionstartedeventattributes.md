# WorkflowExecutionStartedEventAttributes

Provides details of `WorkflowExecutionStarted` event.

## Contents

**childPolicy**

The policy to use for the child workflow executions if this workflow execution is terminated, by calling the
[TerminateWorkflowExecution](api-terminateworkflowexecution.md) action explicitly or due to an expired timeout.

The supported child policies are:

- `TERMINATE` – The child executions are terminated.

- `REQUEST_CANCEL` – A request to cancel is attempted for each child
execution by recording a `WorkflowExecutionCancelRequested` event in its history. It is up to the decider
to take appropriate actions when it receives an execution history with this event.

- `ABANDON` – No action is taken. The child executions continue to run.

Type: String

Valid Values: `TERMINATE | REQUEST_CANCEL | ABANDON`

Required: Yes

**taskList**

The name of the task list for scheduling the decision tasks for this workflow execution.

Type: [TaskList](api-tasklist.md) object

Required: Yes

**workflowType**

The workflow type of this execution.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

**continuedExecutionRunId**

If this workflow execution was started due to a `ContinueAsNewWorkflowExecution` decision, then it
contains the `runId` of the previous workflow execution that was closed and continued as this
execution.

Type: String

Length Constraints: Maximum length of 64.

Required: No

**executionStartToCloseTimeout**

The maximum duration for this workflow execution.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**input**

The input provided to the workflow execution.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**lambdaRole**

The IAM role attached to the workflow execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**parentInitiatedEventId**

The ID of the `StartChildWorkflowExecutionInitiated` event corresponding to the
`StartChildWorkflowExecution` [Decision](api-decision.md) to start this workflow execution. The source event with
this ID can be found in the history of the source workflow execution. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: No

**parentWorkflowExecution**

The source workflow execution that started this workflow execution. The member isn't set if the workflow execution was not started by a workflow.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: No

**tagList**

The list of tags associated with this workflow execution. An execution can have up to 5 tags.

Type: Array of strings

Array Members: Maximum number of 5 items.

Length Constraints: Minimum length of 0. Maximum length of 256.

Required: No

**taskPriority**

The priority of the decision tasks in the workflow execution.

Type: String

Required: No

**taskStartToCloseTimeout**

The maximum duration of decision tasks for this workflow type.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutionstartedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutionstartedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutionstartedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionSignaledEventAttributes

WorkflowExecutionTerminatedEventAttributes

All content copied from https://docs.aws.amazon.com/.
