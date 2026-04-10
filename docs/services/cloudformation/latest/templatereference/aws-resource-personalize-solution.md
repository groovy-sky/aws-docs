This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution

###### Important

By default, all new solutions use automatic training. With automatic training, you incur training costs while
your solution is active. To avoid unnecessary costs, when you are finished you can
[update the solution](../../../personalize/latest/dg/api-updatesolution.md) to turn off automatic training.
For information about training
costs, see [Amazon Personalize pricing](https://aws.amazon.com/personalize/pricing).

An object that provides information about a solution. A solution includes the custom recipe, customized parameters, and
trained models (Solution Versions) that Amazon Personalize uses to generate recommendations.

After you create a solution, you can’t change its configuration. If you need to make changes, you can [clone the solution](../../../personalize/latest/dg/cloning-solution.md) with the Amazon Personalize console
or create a new one.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Personalize::Solution",
  "Properties" : {
      "DatasetGroupArn" : String,
      "EventType" : String,
      "Name" : String,
      "PerformAutoML" : Boolean,
      "PerformHPO" : Boolean,
      "RecipeArn" : String,
      "SolutionConfig" : SolutionConfig
    }
}

```

### YAML

```yaml

Type: AWS::Personalize::Solution
Properties:
  DatasetGroupArn: String
  EventType: String
  Name: String
  PerformAutoML: Boolean
  PerformHPO: Boolean
  RecipeArn: String
  SolutionConfig:
    SolutionConfig

```

## Properties

`DatasetGroupArn`

The Amazon Resource Name (ARN) of the dataset group that provides the training data.

_Required_: Yes

_Type_: String

_Pattern_: `arn:([a-z\d-]+):personalize:.*:.*:.+`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventType`

The event type (for example, 'click' or 'like') that is used for training the model.
If no `eventType` is provided, Amazon Personalize uses all interactions for training with
equal weight regardless of type.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the solution.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9\-_]*`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PerformAutoML`

###### Important

We don't recommend enabling automated machine learning. Instead, match your use case to the available Amazon Personalize
recipes. For more information, see [Determining your use case.](../../../personalize/latest/dg/determining-use-case.md)

When true, Amazon Personalize performs a search for the best USER\_PERSONALIZATION recipe from
the list specified in the solution configuration ( `recipeArn` must not be specified).
When false (the default), Amazon Personalize uses `recipeArn` for training.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PerformHPO`

Whether to perform hyperparameter optimization (HPO) on the chosen recipe. The
default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecipeArn`

The ARN of the recipe used to create the solution. This is required when
`performAutoML` is false.

_Required_: No

_Type_: String

_Pattern_: `arn:([a-z\d-]+):personalize:.*:.*:.+`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SolutionConfig`

Describes the configuration properties for the solution.

_Required_: No

_Type_: [SolutionConfig](aws-properties-personalize-solution-solutionconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the resource.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`SolutionArn`

The Amazon Resource Name (ARN) of the solution.

## Examples

### Creating a solution

The following example creates an Amazon Personalize solution with the User-Personalization recipe
and an event value threshold.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "MySolution": {
            "Type": "AWS::Personalize::Solution",
            "Properties": {
               "Name": "my-solution-name",
               "DatasetGroupArn": "arn:aws:personalize:us-west-2:123456789012:dataset-group/my-dataset-group-name",
               "RecipeArn": "arn:aws:personalize:::recipe/aws-user-personalization",
               "SolutionConfig": {
                  "EventValueThreshold" : ".05"
                }
            }
         }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MySolution:
    Type: 'AWS::Personalize::Solution'
    Properties:
      Name: my-solution-name
      DatasetGroupArn: >-
        arn:aws:personalize:us-west-2:123456789012:dataset-group/my-dataset-group-name
      RecipeArn: 'arn:aws:personalize:::recipe/aws-user-personalization'
      SolutionConfig:
        EventValueThreshold: '.05'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Personalize::Schema

AlgorithmHyperParameterRanges

All content copied from https://docs.aws.amazon.com/.
