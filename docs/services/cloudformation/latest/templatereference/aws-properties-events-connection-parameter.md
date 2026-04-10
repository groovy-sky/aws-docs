This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection Parameter

Any additional query string parameter for the connection. You can include up to 100 additional
query string parameters per request. Each additional parameter counts towards the event
payload size, which cannot exceed 64 KB.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsValueSecret" : Boolean,
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  IsValueSecret: Boolean
  Key: String
  Value: String

```

## Properties

`IsValueSecret`

Specifies whether the value is secret.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The key for a query string parameter.

_Required_: Yes

_Type_: String

_Pattern_: `[^\x00-\x1F\x7F]+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value associated with the key for the query string parameter.

_Required_: Yes

_Type_: String

_Pattern_: `[^\x00-\x09\x0B\x0C\x0E-\x1F\x7F]+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuthParameters

ResourceParameters

All content copied from https://docs.aws.amazon.com/.
