This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget CostCategoryValues

The cost category values used for filtering the costs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "MatchOptions" : [ String, ... ],
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  MatchOptions:
    - String
  Values:
    - String

```

## Properties

`Key`

The unique name of the cost category.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchOptions`

The match options that you can use to filter your results.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The specific value of the cost category.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BudgetData

CostTypes

All content copied from https://docs.aws.amazon.com/.
