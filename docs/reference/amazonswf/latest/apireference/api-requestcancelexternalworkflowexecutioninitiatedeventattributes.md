# RequestCancelExternalWorkflowExecutionInitiatedEventAttributes

Provides the details of the `RequestCancelExternalWorkflowExecutionInitiated` event.

## Contents

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision task that resulted in the
`RequestCancelExternalWorkflowExecution` decision for this cancellation request.
This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**workflowId**

The `workflowId` of the external workflow execution to be canceled.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**control**

Data attached to the event that can be used by the decider in subsequent workflow tasks.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**runId**

The `runId` of the external workflow execution to be canceled.

Type: String

Length Constraints: Maximum length of 64.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/requestcancelexternalworkflowexecutioninitiatedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/requestcancelexternalworkflowexecutioninitiatedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/requestcancelexternalworkflowexecutioninitiatedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestCancelExternalWorkflowExecutionFailedEventAttributes

ResourceTag

All content copied from https://docs.aws.amazon.com/.
