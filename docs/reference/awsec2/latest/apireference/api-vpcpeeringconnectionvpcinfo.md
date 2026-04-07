# VpcPeeringConnectionVpcInfo

Describes a VPC in a VPC peering connection.

## Contents

**cidrBlock**

The IPv4 CIDR block for the VPC.

Type: String

Required: No

**CidrBlockSet.N**

Information about the IPv4 CIDR blocks for the VPC.

Type: Array of [CidrBlock](api-cidrblock.md) objects

Required: No

**Ipv6CidrBlockSet.N**

The IPv6 CIDR block for the VPC.

Type: Array of [Ipv6CidrBlock](api-ipv6cidrblock.md) objects

Required: No

**ownerId**

The ID of the AWS account that owns the VPC.

Type: String

Required: No

**peeringOptions**

Information about the VPC peering connection options for the accepter or requester VPC.

Type: [VpcPeeringConnectionOptionsDescription](api-vpcpeeringconnectionoptionsdescription.md) object

Required: No

**region**

The Region in which the VPC is located.

Type: String

Required: No

**vpcId**

The ID of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpcpeeringconnectionvpcinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpcpeeringconnectionvpcinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpcpeeringconnectionvpcinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpcPeeringConnectionStateReason

VpnConcentrator
