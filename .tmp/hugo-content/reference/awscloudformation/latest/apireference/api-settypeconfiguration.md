# SetTypeConfiguration

Specifies the configuration data for a CloudFormation extension, such as a resource or Hook,
in the given account and Region.

For more information, see [Edit configuration\
data for extensions in your account](../../../../services/cloudformation/latest/userguide/registry-set-configuration.md) in the
_AWS CloudFormation User Guide_.

To view the current configuration data for an extension, refer to the
`ConfigurationSchema` element of [DescribeType](api-describetype.md).

###### Important

It's strongly recommended that you use dynamic references to restrict sensitive
configuration definitions, such as third-party credentials. For more information, see [Specify values stored in other services using dynamic references](../../../../services/cloudformation/latest/userguide/dynamic-references.md) in the
_AWS CloudFormation User Guide_.

For more information about setting the configuration data for resource types, see [Defining the account-level configuration of an extension](../../../../services/cloudformation-cli/latest/userguide/resource-type-model.md#resource-type-howto-configuration) in the
_AWS CloudFormation Command Line Interface (CLI) User Guide_. For more information about setting the configuration
data for Hooks, see the [CloudFormation Hooks User Guide](../../../../services/cloudformation-cli/latest/hooks-userguide/what-is-cloudformation-hooks.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Configuration**

The configuration data for the extension in this account and Region.

The configuration data must be formatted as JSON and validate against the extension's
schema returned in the `Schema` response element of [DescribeType](api-describetype.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 204800.

Pattern: `[\s\S]+`

Required: Yes

**ConfigurationAlias**

An alias by which to refer to this extension configuration data.

Conditional: Specifying a configuration alias is required when setting a configuration for
a resource type extension.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `^[a-zA-Z0-9]{1,256}$`

Required: No

**Type**

The type of extension.

Conditional: You must specify `ConfigurationArn`, or `Type` and
`TypeName`.

Type: String

Valid Values: `RESOURCE | MODULE | HOOK`

Required: No

**TypeArn**

The Amazon Resource Name (ARN) for the extension in this account and Region.

For public extensions, this will be the ARN assigned when you call the [ActivateType](api-activatetype.md) API operation in this account and Region. For private extensions, this
will be the ARN assigned when you call the [RegisterType](api-registertype.md) API
operation in this account and Region.

Do not include the extension versions suffix at the end of the ARN. You can set the
configuration for an extension, but not for a specific extension version.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+`

Required: No

**TypeName**

The name of the extension.

Conditional: You must specify `ConfigurationArn`, or `Type` and
`TypeName`.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 204.

Pattern: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

Required: No

## Response Elements

The following element is returned by the service.

**ConfigurationArn**

The Amazon Resource Name (ARN) for the configuration data in this account and
Region.

Conditional: You must specify `ConfigurationArn`, or `Type` and
`TypeName`.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type-configuration/.+`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/settypeconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/settypeconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SetStackPolicy

SetTypeDefaultVersion

All content copied from https://docs.aws.amazon.com/.
