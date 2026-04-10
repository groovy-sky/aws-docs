This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Recipe ConditionExpression

Represents an individual condition that evaluates to true or false.

Conditions are used with recipe actions. The action is only performed for column values where the
condition evaluates to true.

If a recipe requires more than one condition, then the recipe must specify multiple
`ConditionExpression` elements. Each condition is applied to the rows in a dataset first, before
the recipe action is performed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : String,
  "TargetColumn" : String,
  "Value" : String
}

```

### YAML

```yaml

  Condition: String
  TargetColumn: String
  Value: String

```

## Properties

`Condition`

A specific condition to apply to a recipe action. For more information, see [Recipe\
structure](../../../databrew/latest/dg/recipe-structure.md) in the _AWS Glue DataBrew Developer_
_Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z\_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetColumn`

A column to apply this condition to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A value that the condition must evaluate to for the condition to succeed.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

DataCatalogInputDefinition

All content copied from https://docs.aws.amazon.com/.
