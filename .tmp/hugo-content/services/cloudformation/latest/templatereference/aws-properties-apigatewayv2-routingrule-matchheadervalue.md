This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::RoutingRule MatchHeaderValue

Represents a `MatchHeaderValue`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Header" : String,
  "ValueGlob" : String
}

```

### YAML

```yaml

  Header: String
  ValueGlob: String

```

## Properties

`Header`

The case insensitive header name to be matched. The header name must be less than 40 characters and the only allowed characters are `a-z`, `A-Z`, `0-9`, and the following special characters: ``*?-!#$%&'.^_`|~.``.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueGlob`

The case sensitive header glob value to be matched against entire header value. The header glob value must be less than 128 characters and the only allowed characters are `a-z`, `A-Z`, `0-9`, and the following special characters: ``*?-!#$%&'.^_`|~``. Wildcard matching is supported for header glob values but must be for `*prefix-match`, `suffix-match*`, or `*infix*-match`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MatchHeaders

AWS::ApiGatewayV2::Stage

All content copied from https://docs.aws.amazon.com/.
