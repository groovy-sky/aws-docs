This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow TextPromptTemplateConfiguration

Contains configurations for a text prompt template. To include a variable, enclose a word in double curly braces as in `{{variable}}`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputVariables" : [ PromptInputVariable, ... ],
  "Text" : String
}

```

### YAML

```yaml

  InputVariables:
    - PromptInputVariable
  Text: String

```

## Properties

`InputVariables`

An array of the variables in the prompt template.

_Required_: No

_Type_: Array of [PromptInputVariable](aws-properties-bedrock-flow-promptinputvariable.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

The message for the prompt.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `200000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageFlowNodeServiceConfiguration

VectorSearchBedrockRerankingConfiguration

All content copied from https://docs.aws.amazon.com/.
