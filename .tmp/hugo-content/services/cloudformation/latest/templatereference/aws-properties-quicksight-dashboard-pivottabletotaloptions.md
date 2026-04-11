This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PivotTableTotalOptions

The total options for a pivot table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnSubtotalOptions" : SubtotalOptions,
  "ColumnTotalOptions" : PivotTotalOptions,
  "RowSubtotalOptions" : SubtotalOptions,
  "RowTotalOptions" : PivotTotalOptions
}

```

### YAML

```yaml

  ColumnSubtotalOptions:
    SubtotalOptions
  ColumnTotalOptions:
    PivotTotalOptions
  RowSubtotalOptions:
    SubtotalOptions
  RowTotalOptions:
    PivotTotalOptions

```

## Properties

`ColumnSubtotalOptions`

The column subtotal options.

_Required_: No

_Type_: [SubtotalOptions](aws-properties-quicksight-dashboard-subtotaloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnTotalOptions`

The column total options.

_Required_: No

_Type_: [PivotTotalOptions](aws-properties-quicksight-dashboard-pivottotaloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowSubtotalOptions`

The row subtotal options.

_Required_: No

_Type_: [SubtotalOptions](aws-properties-quicksight-dashboard-subtotaloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowTotalOptions`

The row total options.

_Required_: No

_Type_: [PivotTotalOptions](aws-properties-quicksight-dashboard-pivottotaloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableSortConfiguration

PivotTableVisual

All content copied from https://docs.aws.amazon.com/.
