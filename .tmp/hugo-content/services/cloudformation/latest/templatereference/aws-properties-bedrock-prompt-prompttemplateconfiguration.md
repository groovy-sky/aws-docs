This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt PromptTemplateConfiguration

Contains the message for a prompt. For more information, see [Construct and store reusable prompts with Prompt management in Amazon Bedrock](../../../bedrock/latest/userguide/prompt-management.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Chat" : ChatPromptTemplateConfiguration,
  "Text" : TextPromptTemplateConfiguration
}

```

### YAML

```yaml

  Chat:
    ChatPromptTemplateConfiguration
  Text:
    TextPromptTemplateConfiguration

```

## Properties

`Chat`

Contains configurations to use the prompt in a conversational format.

_Required_: No

_Type_: [ChatPromptTemplateConfiguration](aws-properties-bedrock-prompt-chatprompttemplateconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

Contains configurations for the text in a message for a prompt.

_Required_: No

_Type_: [TextPromptTemplateConfiguration](aws-properties-bedrock-prompt-textprompttemplateconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptModelInferenceConfiguration

PromptVariant

All content copied from https://docs.aws.amazon.com/.
