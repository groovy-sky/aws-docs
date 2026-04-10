This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::AccessPoint PosixUser

Specifies the POSIX identity with uid, gid, and secondary group IDs for user enforcement.

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

_Pattern_: `^[0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecondaryGids`

Secondary POSIX group IDs used for all file system operations using this access point.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Uid`

The POSIX user ID used for all file system operations using this access point.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreationPermissions

RootDirectory

All content copied from https://docs.aws.amazon.com/.
