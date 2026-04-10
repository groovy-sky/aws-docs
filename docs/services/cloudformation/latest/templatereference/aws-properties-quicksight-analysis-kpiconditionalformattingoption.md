This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis KPIConditionalFormattingOption

The conditional formatting options of a KPI visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActualValue" : KPIActualValueConditionalFormatting,
  "ComparisonValue" : KPIComparisonValueConditionalFormatting,
  "PrimaryValue" : KPIPrimaryValueConditionalFormatting,
  "ProgressBar" : KPIProgressBarConditionalFormatting
}

```

### YAML

```yaml

  ActualValue:
    KPIActualValueConditionalFormatting
  ComparisonValue:
    KPIComparisonValueConditionalFormatting
  PrimaryValue:
    KPIPrimaryValueConditionalFormatting
  ProgressBar:
    KPIProgressBarConditionalFormatting

```

## Properties

`ActualValue`

The conditional formatting for the actual value of a KPI visual.

_Required_: No

_Type_: [KPIActualValueConditionalFormatting](aws-properties-quicksight-analysis-kpiactualvalueconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComparisonValue`

The conditional formatting for the comparison value of a KPI visual.

_Required_: No

_Type_: [KPIComparisonValueConditionalFormatting](aws-properties-quicksight-analysis-kpicomparisonvalueconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryValue`

The conditional formatting for the primary value of a KPI visual.

_Required_: No

_Type_: [KPIPrimaryValueConditionalFormatting](aws-properties-quicksight-analysis-kpiprimaryvalueconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgressBar`

The conditional formatting for the progress bar of a KPI visual.

_Required_: No

_Type_: [KPIProgressBarConditionalFormatting](aws-properties-quicksight-analysis-kpiprogressbarconditionalformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KPIConditionalFormatting

KPIConfiguration

All content copied from https://docs.aws.amazon.com/.
