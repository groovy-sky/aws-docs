# SecondaryNetwork

Describes a secondary network.

## Contents

**Ipv4CidrBlockAssociationSet.N**

Information about the IPv4 CIDR blocks associated with the secondary network.

Type: Array of [SecondaryNetworkIpv4CidrBlockAssociation](api-secondarynetworkipv4cidrblockassociation.md) objects

Required: No

**ownerId**

The ID of the AWS account that owns the secondary network.

Type: String

Required: No

**secondaryNetworkArn**

The Amazon Resource Name (ARN) of the secondary network.

Type: String

Required: No

**secondaryNetworkId**

The ID of the secondary network.

Type: String

Required: No

**state**

The state of the secondary network.

Type: String

Valid Values: `create-in-progress | create-complete | create-failed | delete-in-progress | delete-complete | delete-failed`

Required: No

**stateReason**

The reason for the current state of the secondary network.

Type: String

Required: No

**TagSet.N**

The tags assigned to the secondary network.

Type: Array of [Tag](api-tag.md) objects

Required: No

**type**

The type of the secondary network.

Type: String

Valid Values: `rdma`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/secondarynetwork.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/secondarynetwork.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/secondarynetwork.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecondaryInterfacePrivateIpAddressSpecificationRequest

SecondaryNetworkIpv4CidrBlockAssociation
