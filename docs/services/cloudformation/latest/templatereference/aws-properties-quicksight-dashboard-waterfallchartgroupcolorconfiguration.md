This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard WaterfallChartGroupColorConfiguration

The color configuration for individual groups within a waterfall visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NegativeBarColor" : String,
  "PositiveBarColor" : String,
  "TotalBarColor" : String
}

```

### YAML

```yaml

  NegativeBarColor: String
  PositiveBarColor: String
  TotalBarColor: String

```

## Properties

`NegativeBarColor`

Defines the color for the negative bars of a waterfall chart.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PositiveBarColor`

Defines the color for the positive bars of a waterfall chart.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalBarColor`

Defines the color for the total bars of a waterfall chart.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WaterfallChartFieldWells

WaterfallChartOptions

All content copied from https://docs.aws.amazon.com/.
