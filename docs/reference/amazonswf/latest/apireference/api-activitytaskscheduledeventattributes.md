# ActivityTaskScheduledEventAttributes

Provides the details of the `ActivityTaskScheduled` event.

## Contents

**activityId**

The unique ID of the activity task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**activityType**

The type of the activity task.

Type: [ActivityType](api-activitytype.md) object

Required: Yes

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision that resulted in the scheduling of this activity task. This information can be useful for diagnosing problems by tracing back the chain of events leading up to this event.

Type: Long

Required: Yes

**taskList**

The task list in which the activity task has been scheduled.

Type: [TaskList](api-tasklist.md) object

Required: Yes

**control**

Data attached to the event that can be used by the decider in subsequent workflow tasks. This data isn't sent to the activity.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**heartbeatTimeout**

The maximum time before which the worker processing this task must report progress by calling
[RecordActivityTaskHeartbeat](api-recordactivitytaskheartbeat.md). If the timeout is exceeded, the activity task is automatically timed out. If
the worker subsequently attempts to record a heartbeat or return a result, it is ignored.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**input**

The input provided to the activity task.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**scheduleToCloseTimeout**

The maximum amount of time for this activity task.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**scheduleToStartTimeout**

The maximum amount of time the activity task can wait to be assigned to a worker.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**startToCloseTimeout**

The maximum amount of time a worker may take to process the activity task.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**taskPriority**

The priority to assign to the scheduled activity task. If set, this overrides any default
priority value that was assigned when the activity type was registered.

Valid values are integers that range from Java's `Integer.MIN_VALUE`
(-2147483648) to `Integer.MAX_VALUE` (2147483647). Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/activitytaskscheduledeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/activitytaskscheduledeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/activitytaskscheduledeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivityTaskFailedEventAttributes

ActivityTaskStartedEventAttributes

All content copied from https://docs.aws.amazon.com/.
