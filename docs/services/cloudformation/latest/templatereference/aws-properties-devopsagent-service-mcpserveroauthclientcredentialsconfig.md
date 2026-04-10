This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service MCPServerOAuthClientCredentialsConfig

The OAuth client credentials configuration for a custom MCP server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "ClientName" : String,
  "ClientSecret" : String,
  "ExchangeParameters" : Json,
  "ExchangeUrl" : String,
  "Scopes" : [ String, ... ]
}

```

### YAML

```yaml

  ClientId: String
  ClientName: String
  ClientSecret: String
  ExchangeParameters: Json
  ExchangeUrl: String
  Scopes:
    - String

```

## Properties

`ClientId`

The OAuth client ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientName`

A friendly name for the OAuth client.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientSecret`

The OAuth client secret.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExchangeParameters`

Additional parameters for the OAuth token exchange request.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExchangeUrl`

The OAuth token exchange endpoint URL.

_Required_: Yes

_Type_: String

_Pattern_: `^https://[a-zA-Z0-9.-]+(?::[0-9]+)?(?:/.*)?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scopes`

The OAuth scopes to request during authentication.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MCPServerDetails

MCPServerSplunkAuthorizationConfig

All content copied from https://docs.aws.amazon.com/.
