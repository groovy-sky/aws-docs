This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Instance Disk

`Disk` is a property of the [Hardware](../userguide/aws-properties-lightsail-instance-hardware.md) property. It describes a disk attached to an instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttachedTo" : String,
  "AttachmentState" : String,
  "DiskName" : String,
  "IOPS" : Integer,
  "IsSystemDisk" : Boolean,
  "Path" : String,
  "SizeInGb" : String
}

```

### YAML

```yaml

  AttachedTo: String
  AttachmentState: String
  DiskName: String
  IOPS: Integer
  IsSystemDisk: Boolean
  Path: String
  SizeInGb: String

```

## Properties

`AttachedTo`

The resources to which the disk is attached.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`AttachmentState`

(Deprecated) The attachment state of the disk.

###### Note

In releases prior to November 14, 2017, this parameter returned `attached`
for system disks in the API response. It is now deprecated, but still included in the
response. Use `isAttached` instead.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`DiskName`

The unique name of the disk.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][\w\-.]*[a-zA-Z0-9]$`

_Minimum_: `1`

_Maximum_: `254`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IOPS`

The input/output operations per second (IOPS) of the disk.

_Required_: No

_Type_: Integer

_Update requires_: Updates are not supported.

`IsSystemDisk`

A Boolean value indicating whether this disk is a system disk (has an operating system
loaded on it).

_Required_: No

_Type_: Boolean

_Update requires_: Updates are not supported.

`Path`

The disk path.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`SizeInGb`

The size of the disk in GB.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoSnapshotAddOn

Hardware

All content copied from https://docs.aws.amazon.com/.
