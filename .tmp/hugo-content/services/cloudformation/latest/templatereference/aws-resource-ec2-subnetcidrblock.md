This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SubnetCidrBlock

Associates a CIDR block with your subnet. You can associate a single IPv6 CIDR block
with your subnet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SubnetCidrBlock",
  "Properties" : {
      "Ipv6CidrBlock" : String,
      "Ipv6IpamPoolId" : String,
      "Ipv6NetmaskLength" : Integer,
      "SubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SubnetCidrBlock
Properties:
  Ipv6CidrBlock: String
  Ipv6IpamPoolId: String
  Ipv6NetmaskLength: Integer
  SubnetId: String

```

## Properties

`Ipv6CidrBlock`

The IPv6 network range for the subnet, in CIDR notation.

_Required_: Conditional

_Type_: String

_Maximum_: `42`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6IpamPoolId`

An IPv6 IPAM pool ID for the subnet.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6NetmaskLength`

An IPv6 netmask length for the subnet.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the association ID for the subnet’s IPv6 CIDR block.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the association.

`IpSource`

The source that allocated the IP address space. `byoip` or `amazon` indicates public IP address space allocated by Amazon or space that you have allocated with Bring your own IP (BYOIP). `none` indicates private space.

`Ipv6AddressAttribute`

Public IPv6 addresses are those advertised on the internet from AWS. Private IP addresses are not and cannot be advertised on the internet from AWS.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::SubnetNetworkAclAssociation

All content copied from https://docs.aws.amazon.com/.
