# ScheduleLambdaFunctionFailedEventAttributes

Provides the details of the `ScheduleLambdaFunctionFailed` event. It isn't
set for other event types.

## Contents

**cause**

The cause of the failure. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

###### Note

If `cause` is set to `OPERATION_NOT_PERMITTED`, the decision
failed because it lacked sufficient permissions. For details and example IAM policies, see
[Using\
IAM to Manage Access to Amazon SWF Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md) in the
_Amazon SWF Developer Guide_.

Type: String

Valid Values: `ID_ALREADY_IN_USE | OPEN_LAMBDA_FUNCTIONS_LIMIT_EXCEEDED | LAMBDA_FUNCTION_CREATION_RATE_EXCEEDED | LAMBDA_SERVICE_NOT_AVAILABLE_IN_REGION`

Required: Yes

**decisionTaskCompletedEventId**

The ID of the `LambdaFunctionCompleted` event corresponding to the decision
that resulted in scheduling this Lambda task. To help diagnose issues, use this information to trace back the chain of events leading up to this event.

Type: Long

Required: Yes

**id**

The ID provided in the `ScheduleLambdaFunction` decision that failed.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**name**

The name of the Lambda function.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/schedulelambdafunctionfailedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/schedulelambdafunctionfailedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/schedulelambdafunctionfailedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScheduleLambdaFunctionDecisionAttributes

SignalExternalWorkflowExecutionDecisionAttributes

All content copied from https://docs.aws.amazon.com/.
