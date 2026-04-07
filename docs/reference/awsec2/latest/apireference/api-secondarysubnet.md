# SecondarySubnet

Describes a secondary subnet.

## Contents

**availabilityZone**

The Availability Zone of the secondary subnet.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone of the secondary subnet.

Type: String

Required: No

**Ipv4CidrBlockAssociationSet.N**

Information about the IPv4 CIDR blocks associated with the secondary subnet.

Type: Array of [SecondarySubnetIpv4CidrBlockAssociation](api-secondarysubnetipv4cidrblockassociation.md) objects

Required: No

**ownerId**

The ID of the AWS account that owns the secondary subnet.

Type: String

Required: No

**secondaryNetworkId**

The ID of the secondary network.

Type: String

Required: No

**secondaryNetworkType**

The type of the secondary network.

Type: String

Valid Values: `rdma`

Required: No

**secondarySubnetArn**

The Amazon Resource Name (ARN) of the secondary subnet.

Type: String

Required: No

**secondarySubnetId**

The ID of the secondary subnet.

Type: String

Required: No

**state**

The state of the secondary subnet.

Type: String

Valid Values: `create-in-progress | create-complete | create-failed | delete-in-progress | delete-complete | delete-failed`

Required: No

**stateReason**

The reason for the current state of the secondary subnet.

Type: String

Required: No

**TagSet.N**

The tags assigned to the secondary subnet.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/secondarysubnet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/secondarysubnet.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/secondarysubnet.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecondaryNetworkIpv4CidrBlockAssociation

SecondarySubnetIpv4CidrBlockAssociation
