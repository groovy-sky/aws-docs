This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Evaluator NumericalScaleDefinition

The definition of a numerical rating scale option that provides a numeric value with its description for evaluation scoring.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Definition" : String,
  "Label" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Definition: String
  Label: String
  Value: Number

```

## Properties

`Definition`

The description that explains what this numerical rating represents and when it should be used.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

The label or name that describes this numerical rating option.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The numerical value for this rating scale option.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LlmAsAJudgeEvaluatorConfig

RatingScale

All content copied from https://docs.aws.amazon.com/.
