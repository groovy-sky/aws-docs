This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TableConditionalFormattingOption

Conditional formatting options for a `PivotTableVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cell" : TableCellConditionalFormatting,
  "Row" : TableRowConditionalFormatting
}

```

### YAML

```yaml

  Cell:
    TableCellConditionalFormatting
  Row:
    TableRowConditionalFormatting

```

## Properties

`Cell`

The cell conditional formatting option for a table.

_Required_: No

_Type_: [TableCellConditionalFormatting](aws-properties-quicksight-analysis-tablecellconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Row`

The row conditional formatting option for a table.

_Required_: No

_Type_: [TableRowConditionalFormatting](aws-properties-quicksight-analysis-tablerowconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableConditionalFormatting

TableConfiguration

All content copied from https://docs.aws.amazon.com/.
