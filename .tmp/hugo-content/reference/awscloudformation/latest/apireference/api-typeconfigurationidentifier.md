# TypeConfigurationIdentifier

Identifying information for the configuration of a CloudFormation extension.

## Contents

**Type**

The type of extension.

Type: String

Valid Values: `RESOURCE | MODULE | HOOK`

Required: No

**TypeArn**

The ARN for the extension, in this account and Region.

For public extensions, this will be the ARN assigned when you call the [ActivateType](api-activatetype.md) API operation in this account and Region. For private extensions, this
will be the ARN assigned when you call the [RegisterType](api-registertype.md) API
operation in this account and Region.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+`

Required: No

**TypeConfigurationAlias**

The alias specified for this configuration, if one was specified when the configuration was
set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `^[a-zA-Z0-9]{1,256}$`

Required: No

**TypeConfigurationArn**

The ARN for the configuration, in this account and Region.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type-configuration/.+`

Required: No

**TypeName**

The name of the extension type to which this configuration applies.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 204.

Pattern: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/typeconfigurationidentifier.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/typeconfigurationidentifier.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/typeconfigurationidentifier.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TypeConfigurationDetails

TypeFilters

All content copied from https://docs.aws.amazon.com/.
