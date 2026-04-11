# TransitGatewayVpcAttachmentOptions

Describes the VPC attachment options.

## Contents

**applianceModeSupport**

Indicates whether appliance mode support is enabled.

Type: String

Valid Values: `enable | disable`

Required: No

**dnsSupport**

Indicates whether DNS support is enabled.

Type: String

Valid Values: `enable | disable`

Required: No

**ipv6Support**

Indicates whether IPv6 support is disabled.

Type: String

Valid Values: `enable | disable`

Required: No

**securityGroupReferencingSupport**

Enables you to reference a security group across VPCs attached to a transit gateway to simplify security group management.

This option is enabled by default.

For more information about security group referencing, see [Security group referencing](../../../../services/vpc/latest/tgw/tgw-vpc-attachments.md#vpc-attachment-security) in the _AWS Transit Gateways Guide_.

Type: String

Valid Values: `enable | disable`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayvpcattachmentoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayvpcattachmentoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayvpcattachmentoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayVpcAttachment

TrunkInterfaceAssociation
