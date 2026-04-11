This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ColumnTooltipItem

The tooltip item for the columns that are not part of a field well.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregation" : AggregationFunction,
  "Column" : ColumnIdentifier,
  "Label" : String,
  "TooltipTarget" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  Aggregation:
    AggregationFunction
  Column:
    ColumnIdentifier
  Label: String
  TooltipTarget: String
  Visibility: String

```

## Properties

`Aggregation`

The aggregation function of the column tooltip item.

_Required_: No

_Type_: [AggregationFunction](aws-properties-quicksight-analysis-aggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Column`

The target column of the tooltip item.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

The label of the tooltip item.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TooltipTarget`

Determines the target of the column tooltip item in a combo chart visual.

_Required_: No

_Type_: String

_Allowed values_: `BOTH | BAR | LINE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the tooltip item.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnSort

ComboChartAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
