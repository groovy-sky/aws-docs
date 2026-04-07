# LocalGatewayVirtualInterface

Describes a local gateway virtual interface.

## Contents

**configurationState**

The current state of the local gateway virtual interface.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**localAddress**

The local address.

Type: String

Required: No

**localBgpAsn**

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the local gateway.

Type: Integer

Required: No

**localGatewayId**

The ID of the local gateway.

Type: String

Required: No

**localGatewayVirtualInterfaceArn**

The Amazon Resource Number (ARN) of the local gateway virtual interface.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**localGatewayVirtualInterfaceGroupId**

The ID of the local gateway virtual interface group.

Type: String

Required: No

**localGatewayVirtualInterfaceId**

The ID of the virtual interface.

Type: String

Required: No

**outpostLagId**

The Outpost LAG ID.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the local gateway virtual interface.

Type: String

Required: No

**peerAddress**

The peer address.

Type: String

Required: No

**peerBgpAsn**

The peer BGP ASN.

Type: Integer

Required: No

**peerBgpAsnExtended**

The extended 32-bit ASN of the BGP peer for use with larger ASN values.

Type: Long

Required: No

**TagSet.N**

The tags assigned to the virtual interface.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vlan**

The ID of the VLAN.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/localgatewayvirtualinterface.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/localgatewayvirtualinterface.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/localgatewayvirtualinterface.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LocalGatewayRouteTableVpcAssociation

LocalGatewayVirtualInterfaceGroup
