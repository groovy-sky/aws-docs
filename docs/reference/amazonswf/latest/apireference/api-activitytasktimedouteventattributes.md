# ActivityTaskTimedOutEventAttributes

Provides the details of the `ActivityTaskTimedOut` event.

## Contents

**scheduledEventId**

The ID of the `ActivityTaskScheduled` event that was recorded when this activity task was scheduled. This information can be useful for diagnosing problems by tracing back the chain of events leading up to this event.

Type: Long

Required: Yes

**startedEventId**

The ID of the `ActivityTaskStarted` event recorded when this activity task was started. This
information can be useful for diagnosing problems by tracing back the chain of events leading up to this
event.

Type: Long

Required: Yes

**timeoutType**

The type of the timeout that caused this event.

Type: String

Valid Values: `START_TO_CLOSE | SCHEDULE_TO_START | SCHEDULE_TO_CLOSE | HEARTBEAT`

Required: Yes

**details**

Contains the content of the `details` parameter for the last call made by the activity to
`RecordActivityTaskHeartbeat`.

Type: String

Length Constraints: Maximum length of 2048.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/activitytasktimedouteventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/activitytasktimedouteventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/activitytasktimedouteventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivityTaskStartedEventAttributes

ActivityType

All content copied from https://docs.aws.amazon.com/.
