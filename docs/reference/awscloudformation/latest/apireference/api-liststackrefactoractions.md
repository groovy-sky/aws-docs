# ListStackRefactorActions

Lists the stack refactor actions that will be taken after calling the [ExecuteStackRefactor](api-executestackrefactor.md) action.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**MaxResults**

The maximum number of results to be returned with a single call. If the number of
available results exceeds this maximum, the response includes a `NextToken` value
that you can assign to the `NextToken` request parameter to get the next set of
results.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**StackRefactorId**

The ID associated with the stack refactor created from the [CreateStackRefactor](api-createstackrefactor.md) action.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**NextToken**

If the request doesn't return all the remaining results, `NextToken` is set to
a token. To retrieve the next set of results, call this action again and assign that token to
the request object's `NextToken` parameter. If the request returns all results,
`NextToken` is set to `null`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**StackRefactorActions.member.N**

The stack refactor actions.

Type: Array of [StackRefactorAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackRefactorAction.html) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/cloudformation-2010-05-15/ListStackRefactorActions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/cloudformation-2010-05-15/ListStackRefactorActions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListStackInstances

ListStackRefactors
