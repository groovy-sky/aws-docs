# LocalGatewayRouteTableVpcAssociation

Describes an association between a local gateway route table and a VPC.

## Contents

**localGatewayId**

The ID of the local gateway.

Type: String

Required: No

**localGatewayRouteTableArn**

The Amazon Resource Name (ARN) of the local gateway route table for the association.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**localGatewayRouteTableId**

The ID of the local gateway route table.

Type: String

Required: No

**localGatewayRouteTableVpcAssociationId**

The ID of the association.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the local gateway route table for the association.

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

**vpcId**

The ID of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/localgatewayroutetablevpcassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/localgatewayroutetablevpcassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/localgatewayroutetablevpcassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LocalGatewayRouteTableVirtualInterfaceGroupAssociation

LocalGatewayVirtualInterface
