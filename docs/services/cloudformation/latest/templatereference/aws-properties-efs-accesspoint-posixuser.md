This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::AccessPoint PosixUser

The full POSIX identity, including the user ID, group ID, and any secondary group IDs, on the access point that is used for all file system operations performed by
NFS clients using the access point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Gid" : String,
  "SecondaryGids" : [ String, ... ],
  "Uid" : String
}

```

### YAML

```yaml

  Gid: String
  SecondaryGids:
    - String
  Uid: String

```

## Properties

`Gid`

The POSIX group ID used for all file system operations using this access point.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecondaryGids`

Secondary POSIX group IDs used for all file system operations using this access point.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Uid`

The POSIX user ID used for all file system operations using this access point.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreationInfo

RootDirectory

All content copied from https://docs.aws.amazon.com/.
