This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MemoryDB::User AuthenticationMode

Denotes the user's authentication properties, such as whether it requires a password to authenticate. Used in output responses.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Passwords" : [ String, ... ],
  "Type" : String
}

```

### YAML

```yaml

  Passwords:
    - String
  Type: String

```

## Properties

`Passwords`

The password(s) used for authentication

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Indicates whether the user requires a password to authenticate. All newly-created users require a password.

_Required_: No

_Type_: String

_Allowed values_: `password | iam`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MemoryDB::User

Tag
