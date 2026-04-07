# TransitGatewayPolicyTableAssociation

Describes a transit gateway policy table association.

## Contents

**resourceId**

The resource ID of the transit gateway attachment.

Type: String

Required: No

**resourceType**

The resource type for the transit gateway policy table association.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

**state**

The state of the transit gateway policy table association.

Type: String

Valid Values: `associating | associated | disassociating | disassociated`

Required: No

**transitGatewayAttachmentId**

The ID of the transit gateway attachment.

Type: String

Required: No

**transitGatewayPolicyTableId**

The ID of the transit gateway policy table.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaypolicytableassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaypolicytableassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaypolicytableassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayPolicyTable

TransitGatewayPolicyTableEntry
