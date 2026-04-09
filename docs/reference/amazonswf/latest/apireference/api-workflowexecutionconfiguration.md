# WorkflowExecutionConfiguration

The configuration settings for a workflow execution including timeout values, tasklist etc. These configuration settings are determined from the defaults specified when registering the workflow type and those specified when starting the workflow execution.

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

**executionStartToCloseTimeout**

The total duration for this workflow execution.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8.

Required: Yes

**taskList**

The task list used for the decision tasks generated for this workflow execution.

Type: [TaskList](api-tasklist.md) object

Required: Yes

**taskStartToCloseTimeout**

The maximum duration allowed for decision tasks for this workflow execution.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8.

Required: Yes

**lambdaRole**

The IAM role attached to the child workflow execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**taskPriority**

The priority assigned to decision tasks for this workflow execution. Valid values are integers that range from Java's `Integer.MIN_VALUE`
(-2147483648) to `Integer.MAX_VALUE` (2147483647). Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutionconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutionconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutionconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionCompletedEventAttributes

WorkflowExecutionContinuedAsNewEventAttributes

All content copied from https://docs.aws.amazon.com/.
