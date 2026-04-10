This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt SystemContentBlock

Contains configurations for instructions to provide the model for how to handle input.
To learn more, see [Using the Converse\
API](../../../bedrock/latest/userguide/conversation-inference-call.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CachePoint" : CachePointBlock,
  "Text" : String
}

```

### YAML

```yaml

  CachePoint:
    CachePointBlock
  Text: String

```

## Properties

`CachePoint`

CachePoint to include in the system prompt.

_Required_: No

_Type_: [CachePointBlock](aws-properties-bedrock-prompt-cachepointblock.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

A system prompt for the model.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpecificToolChoice

TextPromptTemplateConfiguration

All content copied from https://docs.aws.amazon.com/.
