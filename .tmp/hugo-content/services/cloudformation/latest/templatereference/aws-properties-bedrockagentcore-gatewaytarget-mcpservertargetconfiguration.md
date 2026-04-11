This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget McpServerTargetConfiguration

The configuration for an MCP server target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Endpoint" : String
}

```

### YAML

```yaml

  Endpoint: String

```

## Properties

`Endpoint`

The endpoint URL for the MCP server.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

McpLambdaTargetConfiguration

McpTargetConfiguration

All content copied from https://docs.aws.amazon.com/.
