This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TooltipItem

The tooltip.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnTooltipItem" : ColumnTooltipItem,
  "FieldTooltipItem" : FieldTooltipItem
}

```

### YAML

```yaml

  ColumnTooltipItem:
    ColumnTooltipItem
  FieldTooltipItem:
    FieldTooltipItem

```

## Properties

`ColumnTooltipItem`

The tooltip item for the columns that are not part of a field well.

_Required_: No

_Type_: [ColumnTooltipItem](aws-properties-quicksight-template-columntooltipitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldTooltipItem`

The tooltip item for the fields.

_Required_: No

_Type_: [FieldTooltipItem](aws-properties-quicksight-template-fieldtooltipitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeRangeFilterValue

TooltipOptions

All content copied from https://docs.aws.amazon.com/.
