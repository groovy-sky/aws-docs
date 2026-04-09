# ExternalWorkflowExecutionCancelRequestedEventAttributes

Provides the details of the `ExternalWorkflowExecutionCancelRequested` event.

## Contents

**initiatedEventId**

The ID of the `RequestCancelExternalWorkflowExecutionInitiated` event corresponding to the
`RequestCancelExternalWorkflowExecution` decision to cancel this external workflow execution. This
information can be useful for diagnosing problems by tracing back the chain of events leading up to this
event.

Type: Long

Required: Yes

**workflowExecution**

The external workflow execution to which the cancellation request was delivered.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/externalworkflowexecutioncancelrequestedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/externalworkflowexecutioncancelrequestedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/externalworkflowexecutioncancelrequestedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExecutionTimeFilter

ExternalWorkflowExecutionSignaledEventAttributes

All content copied from https://docs.aws.amazon.com/.
