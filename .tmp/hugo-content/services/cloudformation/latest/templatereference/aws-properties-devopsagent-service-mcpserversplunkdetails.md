This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service MCPServerSplunkDetails

Configuration details for registering a Splunk MCP server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationConfig" : MCPServerSplunkAuthorizationConfig,
  "Description" : String,
  "Endpoint" : String,
  "Name" : String
}

```

### YAML

```yaml

  AuthorizationConfig:
    MCPServerSplunkAuthorizationConfig
  Description: String
  Endpoint: String
  Name: String

```

## Properties

`AuthorizationConfig`

The authorization configuration for the Splunk MCP server.

_Required_: Yes

_Type_: [MCPServerSplunkAuthorizationConfig](aws-properties-devopsagent-service-mcpserversplunkauthorizationconfig.md)

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

MCPServerSplunkAuthorizationConfig

NewRelicApiKeyConfig

All content copied from https://docs.aws.amazon.com/.
