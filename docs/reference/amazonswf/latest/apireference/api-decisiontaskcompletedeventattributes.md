# DecisionTaskCompletedEventAttributes

Provides the details of the `DecisionTaskCompleted` event.

## Contents

**scheduledEventId**

The ID of the `DecisionTaskScheduled` event that was recorded when this decision task was scheduled.
This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**startedEventId**

The ID of the `DecisionTaskStarted` event recorded when this decision task was started.
This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**executionContext**

User defined context for the workflow execution.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**taskList**

Represents a task list.

Type: [TaskList](api-tasklist.md) object

Required: No

**taskListScheduleToStartTimeout**

The maximum amount of time the decision task can wait to be assigned to a worker.

Type: String

Length Constraints: Maximum length of 8.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/decisiontaskcompletedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/decisiontaskcompletedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/decisiontaskcompletedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecisionTask

DecisionTaskScheduledEventAttributes

All content copied from https://docs.aws.amazon.com/.
