This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget ToolDefinition

A tool definition for a gateway target. This structure defines a tool that the target exposes through the Model Context Protocol.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "InputSchema" : SchemaDefinition,
  "Name" : String,
  "OutputSchema" : SchemaDefinition
}

```

### YAML

```yaml

  Description: String
  InputSchema:
    SchemaDefinition
  Name: String
  OutputSchema:
    SchemaDefinition

```

## Properties

`Description`

The description of the tool. This description provides information about the purpose and usage of the tool.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSchema`

The input schema for the tool. This schema defines the structure of the input that the tool accepts.

_Required_: Yes

_Type_: [SchemaDefinition](aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the tool. This name identifies the tool in the Model Context Protocol.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputSchema`

The output schema for the tool. This schema defines the structure of the output that the tool produces.

_Required_: No

_Type_: [SchemaDefinition](aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetConfiguration

ToolSchema

All content copied from https://docs.aws.amazon.com/.
