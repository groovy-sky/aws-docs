This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment OpenZFSPosixFileSystemUser

The FSx for OpenZFS file system user that is used for authorizing all file access requests that are made using the S3 access point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Gid" : Number,
  "SecondaryGids" : [ FileSystemGID, ... ],
  "Uid" : Number
}

```

### YAML

```yaml

  Gid: Number
  SecondaryGids:
    - FileSystemGID
  Uid: Number

```

## Properties

`Gid`

The GID of the file system user.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `4294967295`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecondaryGids`

The list of secondary GIDs for the file system user.

_Required_: No

_Type_: Array of [FileSystemGID](aws-properties-fsx-s3accesspointattachment-filesystemgid.md)

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Uid`

The UID of the file system user.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `4294967295`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenZFSFileSystemIdentity

S3AccessPoint

All content copied from https://docs.aws.amazon.com/.
