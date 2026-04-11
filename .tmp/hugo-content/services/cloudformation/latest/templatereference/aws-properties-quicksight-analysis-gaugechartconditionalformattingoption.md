This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GaugeChartConditionalFormattingOption

Conditional formatting options of a `GaugeChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arc" : GaugeChartArcConditionalFormatting,
  "PrimaryValue" : GaugeChartPrimaryValueConditionalFormatting
}

```

### YAML

```yaml

  Arc:
    GaugeChartArcConditionalFormatting
  PrimaryValue:
    GaugeChartPrimaryValueConditionalFormatting

```

## Properties

`Arc`

The options that determine the presentation of the arc of a `GaugeChartVisual`.

_Required_: No

_Type_: [GaugeChartArcConditionalFormatting](aws-properties-quicksight-analysis-gaugechartarcconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryValue`

The conditional formatting for the primary value of a `GaugeChartVisual`.

_Required_: No

_Type_: [GaugeChartPrimaryValueConditionalFormatting](aws-properties-quicksight-analysis-gaugechartprimaryvalueconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartConditionalFormatting

GaugeChartConfiguration

All content copied from https://docs.aws.amazon.com/.
