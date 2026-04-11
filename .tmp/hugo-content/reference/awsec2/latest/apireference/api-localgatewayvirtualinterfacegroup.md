# LocalGatewayVirtualInterfaceGroup

Describes a local gateway virtual interface group.

## Contents

**configurationState**

The current state of the local gateway virtual interface group.

Type: String

Valid Values: `pending | incomplete | available | deleting | deleted`

Required: No

**localBgpAsn**

The Autonomous System Number(ASN) for the local Border Gateway Protocol (BGP).

Type: Integer

Required: No

**localBgpAsnExtended**

The extended 32-bit ASN for the local BGP configuration.

Type: Long

Required: No

**localGatewayId**

The ID of the local gateway.

Type: String

Required: No

**localGatewayVirtualInterfaceGroupArn**

The Amazon Resource Number (ARN) of the local gateway virtual interface group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**localGatewayVirtualInterfaceGroupId**

The ID of the virtual interface group.

Type: String

Required: No

**LocalGatewayVirtualInterfaceIdSet.N**

The IDs of the virtual interfaces.

Type: Array of strings

Required: No

**ownerId**

The ID of the AWS account that owns the local gateway virtual interface group.

Type: String

Required: No

**TagSet.N**

The tags assigned to the virtual interface group.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/localgatewayvirtualinterfacegroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/localgatewayvirtualinterfacegroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/localgatewayvirtualinterfacegroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LocalGatewayVirtualInterface

LockedSnapshotsInfo
