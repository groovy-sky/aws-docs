This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::AccessPoint RootDirectory

Specifies the root directory path and optional creation permissions for newly created directories.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreationPermissions" : CreationPermissions,
  "Path" : String
}

```

### YAML

```yaml

  CreationPermissions:
    CreationPermissions
  Path: String

```

## Properties

`CreationPermissions`

The POSIX IDs and permissions to apply to the access point's root directory. If the
root directory path specified does not exist, Amazon S3 Files creates the root directory
using these settings when a client connects to the access point. When specifying
`CreationPermissions`, you must provide values for all properties.

_Required_: No

_Type_: [CreationPermissions](aws-properties-s3files-accesspoint-creationpermissions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Path`

The path on the S3 Files file system to expose as the root directory to clients
using the access point. A path can have up to four subdirectories. If the specified
path does not exist, you are required to provide the `CreationPermissions`.

_Required_: No

_Type_: String

_Pattern_: ``^(\/|(\/(?!\.)+[^$#<>;;`|&?{}^*/\n]+){1,4})$``

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PosixUser

AWS::S3Files::FileSystem

All content copied from https://docs.aws.amazon.com/.
