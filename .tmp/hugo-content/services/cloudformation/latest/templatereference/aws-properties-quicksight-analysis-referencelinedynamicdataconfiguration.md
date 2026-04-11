This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ReferenceLineDynamicDataConfiguration

The dynamic configuration of the reference line data configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Calculation" : NumericalAggregationFunction,
  "Column" : ColumnIdentifier,
  "MeasureAggregationFunction" : AggregationFunction
}

```

### YAML

```yaml

  Calculation:
    NumericalAggregationFunction
  Column:
    ColumnIdentifier
  MeasureAggregationFunction:
    AggregationFunction

```

## Properties

`Calculation`

The calculation that is used in the dynamic data.

_Required_: Yes

_Type_: [NumericalAggregationFunction](aws-properties-quicksight-analysis-numericalaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Column`

The column that the dynamic data targets.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureAggregationFunction`

The aggregation function that is used in the dynamic data.

_Required_: No

_Type_: [AggregationFunction](aws-properties-quicksight-analysis-aggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceLineDataConfiguration

ReferenceLineLabelConfiguration

All content copied from https://docs.aws.amazon.com/.
