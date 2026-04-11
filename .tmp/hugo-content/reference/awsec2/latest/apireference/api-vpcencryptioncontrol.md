# VpcEncryptionControl

Describes the configuration and state of VPC encryption controls.

For more information, see [Enforce VPC encryption in transit](../../../../services/vpc/latest/userguide/vpc-encryption-controls.md) in the _Amazon VPC User Guide_.

## Contents

**mode**

The encryption mode for the VPC Encryption Control configuration.

Type: String

Valid Values: `monitor | enforce`

Required: No

**resourceExclusions**

Information about resource exclusions for the VPC Encryption Control configuration.

Type: [VpcEncryptionControlExclusions](api-vpcencryptioncontrolexclusions.md) object

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpcencryptioncontrol.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpcencryptioncontrol.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpcencryptioncontrol.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpcClassicLink

VpcEncryptionControlConfiguration
