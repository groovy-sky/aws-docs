# RouteTable

Describes a route table.

## Contents

**AssociationSet.N**

The associations between the route table and your subnets or gateways.

Type: Array of [RouteTableAssociation](api-routetableassociation.md) objects

Required: No

**ownerId**

The ID of the AWS account that owns the route table.

Type: String

Required: No

**PropagatingVgwSet.N**

Any virtual private gateway (VGW) propagating routes.

Type: Array of [PropagatingVgw](api-propagatingvgw.md) objects

Required: No

**RouteSet.N**

The routes in the route table.

Type: Array of [Route](api-route.md) objects

Required: No

**routeTableId**

The ID of the route table.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the route table.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcId**

The ID of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/routetable.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/routetable.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/routetable.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RouteServerRouteInstallationDetail

RouteTableAssociation
