This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment OpenZFSFileSystemIdentity

Specifies the file system user identity that will be used for authorizing all file access requests that are made using the S3 access point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PosixUser" : OpenZFSPosixFileSystemUser,
  "Type" : String
}

```

### YAML

```yaml

  PosixUser:
    OpenZFSPosixFileSystemUser
  Type: String

```

## Properties

`PosixUser`

Specifies the UID and GIDs of the file system POSIX user.

_Required_: Yes

_Type_: [OpenZFSPosixFileSystemUser](aws-properties-fsx-s3accesspointattachment-openzfsposixfilesystemuser.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

Specifies the FSx for OpenZFS user identity type, accepts only `POSIX`.

_Required_: Yes

_Type_: String

_Allowed values_: `POSIX`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OntapWindowsFileSystemUser

OpenZFSPosixFileSystemUser

All content copied from https://docs.aws.amazon.com/.
