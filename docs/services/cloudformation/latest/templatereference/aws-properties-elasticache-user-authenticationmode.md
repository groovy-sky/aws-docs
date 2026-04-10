This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::User AuthenticationMode

Specifies the authentication mode to use.

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

Specifies the passwords to use for authentication if `Type` is set to
`password`.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the authentication type. Possible options are IAM authentication, password
and no password.

_Required_: Yes

_Type_: String

_Allowed values_: `password | no-password-required | iam`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ElastiCache::User

Tag

All content copied from https://docs.aws.amazon.com/.
