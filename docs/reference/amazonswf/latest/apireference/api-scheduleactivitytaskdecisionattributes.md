# ScheduleActivityTaskDecisionAttributes

Provides the details of the `ScheduleActivityTask` decision.

**Access Control**

You can use IAM policies to control this decision's access to Amazon SWF resources as follows:

- Use a `Resource` element with the domain name to limit the action to only
specified domains.

- Use an `Action` element to allow or deny permission to call this action.

- Constrain the following parameters by using a `Condition` element with the
appropriate keys.

- `activityType.name` – String constraint. The key is `swf:activityType.name`.

- `activityType.version` – String constraint. The key is `swf:activityType.version`.

- `taskList` – String constraint. The key is `swf:taskList.name`.

If the caller doesn't have sufficient permissions to invoke the action, or the
parameter values fall outside the specified constraints, the action fails. The associated event attribute's
`cause` parameter is set to `OPERATION_NOT_PERMITTED`. For details and example IAM policies, see
[Using IAM to Manage Access to Amazon SWF Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md)
in the _Amazon SWF Developer Guide_.

## Contents

**activityId**

The `activityId` of the activity task.

The specified string must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**activityType**

The type of the activity task to schedule.

Type: [ActivityType](api-activitytype.md) object

Required: Yes

**control**

Data attached to the event that can be used by the decider in subsequent workflow tasks. This data isn't sent to the activity.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**heartbeatTimeout**

If set, specifies the maximum time before which a worker processing a task of this type must report progress by
calling [RecordActivityTaskHeartbeat](api-recordactivitytaskheartbeat.md). If the timeout is exceeded, the activity task is automatically timed
out. If the worker subsequently attempts to record a heartbeat or returns a result, it is ignored. This
overrides the default heartbeat timeout specified when registering the activity type using
[RegisterActivityType](api-registeractivitytype.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**input**

The input provided to the activity task.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**scheduleToCloseTimeout**

The maximum duration for this activity task.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

###### Note

A schedule-to-close timeout for this activity task must be specified either as a default for the activity type or through this field. If neither this field is set nor a default schedule-to-close timeout was specified at registration time then a fault is returned.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**scheduleToStartTimeout**

If set, specifies the maximum duration the activity task can wait to be assigned to a worker.
This overrides the default schedule-to-start timeout specified when registering the activity type using
[RegisterActivityType](api-registeractivitytype.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

###### Note

A schedule-to-start timeout for this activity task must be specified either as a default for the activity type or through this field. If neither this field is set nor a default schedule-to-start timeout was specified at registration time then a fault is returned.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**startToCloseTimeout**

If set, specifies the maximum duration a worker may take to process this activity task. This overrides the
default start-to-close timeout specified when registering the activity type using [RegisterActivityType](api-registeractivitytype.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

###### Note

A start-to-close timeout for this activity task must be specified either as a default for the activity type or through this field. If neither this field is set nor a default start-to-close timeout was specified at registration time then a fault is returned.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**taskList**

If set, specifies the name of the task list in which to schedule the activity task. If not specified, the
`defaultTaskList` registered with the activity type is used.

###### Note

A task list for this activity task must be specified either as a default for the activity type or through this field. If neither this field is set nor a default task list was specified at registration time then a fault is returned.

The specified string must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: [TaskList](api-tasklist.md) object

Required: No

**taskPriority**

If set, specifies the priority with which the activity task is to be assigned to a worker. This
overrides the defaultTaskPriority specified when registering the activity type using [RegisterActivityType](api-registeractivitytype.md).
Valid values are integers that range from Java's `Integer.MIN_VALUE`
(-2147483648) to `Integer.MAX_VALUE` (2147483647). Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/scheduleactivitytaskdecisionattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/scheduleactivitytaskdecisionattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/scheduleactivitytaskdecisionattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTag

ScheduleActivityTaskFailedEventAttributes

All content copied from https://docs.aws.amazon.com/.
