# ContinueAsNewWorkflowExecutionDecisionAttributes

Provides the details of the `ContinueAsNewWorkflowExecution` decision.

**Access Control**

You can use IAM policies to control this decision's access to Amazon SWF resources as follows:

- Use a `Resource` element with the domain name to limit the action to only
specified domains.

- Use an `Action` element to allow or deny permission to call this action.

- Constrain the following parameters by using a `Condition` element with the
appropriate keys.

- `tag` – A tag used to identify the workflow execution

- `taskList` – String constraint. The key is `swf:taskList.name`.

- `workflowType.version` – String constraint. The key is `swf:workflowType.version`.

If the caller doesn't have sufficient permissions to invoke the action, or the
parameter values fall outside the specified constraints, the action fails. The associated event attribute's
`cause` parameter is set to `OPERATION_NOT_PERMITTED`. For details and example IAM policies, see
[Using IAM to Manage Access to Amazon SWF Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md) in the _Amazon SWF Developer Guide_.

## Contents

**childPolicy**

If set, specifies the policy to use for the child workflow executions of the new execution if it is terminated
by calling the [TerminateWorkflowExecution](api-terminateworkflowexecution.md) action explicitly or due to an expired timeout. This policy
overrides the default child policy specified when registering the workflow type using
[RegisterWorkflowType](api-registerworkflowtype.md).

The supported child policies are:

- `TERMINATE` – The child executions are terminated.

- `REQUEST_CANCEL` – A request to cancel is attempted for each child
execution by recording a `WorkflowExecutionCancelRequested` event in its history. It is up to the decider
to take appropriate actions when it receives an execution history with this event.

- `ABANDON` – No action is taken. The child executions continue to run.

###### Note

A child policy for this workflow execution must be specified either as a default for the workflow type or through this parameter. If neither this parameter is set nor a default child policy was specified at registration time then a fault is returned.

Type: String

Valid Values: `TERMINATE | REQUEST_CANCEL | ABANDON`

Required: No

**executionStartToCloseTimeout**

If set, specifies the total duration for this workflow execution. This overrides the
`defaultExecutionStartToCloseTimeout` specified when registering the workflow type.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

###### Note

An execution start-to-close timeout for this workflow execution must be specified either as a default for the workflow type or through this field. If neither this field is set nor a default execution start-to-close timeout was specified at registration time then a fault is returned.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**input**

The input provided to the new workflow execution.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**lambdaRole**

The IAM role to attach to the new (continued) execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**tagList**

The list of tags to associate with the new workflow execution. A maximum of 5 tags can be specified. You can
list workflow executions with a specific tag by calling [ListOpenWorkflowExecutions](api-listopenworkflowexecutions.md) or
[ListClosedWorkflowExecutions](api-listclosedworkflowexecutions.md) and specifying a [TagFilter](api-tagfilter.md).

Type: Array of strings

Array Members: Maximum number of 5 items.

Length Constraints: Minimum length of 0. Maximum length of 256.

Required: No

**taskList**

The task list to use for the decisions of the new (continued) workflow
execution.

Type: [TaskList](api-tasklist.md) object

Required: No

**taskPriority**

The task priority that, if set, specifies the priority for the decision tasks for this workflow
execution. This overrides the defaultTaskPriority specified when registering the workflow type.
Valid values are integers that range from Java's `Integer.MIN_VALUE`
(-2147483648) to `Integer.MAX_VALUE` (2147483647). Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

**taskStartToCloseTimeout**

Specifies the maximum duration of decision tasks for the new workflow execution. This parameter overrides the
`defaultTaskStartToCloseTimout` specified when registering the workflow type using
[RegisterWorkflowType](api-registerworkflowtype.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

###### Note

A task start-to-close timeout for the new workflow execution must be specified either as a default for the workflow type or through this parameter. If neither this parameter is set nor a default task start-to-close timeout was specified at registration time then a fault is returned.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**workflowTypeVersion**

The version of the workflow to start.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/continueasnewworkflowexecutiondecisionattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/continueasnewworkflowexecutiondecisionattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/continueasnewworkflowexecutiondecisionattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CompleteWorkflowExecutionFailedEventAttributes

ContinueAsNewWorkflowExecutionFailedEventAttributes

All content copied from https://docs.aws.amazon.com/.
