# VpcEncryptionControl

Describes the configuration and state of VPC encryption controls.

For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the _Amazon VPC User Guide_.

## Contents

**mode**

The encryption mode for the VPC Encryption Control configuration.

Type: String

Valid Values: `monitor | enforce`

Required: No

**resourceExclusions**

Information about resource exclusions for the VPC Encryption Control configuration.

Type: [VpcEncryptionControlExclusions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcEncryptionControlExclusions.html) object

Required: No

**state**

The current state of the VPC Encryption Control configuration.

Type: String

Valid Values: `enforce-in-progress | monitor-in-progress | enforce-failed | monitor-failed | deleting | deleted | available | creating | delete-failed`

Required: No

**stateMessage**

A message providing additional information about the encryption control state.

Type: String

Required: No

**TagSet.N**

The tags assigned to the VPC Encryption Control configuration.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcEncryptionControlId**

The ID of the VPC Encryption Control configuration.

Type: String

Required: No

**vpcId**

The ID of the VPC associated with the encryption control configuration.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/VpcEncryptionControl)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/VpcEncryptionControl)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/VpcEncryptionControl)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcClassicLink

VpcEncryptionControlConfiguration
