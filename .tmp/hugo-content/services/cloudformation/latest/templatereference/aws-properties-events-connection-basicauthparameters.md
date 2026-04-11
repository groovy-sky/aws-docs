This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection BasicAuthParameters

The Basic authorization parameters for the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Password" : String,
  "Username" : String
}

```

### YAML

```yaml

  Password: String
  Username: String

```

## Properties

`Password`

The password associated with the user name to use for Basic authorization.

_Required_: Yes

_Type_: String

_Pattern_: `^[ \t]*[^\x00-\x1F\x7F]+([ \t]+[^\x00-\x1F\x7F]+)*[ \t]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The user name to use for Basic authorization.

_Required_: Yes

_Type_: String

_Pattern_: `^[ \t]*[^\x00-\x1F\x7F]+([ \t]+[^\x00-\x1F\x7F]+)*[ \t]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthParameters

ClientParameters

All content copied from https://docs.aws.amazon.com/.
