This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard BodySectionDynamicNumericDimensionConfiguration

Describes the **Numeric** dataset column and constraints for the dynamic values used to repeat the contents of a section.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "Limit" : Number,
  "SortByMetrics" : [ ColumnSort, ... ]
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  Limit: Number
  SortByMetrics:
    - ColumnSort

```

## Properties

`Column`

Property description not available.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-dashboard-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Limit`

Number of values to use from the column for repetition.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortByMetrics`

Sort criteria on the column values that you use for repetition.

_Required_: No

_Type_: Array of [ColumnSort](aws-properties-quicksight-dashboard-columnsort.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BodySectionDynamicCategoryDimensionConfiguration

BodySectionRepeatConfiguration

All content copied from https://docs.aws.amazon.com/.
