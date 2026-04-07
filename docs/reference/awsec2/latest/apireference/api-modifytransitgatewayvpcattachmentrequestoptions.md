# ModifyTransitGatewayVpcAttachmentRequestOptions

Describes the options for a VPC attachment.

## Contents

**ApplianceModeSupport**

Enable or disable support for appliance mode. If enabled, a traffic flow between a source and destination uses the same Availability Zone for the VPC attachment for the lifetime of that flow. The default is `disable`.

Type: String

Valid Values: `enable | disable`

Required: No

**DnsSupport**

Enable or disable DNS support. The default is `enable`.

Type: String

Valid Values: `enable | disable`

Required: No

**Ipv6Support**

Enable or disable IPv6 support. The default is `enable`.

Type: String

Valid Values: `enable | disable`

Required: No

**SecurityGroupReferencingSupport**

Enables you to reference a security group across VPCs attached to a transit gateway to simplify security group management.

This option is disabled by default.

For more information about security group referencing, see [Security group referencing](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-vpc-attachments.html#vpc-attachment-security) in the _AWS Transit Gateways Guide_.

Type: String

Valid Values: `enable | disable`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ModifyTransitGatewayVpcAttachmentRequestOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ModifyTransitGatewayVpcAttachmentRequestOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ModifyTransitGatewayVpcAttachmentRequestOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModifyTransitGatewayOptions

ModifyVerifiedAccessEndpointCidrOptions
