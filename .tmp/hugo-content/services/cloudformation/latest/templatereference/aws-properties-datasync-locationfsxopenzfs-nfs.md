This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxOpenZFS NFS

Represents the Network File System (NFS) protocol that AWS DataSync uses to
access your Amazon FSx for OpenZFS file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MountOptions" : MountOptions
}

```

### YAML

```yaml

  MountOptions:
    MountOptions

```

## Properties

`MountOptions`

Represents the mount options that are available for DataSync to access an
NFS location.

_Required_: Yes

_Type_: [MountOptions](aws-properties-datasync-locationfsxopenzfs-mountoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MountOptions

Protocol

All content copied from https://docs.aws.amazon.com/.
