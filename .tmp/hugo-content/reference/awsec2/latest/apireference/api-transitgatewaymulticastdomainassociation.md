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

Type: [SubnetAssociation](api-subnetassociation.md) object

Required: No

**transitGatewayAttachmentId**

The ID of the transit gateway attachment.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaymulticastdomainassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaymulticastdomainassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaymulticastdomainassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayMulticastDomain

TransitGatewayMulticastDomainAssociations
