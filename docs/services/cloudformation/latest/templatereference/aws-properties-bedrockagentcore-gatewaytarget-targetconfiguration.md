This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget TargetConfiguration

The configuration for a gateway target. This structure defines how the gateway connects to and interacts with the target endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mcp" : McpTargetConfiguration
}

```

### YAML

```yaml

  Mcp:
    McpTargetConfiguration

```

## Properties

`Mcp`

The Model Context Protocol (MCP) configuration for the target. This configuration defines how the gateway uses MCP to communicate with the target.

_Required_: Yes

_Type_: [McpTargetConfiguration](aws-properties-bedrockagentcore-gatewaytarget-mcptargetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaDefinition

ToolDefinition

All content copied from https://docs.aws.amazon.com/.
