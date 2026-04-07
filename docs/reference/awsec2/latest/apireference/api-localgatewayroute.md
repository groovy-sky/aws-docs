# LocalGatewayRoute

Describes a route for a local gateway route table.

## Contents

**coipPoolId**

The ID of the customer-owned address pool.

Type: String

Required: No

**destinationCidrBlock**

The CIDR block used for destination matches.

Type: String

Required: No

**destinationPrefixListId**

The ID of the prefix list.

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

**localGatewayVirtualInterfaceGroupId**

The ID of the virtual interface group.

Type: String

Required: No

**networkInterfaceId**

The ID of the network interface.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the local gateway route.

Type: String

Required: No

**state**

The state of the route.

Type: String

Valid Values: `pending | active | blackhole | deleting | deleted`

Required: No

**subnetId**

The ID of the subnet.

Type: String

Required: No

**type**

The route type.

Type: String

Valid Values: `static | propagated`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/localgatewayroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/localgatewayroute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/localgatewayroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LocalGateway

LocalGatewayRouteTable
