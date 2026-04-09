# DecisionTaskTimedOutEventAttributes

Provides the details of the `DecisionTaskTimedOut` event.

## Contents

**scheduledEventId**

The ID of the `DecisionTaskScheduled` event that was recorded when this decision task was scheduled.
This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**startedEventId**

The ID of the `DecisionTaskStarted` event recorded when this decision task was started. This
information can be useful for diagnosing problems by tracing back the chain of events leading up to this
event.

Type: Long

Required: Yes

**timeoutType**

The type of timeout that expired before the decision task could be completed.

Type: String

Valid Values: `START_TO_CLOSE | SCHEDULE_TO_START`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/decisiontasktimedouteventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/decisiontasktimedouteventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/decisiontasktimedouteventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecisionTaskStartedEventAttributes

DomainConfiguration

All content copied from https://docs.aws.amazon.com/.
