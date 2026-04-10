This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Evaluator RatingScale

The rating scale that defines how the evaluator should score agent performance, either numerical or categorical.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Categorical" : [ CategoricalScaleDefinition, ... ],
  "Numerical" : [ NumericalScaleDefinition, ... ]
}

```

### YAML

```yaml

  Categorical:
    - CategoricalScaleDefinition
  Numerical:
    - NumericalScaleDefinition

```

## Properties

`Categorical`

The categorical rating scale with named categories and definitions for qualitative evaluation.

_Required_: No

_Type_: Array of [CategoricalScaleDefinition](aws-properties-bedrockagentcore-evaluator-categoricalscaledefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Numerical`

The numerical rating scale with defined score values and descriptions for quantitative evaluation.

_Required_: No

_Type_: Array of [NumericalScaleDefinition](aws-properties-bedrockagentcore-evaluator-numericalscaledefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumericalScaleDefinition

Tag

All content copied from https://docs.aws.amazon.com/.
