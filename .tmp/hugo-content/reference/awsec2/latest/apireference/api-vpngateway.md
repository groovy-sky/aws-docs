# VpnGateway

Describes a virtual private gateway.

## Contents

**amazonSideAsn**

The private Autonomous System Number (ASN) for the Amazon side of a BGP
session.

Type: Long

Required: No

**Attachments.N**

Any VPCs attached to the virtual private gateway.

Type: Array of [VpcAttachment](api-vpcattachment.md) objects

Required: No

**availabilityZone**

The Availability Zone where the virtual private gateway was created, if applicable.
This field may be empty or not returned.

Type: String

Required: No

**state**

The current state of the virtual private gateway.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**TagSet.N**

Any tags assigned to the virtual private gateway.

Type: Array of [Tag](api-tag.md) objects

Required: No

**type**

The type of VPN connection the virtual private gateway supports.

Type: String

Valid Values: `ipsec.1`

Required: No

**vpnGatewayId**

The ID of the virtual private gateway.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpngateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpngateway.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpngateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpnConnectionOptionsSpecification

VpnStaticRoute
