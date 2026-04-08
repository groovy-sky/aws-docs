# OptionGroup

## Contents

###### Note

In the following list, the required parameters are described first.

**AllowsVpcAndNonVpcInstanceMemberships**

Indicates whether this option group can be applied to both VPC
and non-VPC instances. The value `true` indicates the option group
can be applied to both VPC and non-VPC instances.

Type: Boolean

Required: No

**CopyTimestamp**

Indicates when the option group was copied.

Type: Timestamp

Required: No

**EngineName**

Indicates the name of the engine that this option group can be applied to.

Type: String

Required: No

**MajorEngineVersion**

Indicates the major engine version associated with this option group.

Type: String

Required: No

**OptionGroupArn**

Specifies the Amazon Resource Name (ARN) for the option group.

Type: String

Required: No

**OptionGroupDescription**

Provides a description of the option group.

Type: String

Required: No

**OptionGroupName**

Specifies the name of the option group.

Type: String

Required: No

**Options.Option.N**

Indicates what options are available in the option group.

Type: Array of [Option](api-option.md) objects

Required: No

**SourceAccountId**

Specifies the AWS account ID for the option group from which this option group is copied.

Type: String

Required: No

**SourceOptionGroup**

Specifies the name of the option group from which this option group is copied.

Type: String

Required: No

**VpcId**

If **AllowsVpcAndNonVpcInstanceMemberships** is `false`, this field is blank.
If **AllowsVpcAndNonVpcInstanceMemberships** is `true` and this field is blank,
then this option group can be applied to both VPC and non-VPC instances.
If this field contains a value, then this option group can only be
applied to instances that are in the VPC indicated by this field.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/optiongroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/optiongroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/optiongroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OptionConfiguration

OptionGroupMembership

All content copied from https://docs.aws.amazon.com/.
