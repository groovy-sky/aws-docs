This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Route

Specifies a route in a route table. For more information, see [Routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#route-table-routes)
in the _Amazon VPC User Guide_.

You must specify either a destination CIDR block or prefix list ID. You must also
specify exactly one of the resources as the target.

If you create a route that references a transit gateway in the same template where you
create the transit gateway, you must declare a dependency on the transit gateway
attachment. The route table cannot use the transit gateway until it has successfully
attached to the VPC. Add a [DependsOn\
Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) in the `AWS::EC2::Route` resource to explicitly declare a
dependency on the `AWS::EC2::TransitGatewayAttachment` resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::Route",
  "Properties" : {
      "CarrierGatewayId" : String,
      "CoreNetworkArn" : String,
      "DestinationCidrBlock" : String,
      "DestinationIpv6CidrBlock" : String,
      "DestinationPrefixListId" : String,
      "EgressOnlyInternetGatewayId" : String,
      "GatewayId" : String,
      "InstanceId" : String,
      "LocalGatewayId" : String,
      "NatGatewayId" : String,
      "NetworkInterfaceId" : String,
      "RouteTableId" : String,
      "TransitGatewayId" : String,
      "VpcEndpointId" : String,
      "VpcPeeringConnectionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::Route
Properties:
  CarrierGatewayId: String
  CoreNetworkArn: String
  DestinationCidrBlock: String
  DestinationIpv6CidrBlock: String
  DestinationPrefixListId: String
  EgressOnlyInternetGatewayId: String
  GatewayId: String
  InstanceId: String
  LocalGatewayId: String
  NatGatewayId: String
  NetworkInterfaceId: String
  RouteTableId: String
  TransitGatewayId: String
  VpcEndpointId: String
  VpcPeeringConnectionId: String

```

## Properties

`CarrierGatewayId`

The ID of the carrier gateway.

You can only use this option when the VPC contains a subnet which is associated with a Wavelength Zone.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CoreNetworkArn`

The Amazon Resource Name (ARN) of the core network.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationCidrBlock`

The IPv4 CIDR address block used for the destination match. Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify `100.68.0.18/18`, we modify it to `100.68.0.0/18`.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationIpv6CidrBlock`

The IPv6 CIDR block used for the destination match. Routing decisions are based on the most specific match.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPrefixListId`

The ID of a prefix list used for the destination match.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EgressOnlyInternetGatewayId`

\[IPv6 traffic only\] The ID of an egress-only internet gateway.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GatewayId`

The ID of an internet gateway or virtual private gateway attached to your
VPC.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceId`

The ID of a NAT instance in your VPC. The operation fails if you specify an instance ID unless exactly one network interface is attached.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalGatewayId`

The ID of the local gateway.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NatGatewayId`

\[IPv4 traffic only\] The ID of a NAT gateway.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceId`

The ID of a network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouteTableId`

The ID of the route table for the route.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayId`

The ID of a transit gateway.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcEndpointId`

The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcPeeringConnectionId`

The ID of a VPC peering connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the route.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CidrBlock`

The IPv4 CIDR block.

## Examples

- [Create a route to a gateway](#aws-resource-ec2-route--examples--Create_a_route_to_a_gateway)

- [Create a route to a carrier gateway](#aws-resource-ec2-route--examples--Create_a_route_to_a_carrier_gateway)

### Create a route to a gateway

The following example adds a route that is added to an internet gateway.

#### JSON

```json

"myRoute" : {
   "Type" : "AWS::EC2::Route",
   "DependsOn" : "GatewayToInternet",
   "Properties" : {
      "RouteTableId" : { "Ref" : "myRouteTable" },
      "DestinationCidrBlock" : "0.0.0.0/0",
      "GatewayId" : { "Ref" : "myInternetGateway" }
   }
}
```

#### YAML

```yaml

  myRoute:
    Type: AWS::EC2::Route
    DependsOn: GatewayToInternet
    Properties:
       RouteTableId:
         Ref: myRouteTable
       DestinationCidrBlock: 0.0.0.0/0
       GatewayId:
         Ref: myInternetGateway
```

### Create a route to a carrier gateway

The following example creates a route to a carrier gateway.

#### JSON

```json

"myCarrierRoute" : {
   "Type" : "AWS::EC2::Route",
   "DependsOn" : "GatewayToInternetAndCarrierNetwork",
   "Properties" : {
      "RouteTableId" : { "Ref" : "myRouteTable" },
      "DestinationCidrBlock" : "0.0.0.0/0",
      "GatewayId" : { "Ref" : "myCarrierGateway" }
   }
}
```

#### YAML

```yaml

 myCarrierRoute:
    Type: AWS::EC2::Route
    DependsOn: GatewayToInternetAndCarrierNetwork
    Properties:
       RouteTableId:
         Ref: myRouteTable
       DestinationCidrBlock: 0.0.0.0/0
       GatewayId:
         Ref: myCarrierGateway
```

## See also

- [CreateRoute](../../../../reference/awsec2/latest/apireference/api-createroute.md) in the _Amazon EC2 API_
_Reference_

- [Route\
tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html) in the _Amazon VPC User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::RouteServer
