This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway GatewayProtocolConfiguration

The configuration for a gateway protocol. This structure defines how the gateway communicates with external services.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mcp" : MCPGatewayConfiguration
}

```

### YAML

```yaml

  Mcp:
    MCPGatewayConfiguration

```

## Properties

`Mcp`

The configuration for the Model Context Protocol (MCP). This protocol enables communication between Amazon Bedrock Agent and external tools.

_Required_: Yes

_Type_: [MCPGatewayConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-gateway-mcpgatewayconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GatewayPolicyEngineConfiguration

InterceptorConfiguration
