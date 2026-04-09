# DecisionTaskScheduledEventAttributes

Provides details about the `DecisionTaskScheduled` event.

## Contents

**taskList**

The name of the task list in which the decision task was scheduled.

Type: [TaskList](api-tasklist.md) object

Required: Yes

**scheduleToStartTimeout**

The maximum amount of time the decision task can wait to be assigned to a worker.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**startToCloseTimeout**

The maximum duration for this decision task. The task is considered timed out if it doesn't completed within this duration.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**taskPriority**

A task priority that, if set, specifies the priority for this decision task.
Valid values are integers that range from Java's `Integer.MIN_VALUE`
(-2147483648) to `Integer.MAX_VALUE` (2147483647). Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/decisiontaskscheduledeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/decisiontaskscheduledeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/decisiontaskscheduledeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecisionTaskCompletedEventAttributes

DecisionTaskStartedEventAttributes

All content copied from https://docs.aws.amazon.com/.
