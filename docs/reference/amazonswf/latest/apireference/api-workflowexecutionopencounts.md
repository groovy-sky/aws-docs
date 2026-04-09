# WorkflowExecutionOpenCounts

Contains the counts of open tasks, child workflow executions and timers for a workflow execution.

## Contents

**openActivityTasks**

The count of activity tasks whose status is `OPEN`.

Type: Integer

Valid Range: Minimum value of 0.

Required: Yes

**openChildWorkflowExecutions**

The count of child workflow executions whose status is `OPEN`.

Type: Integer

Valid Range: Minimum value of 0.

Required: Yes

**openDecisionTasks**

The count of decision tasks whose status is OPEN. A workflow execution can have at most one open decision task.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 1.

Required: Yes

**openTimers**

The count of timers started by this workflow execution that have not fired yet.

Type: Integer

Valid Range: Minimum value of 0.

Required: Yes

**openLambdaFunctions**

The count of Lambda tasks whose status is `OPEN`.

Type: Integer

Valid Range: Minimum value of 0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutionopencounts.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutionopencounts.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutionopencounts.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionInfo

WorkflowExecutionSignaledEventAttributes

All content copied from https://docs.aws.amazon.com/.
