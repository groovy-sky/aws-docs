# DecisionTaskStartedEventAttributes

Provides the details of the `DecisionTaskStarted` event.

## Contents

**scheduledEventId**

The ID of the `DecisionTaskScheduled` event that was recorded when this decision task was scheduled.
This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**identity**

Identity of the decider making the request. This enables diagnostic tracing when problems arise. The form of this identity is user defined.

Type: String

Length Constraints: Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/decisiontaskstartedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/decisiontaskstartedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/decisiontaskstartedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecisionTaskScheduledEventAttributes

DecisionTaskTimedOutEventAttributes

All content copied from https://docs.aws.amazon.com/.
