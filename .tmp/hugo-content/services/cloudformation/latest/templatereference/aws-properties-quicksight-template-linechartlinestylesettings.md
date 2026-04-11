This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template LineChartLineStyleSettings

Line styles options for a line series in `LineChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LineInterpolation" : String,
  "LineStyle" : String,
  "LineVisibility" : String,
  "LineWidth" : String
}

```

### YAML

```yaml

  LineInterpolation: String
  LineStyle: String
  LineVisibility: String
  LineWidth: String

```

## Properties

`LineInterpolation`

Interpolation style for line series.

- `LINEAR`: Show as default, linear style.

- `SMOOTH`: Show as a smooth curve.

- `STEPPED`: Show steps in line.

_Required_: No

_Type_: String

_Allowed values_: `LINEAR | SMOOTH | STEPPED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineStyle`

Line style for line series.

- `SOLID`: Show as a solid line.

- `DOTTED`: Show as a dotted line.

- `DASHED`: Show as a dashed line.

_Required_: No

_Type_: String

_Allowed values_: `SOLID | DOTTED | DASHED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineVisibility`

Configuration option that determines whether to show the line for the series.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineWidth`

Width that determines the line thickness.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LineChartFieldWells

LineChartMarkerStyleSettings

All content copied from https://docs.aws.amazon.com/.
