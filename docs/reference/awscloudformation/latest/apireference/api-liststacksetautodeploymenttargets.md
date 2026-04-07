# ListStackSetAutoDeploymentTargets

Returns summary information about deployment targets for a StackSet.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CallAs**

Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.

By default, `SELF` is specified. Use `SELF` for StackSets with
self-managed permissions.

- If you are signed in to the management account, specify
`SELF`.

- If you are signed in to a delegated administrator account, specify
`DELEGATED_ADMIN`.

Your AWS account must be registered as a delegated administrator in the management account. For more information, see [Register a\
delegated administrator](../../../../services/cloudformation/latest/userguide/stacksets-orgs-delegated-admin.md) in the _AWS CloudFormation User Guide_.

Type: String

Valid Values: `SELF | DELEGATED_ADMIN`

Required: No

**MaxResults**

The maximum number of results to be returned with a single call. If the number of
available results exceeds this maximum, the response includes a `NextToken` value
that you can assign to the `NextToken` request parameter to get the next set of
results.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**NextToken**

The token for the next set of items to return. (You received this token from a previous call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**StackSetName**

The name or unique ID of the StackSet that you want to get automatic deployment targets
for.

Type: String

Pattern: `[a-zA-Z][-a-zA-Z0-9]*(?::[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12})?`

Required: Yes

## Response Elements

The following elements are returned by the service.

**NextToken**

If the request doesn't return all the remaining results, `NextToken` is set to
a token. To retrieve the next set of results, call [ListStackSetAutoDeploymentTargets](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackSetAutoDeploymentTargets.html) again and use that value for the
`NextToken` parameter. If the request returns all results, `NextToken`
is set to an empty string.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**Summaries.member.N**

An array of summaries of the deployment targets for the StackSet.

Type: Array of [StackSetAutoDeploymentTargetSummary](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetAutoDeploymentTargetSummary.html) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**StackSetNotFound**

The specified StackSet doesn't exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/cloudformation-2010-05-15/ListStackSetAutoDeploymentTargets)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListStacks

ListStackSetOperationResults
