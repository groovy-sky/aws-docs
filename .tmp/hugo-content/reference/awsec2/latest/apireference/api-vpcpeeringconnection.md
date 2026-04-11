# VpcPeeringConnection

Describes a VPC peering connection.

## Contents

**accepterVpcInfo**

Information about the accepter VPC. CIDR block information is only returned when describing an active VPC peering connection.

Type: [VpcPeeringConnectionVpcInfo](api-vpcpeeringconnectionvpcinfo.md) object

Required: No

**expirationTime**

The time that an unaccepted VPC peering connection will expire.

Type: Timestamp

Required: No

**requesterVpcInfo**

Information about the requester VPC. CIDR block information is only returned when describing an active VPC peering connection.

Type: [VpcPeeringConnectionVpcInfo](api-vpcpeeringconnectionvpcinfo.md) object

Required: No

**status**

The status of the VPC peering connection.

Type: [VpcPeeringConnectionStateReason](api-vpcpeeringconnectionstatereason.md) object

Required: No

**TagSet.N**

Any tags assigned to the resource.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcPeeringConnectionId**

The ID of the VPC peering connection.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpcpeeringconnection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpcpeeringconnection.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpcpeeringconnection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpcIpv6CidrBlockAssociation

VpcPeeringConnectionOptionsDescription
