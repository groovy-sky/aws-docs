# WorkflowExecutionSignaledEventAttributes

Provides the details of the `WorkflowExecutionSignaled` event.

## Contents

**signalName**

The name of the signal received. The decider can use the signal name and inputs to determine how to the process the signal.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**externalInitiatedEventId**

The ID of the `SignalExternalWorkflowExecutionInitiated` event corresponding to the
`SignalExternalWorkflow` decision to signal this workflow execution.The source event with this ID can
be found in the history of the source workflow execution. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event. This field is set only if
the signal was initiated by another workflow execution.

Type: Long

Required: No

**externalWorkflowExecution**

The workflow execution that sent the signal. This is set only of the signal was sent by another workflow execution.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: No

**input**

The inputs provided with the signal. The decider can use the signal name and inputs to determine how to process the signal.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowexecutionsignaledeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowexecutionsignaledeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowexecutionsignaledeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowExecutionOpenCounts

WorkflowExecutionStartedEventAttributes

All content copied from https://docs.aws.amazon.com/.
