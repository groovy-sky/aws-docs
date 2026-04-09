# StartChildWorkflowExecutionInitiatedEventAttributes

Provides the details of the `StartChildWorkflowExecutionInitiated` event.

## Contents

**childPolicy**

The policy to use for the child workflow executions if this execution gets terminated by explicitly calling the
[TerminateWorkflowExecution](api-terminateworkflowexecution.md) action or due to an expired timeout.

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
`StartChildWorkflowExecution` [Decision](api-decision.md) to request this child workflow execution. This
information can be useful for diagnosing problems by tracing back the cause of events.

Type: Long

Required: Yes

**taskList**

The name of the task list used for the decision tasks of the child workflow execution.

Type: [TaskList](api-tasklist.md) object

Required: Yes

**workflowId**

The `workflowId` of the child workflow execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**workflowType**

The type of the child workflow execution.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

**control**

Data attached to the event that can be used by the decider in subsequent decision tasks. This data isn't sent to the activity.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**executionStartToCloseTimeout**

The maximum duration for the child workflow execution. If the workflow execution isn't closed within this duration, it is timed out and force-terminated.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**input**

The inputs provided to the child workflow execution.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**lambdaRole**

The IAM role to attach to the child workflow execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**tagList**

The list of tags to associated with the child workflow execution.

Type: Array of strings

Array Members: Maximum number of 5 items.

Length Constraints: Minimum length of 0. Maximum length of 256.

Required: No

**taskPriority**

The priority assigned for the decision tasks for this workflow execution.
Valid values are integers that range from Java's `Integer.MIN_VALUE`
(-2147483648) to `Integer.MAX_VALUE` (2147483647). Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

**taskStartToCloseTimeout**

The maximum duration allowed for the decision tasks for this workflow execution.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/startchildworkflowexecutioninitiatedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/startchildworkflowexecutioninitiatedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/startchildworkflowexecutioninitiatedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartChildWorkflowExecutionFailedEventAttributes

StartLambdaFunctionFailedEventAttributes

All content copied from https://docs.aws.amazon.com/.
