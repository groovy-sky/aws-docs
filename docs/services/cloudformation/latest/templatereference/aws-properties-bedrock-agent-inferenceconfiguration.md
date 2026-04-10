This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent InferenceConfiguration

Base inference parameters to pass to a model in a call to [Converse](../../../../reference/bedrock/latest/apireference/api-runtime-converse.md) or [ConverseStream](../../../../reference/bedrock/latest/apireference/api-runtime-conversestream.md). For more information,
see [Inference parameters for foundation models](../../../bedrock/latest/userguide/model-parameters.md).

If you need to pass additional parameters that the model
supports, use the `additionalModelRequestFields` request field in the call to `Converse`
or `ConverseStream`.
For more information, see [Model parameters](../../../bedrock/latest/userguide/model-parameters.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumLength" : Number,
  "StopSequences" : [ String, ... ],
  "Temperature" : Number,
  "TopK" : Number,
  "TopP" : Number
}

```

### YAML

```yaml

  MaximumLength: Number
  StopSequences:
    - String
  Temperature: Number
  TopK: Number
  TopP: Number

```

## Properties

`MaximumLength`

The maximum number of tokens allowed in the generated response.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StopSequences`

A list of stop sequences. A stop sequence is a sequence of characters that causes the
model to stop generating the response.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Temperature`

The likelihood of the model selecting higher-probability options while generating a
response. A lower value makes the model more likely to choose higher-probability options,
while a higher value makes the model more likely to choose lower-probability
options.

The default value is the default value for the model that
you are using. For more information, see [Inference parameters for foundation\
models](../../../bedrock/latest/userguide/model-parameters.md).

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopK`

While generating a response, the model determines the probability of the following token at each point of generation. The value that you set for `topK` is the number of most-likely candidates from which the model chooses the next token in the sequence. For example, if you set `topK` to 50, the model selects the next token from among the top 50 most likely choices.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopP`

The percentage of most-likely candidates that the model considers for the next token. For
example, if you choose a value of 0.8 for `topP`, the model selects from the top 80% of the
probability distribution of tokens that could be next in the sequence.

The default value is the default value for the model that you are using. For more information, see
[Inference parameters for foundation models](../../../bedrock/latest/userguide/model-parameters.md).

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GuardrailConfiguration

MemoryConfiguration

All content copied from https://docs.aws.amazon.com/.
