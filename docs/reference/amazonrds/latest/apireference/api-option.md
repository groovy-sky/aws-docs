# Option

The details of an option.

## Contents

###### Note

In the following list, the required parameters are described first.

**DBSecurityGroupMemberships.DBSecurityGroup.N**

If the option requires access to a port, then this DB security group allows access to the port.

Type: Array of [DBSecurityGroupMembership](api-dbsecuritygroupmembership.md) objects

Required: No

**OptionDescription**

The description of the option.

Type: String

Required: No

**OptionName**

The name of the option.

Type: String

Required: No

**OptionSettings.OptionSetting.N**

The option settings for this option.

Type: Array of [OptionSetting](api-optionsetting.md) objects

Required: No

**OptionVersion**

The version of the option.

Type: String

Required: No

**Permanent**

Indicates whether this option is permanent.

Type: Boolean

Required: No

**Persistent**

Indicates whether this option is persistent.

Type: Boolean

Required: No

**Port**

If required, the port configured for this option to use.

Type: Integer

Required: No

**VpcSecurityGroupMemberships.VpcSecurityGroupMembership.N**

If the option requires access to a port, then this VPC security group allows access to the port.

Type: Array of [VpcSecurityGroupMembership](api-vpcsecuritygroupmembership.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/option.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/option.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/option.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyAdditionalStorageVolume

OptionConfiguration

All content copied from https://docs.aws.amazon.com/.
