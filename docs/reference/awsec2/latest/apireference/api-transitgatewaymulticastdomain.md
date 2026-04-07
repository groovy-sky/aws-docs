# TransitGatewayMulticastDomain

Describes the transit gateway multicast domain.

## Contents

**creationTime**

The time the transit gateway multicast domain was created.

Type: Timestamp

Required: No

**options**

The options for the transit gateway multicast domain.

Type: [TransitGatewayMulticastDomainOptions](api-transitgatewaymulticastdomainoptions.md) object

Required: No

**ownerId**

The ID of the AWS account that owns the transit gateway multicast domain.

Type: String

Required: No

**state**

The state of the transit gateway multicast domain.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**TagSet.N**

The tags for the transit gateway multicast domain.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

**transitGatewayMulticastDomainArn**

The Amazon Resource Name (ARN) of the transit gateway multicast domain.

Type: String

Required: No

**transitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaymulticastdomain.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaymulticastdomain.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaymulticastdomain.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayMulticastDeregisteredGroupSources

TransitGatewayMulticastDomainAssociation
