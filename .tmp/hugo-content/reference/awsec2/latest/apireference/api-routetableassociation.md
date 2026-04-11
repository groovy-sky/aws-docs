# RouteTableAssociation

Describes an association between a route table and a subnet or gateway.

## Contents

**associationState**

The state of the association.

Type: [RouteTableAssociationState](api-routetableassociationstate.md) object

Required: No

**gatewayId**

The ID of the internet gateway or virtual private gateway.

Type: String

Required: No

**main**

Indicates whether this is the main route table.

Type: Boolean

Required: No

**publicIpv4Pool**

The ID of a public IPv4 pool. A public IPv4 pool is a pool of IPv4 addresses that you've brought to AWS with BYOIP.

Type: String

Required: No

**routeTableAssociationId**

The ID of the association.

Type: String

Required: No

**routeTableId**

The ID of the route table.

Type: String

Required: No

**subnetId**

The ID of the subnet. A subnet ID is not returned for an implicit association.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/routetableassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/routetableassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/routetableassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RouteTable

RouteTableAssociationState
