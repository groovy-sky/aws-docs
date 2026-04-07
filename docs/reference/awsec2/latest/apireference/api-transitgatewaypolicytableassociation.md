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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayPolicyTableAssociation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayPolicyTableAssociation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayPolicyTableAssociation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayPolicyTable

TransitGatewayPolicyTableEntry
