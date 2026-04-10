This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow PromptModelInferenceConfiguration

Contains inference configurations related to model inference for a prompt. For more information, see [Inference parameters](../../../bedrock/latest/userguide/inference-parameters.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxTokens" : Number,
  "StopSequences" : [ String, ... ],
  "Temperature" : Number,
  "TopP" : Number
}

```

### YAML

```yaml

  MaxTokens: Number
  StopSequences:
    - String
  Temperature: Number
  TopP: Number

```

## Properties

`MaxTokens`

The maximum number of tokens to return in the response.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StopSequences`

A list of strings that define sequences after which the model will stop generating.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Temperature`

Controls the randomness of the response. Choose a lower value for more predictable outputs and a higher value for more surprising outputs.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopP`

The percentage of most-likely candidates that the model considers for the next token.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptInputVariable

PromptTemplateConfiguration

All content copied from https://docs.aws.amazon.com/.
