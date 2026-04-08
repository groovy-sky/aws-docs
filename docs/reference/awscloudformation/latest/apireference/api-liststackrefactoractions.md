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

Type: Array of [StackRefactorAction](api-stackrefactoraction.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/liststackrefactoractions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/liststackrefactoractions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListStackInstances

ListStackRefactors
