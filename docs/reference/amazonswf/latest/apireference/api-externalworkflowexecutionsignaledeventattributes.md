# ExternalWorkflowExecutionSignaledEventAttributes

Provides the details of the `ExternalWorkflowExecutionSignaled` event.

## Contents

**initiatedEventId**

The ID of the `SignalExternalWorkflowExecutionInitiated` event corresponding to the
`SignalExternalWorkflowExecution` decision to request this signal. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**workflowExecution**

The external workflow execution that the signal was delivered to.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/externalworkflowexecutionsignaledeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/externalworkflowexecutionsignaledeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/externalworkflowexecutionsignaledeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExternalWorkflowExecutionCancelRequestedEventAttributes

FailWorkflowExecutionDecisionAttributes

All content copied from https://docs.aws.amazon.com/.
