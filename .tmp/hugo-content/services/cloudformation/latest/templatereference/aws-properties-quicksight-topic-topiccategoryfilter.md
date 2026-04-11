This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicCategoryFilter

A structure that represents a category filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryFilterFunction" : String,
  "CategoryFilterType" : String,
  "Constant" : TopicCategoryFilterConstant,
  "Inverse" : Boolean
}

```

### YAML

```yaml

  CategoryFilterFunction: String
  CategoryFilterType: String
  Constant:
    TopicCategoryFilterConstant
  Inverse: Boolean

```

## Properties

`CategoryFilterFunction`

The category filter function. Valid values for this structure are `EXACT` and `CONTAINS`.

_Required_: No

_Type_: String

_Allowed values_: `EXACT | CONTAINS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryFilterType`

The category filter type. This element is used to specify whether a filter is a simple category filter or an inverse category filter.

_Required_: No

_Type_: String

_Allowed values_: `CUSTOM_FILTER | CUSTOM_FILTER_LIST | FILTER_LIST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Constant`

The constant used in a category filter.

_Required_: No

_Type_: [TopicCategoryFilterConstant](aws-properties-quicksight-topic-topiccategoryfilterconstant.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Inverse`

A Boolean value that indicates if the filter is inverse.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicCalculatedField

TopicCategoryFilterConstant

All content copied from https://docs.aws.amazon.com/.
