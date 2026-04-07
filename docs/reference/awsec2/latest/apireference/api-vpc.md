# Vpc

Describes a VPC.

## Contents

**blockPublicAccessStates**

The state of VPC Block Public Access (BPA).

Type: [BlockPublicAccessStates](api-blockpublicaccessstates.md) object

Required: No

**cidrBlock**

The primary IPv4 CIDR block for the VPC.

Type: String

Required: No

**CidrBlockAssociationSet.N**

Information about the IPv4 CIDR blocks associated with the VPC.

Type: Array of [VpcCidrBlockAssociation](api-vpccidrblockassociation.md) objects

Required: No

**dhcpOptionsId**

The ID of the set of DHCP options you've associated with the VPC.

Type: String

Required: No

**instanceTenancy**

The allowed tenancy of instances launched into the VPC.

Type: String

Valid Values: `default | dedicated | host`

Required: No

**Ipv6CidrBlockAssociationSet.N**

Information about the IPv6 CIDR blocks associated with the VPC.

Type: Array of [VpcIpv6CidrBlockAssociation](api-vpcipv6cidrblockassociation.md) objects

Required: No

**isDefault**

Indicates whether the VPC is the default VPC.

Type: Boolean

Required: No

**ownerId**

The ID of the AWS account that owns the VPC.

Type: String

Required: No

**state**

The current state of the VPC.

Type: String

Valid Values: `pending | available`

Required: No

**TagSet.N**

Any tags assigned to the VPC.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcId**

The ID of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpc.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpc.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpc.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VolumeStatusItem

VpcAttachment
