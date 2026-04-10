This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment OntapFileSystemIdentity

Specifies the file system user identity that will be used for authorizing all file access requests that are made using the S3 access point. The identity can be either a UNIX user or a Windows user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "UnixUser" : OntapUnixFileSystemUser,
  "WindowsUser" : OntapWindowsFileSystemUser
}

```

### YAML

```yaml

  Type: String
  UnixUser:
    OntapUnixFileSystemUser
  WindowsUser:
    OntapWindowsFileSystemUser

```

## Properties

`Type`

Specifies the FSx for ONTAP user identity type. Valid values are `UNIX` and `WINDOWS`.

_Required_: Yes

_Type_: String

_Allowed values_: `UNIX | WINDOWS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UnixUser`

Specifies the UNIX user identity for file system operations.

_Required_: No

_Type_: [OntapUnixFileSystemUser](aws-properties-fsx-s3accesspointattachment-ontapunixfilesystemuser.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WindowsUser`

Specifies the Windows user identity for file system operations.

_Required_: No

_Type_: [OntapWindowsFileSystemUser](aws-properties-fsx-s3accesspointattachment-ontapwindowsfilesystemuser.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FileSystemGID

OntapUnixFileSystemUser

All content copied from https://docs.aws.amazon.com/.
