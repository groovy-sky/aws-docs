# TransitGatewayRoute

Describes a route for a transit gateway route table.

## Contents

**destinationCidrBlock**

The CIDR block used for destination matches.

Type: String

Required: No

**prefixListId**

The ID of the prefix list used for destination matches.

Type: String

Required: No

**state**

The state of the route.

Type: String

Valid Values: `pending | active | blackhole | deleting | deleted`

Required: No

**TransitGatewayAttachments.N**

The attachments.

Type: Array of [TransitGatewayRouteAttachment](api-transitgatewayrouteattachment.md) objects

Required: No

**transitGatewayRouteTableAnnouncementId**

The ID of the transit gateway route table announcement.

Type: String

Required: No

**type**

The route type.

Type: String

Valid Values: `static | propagated`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayroute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayRequestOptions

TransitGatewayRouteAttachment
