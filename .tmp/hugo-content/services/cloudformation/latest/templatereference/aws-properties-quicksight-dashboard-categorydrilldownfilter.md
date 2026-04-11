This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard CategoryDrillDownFilter

The category drill down filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryValues" : [ String, ... ],
  "Column" : ColumnIdentifier
}

```

### YAML

```yaml

  CategoryValues:
    - String
  Column:
    ColumnIdentifier

```

## Properties

`CategoryValues`

A list of the string inputs that are the values of the category drill down filter.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0 | 0`

_Maximum_: `512 | 100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Column`

The column that the filter is applied to.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-dashboard-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CategoricalMeasureField

CategoryFilter

All content copied from https://docs.aws.amazon.com/.
