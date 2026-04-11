This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ArcConfiguration

The arc configuration of a `GaugeChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArcAngle" : Number,
  "ArcThickness" : String
}

```

### YAML

```yaml

  ArcAngle: Number
  ArcThickness: String

```

## Properties

`ArcAngle`

The option that determines the arc angle of a `GaugeChartVisual`.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ArcThickness`

The options that determine the arc thickness of a `GaugeChartVisual`.

_Required_: No

_Type_: String

_Allowed values_: `SMALL | MEDIUM | LARGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArcAxisDisplayRange

ArcOptions

All content copied from https://docs.aws.amazon.com/.
