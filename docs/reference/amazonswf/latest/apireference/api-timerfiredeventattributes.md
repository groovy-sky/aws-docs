# TimerFiredEventAttributes

Provides the details of the `TimerFired` event.

## Contents

**startedEventId**

The ID of the `TimerStarted` event that was recorded when this timer was started.
This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**timerId**

The unique ID of the timer that fired.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/timerfiredeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/timerfiredeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/timerfiredeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimerCanceledEventAttributes

TimerStartedEventAttributes

All content copied from https://docs.aws.amazon.com/.
