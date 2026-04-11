This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GaugeChartColorConfiguration

The color configuration of a `GaugeChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : String,
  "ForegroundColor" : String
}

```

### YAML

```yaml

  BackgroundColor: String
  ForegroundColor: String

```

## Properties

`BackgroundColor`

The background color configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForegroundColor`

The foreground color configuration of a `GaugeChartVisual`.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartArcConditionalFormatting

GaugeChartConditionalFormatting

All content copied from https://docs.aws.amazon.com/.
