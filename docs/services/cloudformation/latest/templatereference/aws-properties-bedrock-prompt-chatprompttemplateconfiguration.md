This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt ChatPromptTemplateConfiguration

Contains configurations to use a prompt in a conversational format. For more information, see [Create a prompt using Prompt management](../../../bedrock/latest/userguide/prompt-management-create.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputVariables" : [ PromptInputVariable, ... ],
  "Messages" : [ Message, ... ],
  "System" : [ SystemContentBlock, ... ],
  "ToolConfiguration" : ToolConfiguration
}

```

### YAML

```yaml

  InputVariables:
    - PromptInputVariable
  Messages:
    - Message
  System:
    - SystemContentBlock
  ToolConfiguration:
    ToolConfiguration

```

## Properties

`InputVariables`

An array of the variables in the prompt template.

_Required_: No

_Type_: Array of [PromptInputVariable](aws-properties-bedrock-prompt-promptinputvariable.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Messages`

Contains messages in the chat for the prompt.

_Required_: Yes

_Type_: Array of [Message](aws-properties-bedrock-prompt-message.md)

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`System`

Contains system prompts to provide context to the model or to describe how it should behave.

_Required_: No

_Type_: Array of [SystemContentBlock](aws-properties-bedrock-prompt-systemcontentblock.md)

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolConfiguration`

Configuration information for the tools that the model can use when generating a response.

_Required_: No

_Type_: [ToolConfiguration](aws-properties-bedrock-prompt-toolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CachePointBlock

ContentBlock

All content copied from https://docs.aws.amazon.com/.
