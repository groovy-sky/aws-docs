# TransitGatewayRouteTable

Describes a transit gateway route table.

## Contents

**creationTime**

The creation time.

Type: Timestamp

Required: No

**defaultAssociationRouteTable**

Indicates whether this is the default association route table for the transit gateway.

Type: Boolean

Required: No

**defaultPropagationRouteTable**

Indicates whether this is the default propagation route table for the transit gateway.

Type: Boolean

Required: No

**state**

The state of the transit gateway route table.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**TagSet.N**

Any tags assigned to the route table.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

**transitGatewayRouteTableId**

The ID of the transit gateway route table.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayroutetable.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayroutetable.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayroutetable.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayRouteAttachment

TransitGatewayRouteTableAnnouncement
