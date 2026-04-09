# WorkflowExecutionContinuedAsNewEventAttributes

Provides the details of the `WorkflowExecutionContinuedAsNew` event.

## Contents

**childPolicy**

The policy to use for the child workflow executions of the new execution if it is terminated by calling the
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

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision task that resulted in the
`ContinueAsNewWorkflowExecution` decision that started this execution. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**newExecutionRunId**

The `runId` of the new workflow execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**taskList**

The task list to use for the decisions of the new (continued) workflow
execution.

Type: [TaskList](api-tasklist.md) object

Required: Yes

**workflowType**

The workflow type of this execution.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

**executionStartToCloseTimeout**

The total duration allowed for the new workflow execution.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**input**

The input provided to the new workflow execution.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**lambdaRole**

The IAM role to attach to the new (continued) workflow execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**tagList**

The list of tags associated with the new workflow execution.

Type: Array of strings

Array Members: Maximum number of 5 items.

Length Constraints: Minimum length of 0. Maximum length of 256.

Required: No

**taskPriority**

The priority of the task to use for the decisions of the new (continued) workflow
execution.

Type: String

Required: No

**taskStartToCloseTimeout**

The maximum duration of decision tasks for the new workflow execution.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutioncontinuedasneweventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutioncontinuedasneweventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutioncontinuedasneweventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionConfiguration

WorkflowExecutionFailedEventAttributes

All content copied from https://docs.aws.amazon.com/.
