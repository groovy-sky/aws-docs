This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt ToolConfiguration

Configuration information for the tools that you pass to a model. For more information,
see [Tool use\
(function calling)](../../../bedrock/latest/userguide/tool-use.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ToolChoice" : ToolChoice,
  "Tools" : [ Tool, ... ]
}

```

### YAML

```yaml

  ToolChoice:
    ToolChoice
  Tools:
    - Tool

```

## Properties

`ToolChoice`

If supported by model, forces the model to request a tool.

_Required_: No

_Type_: [ToolChoice](aws-properties-bedrock-prompt-toolchoice.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tools`

An array of tools that you want to pass to a model.

_Required_: Yes

_Type_: Array of [Tool](aws-properties-bedrock-prompt-tool.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolChoice

ToolInputSchema

All content copied from https://docs.aws.amazon.com/.
