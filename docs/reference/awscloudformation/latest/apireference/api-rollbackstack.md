# RollbackStack

When specifying `RollbackStack`, you preserve the state of previously
provisioned resources when an operation fails. You can check the status of the stack through
the [DescribeStacks](api-describestacks.md) operation.

Rolls back the specified stack to the last known stable state from
`CREATE_FAILED` or `UPDATE_FAILED` stack statuses.

This operation will delete a stack if it doesn't contain a last known stable state. A last
known stable state includes any status in a `*_COMPLETE`. This includes the
following stack statuses.

- `CREATE_COMPLETE`

- `UPDATE_COMPLETE`

- `UPDATE_ROLLBACK_COMPLETE`

- `IMPORT_COMPLETE`

- `IMPORT_ROLLBACK_COMPLETE`

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ClientRequestToken**

A unique identifier for this `RollbackStack` request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9][-a-zA-Z0-9]*`

Required: No

**RetainExceptOnCreate**

When set to `true`, newly created resources are deleted when the operation
rolls back. This includes newly created resources marked with a deletion policy of
`Retain`.

Default: `false`

Type: Boolean

Required: No

**RoleARN**

The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to rollback the
stack.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Required: No

**StackName**

The name that's associated with the stack.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `([a-zA-Z][-a-zA-Z0-9]*)|(arn:\b(aws|aws-us-gov|aws-cn)\b:[-a-zA-Z0-9:/._+]*)`

Required: Yes

## Response Elements

The following elements are returned by the service.

**OperationId**

A unique identifier for this rollback operation that can be used to track the operation's
progress and events.

Type: String

**StackId**

Unique identifier of the stack.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**TokenAlreadyExists**

A client request token already exists.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/rollbackstack.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/rollbackstack.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RegisterType

SetStackPolicy
