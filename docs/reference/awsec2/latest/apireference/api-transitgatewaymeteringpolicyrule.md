# TransitGatewayMeteringPolicyRule

Describes the traffic matching criteria for a transit gateway metering policy rule.

## Contents

**destinationCidrBlock**

The destination CIDR block for the rule.

Type: String

Required: No

**destinationPortRange**

The destination port range for the rule.

Type: String

Required: No

**destinationTransitGatewayAttachmentId**

The ID of the destination transit gateway attachment.

Type: String

Required: No

**destinationTransitGatewayAttachmentType**

The type of the destination transit gateway attachment. Note that the `tgw-peering` resource type has been deprecated. To configure metering policies for Connect, use the transport attachment type.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

**protocol**

The protocol for the rule (1, 6, 17, etc.).

Type: String

Required: No

**sourceCidrBlock**

The source CIDR block for the rule.

Type: String

Required: No

**sourcePortRange**

The source port range for the rule.

Type: String

Required: No

**sourceTransitGatewayAttachmentId**

The ID of the source transit gateway attachment.

Type: String

Required: No

**sourceTransitGatewayAttachmentType**

The type of the source transit gateway attachment. Note that the `tgw-peering` resource type has been deprecated. To configure metering policies for Connect, use the transport attachment type.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaymeteringpolicyrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaymeteringpolicyrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaymeteringpolicyrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayMeteringPolicyEntry

TransitGatewayMulticastDeregisteredGroupMembers
