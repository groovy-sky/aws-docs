# SignalExternalWorkflowExecutionInitiatedEventAttributes

Provides the details of the `SignalExternalWorkflowExecutionInitiated` event.

## Contents

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision task that resulted in the
`SignalExternalWorkflowExecution` decision for this signal. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**signalName**

The name of the signal.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**workflowId**

The `workflowId` of the external workflow execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**control**

Data attached to the event that can be used by the decider in subsequent decision tasks.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**input**

The input provided to the signal.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**runId**

The `runId` of the external workflow execution to send the signal to.

Type: String

Length Constraints: Maximum length of 64.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/signalexternalworkflowexecutioninitiatedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/signalexternalworkflowexecutioninitiatedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/signalexternalworkflowexecutioninitiatedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SignalExternalWorkflowExecutionFailedEventAttributes

StartChildWorkflowExecutionDecisionAttributes

All content copied from https://docs.aws.amazon.com/.
