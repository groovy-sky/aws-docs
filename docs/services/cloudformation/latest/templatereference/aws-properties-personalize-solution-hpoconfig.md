This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution HpoConfig

Describes the properties for hyperparameter optimization (HPO).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlgorithmHyperParameterRanges" : AlgorithmHyperParameterRanges,
  "HpoObjective" : HpoObjective,
  "HpoResourceConfig" : HpoResourceConfig
}

```

### YAML

```yaml

  AlgorithmHyperParameterRanges:
    AlgorithmHyperParameterRanges
  HpoObjective:
    HpoObjective
  HpoResourceConfig:
    HpoResourceConfig

```

## Properties

`AlgorithmHyperParameterRanges`

The hyperparameters and their allowable ranges.

_Required_: No

_Type_: [AlgorithmHyperParameterRanges](aws-properties-personalize-solution-algorithmhyperparameterranges.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HpoObjective`

The metric to optimize during HPO.

###### Note

Amazon Personalize doesn't support configuring the `hpoObjective`
at this time.

_Required_: No

_Type_: [HpoObjective](aws-properties-personalize-solution-hpoobjective.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HpoResourceConfig`

Describes the resource configuration for HPO.

_Required_: No

_Type_: [HpoResourceConfig](aws-properties-personalize-solution-hporesourceconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContinuousHyperParameterRange

HpoObjective

All content copied from https://docs.aws.amazon.com/.
