This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicCategoryFilterConstant

A constant used in a category filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CollectiveConstant" : CollectiveConstant,
  "ConstantType" : String,
  "SingularConstant" : String
}

```

### YAML

```yaml

  CollectiveConstant:
    CollectiveConstant
  ConstantType: String
  SingularConstant: String

```

## Properties

`CollectiveConstant`

A collective constant used in a category filter. This element is used to specify a list of values for the constant.

_Required_: No

_Type_: [CollectiveConstant](aws-properties-quicksight-topic-collectiveconstant.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConstantType`

The type of category filter constant. This element is used to specify whether a constant is a singular or collective. Valid values are `SINGULAR` and `COLLECTIVE`.

_Required_: No

_Type_: String

_Allowed values_: `SINGULAR | RANGE | COLLECTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingularConstant`

A singular constant used in a category filter. This element is used to specify a single value for the constant.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicCategoryFilter

TopicColumn

All content copied from https://docs.aws.amazon.com/.
