# LocalGatewayRouteTable

Describes a local gateway route table.

## Contents

**localGatewayId**

The ID of the local gateway.

Type: String

Required: No

**localGatewayRouteTableArn**

The Amazon Resource Name (ARN) of the local gateway route table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**localGatewayRouteTableId**

The ID of the local gateway route table.

Type: String

Required: No

**mode**

The mode of the local gateway route table.

Type: String

Valid Values: `direct-vpc-routing | coip`

Required: No

**outpostArn**

The Amazon Resource Name (ARN) of the Outpost.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the local gateway route table.

Type: String

Required: No

**state**

The state of the local gateway route table.

Type: String

Required: No

**stateReason**

Information about the state change.

Type: [StateReason](api-statereason.md) object

Required: No

**TagSet.N**

The tags assigned to the local gateway route table.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/localgatewayroutetable.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/localgatewayroutetable.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/localgatewayroutetable.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LocalGatewayRoute

LocalGatewayRouteTableVirtualInterfaceGroupAssociation
