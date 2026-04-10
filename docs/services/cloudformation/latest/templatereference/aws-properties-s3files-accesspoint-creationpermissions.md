This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::AccessPoint CreationPermissions

Specifies the POSIX IDs and permissions to apply when creating the root directory
for an access point.

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

The POSIX group ID to apply to the root directory. Accepts values from 0 to
4294967295.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OwnerUid`

The POSIX user ID to apply to the root directory. Accepts values from 0 to
4294967295.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permissions`

The POSIX permissions to apply to the root directory, in the format of an octal
number representing the file's mode bits.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-7]{3,4}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessPointTag

PosixUser

All content copied from https://docs.aws.amazon.com/.
