# ListHookResults

Returns summaries of invoked Hooks. For more information, see [View invocation\
summaries for CloudFormation Hooks](../../../../services/cloudformation-cli/latest/hooks-userguide/hooks-view-invocations.md) in the _CloudFormation Hooks User Guide_.

This operation supports the following parameter combinations:

- No parameters: Returns all Hook invocation summaries.

- `TypeArn` only: Returns summaries for a specific Hook.

- `TypeArn` and `Status`: Returns summaries for a specific Hook
filtered by status.

- `TargetId` and `TargetType`: Returns summaries for a specific
Hook invocation target.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**Status**

Filters results by the status of Hook invocations. Can only be used in combination with
`TypeArn`. Valid values are:

- `HOOK_IN_PROGRESS`: The Hook is currently running.

- `HOOK_COMPLETE_SUCCEEDED`: The Hook completed successfully.

- `HOOK_COMPLETE_FAILED`: The Hook completed but failed validation.

- `HOOK_FAILED`: The Hook encountered an error during execution.

Type: String

Valid Values: `HOOK_IN_PROGRESS | HOOK_COMPLETE_SUCCEEDED | HOOK_COMPLETE_FAILED | HOOK_FAILED`

Required: No

**TargetId**

Filters results by the unique identifier of the target the Hook was invoked
against.

For change sets, this is the change set ARN. When the target is a Cloud Control API operation, this
value must be the `HookRequestToken` returned by the Cloud Control API request. For more
information on the `HookRequestToken`, see [ProgressEvent](../../../cloudcontrolapi/latest/apireference/api-progressevent.md).

Required when `TargetType` is specified and cannot be used otherwise.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Pattern: `[a-zA-Z][-a-zA-Z0-9]*|arn:[-a-zA-Z0-9:/]*|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**TargetType**

Filters results by target type. Currently, only `CHANGE_SET` and
`CLOUD_CONTROL` are supported filter options.

Required when `TargetId` is specified and cannot be used otherwise.

Type: String

Valid Values: `CHANGE_SET | STACK | RESOURCE | CLOUD_CONTROL`

Required: No

**TypeArn**

Filters results by the ARN of the Hook. Can be used alone or in combination with
`Status`.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/[A-Za-z0-9-]+/?`

Required: No

## Response Elements

The following elements are returned by the service.

**HookResults.member.N**

A list of `HookResultSummary` structures that provides the status and Hook
status reason for each Hook invocation for the specified target.

Type: Array of [HookResultSummary](api-hookresultsummary.md) objects

**NextToken**

Pagination token, `null` or empty if no more results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**TargetId**

The unique identifier of the Hook invocation target.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Pattern: `[a-zA-Z][-a-zA-Z0-9]*|arn:[-a-zA-Z0-9:/]*|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

**TargetType**

The target type.

Type: String

Valid Values: `CHANGE_SET | STACK | RESOURCE | CLOUD_CONTROL`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**HookResultNotFound**

The specified target doesn't have any requested Hook invocations.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/listhookresults.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/listhookresults.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListGeneratedTemplates

ListImports
