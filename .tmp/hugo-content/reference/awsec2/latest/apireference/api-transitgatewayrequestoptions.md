# TransitGatewayRequestOptions

Describes the options for a transit gateway.

## Contents

**AmazonSideAsn**

A private Autonomous System Number (ASN) for the Amazon side of a BGP session.
The range is 64512 to 65534 for 16-bit ASNs and 4200000000 to 4294967294 for 32-bit ASNs. The default is `64512`.

Type: Long

Required: No

**AutoAcceptSharedAttachments**

Enable or disable automatic acceptance of attachment requests. Disabled by default.

Type: String

Valid Values: `enable | disable`

Required: No

**DefaultRouteTableAssociation**

Enable or disable automatic association with the default association route table. Enabled by default.

Type: String

Valid Values: `enable | disable`

Required: No

**DefaultRouteTablePropagation**

Enable or disable automatic propagation of routes to the default propagation route table. Enabled by default.

Type: String

Valid Values: `enable | disable`

Required: No

**DnsSupport**

Enable or disable DNS support. Enabled by default.

Type: String

Valid Values: `enable | disable`

Required: No

**MulticastSupport**

Indicates whether multicast is enabled on the transit gateway

Type: String

Valid Values: `enable | disable`

Required: No

**SecurityGroupReferencingSupport**

Enables you to reference a security group across VPCs attached to a transit gateway to simplify security group management.

This option is disabled by default.

For more information about security group referencing, see [Security group referencing](../../../../services/vpc/latest/tgw/tgw-vpc-attachments.md#vpc-attachment-security) in the _AWS Transit Gateways Guide_.

Type: String

Valid Values: `enable | disable`

Required: No

**TransitGatewayCidrBlocks.N**

One or more IPv4 or IPv6 CIDR blocks for the transit gateway. Must be a size /24 CIDR block or larger for IPv4, or a size /64 CIDR block or larger for IPv6.

Type: Array of strings

Required: No

**VpnEcmpSupport**

Enable or disable Equal Cost Multipath Protocol support. Enabled by default.

Type: String

Valid Values: `enable | disable`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayrequestoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayrequestoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayrequestoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayPropagation

TransitGatewayRoute
