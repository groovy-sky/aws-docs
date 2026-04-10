This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service MCPServerDetails

Configuration details for registering a custom MCP server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationConfig" : MCPServerAuthorizationConfig,
  "Description" : String,
  "Endpoint" : String,
  "Name" : String
}

```

### YAML

```yaml

  AuthorizationConfig:
    MCPServerAuthorizationConfig
  Description: String
  Endpoint: String
  Name: String

```

## Properties

`AuthorizationConfig`

The authorization configuration for the MCP server.

_Required_: Yes

_Type_: [MCPServerAuthorizationConfig](aws-properties-devopsagent-service-mcpserverauthorizationconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the MCP server. Maximum 500 characters.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Endpoint`

The HTTPS endpoint URL of the MCP server.

_Required_: Yes

_Type_: String

_Pattern_: `^https://[a-zA-Z0-9.-]+(?::[0-9]+)?(?:/.*)?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the MCP server.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MCPServerAuthorizationConfig

MCPServerOAuthClientCredentialsConfig

All content copied from https://docs.aws.amazon.com/.
