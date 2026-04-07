This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPC

Specifies a virtual private cloud (VPC).

A VPC must have an associated IPv4 CIDR block. You can specify an IPv4 CIDR block
or an IPAM-allocated IPv4 CIDR block. To associate an IPv6 CIDR block with the VPC,
see [AWS::EC2::VPCCidrBlock](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html).

For more information, see [Virtual private clouds (VPC)](https://docs.aws.amazon.com/vpc/latest/userguide/configure-your-vpc.html)
in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPC",
  "Properties" : {
      "CidrBlock" : String,
      "EnableDnsHostnames" : Boolean,
      "EnableDnsSupport" : Boolean,
      "InstanceTenancy" : String,
      "Ipv4IpamPoolId" : String,
      "Ipv4NetmaskLength" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPC
Properties:
  CidrBlock: String
  EnableDnsHostnames: Boolean
  EnableDnsSupport: Boolean
  InstanceTenancy: String
  Ipv4IpamPoolId: String
  Ipv4NetmaskLength: Integer
  Tags:
    - Tag

```

## Properties

`CidrBlock`

The IPv4 network range for the VPC, in CIDR notation. For example,
`10.0.0.0/16`. We modify the specified CIDR block to its canonical form; for example, if you specify `100.68.0.18/18`, we modify it to `100.68.0.0/18`.

You must specify either `CidrBlock` or `Ipv4IpamPoolId`.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableDnsHostnames`

Indicates whether the instances launched in the VPC get DNS hostnames. If enabled,
instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for
nondefault VPCs. For more information, see [DNS attributes in your\
VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support).

You can only enable DNS hostnames if you've enabled DNS support.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableDnsSupport`

Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to
the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP
address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon
provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not
enabled. Enabled by default. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceTenancy`

The allowed tenancy of instances launched into the VPC.

- `default`: An instance launched into the VPC runs on shared hardware
by default, unless you explicitly specify a different tenancy during instance
launch.

- `dedicated`: An instance launched into the VPC runs on dedicated
hardware by default, unless you explicitly specify a tenancy of `host`
during instance launch. You cannot specify a tenancy of `default` during
instance launch.

Updating `InstanceTenancy` requires no replacement only if you are updating
its value from `dedicated` to `default`. Updating
`InstanceTenancy` from `default` to `dedicated`
requires replacement.

_Required_: No

_Type_: String

_Allowed values_: `default | dedicated | host`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Ipv4IpamPoolId`

The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see
[What is IPAM?](https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the _Amazon VPC IPAM User Guide_.

You must specify either `CidrBlock` or `Ipv4IpamPoolId`.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv4NetmaskLength`

The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the _Amazon VPC IPAM User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the VPC.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-vpc-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPC.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CidrBlock`

The primary IPv4 CIDR block for the VPC. For example, 10.0.0.0/16.

`CidrBlockAssociations`

The association IDs of the IPv4 CIDR blocks for the VPC. For example,
\[ vpc-cidr-assoc-0280ab6b \].

`DefaultNetworkAcl`

The ID of the default network ACL for the VPC. For example, acl-814dafe3.

`DefaultSecurityGroup`

The ID of the default security group for the VPC. For example, sg-b178e0d3.

`Ipv6CidrBlocks`

The IPv6 CIDR blocks for the VPC. For example, \[ 2001:db8:1234:1a00::/56 \].

`VpcId`

The ID of the VPC.

## Examples

- [Create a VPC with an IPv4 CIDR block](#aws-resource-ec2-vpc--examples--Create_a_VPC_with_an_IPv4_CIDR_block)

- [Create a VPC with an IPv4 CIDR block and an IPv6 CIDR block](#aws-resource-ec2-vpc--examples--Create_a_VPC_with_an_IPv4_CIDR_block_and_an_IPv6_CIDR_block)

### Create a VPC with an IPv4 CIDR block

The following example specifies a VPC with an IPv4 address.

#### JSON

```json

{
   "Resources": {
       "myVPC" : {
           "Type" : "AWS::EC2::VPC",
           "Properties" : {
               "CidrBlock" : "10.0.0.0/16",
               "EnableDnsSupport" : "true",
               "EnableDnsHostnames" : "true",
               "Tags" : [
                   {"Key" : "stack", "Value" : "production"}
               ]
           }
       }
   }
}
```

#### YAML

```yaml

Resources:
  myVPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 10.0.0.0/16
      EnableDnsSupport: 'true'
      EnableDnsHostnames: 'true'
      Tags:
       - Key: stack
         Value: production
```

### Create a VPC with an IPv4 CIDR block and an IPv6 CIDR block

The following example specifies a VPC with an IPv4 address range and an IPv6 address range.

#### JSON

```json

{
   "Resources": {
       "myVPC" : {
           "Type" : "AWS::EC2::VPC",
           "Properties" : {
               "CidrBlock" : "10.0.0.0/16",
               "EnableDnsSupport" : "true",
               "EnableDnsHostnames" : "true",
               "Tags" : [
                   {"Key" : "stack", "Value" : "production"}
               ]
           }
       },
       "ipv6CidrBlock": {
           "Type": "AWS::EC2::VPCCidrBlock",
           "Properties": {
                "VpcId": {
                    "Ref": "myVPC"
                },
                "AmazonProvidedIpv6CidrBlock": true
            }
       }
   }
}
```

#### YAML

```yaml

Resources:
  myVPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 10.0.0.0/16
      EnableDnsSupport: 'true'
      EnableDnsHostnames: 'true'
      Tags:
       - Key: stack
         Value: production
  ipv6CidrBlock:
    Type: AWS::EC2::VPCCidrBlock
    Properties:
      VpcId: !Ref myVPC
      AmazonProvidedIpv6CidrBlock: true
```

## See also

- [CreateVpc](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpc.html) in the _Amazon EC2 API Reference_

- [VPC and\
subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the _Amazon VPC User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::VolumeAttachment

Tag
