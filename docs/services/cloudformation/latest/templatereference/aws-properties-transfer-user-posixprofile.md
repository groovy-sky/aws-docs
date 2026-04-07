This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::User PosixProfile

The full POSIX identity, including user ID ( `Uid`), group ID
( `Gid`), and any secondary groups IDs ( `SecondaryGids`), that
controls your users' access to your Amazon EFS file systems. The POSIX permissions that
are set on files and directories in your file system determine the level of access your
users get when transferring files into and out of your Amazon EFS file systems.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Gid" : Number,
  "SecondaryGids" : [ Number, ... ],
  "Uid" : Number
}

```

### YAML

```yaml

  Gid: Number
  SecondaryGids:
    - Number
  Uid: Number

```

## Properties

`Gid`

The POSIX group ID used for all EFS operations by this user.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `4294967295`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryGids`

The secondary POSIX group IDs used for all EFS operations by this user.

_Required_: No

_Type_: Array of Number

_Minimum_: `0 | 0`

_Maximum_: `16 | 4294967295`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uid`

The POSIX user ID used for all EFS operations by this user.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `4294967295`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HomeDirectoryMapEntry

Tag
