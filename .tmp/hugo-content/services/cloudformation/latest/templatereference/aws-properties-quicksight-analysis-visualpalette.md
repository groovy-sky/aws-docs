This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis VisualPalette

The visual display options for the visual palette.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChartColor" : String,
  "ColorMap" : [ DataPathColor, ... ]
}

```

### YAML

```yaml

  ChartColor: String
  ColorMap:
    - DataPathColor

```

## Properties

`ChartColor`

The chart color options for the visual palette.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorMap`

The color map options for the visual palette.

_Required_: No

_Type_: Array of [DataPathColor](aws-properties-quicksight-analysis-datapathcolor.md)

_Minimum_: `0`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisualMenuOption

VisualSubtitleLabelOptions

All content copied from https://docs.aws.amazon.com/.
