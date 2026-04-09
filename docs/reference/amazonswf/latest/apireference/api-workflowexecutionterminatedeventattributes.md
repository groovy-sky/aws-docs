# WorkflowExecutionTerminatedEventAttributes

Provides the details of the `WorkflowExecutionTerminated` event.

## Contents

**childPolicy**

The policy used for the child workflow executions of this workflow execution.

The supported child policies are:

- `TERMINATE` – The child executions are terminated.

- `REQUEST_CANCEL` – A request to cancel is attempted for each child
execution by recording a `WorkflowExecutionCancelRequested` event in its history. It is up to the decider
to take appropriate actions when it receives an execution history with this event.

- `ABANDON` – No action is taken. The child executions continue to run.

Type: String

Valid Values: `TERMINATE | REQUEST_CANCEL | ABANDON`

Required: Yes

**cause**

If set, indicates that the workflow execution was automatically terminated, and specifies the cause. This happens if the parent workflow execution times out or is terminated and the child policy is set to terminate child executions.

Type: String

Valid Values: `CHILD_POLICY_APPLIED | EVENT_LIMIT_EXCEEDED | OPERATOR_INITIATED`

Required: No

**details**

The details provided for the termination.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**reason**

The reason provided for the termination.

Type: String

Length Constraints: Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutionterminatedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutionterminatedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutionterminatedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionStartedEventAttributes

WorkflowExecutionTimedOutEventAttributes

All content copied from https://docs.aws.amazon.com/.
