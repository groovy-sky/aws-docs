This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMAllocation

In IPAM, an allocation is a CIDR assignment from an IPAM pool to another IPAM pool or to a resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::IPAMAllocation",
  "Properties" : {
      "Cidr" : String,
      "Description" : String,
      "IpamPoolId" : String,
      "NetmaskLength" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EC2::IPAMAllocation
Properties:
  Cidr: String
  Description: String
  IpamPoolId: String
  NetmaskLength: Integer

```

## Properties

`Cidr`

The CIDR you would like to allocate from the IPAM pool. Note the following:

- If there is no DefaultNetmaskLength allocation rule set on the pool, you must specify either the NetmaskLength or the CIDR.

- If the DefaultNetmaskLength allocation rule is set on the pool, you can specify either the NetmaskLength or the CIDR and the DefaultNetmaskLength allocation rule will be ignored.

Possible values: Any available IPv4 or IPv6 CIDR.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description for the allocation.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpamPoolId`

The ID of the IPAM pool from which you would like to allocate a CIDR.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetmaskLength`

The netmask length of the CIDR you would like to allocate from the IPAM pool. Note the following:

- If there is no DefaultNetmaskLength allocation rule set on the pool, you must specify either the NetmaskLength or the CIDR.

- If the DefaultNetmaskLength allocation rule is set on the pool, you can specify either the NetmaskLength or the CIDR and the DefaultNetmaskLength allocation rule will be ignored.

Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the pool ID, allocation ID, and CIDR joined with "\|" as a separator.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IpamPoolAllocationId`

The ID of an allocation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::IPAMPool

All content copied from https://docs.aws.amazon.com/.
