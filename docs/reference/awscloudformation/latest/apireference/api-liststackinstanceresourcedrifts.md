# ListStackInstanceResourceDrifts

Returns drift information for resources in a stack instance.

###### Note

`ListStackInstanceResourceDrifts` returns drift information for the most
recent drift detection operation. If an operation is in progress, it may only return partial
results.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CallAs**

\[Service-managed permissions\] Specifies whether you are acting as an account administrator
in the organization's management account or as a delegated administrator in a
member account.

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

**OperationId**

The unique ID of the drift operation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9][-a-zA-Z0-9]*`

Required: Yes

**StackInstanceAccount**

The name of the AWS account that you want to list resource drifts for.

Type: String

Pattern: `^[0-9]{12}$`

Required: Yes

**StackInstanceRegion**

The name of the Region where you want to list resource drifts.

Type: String

Pattern: `^[a-zA-Z0-9-]{1,128}$`

Required: Yes

**StackInstanceResourceDriftStatuses.member.N**

The resource drift status of the stack instance.

- `DELETED`: The resource differs from its expected template configuration in
that the resource has been deleted.

- `MODIFIED`: One or more resource properties differ from their expected
template values.

- `IN_SYNC`: The resource's actual configuration matches its expected
template configuration.

- `NOT_CHECKED`: CloudFormation doesn't currently return this value.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 4 items.

Valid Values: `IN_SYNC | MODIFIED | DELETED | NOT_CHECKED | UNKNOWN | UNSUPPORTED`

Required: No

**StackSetName**

The name or unique ID of the StackSet that you want to list drifted resources for.

Type: String

Pattern: `[a-zA-Z][-a-zA-Z0-9]*(?::[a-zA-Z0-9]{8}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{12})?`

Required: Yes

## Response Elements

The following elements are returned by the service.

**NextToken**

If the previous paginated request didn't return all of the remaining results, the response
object's `NextToken` parameter value is set to a token. To retrieve the next set of
results, call this action again and assign that token to the request object's
`NextToken` parameter. If there are no remaining results, the previous response
object's `NextToken` parameter is set to `null`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**Summaries.member.N**

A list of `StackInstanceResourceDriftsSummary` structures that contain
information about the specified stack instances.

Type: Array of [StackInstanceResourceDriftsSummary](api-stackinstanceresourcedriftssummary.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**OperationNotFound**

The specified ID refers to an operation that doesn't exist.

HTTP Status Code: 404

**StackInstanceNotFound**

The specified stack instance doesn't exist.

HTTP Status Code: 404

**StackSetNotFound**

The specified StackSet doesn't exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/liststackinstanceresourcedrifts.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListResourceScans

ListStackInstances
