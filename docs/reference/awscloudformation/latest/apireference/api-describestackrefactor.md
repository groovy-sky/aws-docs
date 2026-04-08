# DescribeStackRefactor

Describes the stack refactor status.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**StackRefactorId**

The ID associated with the stack refactor created from the [CreateStackRefactor](api-createstackrefactor.md) action.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**Description**

A description to help you identify the refactor.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**ExecutionStatus**

The stack refactor execution operation status that's provided after calling the [ExecuteStackRefactor](api-executestackrefactor.md) action.

Type: String

Valid Values: `UNAVAILABLE | AVAILABLE | OBSOLETE | EXECUTE_IN_PROGRESS | EXECUTE_COMPLETE | EXECUTE_FAILED | ROLLBACK_IN_PROGRESS | ROLLBACK_COMPLETE | ROLLBACK_FAILED`

**ExecutionStatusReason**

A detailed explanation for the stack refactor `ExecutionStatus`.

Type: String

**StackIds.member.N**

The unique ID for each stack.

Type: Array of strings

**StackRefactorId**

The ID associated with the stack refactor created from the [CreateStackRefactor](api-createstackrefactor.md) action.

Type: String

**Status**

The stack refactor operation status that's provided after calling the [CreateStackRefactor](api-createstackrefactor.md) action.

Type: String

Valid Values: `CREATE_IN_PROGRESS | CREATE_COMPLETE | CREATE_FAILED | DELETE_IN_PROGRESS | DELETE_COMPLETE | DELETE_FAILED`

**StatusReason**

A detailed explanation for the stack refactor operation `Status`.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**StackRefactorNotFound**

The specified stack refactor can't be found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/describestackrefactor.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/describestackrefactor.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeStackInstance

DescribeStackResource
