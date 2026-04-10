This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::MountTarget

The `AWS::S3Files::MountTarget` resource specifies a mount target for
an Amazon S3 Files file system. Mount targets provide network access to the file system
from compute resources in a specific Availability Zone and VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Files::MountTarget",
  "Properties" : {
      "FileSystemId" : String,
      "IpAddressType" : String,
      "Ipv4Address" : String,
      "Ipv6Address" : String,
      "SecurityGroups" : [ String, ... ],
      "SubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Files::MountTarget
Properties:
  FileSystemId: String
  IpAddressType: String
  Ipv4Address: String
  Ipv6Address: String
  SecurityGroups:
    - String
  SubnetId: String

```

## Properties

`FileSystemId`

The ID of the S3 Files file system for which to create the mount target.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddressType`

The type of IP address used by the mount target. Valid values are
`IPV4_ONLY`, `IPV6_ONLY`, and `DUAL_STACK`.

_Required_: No

_Type_: String

_Allowed values_: `IPV4_ONLY | IPV6_ONLY | DUAL_STACK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv4Address`

The IPv4 address of the mount target.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`

_Minimum_: `7`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6Address`

The IPv6 address of the mount target.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `39`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroups`

Up to 100 VPC security group IDs that control network access to the mount target.

_Required_: No

_Type_: Array of String

_Minimum_: `11`

_Maximum_: `43 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet where the mount target is located.

_Required_: Yes

_Type_: String

_Pattern_: `^subnet-[0-9a-f]{8,40}$`

_Minimum_: `15`

_Maximum_: `47`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the mount target ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AvailabilityZoneId`

The ID of the Availability Zone in which the mount target is located.

`MountTargetId`

The ID of the mount target.

`NetworkInterfaceId`

The ID of the network interface that Amazon S3 Files created when it created the mount target.

`OwnerId`

The AWS account ID of the mount target owner.

`Status`

The current status of the mount target.

`StatusMessage`

Additional information about the mount target status.

`VpcId`

The ID of the VPC in which the mount target is located.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3Files::FileSystemPolicy

Next

All content copied from https://docs.aws.amazon.com/.
