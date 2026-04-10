This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard VisualTitleLabelOptions

The title label options for a visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FormatText" : ShortFormatText,
  "Visibility" : String
}

```

### YAML

```yaml

  FormatText:
    ShortFormatText
  Visibility: String

```

## Properties

`FormatText`

The short text format of the title label, such as plain text or rich text.

_Required_: No

_Type_: [ShortFormatText](aws-properties-quicksight-dashboard-shortformattext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the title label.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisualSubtitleLabelOptions

WaterfallChartAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
