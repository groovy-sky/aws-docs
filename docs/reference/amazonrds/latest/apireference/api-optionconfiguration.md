# OptionConfiguration

A list of all available options for an option group.

## Contents

###### Note

In the following list, the required parameters are described first.

**OptionName**

The configuration of options to include in a group.

Type: String

Required: Yes

**DBSecurityGroupMemberships.DBSecurityGroupName.N**

A list of DB security groups used for this option.

Type: Array of strings

Required: No

**OptionSettings.OptionSetting.N**

The option settings to include in an option group.

Type: Array of [OptionSetting](api-optionsetting.md) objects

Required: No

**OptionVersion**

The version for the option.

Type: String

Required: No

**Port**

The optional port for the option.

Type: Integer

Required: No

**VpcSecurityGroupMemberships.VpcSecurityGroupId.N**

A list of VPC security group names used for this option.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/optionconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/optionconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/optionconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Option

OptionGroup

All content copied from https://docs.aws.amazon.com/.
