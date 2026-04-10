This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection ClientParameters

The OAuth authorization parameters to use for the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientID" : String,
  "ClientSecret" : String
}

```

### YAML

```yaml

  ClientID: String
  ClientSecret: String

```

## Properties

`ClientID`

The client ID to use for OAuth authorization.

_Required_: Yes

_Type_: String

_Pattern_: `^[ \t]*[^\x00-\x1F\x7F]+([ \t]+[^\x00-\x1F\x7F]+)*[ \t]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret assciated with the client ID to use for OAuth authorization.

_Required_: Yes

_Type_: String

_Pattern_: `^[ \t]*[^\x00-\x1F\x7F]+([ \t]+[^\x00-\x1F\x7F]+)*[ \t]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicAuthParameters

ConnectionHttpParameters

All content copied from https://docs.aws.amazon.com/.
