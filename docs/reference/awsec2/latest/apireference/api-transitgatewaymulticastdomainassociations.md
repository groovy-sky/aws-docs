# TransitGatewayMulticastDomainAssociations

Describes the multicast domain associations.

## Contents

**resourceId**

The ID of the resource.

Type: String

Required: No

**resourceOwnerId**

The ID of the AWS account that owns the resource.

Type: String

Required: No

**resourceType**

The type of resource, for example a VPC attachment.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

**Subnets.N**

The subnets associated with the multicast domain.

Type: Array of [SubnetAssociation](api-subnetassociation.md) objects

Required: No

**transitGatewayAttachmentId**

The ID of the transit gateway attachment.

Type: String

Required: No

**transitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaymulticastdomainassociations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaymulticastdomainassociations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaymulticastdomainassociations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayMulticastDomainAssociation

TransitGatewayMulticastDomainOptions
