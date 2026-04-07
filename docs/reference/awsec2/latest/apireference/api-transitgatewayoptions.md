# TransitGatewayOptions

Describes the options for a transit gateway.

## Contents

**amazonSideAsn**

A private Autonomous System Number (ASN) for the Amazon side of a BGP session.
The range is 64512 to 65534 for 16-bit ASNs and 4200000000 to 4294967294 for 32-bit ASNs.

Type: Long

Required: No

**associationDefaultRouteTableId**

The ID of the default association route table.

Type: String

Required: No

**autoAcceptSharedAttachments**

Indicates whether attachment requests are automatically accepted.

Type: String

Valid Values: `enable | disable`

Required: No

**defaultRouteTableAssociation**

Indicates whether resource attachments are automatically associated with the default
association route table. Enabled by default. Either `defaultRouteTableAssociation` or `defaultRouteTablePropagation` must be set to `enable` for AWS Transit Gateway to create the default transit gateway route table.

Type: String

Valid Values: `enable | disable`

Required: No

**defaultRouteTablePropagation**

Indicates whether resource attachments automatically propagate routes to the default
propagation route table. Enabled by default. If `defaultRouteTablePropagation`
is set to `enable`,
AWS Transit Gateway creates the default transit gateway route
table.

Type: String

Valid Values: `enable | disable`

Required: No

**dnsSupport**

Indicates whether DNS support is enabled.

Type: String

Valid Values: `enable | disable`

Required: No

**encryptionSupport**

Defines if the Transit Gateway supports VPC Encryption Control.

Type: [EncryptionSupport](api-encryptionsupport.md) object

Required: No

**multicastSupport**

Indicates whether multicast is enabled on the transit gateway

Type: String

Valid Values: `enable | disable`

Required: No

**propagationDefaultRouteTableId**

The ID of the default propagation route table.

Type: String

Required: No

**securityGroupReferencingSupport**

Enables you to reference a security group across VPCs attached to a transit gateway to simplify security group management.

This option is disabled by default.

Type: String

Valid Values: `enable | disable`

Required: No

**TransitGatewayCidrBlocks.N**

The transit gateway CIDR blocks.

Type: Array of strings

Required: No

**vpnEcmpSupport**

Indicates whether Equal Cost Multipath Protocol support is enabled.

Type: String

Valid Values: `enable | disable`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayMulticastRegisteredGroupSources

TransitGatewayPeeringAttachment
