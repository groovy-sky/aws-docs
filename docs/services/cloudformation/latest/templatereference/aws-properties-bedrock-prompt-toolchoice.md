This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt ToolChoice

Determines which tools the model should request in a call to `Converse` or
`ConverseStream`. For more information, see [Call a tool with the Converse API](../../../bedrock/latest/userguide/tool-use.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Any" : Json,
  "Auto" : Json,
  "Tool" : SpecificToolChoice
}

```

### YAML

```yaml

  Any: Json
  Auto: Json
  Tool:
    SpecificToolChoice

```

## Properties

`Any`

The model must request at least one tool (no text is generated).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Auto`

(Default). The Model automatically decides if a tool should be called or whether to generate text instead.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tool`

The Model must request the specified tool. Only supported by Anthropic Claude 3 and Amazon Nova models.

_Required_: No

_Type_: [SpecificToolChoice](aws-properties-bedrock-prompt-specifictoolchoice.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tool

ToolConfiguration

All content copied from https://docs.aws.amazon.com/.
