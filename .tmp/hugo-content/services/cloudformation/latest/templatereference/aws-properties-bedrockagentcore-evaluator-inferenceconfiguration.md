This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Evaluator InferenceConfiguration

The inference configuration parameters that control model behavior during evaluation, including temperature, token limits, and sampling settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxTokens" : Integer,
  "Temperature" : Number,
  "TopP" : Number
}

```

### YAML

```yaml

  MaxTokens: Integer
  Temperature: Number
  TopP: Number

```

## Properties

`MaxTokens`

The maximum number of tokens to generate in the model response during evaluation.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Temperature`

The temperature value that controls randomness in the model's responses. Lower values produce more deterministic outputs.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopP`

The top-p sampling parameter that controls the diversity of the model's responses by limiting the cumulative probability of token choices.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluatorModelConfig

LambdaEvaluatorConfig

All content copied from https://docs.aws.amazon.com/.
