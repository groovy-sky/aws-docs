# LambdaFunctionTimedOutEventAttributes

Provides details of the `LambdaFunctionTimedOut` event.

## Contents

**scheduledEventId**

The ID of the `LambdaFunctionScheduled` event that was recorded when this
activity task was scheduled. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: Yes

**startedEventId**

The ID of the `ActivityTaskStarted` event that was recorded when this
activity task started. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: Yes

**timeoutType**

The type of the timeout that caused this event.

Type: String

Valid Values: `START_TO_CLOSE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/lambdafunctiontimedouteventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/lambdafunctiontimedouteventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/lambdafunctiontimedouteventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaFunctionStartedEventAttributes

MarkerRecordedEventAttributes

All content copied from https://docs.aws.amazon.com/.
