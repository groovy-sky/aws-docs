This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::DiskSnapshot

Describes a block storage disk snapshot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::DiskSnapshot",
  "Properties" : {
      "DiskName" : String,
      "DiskSnapshotName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::DiskSnapshot
Properties:
  DiskName: String
  DiskSnapshotName: String
  Tags:
    - Tag

```

## Properties

`DiskName`

The unique name of the disk.

_Required_: Yes

_Type_: String

_Pattern_: `^\w[\w\-]*\w$`

_Minimum_: `2`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DiskSnapshotName`

The name of the disk snapshot ( `my-disk-snapshot`).

_Required_: Yes

_Type_: String

_Pattern_: `^\w[\w\-]*\w$`

_Minimum_: `2`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tag keys and optional values for the resource. For more information about tags in
Lightsail, see the [Amazon Lightsail Developer Guide](../../../lightsail/latest/userguide/amazon-lightsail-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-lightsail-disksnapshot-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The date when the disk snapshot was created.

`DiskSnapshotArn`

The Amazon Resource Name (ARN) of the disk snapshot.

`FromDiskName`

The unique name of the source disk from which the disk snapshot was created.

`IsFromAutoSnapshot`

A Boolean value indicating whether the snapshot was created from an automatic
snapshot.

`Progress`

The progress of the snapshot.

`ResourceType`

The Lightsail resource type ( `DiskSnapshot`).

`SizeInGb`

The size of the disk in GB.

`State`

The status of the disk snapshot operation.

`SupportCode`

The support code. Include this code in your email to support when you have questions about
an instance or another resource in Lightsail. This code enables our support team to look up
your Lightsail information more easily.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Location

All content copied from https://docs.aws.amazon.com/.
