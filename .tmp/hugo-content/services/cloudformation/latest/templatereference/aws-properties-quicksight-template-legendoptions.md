This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template LegendOptions

The options for the legend setup of a visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Height" : String,
  "Position" : String,
  "Title" : LabelOptions,
  "ValueFontConfiguration" : FontConfiguration,
  "Visibility" : String,
  "Width" : String
}

```

### YAML

```yaml

  Height: String
  Position: String
  Title:
    LabelOptions
  ValueFontConfiguration:
    FontConfiguration
  Visibility: String
  Width: String

```

## Properties

`Height`

The height of the legend. If this value is omitted, a default height is used when
rendering.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

The positions for the legend. Choose one of the following
options:

- `AUTO`

- `RIGHT`

- `BOTTOM`

- `LEFT`

_Required_: No

_Type_: String

_Allowed values_: `AUTO | RIGHT | BOTTOM | TOP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The custom title for the legend.

_Required_: No

_Type_: [LabelOptions](aws-properties-quicksight-template-labeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueFontConfiguration`

Property description not available.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-template-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

Determines whether or not the legend is visible.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Width`

The width of the legend. If this value is omitted, a default width is used when rendering.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LayoutConfiguration

LineChartAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
