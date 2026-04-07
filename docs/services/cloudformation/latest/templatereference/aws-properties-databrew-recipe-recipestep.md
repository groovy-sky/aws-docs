This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Recipe RecipeStep

Represents a single step from a DataBrew recipe to be performed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : Action,
  "ConditionExpressions" : [ ConditionExpression, ... ]
}

```

### YAML

```yaml

  Action:
    Action
  ConditionExpressions:
    - ConditionExpression

```

## Properties

`Action`

The particular action to be performed in the recipe step.

_Required_: Yes

_Type_: [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-recipe-action.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConditionExpressions`

One or more conditions that must be met for the recipe step to succeed.

###### Note

All of the conditions in the array must be met. In other words, all of the
conditions must be combined using a logical AND operation.

_Required_: No

_Type_: Array of [ConditionExpression](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-recipe-conditionexpression.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RecipeParameters

S3Location
