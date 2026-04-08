# OptionSetting

Option settings are the actual settings being applied or configured for that option. It is used when you modify an option group or describe option groups. For example, the NATIVE\_NETWORK\_ENCRYPTION option has a setting called SQLNET.ENCRYPTION\_SERVER that can have several different values.

## Contents

###### Note

In the following list, the required parameters are described first.

**AllowedValues**

The allowed values of the option setting.

Type: String

Required: No

**ApplyType**

The DB engine specific parameter type.

Type: String

Required: No

**DataType**

The data type of the option setting.

Type: String

Required: No

**DefaultValue**

The default value of the option setting.

Type: String

Required: No

**Description**

The description of the option setting.

Type: String

Required: No

**IsCollection**

Indicates whether the option setting is part of a collection.

Type: Boolean

Required: No

**IsModifiable**

Indicates whether the option setting can be modified from the default.

Type: Boolean

Required: No

**Name**

The name of the option that has settings that you can set.

Type: String

Required: No

**Value**

The current value of the option setting.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/optionsetting.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/optionsetting.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/optionsetting.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OptionGroupOptionSetting

OptionVersion

All content copied from https://docs.aws.amazon.com/.
