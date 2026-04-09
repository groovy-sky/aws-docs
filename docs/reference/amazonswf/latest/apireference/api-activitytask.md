# ActivityTask

Unit of work sent to an activity worker.

## Contents

**activityId**

The unique ID of the task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**activityType**

The type of this activity task.

Type: [ActivityType](api-activitytype.md) object

Required: Yes

**startedEventId**

The ID of the `ActivityTaskStarted` event recorded in the history.

Type: Long

Required: Yes

**taskToken**

The opaque string used as a handle on the task. This token is used by workers to communicate progress and response information back to the system about the task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**workflowExecution**

The workflow execution that started this activity task.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: Yes

**input**

The inputs provided when the activity task was scheduled. The form of the input is user defined and should be meaningful to the activity implementation.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/activitytask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/activitytask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/activitytask.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

ActivityTaskCanceledEventAttributes

All content copied from https://docs.aws.amazon.com/.
