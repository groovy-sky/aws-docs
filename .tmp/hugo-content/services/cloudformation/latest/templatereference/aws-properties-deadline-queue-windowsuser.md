This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Queue WindowsUser

The Windows user details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PasswordArn" : String,
  "User" : String
}

```

### YAML

```yaml

  PasswordArn: String
  User: String

```

## Properties

`PasswordArn`

The password ARN for the Windows user.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):secretsmanager:[a-z]{2}((-gov)|(-iso(b?)))?-[a-z]+-\d{1}:\d{12}:secret:[a-zA-Z0-9-/_+=.@]{1,2028}$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

The user.

_Required_: Yes

_Type_: String

_Pattern_: `^[^"'/\[\]:;|=,+*?<>\s]*$`

_Minimum_: `0`

_Maximum_: `111`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Deadline::QueueEnvironment

All content copied from https://docs.aws.amazon.com/.
