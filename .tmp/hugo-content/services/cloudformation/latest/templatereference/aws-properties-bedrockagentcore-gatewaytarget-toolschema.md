This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget ToolSchema

A tool schema for a gateway target. This structure defines the schema for a tool that the target exposes through the Model Context Protocol.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InlinePayload" : [ ToolDefinition, ... ],
  "S3" : S3Configuration
}

```

### YAML

```yaml

  InlinePayload:
    - ToolDefinition
  S3:
    S3Configuration

```

## Properties

`InlinePayload`

The inline payload of the tool schema. This payload contains the schema definition directly in the request.

_Required_: No

_Type_: Array of [ToolDefinition](aws-properties-bedrockagentcore-gatewaytarget-tooldefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

The Amazon S3 location of the tool schema. This location contains the schema definition file.

_Required_: No

_Type_: [S3Configuration](aws-properties-bedrockagentcore-gatewaytarget-s3configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolDefinition

AWS::BedrockAgentCore::Memory

All content copied from https://docs.aws.amazon.com/.
