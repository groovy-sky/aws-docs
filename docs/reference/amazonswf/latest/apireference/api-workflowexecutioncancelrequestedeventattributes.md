# WorkflowExecutionCancelRequestedEventAttributes

Provides the details of the `WorkflowExecutionCancelRequested` event.

## Contents

**cause**

If set, indicates that the request to cancel the workflow execution was automatically generated, and specifies the cause. This happens if the parent workflow execution times out or is terminated, and the child policy is set to cancel child executions.

Type: String

Valid Values: `CHILD_POLICY_APPLIED`

Required: No

**externalInitiatedEventId**

The ID of the `RequestCancelExternalWorkflowExecutionInitiated` event corresponding to the
`RequestCancelExternalWorkflowExecution` decision to cancel this workflow execution.The source event
with this ID can be found in the history of the source workflow execution. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: No

**externalWorkflowExecution**

The external workflow execution for which the cancellation was requested.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutioncancelrequestedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutioncancelrequestedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutioncancelrequestedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionCanceledEventAttributes

WorkflowExecutionCompletedEventAttributes

All content copied from https://docs.aws.amazon.com/.
