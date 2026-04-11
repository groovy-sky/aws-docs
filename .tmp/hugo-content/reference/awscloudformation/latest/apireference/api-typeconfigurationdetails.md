# TypeConfigurationDetails

Detailed information concerning the specification of a CloudFormation extension in a given
account and Region.

For more information, see [Edit configuration data\
for extensions in your account](../../../../services/cloudformation/latest/userguide/registry-set-configuration.md) in the _AWS CloudFormation User Guide_.

## Contents

**Alias**

The alias specified for this configuration, if one was specified when the configuration was
set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `^[a-zA-Z0-9]{1,256}$`

Required: No

**Arn**

The ARN for the configuration data, in this account and Region.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type-configuration/.+`

Required: No

**Configuration**

A JSON string specifying the configuration data for the extension, in this account and
Region.

If a configuration hasn't been set for a specified extension, CloudFormation returns
`{}`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 204800.

Pattern: `[\s\S]+`

Required: No

**IsDefaultConfiguration**

Whether this configuration data is the default configuration for the extension.

Type: Boolean

Required: No

**LastUpdated**

When the configuration data was last updated for this extension.

If a configuration hasn't been set for a specified extension, CloudFormation returns
`null`.

Type: Timestamp

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

**TypeName**

The name of the extension.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 204.

Pattern: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/typeconfigurationdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/typeconfigurationdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/typeconfigurationdetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateSummaryConfig

TypeConfigurationIdentifier

All content copied from https://docs.aws.amazon.com/.
