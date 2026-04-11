This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Evaluator EvaluatorConfig

The configuration of the evaluator, including LLM-as-a-Judge settings for custom evaluators.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeBased" : CodeBasedEvaluatorConfig,
  "LlmAsAJudge" : LlmAsAJudgeEvaluatorConfig
}

```

### YAML

```yaml

  CodeBased:
    CodeBasedEvaluatorConfig
  LlmAsAJudge:
    LlmAsAJudgeEvaluatorConfig

```

## Properties

`CodeBased`

Property description not available.

_Required_: No

_Type_: [CodeBasedEvaluatorConfig](aws-properties-bedrockagentcore-evaluator-codebasedevaluatorconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LlmAsAJudge`

The LLM-as-a-Judge configuration that uses a language model to evaluate agent performance based on custom instructions and rating scales.

_Required_: No

_Type_: [LlmAsAJudgeEvaluatorConfig](aws-properties-bedrockagentcore-evaluator-llmasajudgeevaluatorconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeBasedEvaluatorConfig

EvaluatorModelConfig

All content copied from https://docs.aws.amazon.com/.
