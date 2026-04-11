This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIPrompt TextFullAIPromptEditTemplateConfiguration

The configuration for a prompt template that supports full textual prompt
configuration using a YAML prompt.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Text" : String
}

```

### YAML

```yaml

  Text: String

```

## Properties

`Text`

The YAML text for the AI Prompt template.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `200000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AIPromptTemplateConfiguration

AWS::Wisdom::AIPromptVersion

All content copied from https://docs.aws.amazon.com/.
