# TestType

Tests a registered extension to make sure it meets all necessary requirements for being
published in the CloudFormation registry.

- For resource types, this includes passing all contracts tests defined for the
type.

- For modules, this includes determining if the module's model meets all necessary
requirements.

For more information, see [Testing your public extension before publishing](../../../../services/cloudformation-cli/latest/userguide/publish-extension-publish-extension-testing.md) in the
_AWS CloudFormation Command Line Interface (CLI) User Guide_.

If you don't specify a version, CloudFormation uses the default version of the extension in
your account and Region for testing.

To perform testing, CloudFormation assumes the execution role specified when the type was
registered. For more information, see [RegisterType](api-registertype.md).

Once you've initiated testing on an extension using `TestType`, you can pass
the returned `TypeVersionArn` into [DescribeType](api-describetype.md) to
monitor the current test status and test status description for the extension.

An extension must have a test status of `PASSED` before it can be published.
For more information, see [Publishing extensions\
to make them available for public use](../../../../services/cloudformation-cli/latest/userguide/resource-type-publish.md) in the
_AWS CloudFormation Command Line Interface (CLI) User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Arn**

The Amazon Resource Name (ARN) of the extension.

Conditional: You must specify `Arn`, or `TypeName` and
`Type`.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+`

Required: No

**LogDeliveryBucket**

The S3 bucket to which CloudFormation delivers the contract test execution logs.

CloudFormation delivers the logs by the time contract testing has completed and the extension
has been assigned a test type status of `PASSED` or `FAILED`.

The user calling `TestType` must be able to access items in the specified S3
bucket. Specifically, the user needs the following permissions:

- `GetObject`

- `PutObject`

For more information, see [Actions, Resources, and\
Condition Keys for Amazon S3](../../../../services/service-authorization/latest/reference/list-amazons3.md) in the _AWS Identity and Access Management User Guide_.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `[\s\S]+`

Required: No

**Type**

The type of the extension to test.

Conditional: You must specify `Arn`, or `TypeName` and
`Type`.

Type: String

Valid Values: `RESOURCE | MODULE | HOOK`

Required: No

**TypeName**

The name of the extension to test.

Conditional: You must specify `Arn`, or `TypeName` and
`Type`.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 204.

Pattern: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

Required: No

**VersionId**

The version of the extension to test.

You can specify the version id with either `Arn`, or with `TypeName`
and `Type`.

If you don't specify a version, CloudFormation uses the default version of the extension in
this account and Region for testing.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[A-Za-z0-9-]+`

Required: No

## Response Elements

The following element is returned by the service.

**TypeVersionArn**

The Amazon Resource Name (ARN) of the extension.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CFNRegistry**

An error occurred during a CloudFormation registry operation.

**Message**

A message with details about the error that occurred.

HTTP Status Code: 400

**TypeNotFound**

The specified extension doesn't exist in the CloudFormation registry.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/testtype.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/testtype.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StopStackSetOperation

UpdateGeneratedTemplate
