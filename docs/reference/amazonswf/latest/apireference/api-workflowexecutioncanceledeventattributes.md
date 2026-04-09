# WorkflowExecutionCanceledEventAttributes

Provides the details of the `WorkflowExecutionCanceled` event.

## Contents

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision task that resulted in the
`CancelWorkflowExecution` decision for this cancellation request. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**details**

The details of the cancellation.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutioncanceledeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutioncanceledeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutioncanceledeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecution

WorkflowExecutionCancelRequestedEventAttributes

All content copied from https://docs.aws.amazon.com/.
