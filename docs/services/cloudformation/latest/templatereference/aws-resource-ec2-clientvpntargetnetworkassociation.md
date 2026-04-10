This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnTargetNetworkAssociation

Specifies a target network to associate with a Client VPN endpoint. A target network is
a subnet in a VPC. You can associate multiple subnets from the same VPC with a Client VPN
endpoint. You can associate only one subnet in each Availability Zone. We recommend that
you associate at least two subnets to provide Availability Zone redundancy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::ClientVpnTargetNetworkAssociation",
  "Properties" : {
      "ClientVpnEndpointId" : String,
      "SubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::ClientVpnTargetNetworkAssociation
Properties:
  ClientVpnEndpointId: String
  SubnetId: String

```

## Properties

`ClientVpnEndpointId`

The ID of the Client VPN endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet to associate with the Client VPN endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the association ID. For example:
`cvpn-assoc-1234567890abcdef0`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Associate a target subnet with a client VPN endpoint

The following example associates a target network with a client VPN
endpoint.

#### YAML

```yaml

myNetworkAssociation:
  Type: "AWS::EC2::ClientVpnTargetNetworkAssociation"
  Properties:
    ClientVpnEndpointId:
      Ref: myClientVpnEndpoint
    SubnetId:
      Ref: mySubnet
```

#### JSON

```json

"myNetworkAssociation": {
    "Type": "AWS::EC2::ClientVpnTargetNetworkAssociation",
    "Properties": {
        "ClientVpnEndpointId": {
            "Ref": "myClientVpnEndpoint"
        },
        "SubnetId": {
            "Ref": "mySubnet"
        }
    }
}
```

## See also

- [Getting Started\
with Client VPN](../../../vpn/latest/clientvpn-admin/cvpn-getting-started.md) in the _AWS Client VPN_
_Administrator Guide_

- [Target\
Networks](../../../vpn/latest/clientvpn-admin/cvpn-working-target.md) in the _AWS Client VPN Administrator_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::ClientVpnRoute

AWS::EC2::CustomerGateway

All content copied from https://docs.aws.amazon.com/.
