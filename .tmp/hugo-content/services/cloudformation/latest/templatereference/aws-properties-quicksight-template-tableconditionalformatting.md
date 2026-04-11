This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TableConditionalFormatting

The conditional formatting for a `PivotTableVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionalFormattingOptions" : [ TableConditionalFormattingOption, ... ]
}

```

### YAML

```yaml

  ConditionalFormattingOptions:
    - TableConditionalFormattingOption

```

## Properties

`ConditionalFormattingOptions`

Conditional formatting options for a `PivotTableVisual`.

_Required_: No

_Type_: Array of [TableConditionalFormattingOption](aws-properties-quicksight-template-tableconditionalformattingoption.md)

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableCellStyle

TableConditionalFormattingOption

All content copied from https://docs.aws.amazon.com/.
