# LambdaFunctionCompletedEventAttributes

Provides the details of the `LambdaFunctionCompleted` event. It isn't set
for other event types.

## Contents

**scheduledEventId**

The ID of the `LambdaFunctionScheduled` event that was recorded when this
Lambda task was scheduled. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: Yes

**startedEventId**

The ID of the `LambdaFunctionStarted` event recorded when this activity task
started. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: Yes

**result**

The results of the Lambda task.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/lambdafunctioncompletedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/lambdafunctioncompletedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/lambdafunctioncompletedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HistoryEvent

LambdaFunctionFailedEventAttributes

All content copied from https://docs.aws.amazon.com/.
