# TimerStartedEventAttributes

Provides the details of the `TimerStarted` event.

## Contents

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision task that resulted in the
`StartTimer` decision for this activity task. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**startToFireTimeout**

The duration of time after which the timer fires.

The duration is specified in seconds, an integer greater than or equal to `0`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8.

Required: Yes

**timerId**

The unique ID of the timer that was started.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**control**

Data attached to the event that can be used by the decider in subsequent workflow tasks.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/timerstartedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/timerstartedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/timerstartedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimerFiredEventAttributes

WorkflowExecution

All content copied from https://docs.aws.amazon.com/.
