This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Evaluator LlmAsAJudgeEvaluatorConfig

The LLM-as-a-Judge configuration that uses a language model to evaluate agent performance based on custom instructions and rating scales.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Instructions" : String,
  "ModelConfig" : EvaluatorModelConfig,
  "RatingScale" : RatingScale
}

```

### YAML

```yaml

  Instructions: String
  ModelConfig:
    EvaluatorModelConfig
  RatingScale:
    RatingScale

```

## Properties

`Instructions`

The evaluation instructions that guide the language model in assessing agent performance, including criteria and evaluation guidelines.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelConfig`

The model configuration that specifies which foundation model to use and how to configure it for evaluation.

_Required_: Yes

_Type_: [EvaluatorModelConfig](aws-properties-bedrockagentcore-evaluator-evaluatormodelconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RatingScale`

The rating scale that defines how the evaluator should score agent performance, either numerical or categorical.

_Required_: Yes

_Type_: [RatingScale](aws-properties-bedrockagentcore-evaluator-ratingscale.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaEvaluatorConfig

NumericalScaleDefinition

All content copied from https://docs.aws.amazon.com/.
