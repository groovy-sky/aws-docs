This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service ApiKeyDetails

API key authentication details for an MCP server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKeyHeader" : String,
  "ApiKeyName" : String,
  "ApiKeyValue" : String
}

```

### YAML

```yaml

  ApiKeyHeader: String
  ApiKeyName: String
  ApiKeyValue: String

```

## Properties

`ApiKeyHeader`

The HTTP header name used to send the API key.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ApiKeyName`

A friendly name for the API key.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\s-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ApiKeyValue`

The API key value.

_Required_: Yes

_Type_: String

_Pattern_: `^[!-~]([ \t]*[!-~])*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdditionalServiceDetails

BearerTokenDetails

All content copied from https://docs.aws.amazon.com/.
