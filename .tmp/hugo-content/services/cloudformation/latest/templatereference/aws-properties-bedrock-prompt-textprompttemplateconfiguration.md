This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt TextPromptTemplateConfiguration

Contains configurations for a text prompt template. To include a variable, enclose a word in double curly braces as in `{{variable}}`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CachePoint" : CachePointBlock,
  "InputVariables" : [ PromptInputVariable, ... ],
  "Text" : String,
  "TextS3Location" : TextS3Location
}

```

### YAML

```yaml

  CachePoint:
    CachePointBlock
  InputVariables:
    - PromptInputVariable
  Text: String
  TextS3Location:
    TextS3Location

```

## Properties

`CachePoint`

A cache checkpoint within a template configuration.

_Required_: No

_Type_: [CachePointBlock](aws-properties-bedrock-prompt-cachepointblock.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputVariables`

An array of the variables in the prompt template.

_Required_: No

_Type_: Array of [PromptInputVariable](aws-properties-bedrock-prompt-promptinputvariable.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

The message for the prompt.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextS3Location`

The Amazon S3location of the prompt text.

_Required_: No

_Type_: [TextS3Location](aws-properties-bedrock-prompt-texts3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SystemContentBlock

TextS3Location

All content copied from https://docs.aws.amazon.com/.
