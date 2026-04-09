# WorkflowExecutionFailedEventAttributes

Provides the details of the `WorkflowExecutionFailed` event.

## Contents

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision task that resulted in the
`FailWorkflowExecution` decision to fail this execution. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**details**

The details of the failure.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**reason**

The descriptive reason provided for the failure.

Type: String

Length Constraints: Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutionfailedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutionfailedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutionfailedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionContinuedAsNewEventAttributes

WorkflowExecutionFilter

All content copied from https://docs.aws.amazon.com/.
