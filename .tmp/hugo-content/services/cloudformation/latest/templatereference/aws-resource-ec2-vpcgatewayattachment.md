This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCGatewayAttachment

Attaches an internet gateway, or a virtual private gateway to a VPC, enabling
connectivity between the internet and the VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCGatewayAttachment",
  "Properties" : {
      "InternetGatewayId" : String,
      "VpcId" : String,
      "VpnGatewayId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCGatewayAttachment
Properties:
  InternetGatewayId: String
  VpcId: String
  VpnGatewayId: String

```

## Properties

`InternetGatewayId`

The ID of the internet gateway.

You must specify either `InternetGatewayId` or `VpnGatewayId`, but
not both.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpnGatewayId`

The ID of the virtual private gateway.

You must specify either `InternetGatewayId` or `VpnGatewayId`, but
not both.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPC gateway attachment.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### VPN gateway attachment

To attach both an Internet gateway and a VPN gateway to a VPC, you must specify
two separate AWS::EC2::VPCGatewayAttachment resources:

#### JSON

```json

"AttachGateway" : {
   "Type" : "AWS::EC2::VPCGatewayAttachment",
   "Properties" : {
      "VpcId" : { "Ref" : "VPC" },
      "InternetGatewayId" : { "Ref" : "myInternetGateway" }
    }
},

"AttachVpnGateway" : {
   "Type" : "AWS::EC2::VPCGatewayAttachment",
   "Properties" : {
      "VpcId" : { "Ref" : "VPC" },
      "VpnGatewayId" : { "Ref" : "myVPNGateway" }
   }
}
```

#### YAML

```yaml

AttachGateway:
  Type: AWS::EC2::VPCGatewayAttachment
  Properties:
    VpcId:
       Ref: VPC
    InternetGatewayId:
       Ref: myInternetGateway
AttachVpnGateway:
  Type: AWS::EC2::VPCGatewayAttachment
  Properties:
    VpcId:
       Ref: VPC
    VpnGatewayId:
       Ref: myVPNGateway
```

## See also

- [AttachVpnGateway](../../../../reference/awsec2/latest/apireference/api-attachvpngateway.md) in the _Amazon EC2 API_
_Reference_

- [Internet gateways](../../../vpc/latest/userguide/vpc-internet-gateway.md)
in the _Amazon VPC User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VPCEndpointServicePermissions

AWS::EC2::VPCPeeringConnection

All content copied from https://docs.aws.amazon.com/.
