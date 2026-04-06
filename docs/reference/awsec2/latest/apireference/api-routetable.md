# RouteTable

Describes a route table.

## Contents

**AssociationSet.N**

The associations between the route table and your subnets or gateways.

Type: Array of [RouteTableAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RouteTableAssociation.html) objects

Required: No

**ownerId**

The ID of the AWS account that owns the route table.

Type: String

Required: No

**PropagatingVgwSet.N**

Any virtual private gateway (VGW) propagating routes.

Type: Array of [PropagatingVgw](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PropagatingVgw.html) objects

Required: No

**RouteSet.N**

The routes in the route table.

Type: Array of [Route](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Route.html) objects

Required: No

**routeTableId**

The ID of the route table.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the route table.

Type: Array of [Tag](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Tag.html) objects

Required: No

**vpcId**

The ID of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/RouteTable)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/RouteTable)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/RouteTable)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RouteServerRouteInstallationDetail

RouteTableAssociation
