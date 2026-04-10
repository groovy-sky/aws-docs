This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard LineChartMarkerStyleSettings

Marker styles options for a line series in `LineChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MarkerColor" : String,
  "MarkerShape" : String,
  "MarkerSize" : String,
  "MarkerVisibility" : String
}

```

### YAML

```yaml

  MarkerColor: String
  MarkerShape: String
  MarkerSize: String
  MarkerVisibility: String

```

## Properties

`MarkerColor`

Color of marker in the series.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MarkerShape`

Shape option for markers in the series.

- `CIRCLE`: Show marker as a circle.

- `TRIANGLE`: Show marker as a triangle.

- `SQUARE`: Show marker as a square.

- `DIAMOND`: Show marker as a diamond.

- `ROUNDED_SQUARE`: Show marker as a rounded square.

_Required_: No

_Type_: String

_Allowed values_: `CIRCLE | TRIANGLE | SQUARE | DIAMOND | ROUNDED_SQUARE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MarkerSize`

Size of marker in the series.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MarkerVisibility`

Configuration option that determines whether to show the markers in the series.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LineChartLineStyleSettings

LineChartSeriesSettings

All content copied from https://docs.aws.amazon.com/.
