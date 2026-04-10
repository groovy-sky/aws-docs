This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnRoute

Specifies a network route to add to a Client VPN endpoint. Each Client VPN endpoint has
a route table that describes the available destination network routes. Each route in the
route table specifies the path for traffic to specific resources or networks.

A target network association must be created before you can specify a route. If you're
setting up all the components of a Client VPN endpoint at the same time, you must use the
[DependsOn\
Attribute](../userguide/aws-attribute-dependson.md) to declare a dependency on the
`AWS::EC2::ClientVpnTargetNetworkAssociation` resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::ClientVpnRoute",
  "Properties" : {
      "ClientVpnEndpointId" : String,
      "Description" : String,
      "DestinationCidrBlock" : String,
      "TargetVpcSubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::ClientVpnRoute
Properties:
  ClientVpnEndpointId: String
  Description: String
  DestinationCidrBlock: String
  TargetVpcSubnetId: String

```

## Properties

`ClientVpnEndpointId`

The ID of the Client VPN endpoint to which to add the route.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A brief description of the route.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationCidrBlock`

The IPv4 address range, in CIDR notation, of the route destination. For example:

- To add a route for Internet access, enter `0.0.0.0/0`

- To add a route for a peered VPC, enter the peered VPC's IPv4 CIDR range

- To add a route for an on-premises network, enter the AWS Site-to-Site VPN connection's IPv4 CIDR range

- To add a route for the local network, enter the client CIDR range

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetVpcSubnetId`

The ID of the subnet through which you want to route traffic. The specified subnet must be
an existing target network of the Client VPN endpoint.

Alternatively, if you're adding a route for the local network, specify `local`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Add a route to a client VPN endpoint

The following example adds a route for internet access to a client VPN
endpoint.

#### YAML

```yaml

myRoute:
  Type: "AWS::EC2::ClientVpnRoute"
  Properties:
    ClientVpnEndpointId:
      Ref: myClientVpnEndpoint
    TargetVpcSubnetId:
      Ref: mySubnet
    DestinationCidrBlock: "0.0.0.0/0"
    Description: "myRoute"
```

#### JSON

```json

"myRoute": {
    "Type": "AWS::EC2::ClientVpnRoute",
    "Properties": {
        "ClientVpnEndpointId": {
            "Ref": "myClientVpnEndpoint"
        },
        "TargetVpcSubnetId": {
            "Ref": "mySubnet"
        },
        "DestinationCidrBlock": "0.0.0.0/0",
        "Description": "myRoute"
    }
}
```

## See also

- [Getting Started with\
Client VPN](../../../vpn/latest/clientvpn-admin/cvpn-getting-started.md) in the _AWS Client VPN Administrator_
_Guide_

- [Routes](../../../vpn/latest/clientvpn-admin/cvpn-working-routes.md) in the
_AWS Client VPN Administrator Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagSpecification

AWS::EC2::ClientVpnTargetNetworkAssociation

All content copied from https://docs.aws.amazon.com/.
