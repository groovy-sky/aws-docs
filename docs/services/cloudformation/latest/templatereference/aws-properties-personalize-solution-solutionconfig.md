This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution SolutionConfig

Describes the configuration properties for the solution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlgorithmHyperParameters" : {Key: Value, ...},
  "AutoMLConfig" : AutoMLConfig,
  "EventValueThreshold" : String,
  "FeatureTransformationParameters" : {Key: Value, ...},
  "HpoConfig" : HpoConfig
}

```

### YAML

```yaml

  AlgorithmHyperParameters:
    Key: Value
  AutoMLConfig:
    AutoMLConfig
  EventValueThreshold: String
  FeatureTransformationParameters:
    Key: Value
  HpoConfig:
    HpoConfig

```

## Properties

`AlgorithmHyperParameters`

Lists the algorithm hyperparameters and their values.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AutoMLConfig`

The [AutoMLConfig](https://docs.aws.amazon.com/personalize/latest/dg/API_AutoMLConfig.html) object containing a list of recipes to search
when AutoML is performed.

_Required_: No

_Type_: [AutoMLConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-personalize-solution-automlconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventValueThreshold`

Only events with a value greater than or equal to this threshold are
used for training a model.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeatureTransformationParameters`

Lists the feature transformation parameters.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HpoConfig`

Describes the properties for hyperparameter optimization (HPO).

_Required_: No

_Type_: [HpoConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-personalize-solution-hpoconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IntegerHyperParameterRange

Next
