# LambdaFunctionScheduledEventAttributes

Provides the details of the `LambdaFunctionScheduled` event. It isn't set
for other event types.

## Contents

**decisionTaskCompletedEventId**

The ID of the `LambdaFunctionCompleted` event corresponding to the decision
that resulted in scheduling this activity task. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: Yes

**id**

The unique ID of the Lambda task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**name**

The name of the Lambda function.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**control**

Data attached to the event that the decider can use in subsequent workflow tasks. This
data isn't sent to the Lambda task.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**input**

The input provided to the Lambda task.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 32768.

Required: No

**startToCloseTimeout**

The maximum amount of time a worker can take to process the Lambda task.

Type: String

Length Constraints: Maximum length of 8.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/lambdafunctionscheduledeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/lambdafunctionscheduledeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/lambdafunctionscheduledeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaFunctionFailedEventAttributes

LambdaFunctionStartedEventAttributes

All content copied from https://docs.aws.amazon.com/.
