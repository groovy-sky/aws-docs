# LambdaFunctionFailedEventAttributes

Provides the details of the `LambdaFunctionFailed` event. It isn't set for
other event types.

## Contents

**scheduledEventId**

The ID of the `LambdaFunctionScheduled` event that was recorded when this
activity task was scheduled. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: Yes

**startedEventId**

The ID of the `LambdaFunctionStarted` event recorded when this activity task
started. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: Yes

**details**

The details of the failure.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**reason**

The reason provided for the failure.

Type: String

Length Constraints: Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/lambdafunctionfailedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/lambdafunctionfailedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/lambdafunctionfailedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaFunctionCompletedEventAttributes

LambdaFunctionScheduledEventAttributes

All content copied from https://docs.aws.amazon.com/.
