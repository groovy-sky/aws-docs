This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Disk

The `AWS::Lightsail::Disk` resource specifies a disk that can be attached to
an Amazon Lightsail instance that is in the same AWS Region
and Availability Zone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Disk",
  "Properties" : {
      "AddOns" : [ AddOn, ... ],
      "AvailabilityZone" : String,
      "DiskName" : String,
      "Location" : Location,
      "SizeInGb" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Disk
Properties:
  AddOns:
    - AddOn
  AvailabilityZone: String
  DiskName: String
  Location:
    Location
  SizeInGb: Integer
  Tags:
    - Tag

```

## Properties

`AddOns`

An array of add-ons for the disk.

###### Note

If the disk has an add-on enabled when performing a delete disk request, the add-on
is automatically disabled before the disk is deleted.

_Required_: No

_Type_: Array of [AddOn](aws-properties-lightsail-disk-addon.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The AWS Region and Availability Zone location for the disk (for
example, `us-east-1a`).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`DiskName`

The name of the disk.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][\w\-.]*[a-zA-Z0-9]$`

_Minimum_: `1`

_Maximum_: `254`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Location`

The AWS Region and Availability Zone where the disk is located.

_Required_: No

_Type_: [Location](aws-properties-lightsail-disk-location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeInGb`

The size of the disk in GB.

_Required_: Yes

_Type_: Integer

_Update requires_: Updates are not supported.

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md)
in the _AWS CloudFormation User Guide_.

###### Note

The `Value` of `Tags` is optional for Lightsail resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-lightsail-disk-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttachedTo`

The resources to which the disk is attached.

`AttachmentState`

(Deprecated) The attachment state of the disk.

###### Note

In releases prior to November 14, 2017, this parameter returned `attached`
for system disks in the API response. It is now deprecated, but still included in the
response. Use `isAttached` instead.

`DiskArn`

The Amazon Resource Name (ARN) of the disk.

`Iops`

The input/output operations per second (IOPS) of the disk.

`IsAttached`

A Boolean value indicating whether the disk is attached.

`Location.AvailabilityZone`

The Availability Zone where the disk is located.

`Location.RegionName`

The AWS Region where the disk is located.

`Path`

The disk path.

`ResourceType`

The resource type of the disk (for example, `Disk`).

`State`

The state of the disk (for example, `in-use`).

`SupportCode`

The support code of the disk.

Include this code in your email to support when you have questions about a disk or
another resource in Lightsail. This code helps our support team to look up
your Lightsail information.

## Remarks

_Availability Zone_

You can specify an Availability Zone when you perform a create disk request. If you
don’t specify one, the disk is created in the same Availability Zone as the last Lightsail resource you created.

_Disk state_

Disks can be deleted only when they're in an `available` state. If the
disk is in an `attached` state when performing a delete disk request, the
service will wait to check if the disk state changes to `available`. The
delete disk request times out if the disk state doesn't change to `available`
within 15 minutes.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AddOn

All content copied from https://docs.aws.amazon.com/.
