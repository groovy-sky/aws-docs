This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution AutoMLConfig

When the solution performs AutoML ( `performAutoML` is true in
[CreateSolution](../../../personalize/latest/dg/api-createsolution.md)), Amazon Personalize
determines which recipe, from the specified list, optimizes the given metric.
Amazon Personalize then uses that recipe for the solution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricName" : String,
  "RecipeList" : [ String, ... ]
}

```

### YAML

```yaml

  MetricName: String
  RecipeList:
    - String

```

## Properties

`MetricName`

The metric to optimize.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecipeList`

The list of candidate recipes.

_Required_: No

_Type_: Array of String

_Maximum_: `256 | 100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlgorithmHyperParameterRanges

CategoricalHyperParameterRange

All content copied from https://docs.aws.amazon.com/.
