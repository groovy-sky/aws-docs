This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpaces::WorkspacesPool Capacity

Describes the user capacity for the pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DesiredUserSessions" : Integer
}

```

### YAML

```yaml

  DesiredUserSessions: Integer

```

## Properties

`DesiredUserSessions`

The desired number of user sessions for the WorkSpaces in the pool.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationSettings

TimeoutSettings

All content copied from https://docs.aws.amazon.com/.
