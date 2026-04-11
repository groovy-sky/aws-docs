# TransitGatewayRouteTableAnnouncement

Describes a transit gateway route table announcement.

## Contents

**announcementDirection**

The direction for the route table announcement.

Type: String

Valid Values: `outgoing | incoming`

Required: No

**coreNetworkId**

The ID of the core network for the transit gateway route table announcement.

Type: String

Required: No

**creationTime**

The timestamp when the transit gateway route table announcement was created.

Type: Timestamp

Required: No

**peerCoreNetworkId**

The ID of the core network ID for the peer.

Type: String

Required: No

**peeringAttachmentId**

The ID of the peering attachment.

Type: String

Required: No

**peerTransitGatewayId**

The ID of the peer transit gateway.

Type: String

Required: No

**state**

The state of the transit gateway announcement.

Type: String

Valid Values: `available | pending | failing | failed | deleting | deleted`

Required: No

**TagSet.N**

The key-value pairs associated with the route table announcement.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

**transitGatewayRouteTableAnnouncementId**

The ID of the transit gateway route table announcement.

Type: String

Required: No

**transitGatewayRouteTableId**

The ID of the transit gateway route table.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayroutetableannouncement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayroutetableannouncement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayroutetableannouncement.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayRouteTable

TransitGatewayRouteTableAssociation
