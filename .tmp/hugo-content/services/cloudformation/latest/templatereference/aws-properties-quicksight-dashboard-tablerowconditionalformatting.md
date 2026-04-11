This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TableRowConditionalFormatting

The conditional formatting of a table row.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : ConditionalFormattingColor,
  "TextColor" : ConditionalFormattingColor
}

```

### YAML

```yaml

  BackgroundColor:
    ConditionalFormattingColor
  TextColor:
    ConditionalFormattingColor

```

## Properties

`BackgroundColor`

The conditional formatting color (solid, gradient) of the background for a table row.

_Required_: No

_Type_: [ConditionalFormattingColor](aws-properties-quicksight-dashboard-conditionalformattingcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextColor`

The conditional formatting color (solid, gradient) of the text for a table row.

_Required_: No

_Type_: [ConditionalFormattingColor](aws-properties-quicksight-dashboard-conditionalformattingcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TablePinnedFieldOptions

TableSideBorderOptions

All content copied from https://docs.aws.amazon.com/.
