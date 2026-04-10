This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association MCPServerNewRelicConfiguration

Configuration for New Relic MCP server integration. Defines the New Relic account ID and MCP server endpoint URL
required for the Agent Space to authenticate and query observability data from New Relic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "Endpoint" : String
}

```

### YAML

```yaml

  AccountId: String
  Endpoint: String

```

## Properties

`AccountId`

The New Relic account ID. Must be a numeric string of at least 6 characters.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+$`

_Minimum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The MCP server endpoint URL. Must be an HTTPS URL (for example, `https://mcp.newrelic.com/mcp/`).

_Required_: Yes

_Type_: String

_Pattern_: `^https://[a-zA-Z0-9.-]+(?::[0-9]+)?(?:/.*)?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MCPServerDatadogConfiguration

MCPServerSplunkConfiguration

All content copied from https://docs.aws.amazon.com/.
