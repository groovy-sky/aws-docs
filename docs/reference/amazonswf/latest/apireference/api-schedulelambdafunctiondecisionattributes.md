# ScheduleLambdaFunctionDecisionAttributes

Decision attributes specified in `scheduleLambdaFunctionDecisionAttributes` within the list of
decisions `decisions` passed to [RespondDecisionTaskCompleted](api-responddecisiontaskcompleted.md).

## Contents

**id**

A string that identifies the Lambda function execution in the event history.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**name**

The name, or ARN, of the Lambda function to schedule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**control**

The data attached to the event that the decider can use in subsequent workflow tasks.
This data isn't sent to the Lambda task.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**input**

The optional input data to be supplied to the Lambda function.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 32768.

Required: No

**startToCloseTimeout**

The timeout value, in seconds, after which the Lambda function is considered to be
failed once it has started. This can be any integer from 1-900 (1s-15m).

If no value is supplied, then a default value of 300s is assumed.

Type: String

Length Constraints: Maximum length of 8.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/schedulelambdafunctiondecisionattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/schedulelambdafunctiondecisionattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/schedulelambdafunctiondecisionattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScheduleActivityTaskFailedEventAttributes

ScheduleLambdaFunctionFailedEventAttributes

All content copied from https://docs.aws.amazon.com/.
