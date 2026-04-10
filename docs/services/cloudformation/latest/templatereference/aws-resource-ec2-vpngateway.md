This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNGateway

Specifies a virtual private gateway. A virtual private gateway is the endpoint on the
VPC side of your VPN connection. You can create a virtual private gateway before creating
the VPC itself.

For more information, see [AWS Site-to-Site VPN](../../../vpn/latest/s2svpn/vpc-vpn.md) in the
_AWS Site-to-Site VPN User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPNGateway",
  "Properties" : {
      "AmazonSideAsn" : Integer,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPNGateway
Properties:
  AmazonSideAsn: Integer
  Tags:
    - Tag
  Type: String

```

## Properties

`AmazonSideAsn`

The private Autonomous System Number (ASN) for the Amazon side of a BGP
session.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Any tags assigned to the virtual private gateway.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-vpngateway-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of VPN connection the virtual private gateway supports.

_Required_: Yes

_Type_: String

_Allowed values_: `ipsec.1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPN gateway.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`VPNGatewayId`

The ID of the VPN gateway.

## Examples

### VPN gateway

The following example declares a VPN gateway that uses IPSec 1.

#### JSON

```json

"myVPNGateway" : {
   "Type" : "AWS::EC2::VPNGateway",
   "Properties" : {
      "Type" : "ipsec.1",
      "Tags" : [ { "Key" : "Use", "Value" : "Test" } ]
  }
}
```

#### YAML

```yaml

  myVPNGateway:
   Type: AWS::EC2::VPNGateway
   Properties:
      Type: ipsec.1
      Tags:
      - Key: Use
        Value: Test
```

## See also

- [CreateVPNGateway](../../../../reference/awsec2/latest/apireference/api-createvpngateway.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VPNConnectionRoute

Tag

All content copied from https://docs.aws.amazon.com/.
