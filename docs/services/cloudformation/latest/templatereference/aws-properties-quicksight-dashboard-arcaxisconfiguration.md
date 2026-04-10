This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ArcAxisConfiguration

The arc axis configuration of a `GaugeChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Range" : ArcAxisDisplayRange,
  "ReserveRange" : Number
}

```

### YAML

```yaml

  Range:
    ArcAxisDisplayRange
  ReserveRange: Number

```

## Properties

`Range`

The arc axis range of a `GaugeChartVisual`.

_Required_: No

_Type_: [ArcAxisDisplayRange](aws-properties-quicksight-dashboard-arcaxisdisplayrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReserveRange`

The reserved range of the arc axis.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnchorDateConfiguration

ArcAxisDisplayRange

All content copied from https://docs.aws.amazon.com/.
