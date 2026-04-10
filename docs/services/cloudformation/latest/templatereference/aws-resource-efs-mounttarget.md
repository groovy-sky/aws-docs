This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::MountTarget

The `AWS::EFS::MountTarget` resource is an Amazon EFS resource that creates a mount target for an EFS
file system. You can then mount the file system on Amazon EC2 instances or other resources by using the mount target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EFS::MountTarget",
  "Properties" : {
      "FileSystemId" : String,
      "IpAddress" : String,
      "IpAddressType" : String,
      "Ipv6Address" : String,
      "SecurityGroups" : [ String, ... ],
      "SubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EFS::MountTarget
Properties:
  FileSystemId: String
  IpAddress: String
  IpAddressType: String
  Ipv6Address: String
  SecurityGroups:
    - String
  SubnetId: String

```

## Properties

`FileSystemId`

The ID of the file system for which to create the mount target.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws[-a-z]*:elasticfilesystem:[0-9a-z-:]+:file-system/fs-[0-9a-f]{8,40}|fs-[0-9a-f]{8,40})$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddress`

If the `IpAddressType` for the mount target is IPv4 ( `IPV4_ONLY` or
`DUAL_STACK`), then specify the IPv4 address to use. If you do not specify an
`IpAddress`, then Amazon EFS selects an unused IP address from the subnet
specified for `SubnetId`.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`

_Minimum_: `7`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddressType`

The IP address type for the mount target. The possible values are `IPV4_ONLY`
(only IPv4 addresses), `IPV6_ONLY` (only IPv6 addresses), and
`DUAL_STACK` (dual-stack, both IPv4 and IPv6 addresses). If you don’t specify an
`IpAddressType`, then `IPV4_ONLY` is used.

###### Note

The `IPAddressType` must match the IP type of the subnet. Additionally, the
`IPAddressType` parameter overrides the value set as the default IP address for
the subnet in the VPC. For example, if the `IPAddressType` is
`IPV4_ONLY` and `AssignIpv6AddressOnCreation` is `true`,
then IPv4 is used for the mount target. For more information, see [Modify the IP\
addressing attributes of your subnet](../../../vpc/latest/userguide/subnet-public-ip.md).

_Required_: No

_Type_: String

_Allowed values_: `IPV4_ONLY | IPV6_ONLY | DUAL_STACK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6Address`

If the `IPAddressType` for the mount target is IPv6 ( `IPV6_ONLY` or
`DUAL_STACK`), then specify the IPv6 address to use. If you do not specify an
`Ipv6Address`, then Amazon EFS selects an unused IP address from the
subnet specified for `SubnetId`.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `39`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroups`

VPC security group IDs, of the form `sg-xxxxxxxx`. These must be for the same
VPC as the subnet specified. The maximum number of security groups depends on account quota.
For more information, see [Amazon VPC Quotas](../../../vpc/latest/userguide/amazon-vpc-limits.md) in the
_Amazon VPC User Guide_ (see the **Security**
**Groups** table). If you don't specify a security group, then Amazon EFS
uses the default security group for the subnet's VPC.

_Required_: Yes

_Type_: Array of String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet to add the mount target in. For One Zone file systems, use the subnet
that is associated with the file system's Availability Zone. The subnet type must be the same type
as the `IpAddressType`.

_Required_: Yes

_Type_: String

_Pattern_: `^subnet-[0-9a-f]{8,40}$`

_Minimum_: `15`

_Maximum_: `47`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the MountTarget ID. For example:

`{"Ref":"logical_mount_target_id"}` returns

`fsmt-0123456789abcdef8`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the Amazon EFS file system that the mount target provides access to.

Example: `fs-0123456789111222a`

`IpAddress`

The IPv4 address of the mount target.

Example: 192.0.2.0

## Examples

### Declare a Mount Target for an EFS File System

The following example declares a mount target that is associated with a file system,
subnet, and security group, which are all declared in the same template. EC2 instances that
are in the same Availability Zone (AZ) as the mount target can use the mount target to connect to the
associated file system. For information about mounting file systems on EC2 instances, see
[Mounting File Systems](../../../efs/latest/ug/mounting-fs.md) in the
_EFS User Guide_.

#### JSON

```json

"MountTarget": {
  "Type": "AWS::EFS::MountTarget",
  "Properties": {
    "FileSystemId": { "Ref": "FileSystem" },
    "SubnetId": { "Ref": "Subnet" },
    "SecurityGroups": [ { "Ref": "MountTargetSecurityGroup" } ]
  }
}
```

#### YAML

```yaml

MountTarget:
  Type: AWS::EFS::MountTarget
  Properties:
    FileSystemId:
      Ref: "FileSystem"
    SubnetId:
      Ref: "Subnet"
    SecurityGroups:
      - Ref: "MountTargetSecurityGroup"
```

## See also

- [Amazon EFS: How it works](../../../efs/latest/ug/how-it-works.md)

- [Creating mount targets](../../../efs/latest/ug/accessing-fs.md)

- [Walkthrough: Mounting a file system on-premises](../../../efs/latest/ug/efs-onpremises.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationDestination

Next

All content copied from https://docs.aws.amazon.com/.
