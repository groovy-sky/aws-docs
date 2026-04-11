# ActivateType

Activates a public third-party extension, such as a resource or module, to make it
available for use in stack templates in your current account and Region. It can also create
CloudFormation Hooks, which allow you to evaluate resource configurations before CloudFormation
provisions them. Hooks integrate with both CloudFormation and Cloud Control API operations.

After you activate an extension, you can use [SetTypeConfiguration](api-settypeconfiguration.md) to set specific properties for the extension.

To see which extensions have been activated, use [ListTypes](api-listtypes.md). To see
configuration details for an extension, use [DescribeType](api-describetype.md).

For more information, see [Activate a\
third-party public extension in your account](../../../../services/cloudformation/latest/userguide/registry-public-activate-extension.md) in the
_AWS CloudFormation User Guide_. For information about creating Hooks, see the
[CloudFormation Hooks User Guide](../../../../services/cloudformation-cli/latest/hooks-userguide/what-is-cloudformation-hooks.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**AutoUpdate**

Whether to automatically update the extension in this account and Region when a new
_minor_ version is published by the extension publisher. Major versions
released by the publisher must be manually updated.

The default is `true`.

Type: Boolean

Required: No

**ExecutionRoleArn**

The name of the IAM execution role to use to activate the extension.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `arn:.+:iam::[0-9]{12}:role/.+`

Required: No

**LoggingConfig**

Contains logging configuration information for an extension.

Type: [LoggingConfig](api-loggingconfig.md) object

Required: No

**MajorVersion**

The major version of this extension you want to activate, if multiple major versions are
available. The default is the latest major version. CloudFormation uses the latest available
_minor_ version of the major version selected.

You can specify `MajorVersion` or `VersionBump`, but not
both.

Type: Long

Valid Range: Minimum value of 1. Maximum value of 100000.

Required: No

**PublicTypeArn**

The Amazon Resource Name (ARN) of the public extension.

Conditional: You must specify `PublicTypeArn`, or `TypeName`,
`Type`, and `PublisherId`.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}::type/.+/[0-9a-zA-Z]{12,40}/.+`

Required: No

**PublisherId**

The ID of the extension publisher.

Conditional: You must specify `PublicTypeArn`, or `TypeName`,
`Type`, and `PublisherId`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 40.

Pattern: `[0-9a-zA-Z]{12,40}`

Required: No

**Type**

The extension type.

Conditional: You must specify `PublicTypeArn`, or `TypeName`,
`Type`, and `PublisherId`.

Type: String

Valid Values: `RESOURCE | MODULE | HOOK`

Required: No

**TypeName**

The name of the extension.

Conditional: You must specify `PublicTypeArn`, or `TypeName`,
`Type`, and `PublisherId`.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 204.

Pattern: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

Required: No

**TypeNameAlias**

An alias to assign to the public extension in this account and Region. If you specify an
alias for the extension, CloudFormation treats the alias as the extension type name within this
account and Region. You must use the alias to refer to the extension in your templates, API
calls, and CloudFormation console.

An extension alias must be unique within a given account and Region. You can activate the
same public resource multiple times in the same account and Region, using different type name
aliases.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 204.

Pattern: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

Required: No

**VersionBump**

Manually updates a previously-activated type to a new major or minor version, if
available. You can also use this parameter to update the value of
`AutoUpdate`.

- `MAJOR`: CloudFormation updates the extension to the newest major version, if
one is available.

- `MINOR`: CloudFormation updates the extension to the newest minor version, if
one is available.

Type: String

Valid Values: `MAJOR | MINOR`

Required: No

## Response Elements

The following element is returned by the service.

**Arn**

The Amazon Resource Name (ARN) of the activated extension in this account and
Region.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:[0-9]{12}:type/.+`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/activatetype.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/activatetype.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivateOrganizationsAccess

BatchDescribeTypeConfigurations

All content copied from https://docs.aws.amazon.com/.
