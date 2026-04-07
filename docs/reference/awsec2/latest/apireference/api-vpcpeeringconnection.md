# VpcPeeringConnection

Describes a VPC peering connection.

## Contents

**accepterVpcInfo**

Information about the accepter VPC. CIDR block information is only returned when describing an active VPC peering connection.

Type: [VpcPeeringConnectionVpcInfo](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcPeeringConnectionVpcInfo.html) object

Required: No

**expirationTime**

The time that an unaccepted VPC peering connection will expire.

Type: Timestamp

Required: No

**requesterVpcInfo**

Information about the requester VPC. CIDR block information is only returned when describing an active VPC peering connection.

Type: [VpcPeeringConnectionVpcInfo](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcPeeringConnectionVpcInfo.html) object

Required: No

**status**

The status of the VPC peering connection.

Type: [VpcPeeringConnectionStateReason](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpcPeeringConnectionStateReason.html) object

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/VpcPeeringConnection)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/VpcPeeringConnection)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/VpcPeeringConnection)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcIpv6CidrBlockAssociation

VpcPeeringConnectionOptionsDescription
