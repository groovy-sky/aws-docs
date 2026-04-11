This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::UserProfile CustomPosixUserConfig

Details about the POSIX identity that is used for file system operations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Gid" : Integer,
  "Uid" : Integer
}

```

### YAML

```yaml

  Gid: Integer
  Uid: Integer

```

## Properties

`Gid`

The POSIX group ID.

_Required_: Yes

_Type_: Integer

_Minimum_: `1001`

_Maximum_: `4000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uid`

The POSIX user ID.

_Required_: Yes

_Type_: Integer

_Minimum_: `10000`

_Maximum_: `4000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomImage

DefaultEbsStorageSettings

All content copied from https://docs.aws.amazon.com/.
