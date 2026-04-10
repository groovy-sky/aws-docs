This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution AlgorithmHyperParameterRanges

Specifies the hyperparameters and their ranges.
Hyperparameters can be categorical, continuous, or integer-valued.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoricalHyperParameterRanges" : [ CategoricalHyperParameterRange, ... ],
  "ContinuousHyperParameterRanges" : [ ContinuousHyperParameterRange, ... ],
  "IntegerHyperParameterRanges" : [ IntegerHyperParameterRange, ... ]
}

```

### YAML

```yaml

  CategoricalHyperParameterRanges:
    - CategoricalHyperParameterRange
  ContinuousHyperParameterRanges:
    - ContinuousHyperParameterRange
  IntegerHyperParameterRanges:
    - IntegerHyperParameterRange

```

## Properties

`CategoricalHyperParameterRanges`

Provides the name and range of a categorical hyperparameter.

_Required_: No

_Type_: Array of [CategoricalHyperParameterRange](aws-properties-personalize-solution-categoricalhyperparameterrange.md)

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContinuousHyperParameterRanges`

Provides the name and range of a continuous hyperparameter.

_Required_: No

_Type_: Array of [ContinuousHyperParameterRange](aws-properties-personalize-solution-continuoushyperparameterrange.md)

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IntegerHyperParameterRanges`

Provides the name and range of an integer-valued hyperparameter.

_Required_: No

_Type_: Array of [IntegerHyperParameterRange](aws-properties-personalize-solution-integerhyperparameterrange.md)

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Personalize::Solution

AutoMLConfig

All content copied from https://docs.aws.amazon.com/.
