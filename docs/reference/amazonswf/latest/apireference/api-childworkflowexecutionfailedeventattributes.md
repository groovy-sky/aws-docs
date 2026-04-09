# ChildWorkflowExecutionFailedEventAttributes

Provides the details of the `ChildWorkflowExecutionFailed` event.

## Contents

**initiatedEventId**

The ID of the `StartChildWorkflowExecutionInitiated` event corresponding to the
`StartChildWorkflowExecution` [Decision](api-decision.md) to start this child workflow execution.
This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**startedEventId**

The ID of the `ChildWorkflowExecutionStarted` event recorded when this child workflow execution was
started. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**workflowExecution**

The child workflow execution that failed.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: Yes

**workflowType**

The type of the child workflow execution.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

**details**

The details of the failure (if provided).

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**reason**

The reason for the failure (if provided).

Type: String

Length Constraints: Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/childworkflowexecutionfailedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/childworkflowexecutionfailedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/childworkflowexecutionfailedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChildWorkflowExecutionCompletedEventAttributes

ChildWorkflowExecutionStartedEventAttributes

All content copied from https://docs.aws.amazon.com/.
