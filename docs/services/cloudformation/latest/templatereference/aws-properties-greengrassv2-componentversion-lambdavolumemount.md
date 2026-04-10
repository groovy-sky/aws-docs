This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion LambdaVolumeMount

Contains information about a volume that Linux processes in a container can access. When
you define a volume, the AWS IoT Greengrass Core software mounts the source files to the
destination inside the container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddGroupOwner" : Boolean,
  "DestinationPath" : String,
  "Permission" : String,
  "SourcePath" : String
}

```

### YAML

```yaml

  AddGroupOwner: Boolean
  DestinationPath: String
  Permission: String
  SourcePath: String

```

## Properties

`AddGroupOwner`

Whether or not to add the AWS IoT Greengrass user group as an owner of the
volume.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPath`

The path to the logical volume in the file system.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permission`

The permission to access the volume: read/only ( `ro`) or read/write
( `rw`).

Default: `ro`

_Required_: No

_Type_: String

_Allowed values_: `ro | rw`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePath`

The path to the physical volume in the file system.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaLinuxProcessParams

AWS::GreengrassV2::Deployment

All content copied from https://docs.aws.amazon.com/.
