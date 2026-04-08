# RecordHandlerProgress

Reports progress of a resource handler to CloudFormation.

Reserved for use by the [CloudFormation\
CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md). Don't use this API in your code.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**BearerToken**

Reserved for use by the [CloudFormation\
CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**ClientRequestToken**

Reserved for use by the [CloudFormation\
CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9][-a-zA-Z0-9]*`

Required: No

**CurrentOperationStatus**

Reserved for use by the [CloudFormation\
CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md).

Type: String

Valid Values: `PENDING | IN_PROGRESS | SUCCESS | FAILED`

Required: No

**ErrorCode**

Reserved for use by the [CloudFormation\
CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md).

Type: String

Valid Values: `NotUpdatable | InvalidRequest | AccessDenied | InvalidCredentials | AlreadyExists | NotFound | ResourceConflict | Throttling | ServiceLimitExceeded | NotStabilized | GeneralServiceException | ServiceInternalError | NetworkFailure | InternalFailure | InvalidTypeConfiguration | HandlerInternalFailure | NonCompliant | Unknown | UnsupportedTarget`

Required: No

**OperationStatus**

Reserved for use by the [CloudFormation\
CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md).

Type: String

Valid Values: `PENDING | IN_PROGRESS | SUCCESS | FAILED`

Required: Yes

**ResourceModel**

Reserved for use by the [CloudFormation\
CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 16384.

Required: No

**StatusMessage**

Reserved for use by the [CloudFormation\
CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md).

Type: String

Length Constraints: Maximum length of 1024.

Required: No

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConditionalCheckFailed**

Error reserved for use by the [CloudFormation CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md). CloudFormation doesn't
return this error to users.

HTTP Status Code: 400

**InvalidStateTransition**

Error reserved for use by the [CloudFormation CLI](../../../../services/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md). CloudFormation doesn't
return this error to users.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/recordhandlerprogress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/recordhandlerprogress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PublishType

RegisterPublisher
