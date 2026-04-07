# TransitGatewayMulticastDomainAssociation

Describes the resources associated with the transit gateway multicast domain.

## Contents

**resourceId**

The ID of the resource.

Type: String

Required: No

**resourceOwnerId**

The ID of the AWS account that owns the transit gateway multicast domain association resource.

Type: String

Required: No

**resourceType**

The type of resource, for example a VPC attachment.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

**subnet**

The subnet associated with the transit gateway multicast domain.

Type: [SubnetAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SubnetAssociation.html) object

Required: No

**transitGatewayAttachmentId**

The ID of the transit gateway attachment.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayMulticastDomainAssociation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayMulticastDomainAssociation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayMulticastDomainAssociation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayMulticastDomain

TransitGatewayMulticastDomainAssociations
