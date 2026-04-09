# StartLambdaFunctionFailedEventAttributes

Provides the details of the `StartLambdaFunctionFailed` event. It isn't set
for other event types.

## Contents

**cause**

The cause of the failure. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

###### Note

If `cause` is set to `OPERATION_NOT_PERMITTED`, the decision
failed because the IAM role attached to the execution lacked sufficient permissions. For
details and example IAM policies, see [Lambda Tasks](../../../../services/amazonswf/latest/developerguide/lambda-task.md) in the
_Amazon SWF Developer Guide_.

Type: String

Valid Values: `ASSUME_ROLE_FAILED`

Required: No

**message**

A description that can help diagnose the cause of the fault.

Type: String

Length Constraints: Maximum length of 1728.

Required: No

**scheduledEventId**

The ID of the `ActivityTaskScheduled` event that was recorded when this
activity task was scheduled. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/startlambdafunctionfailedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/startlambdafunctionfailedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/startlambdafunctionfailedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartChildWorkflowExecutionInitiatedEventAttributes

StartTimerDecisionAttributes

All content copied from https://docs.aws.amazon.com/.
