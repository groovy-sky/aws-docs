# LocalGatewayRouteTableVirtualInterfaceGroupAssociation

Describes an association between a local gateway route table and a virtual interface group.

## Contents

**localGatewayId**

The ID of the local gateway.

Type: String

Required: No

**localGatewayRouteTableArn**

The Amazon Resource Name (ARN) of the local gateway route table for the virtual interface group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**localGatewayRouteTableId**

The ID of the local gateway route table.

Type: String

Required: No

**localGatewayRouteTableVirtualInterfaceGroupAssociationId**

The ID of the association.

Type: String

Required: No

**localGatewayVirtualInterfaceGroupId**

The ID of the virtual interface group.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the local gateway virtual interface group association.

Type: String

Required: No

**state**

The state of the association.

Type: String

Required: No

**TagSet.N**

The tags assigned to the association.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/localgatewayroutetablevirtualinterfacegroupassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/localgatewayroutetablevirtualinterfacegroupassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/localgatewayroutetablevirtualinterfacegroupassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LocalGatewayRouteTable

LocalGatewayRouteTableVpcAssociation
