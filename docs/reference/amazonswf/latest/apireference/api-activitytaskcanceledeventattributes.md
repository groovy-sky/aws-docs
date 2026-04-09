# ActivityTaskCanceledEventAttributes

Provides the details of the `ActivityTaskCanceled` event.

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

**details**

Details of the cancellation.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**latestCancelRequestedEventId**

If set, contains the ID of the last `ActivityTaskCancelRequested` event recorded for this activity task. This information can be useful for diagnosing problems by tracing back the chain of events leading up to this event.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/activitytaskcanceledeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/activitytaskcanceledeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/activitytaskcanceledeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivityTask

ActivityTaskCancelRequestedEventAttributes

All content copied from https://docs.aws.amazon.com/.
