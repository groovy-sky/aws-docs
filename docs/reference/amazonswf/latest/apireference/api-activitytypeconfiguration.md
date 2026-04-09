# ActivityTypeConfiguration

Configuration settings registered with the activity type.

## Contents

**defaultTaskHeartbeatTimeout**

The default maximum time, in seconds, before which a worker processing a task must report
progress by calling [RecordActivityTaskHeartbeat](api-recordactivitytaskheartbeat.md).

You can specify this value only when _registering_ an activity type. The registered default value can be
overridden when you schedule a task through the `ScheduleActivityTask` [Decision](api-decision.md). If the activity
worker subsequently attempts to record a heartbeat or returns a result, the activity worker receives an
`UnknownResource` fault. In this case, Amazon SWF no longer considers the activity task to be valid;
the activity worker should clean up the activity task.

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**defaultTaskList**

The default task list specified for this activity type at registration. This default is used if
a task list isn't provided when a task is scheduled through the `ScheduleActivityTask` [Decision](api-decision.md). You can override the default registered task list when scheduling a task through the
`ScheduleActivityTask` [Decision](api-decision.md).

Type: [TaskList](api-tasklist.md) object

Required: No

**defaultTaskPriority**

The default task priority for tasks of this activity type, specified at registration. If not
set, then `0` is used as the default priority. This default can be overridden when scheduling an activity
task.

Valid values are integers that range from Java's `Integer.MIN_VALUE`
(-2147483648) to `Integer.MAX_VALUE` (2147483647). Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

**defaultTaskScheduleToCloseTimeout**

The default maximum duration, specified when registering the activity type, for tasks of this activity
type. You can override this default when scheduling a task through the `ScheduleActivityTask` [Decision](api-decision.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**defaultTaskScheduleToStartTimeout**

The default maximum duration, specified when registering the activity type, that a task of an
activity type can wait before being assigned to a worker. You can override this default when scheduling a task
through the `ScheduleActivityTask` [Decision](api-decision.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**defaultTaskStartToCloseTimeout**

The default maximum duration for tasks of an activity type specified when registering the activity
type. You can override this default when scheduling a task through the `ScheduleActivityTask` [Decision](api-decision.md).

The duration is specified in seconds, an integer greater than or equal to `0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/activitytypeconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/activitytypeconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/activitytypeconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivityType

ActivityTypeInfo

All content copied from https://docs.aws.amazon.com/.
