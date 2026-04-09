# ActivityTaskCancelRequestedEventAttributes

Provides the details of the `ActivityTaskCancelRequested` event.

## Contents

**activityId**

The unique ID of the task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision task that resulted in the
`RequestCancelActivityTask` decision for this cancellation request. This information can be useful for diagnosing problems by tracing back the chain of events leading up to this event.

Type: Long

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/activitytaskcancelrequestedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/activitytaskcancelrequestedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/activitytaskcancelrequestedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivityTaskCanceledEventAttributes

ActivityTaskCompletedEventAttributes

All content copied from https://docs.aws.amazon.com/.
