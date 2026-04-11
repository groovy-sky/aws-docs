This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection ApiKeyAuthParameters

The API key authorization parameters for the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKeyName" : String,
  "ApiKeyValue" : String
}

```

### YAML

```yaml

  ApiKeyName: String
  ApiKeyValue: String

```

## Properties

`ApiKeyName`

The name of the API key to use for authorization.

_Required_: Yes

_Type_: String

_Pattern_: `^[ \t]*[^\x00-\x1F\x7F]+([ \t]+[^\x00-\x1F\x7F]+)*[ \t]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApiKeyValue`

The value for the API key to use for authorization.

_Required_: Yes

_Type_: String

_Pattern_: `^[ \t]*[^\x00-\x1F\x7F]+([ \t]+[^\x00-\x1F\x7F]+)*[ \t]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Events::Connection

AuthParameters

All content copied from https://docs.aws.amazon.com/.
