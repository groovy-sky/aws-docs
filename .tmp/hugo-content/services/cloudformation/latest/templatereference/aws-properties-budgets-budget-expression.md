This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget Expression

Use Expression to filter in various Budgets APIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "And" : [ Expression, ... ],
  "CostCategories" : CostCategoryValues,
  "Dimensions" : ExpressionDimensionValues,
  "Not" : Expression,
  "Or" : [ Expression, ... ],
  "Tags" : TagValues
}

```

### YAML

```yaml

  And:
    - Expression
  CostCategories:
    CostCategoryValues
  Dimensions:
    ExpressionDimensionValues
  Not:
    Expression
  Or:
    - Expression
  Tags:
    TagValues

```

## Properties

`And`

Return results that match both Dimension objects.

_Required_: No

_Type_: Array of [Expression](aws-properties-budgets-budget-expression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CostCategories`

The filter that's based on CostCategoryValues.

_Required_: No

_Type_: [CostCategoryValues](aws-properties-budgets-budget-costcategoryvalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dimensions`

The specific Dimension to use for Expression.

_Required_: No

_Type_: [ExpressionDimensionValues](aws-properties-budgets-budget-expressiondimensionvalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Not`

Return results that don't match a Dimension object.

_Required_: No

_Type_: [Expression](aws-properties-budgets-budget-expression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Or`

Return results that match either Dimension object.

_Required_: No

_Type_: Array of [Expression](aws-properties-budgets-budget-expression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The specific Tag to use for Expression.

_Required_: No

_Type_: [TagValues](aws-properties-budgets-budget-tagvalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CostTypes

ExpressionDimensionValues

All content copied from https://docs.aws.amazon.com/.
