This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxOpenZFS Protocol

Represents the protocol that AWS DataSync uses to access your Amazon FSx for
OpenZFS file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NFS" : NFS
}

```

### YAML

```yaml

  NFS:
    NFS

```

## Properties

`NFS`

Represents the Network File System (NFS) protocol that DataSync uses to
access your FSx for OpenZFS file system.

_Required_: No

_Type_: [NFS](aws-properties-datasync-locationfsxopenzfs-nfs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NFS

Tag

All content copied from https://docs.aws.amazon.com/.
