This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service BearerTokenDetails

Bearer token authentication details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationHeader" : String,
  "TokenName" : String,
  "TokenValue" : String
}

```

### YAML

```yaml

  AuthorizationHeader: String
  TokenName: String
  TokenValue: String

```

## Properties

`AuthorizationHeader`

The HTTP header name used to send the bearer token. Defaults to `Authorization`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TokenName`

A friendly name for the bearer token.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\s-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TokenValue`

The bearer token value.

_Required_: Yes

_Type_: String

_Pattern_: `^[\S]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiKeyDetails

DynatraceAuthorizationConfig

All content copied from https://docs.aws.amazon.com/.
