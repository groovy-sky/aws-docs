# StartTimerDecisionAttributes

Provides the details of the `StartTimer` decision.

**Access Control**

You can use IAM policies to control this decision's access to Amazon SWF resources as follows:

- Use a `Resource` element with the domain name to limit the action to only
specified domains.

- Use an `Action` element to allow or deny permission to call this action.

- You cannot use an IAM policy to constrain this action's parameters.

If the caller doesn't have sufficient permissions to invoke the action, or the
parameter values fall outside the specified constraints, the action fails. The associated event attribute's
`cause` parameter is set to `OPERATION_NOT_PERMITTED`. For details and example IAM policies, see
[Using IAM to Manage Access to Amazon SWF Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md) in the _Amazon SWF Developer Guide_.

## Contents

**startToFireTimeout**

The duration to wait before firing the timer.

The duration is specified in seconds, an integer greater than or equal to `0`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8.

Required: Yes

**timerId**

The unique ID of the timer.

The specified string must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**control**

The data attached to the event that can be used by the decider in subsequent workflow tasks.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/starttimerdecisionattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/starttimerdecisionattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/starttimerdecisionattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartLambdaFunctionFailedEventAttributes

StartTimerFailedEventAttributes

All content copied from https://docs.aws.amazon.com/.
