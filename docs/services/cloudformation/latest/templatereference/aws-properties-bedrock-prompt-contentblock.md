This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt ContentBlock

A block of content for a message that you pass to, or receive from, a model with the [Converse](../../../../reference/bedrock/latest/apireference/api-runtime-converse.md) or [ConverseStream](../../../../reference/bedrock/latest/apireference/api-runtime-conversestream.md) API operations.

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

CachePoint to include in the message.

_Required_: No

_Type_: [CachePointBlock](aws-properties-bedrock-prompt-cachepointblock.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

Text to include in the message.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChatPromptTemplateConfiguration

Message

All content copied from https://docs.aws.amazon.com/.
