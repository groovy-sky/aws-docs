# WorkflowTypeConfiguration

The configuration settings of a workflow type.

## Contents

**defaultChildPolicy**

The default policy to use for the child workflow executions when a workflow execution of this
type is terminated, by calling the [TerminateWorkflowExecution](api-terminateworkflowexecution.md) action explicitly or due to an expired
timeout. This default can be overridden when starting a workflow execution using the [StartWorkflowExecution](api-startworkflowexecution.md)
action or the `StartChildWorkflowExecution` [Decision](api-decision.md).

The supported child policies are:

- `TERMINATE` – The child executions are terminated.

- `REQUEST_CANCEL` – A request to cancel is attempted for each child
execution by recording a `WorkflowExecutionCancelRequested` event in its history. It is up to the decider
to take appropriate actions when it receives an execution history with this event.

- `ABANDON` – No action is taken. The child executions continue to run.

Type: String

Valid Values: `TERMINATE | REQUEST_CANCEL | ABANDON`

Required: No

**defaultExecutionStartToCloseTimeout**

The default maximum duration, specified when registering the workflow type, for executions of
this workflow type. This default can be overridden when starting a workflow execution using the
[StartWorkflowExecution](api-startworkflowexecution.md) action or the `StartChildWorkflowExecution` [Decision](api-decision.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**defaultLambdaRole**

The default IAM role attached to this workflow type.

###### Note

Executions of this workflow type need IAM roles to invoke Lambda functions. If you
don't specify an IAM role when starting this workflow type, the default Lambda role is
attached to the execution. For more information, see [https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html](../../../../services/amazonswf/latest/developerguide/lambda-task.md) in the
_Amazon SWF Developer Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**defaultTaskList**

The default task list, specified when registering the workflow type, for decisions tasks
scheduled for workflow executions of this type. This default can be overridden when starting a workflow execution
using the [StartWorkflowExecution](api-startworkflowexecution.md) action or the `StartChildWorkflowExecution` [Decision](api-decision.md).

Type: [TaskList](api-tasklist.md) object

Required: No

**defaultTaskPriority**

The default task priority, specified when registering the workflow type, for all decision tasks
of this workflow type. This default can be overridden when starting a workflow execution using the
[StartWorkflowExecution](api-startworkflowexecution.md) action or the `StartChildWorkflowExecution` decision.

Valid values are integers that range from Java's `Integer.MIN_VALUE`
(-2147483648) to `Integer.MAX_VALUE` (2147483647). Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

**defaultTaskStartToCloseTimeout**

The default maximum duration, specified when registering the workflow type, that a decision task
for executions of this workflow type might take before returning completion or failure. If the task doesn'tdo close
in the specified time then the task is automatically timed out and rescheduled. If the decider eventually reports
a completion or failure, it is ignored. This default can be overridden when starting a workflow execution using
the [StartWorkflowExecution](api-startworkflowexecution.md) action or the `StartChildWorkflowExecution` [Decision](api-decision.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowtypeconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowtypeconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowtypeconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowType

WorkflowTypeFilter

All content copied from https://docs.aws.amazon.com/.
