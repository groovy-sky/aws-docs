This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMPrefixListResolver

An IPAM prefix list resolver is a component that manages the synchronization between IPAM's CIDR selection rules and customer-managed prefix lists. It automates connectivity configurations by selecting CIDRs from IPAM's database based on your business logic and synchronizing them with prefix lists used in resources such as VPC route tables and security groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::IPAMPrefixListResolver",
  "Properties" : {
      "AddressFamily" : String,
      "Description" : String,
      "IpamId" : String,
      "Rules" : [ IpamPrefixListResolverRule, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::IPAMPrefixListResolver
Properties:
  AddressFamily: String
  Description: String
  IpamId: String
  Rules:
    - IpamPrefixListResolverRule
  Tags:
    - Tag

```

## Properties

`AddressFamily`

The address family (IPv4 or IPv6) for the IPAM prefix list resolver.

_Required_: Yes

_Type_: String

_Allowed values_: `ipv4 | ipv6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the IPAM prefix list resolver.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamId`

The ID of the IPAM associated with this resolver.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rules`

CIDR selection rules for this resolver.

_Required_: No

_Type_: Array of [IpamPrefixListResolverRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags assigned to the IPAM prefix list resolver.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipamprefixlistresolver-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the prefix list resolver ID. For example: `ipam-plr-111122223333`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IpamArn`

The Amazon Resource Name (ARN) of the IPAM associated with this resolver.

`IpamPrefixListResolverArn`

The Amazon Resource Name (ARN) of the IPAM prefix list resolver.

`IpamPrefixListResolverId`

The ID of the IPAM prefix list resolver.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::IPAMPoolCidr

IpamPrefixListResolverRule
