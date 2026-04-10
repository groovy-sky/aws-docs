This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNGatewayRoutePropagation

Enables a virtual private gateway (VGW) to propagate routes to the specified route table
of a VPC.

If you reference a VPN gateway that is in the same template as your VPN gateway route
propagation, you must explicitly declare a dependency on the VPN gateway attachment. The
`AWS::EC2::VPNGatewayRoutePropagation` resource cannot use the VPN gateway
until it has successfully attached to the VPC. Add a [DependsOn\
Attribute](../userguide/aws-attribute-dependson.md) in the `AWS::EC2::VPNGatewayRoutePropagation` resource to
explicitly declare a dependency on the VPN gateway attachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPNGatewayRoutePropagation",
  "Properties" : {
      "RouteTableIds" : [ String, ... ],
      "VpnGatewayId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPNGatewayRoutePropagation
Properties:
  RouteTableIds:
    - String
  VpnGatewayId: String

```

## Properties

`RouteTableIds`

The ID of the route table. The routing table must be associated with the same VPC that
the virtual private gateway is attached to.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpnGatewayId`

The ID of the virtual private gateway that is attached to a VPC. The virtual private
gateway must be attached to the same VPC that the routing tables are associated with.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPN gateway.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the VPN gateway.

## Examples

### VPN gateway route propagation

The following example enables route propagation for the private route table named
PrivateRouteTable .

#### JSON

```json

"myVPNGatewayRouteProp" : {
   "Type" : "AWS::EC2::VPNGatewayRoutePropagation",
   "Properties" : {
      "RouteTableIds" : [{"Ref" : "PrivateRouteTable"}],
      "VpnGatewayId" : {"Ref" : "VPNGateway"}
   }
}
```

#### YAML

```yaml

myVPNGatewayRouteProp:
   Type: AWS::EC2::VPNGatewayRoutePropagation
   Properties:
       RouteTableIds:
        - !Ref PrivateRouteTable
       VpnGatewayId: !Ref VPNGateway
```

## See also

- [EnableVgwRoutePropagation](../../../../reference/awsec2/latest/apireference/api-enablevgwroutepropagation.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
