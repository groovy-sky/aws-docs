# WorkflowExecutionTimedOutEventAttributes

Provides the details of the `WorkflowExecutionTimedOut` event.

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

**timeoutType**

The type of timeout that caused this event.

Type: String

Valid Values: `START_TO_CLOSE`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutiontimedouteventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutiontimedouteventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutiontimedouteventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionTerminatedEventAttributes

WorkflowType

All content copied from https://docs.aws.amazon.com/.
