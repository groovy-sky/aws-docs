This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GaugeChartPrimaryValueConditionalFormatting

The conditional formatting for the primary value of a `GaugeChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Icon" : ConditionalFormattingIcon,
  "TextColor" : ConditionalFormattingColor
}

```

### YAML

```yaml

  Icon:
    ConditionalFormattingIcon
  TextColor:
    ConditionalFormattingColor

```

## Properties

`Icon`

The conditional formatting of the primary value icon.

_Required_: No

_Type_: [ConditionalFormattingIcon](aws-properties-quicksight-dashboard-conditionalformattingicon.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextColor`

The conditional formatting of the primary value text color.

_Required_: No

_Type_: [ConditionalFormattingColor](aws-properties-quicksight-dashboard-conditionalformattingcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartOptions

GaugeChartVisual

All content copied from https://docs.aws.amazon.com/.
