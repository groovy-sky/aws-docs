This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::AccessPoint CreationInfo

Required if the `RootDirectory` \> `Path` specified does not exist.
Specifies the POSIX IDs and permissions to apply to the access point's `RootDirectory` \> `Path`.
If the access point root directory does not exist, EFS creates it with these settings when a client connects to the access point.
When specifying `CreationInfo`, you must include values for all properties.

Amazon EFS creates a root directory only if you have provided the CreationInfo: OwnUid, OwnGID, and permissions for the directory.
If you do not provide this information, Amazon EFS does not create the root directory. If the root directory does not exist, attempts to mount
using the access point will fail.

###### Important

If you do not provide `CreationInfo` and the specified `RootDirectory` does not exist,
attempts to mount the file system using the access point will fail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OwnerGid" : String,
  "OwnerUid" : String,
  "Permissions" : String
}

```

### YAML

```yaml

  OwnerGid: String
  OwnerUid: String
  Permissions: String

```

## Properties

`OwnerGid`

Specifies the POSIX group ID to apply to the `RootDirectory`. Accepts values from 0 to 2^32 (4294967295).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OwnerUid`

Specifies the POSIX user ID to apply to the `RootDirectory`. Accepts values from 0 to 2^32 (4294967295).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permissions`

Specifies the POSIX permissions to apply to the `RootDirectory`, in the format of an octal number representing the file's mode bits.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-7]{3,4}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessPointTag

PosixUser

All content copied from https://docs.aws.amazon.com/.
