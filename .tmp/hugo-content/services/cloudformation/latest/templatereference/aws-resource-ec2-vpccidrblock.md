This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCCidrBlock

Associates a CIDR block with your VPC.

A VPC must have an associated IPv4 CIDR block. You can optionally associate additional IPv4 CIDR blocks with a VPC.
You can optionally associate an IPv6 CIDR block with a VPC. You can request an Amazon-provided
IPv6 CIDR block from Amazon's pool of IPv6 addresses, or an IPv6 CIDR block from an IPv6 address
pool that you provisioned through bring your own IP addresses (BYOIP).

For more information, see [VPC CIDR blocks](../../../vpc/latest/userguide/vpc-cidr-blocks.md)
in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCCidrBlock",
  "Properties" : {
      "AmazonProvidedIpv6CidrBlock" : Boolean,
      "CidrBlock" : String,
      "Ipv4IpamPoolId" : String,
      "Ipv4NetmaskLength" : Integer,
      "Ipv6CidrBlock" : String,
      "Ipv6CidrBlockNetworkBorderGroup" : String,
      "Ipv6IpamPoolId" : String,
      "Ipv6NetmaskLength" : Integer,
      "Ipv6Pool" : String,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCCidrBlock
Properties:
  AmazonProvidedIpv6CidrBlock: Boolean
  CidrBlock: String
  Ipv4IpamPoolId: String
  Ipv4NetmaskLength: Integer
  Ipv6CidrBlock: String
  Ipv6CidrBlockNetworkBorderGroup: String
  Ipv6IpamPoolId: String
  Ipv6NetmaskLength: Integer
  Ipv6Pool: String
  VpcId: String

```

## Properties

`AmazonProvidedIpv6CidrBlock`

Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You
cannot specify the range of IPv6 addresses or the size of the CIDR block.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CidrBlock`

An IPv4 CIDR block to associate with the VPC.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv4IpamPoolId`

Associate a CIDR allocated from an IPv4 IPAM pool to a VPC. For more information about Amazon VPC IP Address Manager (IPAM), see [What is IPAM?](../../../vpc/latest/ipam/what-is-it-ipam.md) in the _Amazon VPC IPAM User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv4NetmaskLength`

The netmask length of the IPv4 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see [What is IPAM?](../../../vpc/latest/ipam/what-is-it-ipam.md) in the _Amazon VPC IPAM User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6CidrBlock`

An IPv6 CIDR block from the IPv6 address pool. You must also specify `Ipv6Pool` in the request.

To let Amazon choose the IPv6 CIDR block for you, omit this parameter.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6CidrBlockNetworkBorderGroup`

The name of the location from which we advertise the IPV6 CIDR block. Use this parameter
to limit the CIDR block to this location.

You must set `AmazonProvidedIpv6CidrBlock` to `true` to use this parameter.

You can have one IPv6 CIDR block association per network border group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6IpamPoolId`

Associates a CIDR allocated from an IPv6 IPAM pool to a VPC. For more information about Amazon VPC IP Address Manager (IPAM), see [What is IPAM?](../../../vpc/latest/ipam/what-is-it-ipam.md) in the _Amazon VPC IPAM User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6NetmaskLength`

The netmask length of the IPv6 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see [What is IPAM?](../../../vpc/latest/ipam/what-is-it-ipam.md) in the _Amazon VPC IPAM User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6Pool`

The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The ID of the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the association ID for the VPC CIDR block.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`IpSource`

The source that allocated the IP address space. `byoip` or `amazon` indicates public IP address space allocated by Amazon or space that you have allocated with Bring your own IP (BYOIP). `none` indicates private space.

`Ipv6AddressAttribute`

Public IPv6 addresses are those advertised on the internet from AWS. Private IP addresses are not and cannot be advertised on the internet from AWS.

## Examples

- [Associate an Amazon-provided IPv6 CIDR block](#aws-resource-ec2-vpccidrblock--examples--Associate_an_Amazon-provided_IPv6_CIDR_block)

- [Associate an IPv4 CIDR block and an Amazon-provided IPv6 CIDR block](#aws-resource-ec2-vpccidrblock--examples--Associate_an_IPv4_CIDR_block_and_an_Amazon-provided_IPv6_CIDR_block)

### Associate an Amazon-provided IPv6 CIDR block

The following example associates an Amazon-provided IPv6 CIDR block (with a prefix
length of /56) with the TestVPCIpv6 VPC.

#### JSON

```json

"Ipv6VPCCidrBlock": {
   "Type": "AWS::EC2::VPCCidrBlock",
   "Properties": {
      "AmazonProvidedIpv6CidrBlock": true,
      "VpcId": { "Ref" : "TestVPCIpv6" }
   }
}
```

#### YAML

```yaml

Ipv6VPCCidrBlock:
   Type: AWS::EC2::VPCCidrBlock
   Properties:
      AmazonProvidedIpv6CidrBlock: true
      VpcId: !Ref TestVPCIpv6
```

### Associate an IPv4 CIDR block and an Amazon-provided IPv6 CIDR block

The following example associates an IPv4 CIDR block and an Amazon-provided IPv6
CIDR block with a VPC. It also outputs the list of IPv4 CIDR block association IDs
and IPv6 CIDR blocks that are associated with the VPC.

#### JSON

```json

{
    "Resources": {
        "VPC": {
            "Type": "AWS::EC2::VPC",
            "Properties": {
                "CidrBlock": "10.0.0.0/24"
            }
        },
        "VpcCidrBlock": {
            "Type": "AWS::EC2::VPCCidrBlock",
            "Properties": {
                "VpcId": {
                    "Ref": "VPC"
                },
                "CidrBlock": "192.0.0.0/24"
            }
        },
        "VpcCidrBlockIpv6": {
            "Type": "AWS::EC2::VPCCidrBlock",
            "Properties": {
                "VpcId": {
                    "Ref": "VPC"
                },
                "AmazonProvidedIpv6CidrBlock": true
            }
        }
    },
    "Outputs": {
        "VpcId": {
            "Value": {
                "Ref": "VPC"
            }
        },
        "PrimaryCidrBlock": {
            "Value": {
                "Fn::GetAtt": [
                    "VPC",
                    "CidrBlock"
                ]
            }
        },
        "Ipv6CidrBlock": {
            "Value": {
                "Fn::Select": [
                    0,
                    {
                        "Fn::GetAtt": [
                            "VPC",
                            "Ipv6CidrBlocks"
                        ]
                    }
                ]
            }
        },
        "CidrBlockAssociation": {
            "Value": {
                "Fn::Select": [
                    0,
                    {
                        "Fn::GetAtt": [
                            "VPC",
                            "CidrBlockAssociations"
                        ]
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  VPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 10.0.0.0/24
  VpcCidrBlock:
    Type: AWS::EC2::VPCCidrBlock
    Properties:
      VpcId: !Ref VPC
      CidrBlock: 192.0.0.0/24
  VpcCidrBlockIpv6:
    Type: AWS::EC2::VPCCidrBlock
    Properties:
      VpcId: !Ref VPC
      AmazonProvidedIpv6CidrBlock: true

Outputs:
  VpcId:
    Value: !Ref VPC
  PrimaryCidrBlock:
    Value: !GetAtt VPC.CidrBlock
  Ipv6CidrBlock:
    Value: !Select [ 0, !GetAtt VPC.Ipv6CidrBlocks ]
  CidrBlockAssociation:
    Value: !Select [ 0, !GetAtt VPC.CidrBlockAssociations ]
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VPCBlockPublicAccessOptions

AWS::EC2::VPCDHCPOptionsAssociation

All content copied from https://docs.aws.amazon.com/.
