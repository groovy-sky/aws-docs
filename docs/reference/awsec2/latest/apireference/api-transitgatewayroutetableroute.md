# TransitGatewayRouteTableRoute

Describes a route in a transit gateway route table.

## Contents

**attachmentId**

The ID of the route attachment.

Type: String

Required: No

**destinationCidr**

The CIDR block used for destination matches.

Type: String

Required: No

**prefixListId**

The ID of the prefix list.

Type: String

Required: No

**resourceId**

The ID of the resource for the route attachment.

Type: String

Required: No

**resourceType**

The resource type for the route attachment.

Type: String

Required: No

**routeOrigin**

The route origin. The following are the possible values:

- static

- propagated

Type: String

Required: No

**state**

The state of the route.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayroutetableroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayroutetableroute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayroutetableroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayRouteTablePropagation

TransitGatewayVpcAttachment
