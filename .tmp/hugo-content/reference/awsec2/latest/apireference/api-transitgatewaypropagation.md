# TransitGatewayPropagation

Describes route propagation.

## Contents

**resourceId**

The ID of the resource.

Type: String

Required: No

**resourceType**

The resource type. Note that the `tgw-peering` resource type has been deprecated.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

**state**

The state.

Type: String

Valid Values: `enabling | enabled | disabling | disabled`

Required: No

**transitGatewayAttachmentId**

The ID of the attachment.

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaypropagation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaypropagation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaypropagation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayPrefixListReference

TransitGatewayRequestOptions
