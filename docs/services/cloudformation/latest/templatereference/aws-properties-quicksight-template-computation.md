This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template Computation

The computation union that is used in an insight visual.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Forecast" : ForecastComputation,
  "GrowthRate" : GrowthRateComputation,
  "MaximumMinimum" : MaximumMinimumComputation,
  "MetricComparison" : MetricComparisonComputation,
  "PeriodOverPeriod" : PeriodOverPeriodComputation,
  "PeriodToDate" : PeriodToDateComputation,
  "TopBottomMovers" : TopBottomMoversComputation,
  "TopBottomRanked" : TopBottomRankedComputation,
  "TotalAggregation" : TotalAggregationComputation,
  "UniqueValues" : UniqueValuesComputation
}

```

### YAML

```yaml

  Forecast:
    ForecastComputation
  GrowthRate:
    GrowthRateComputation
  MaximumMinimum:
    MaximumMinimumComputation
  MetricComparison:
    MetricComparisonComputation
  PeriodOverPeriod:
    PeriodOverPeriodComputation
  PeriodToDate:
    PeriodToDateComputation
  TopBottomMovers:
    TopBottomMoversComputation
  TopBottomRanked:
    TopBottomRankedComputation
  TotalAggregation:
    TotalAggregationComputation
  UniqueValues:
    UniqueValuesComputation

```

## Properties

`Forecast`

The forecast computation configuration.

_Required_: No

_Type_: [ForecastComputation](aws-properties-quicksight-template-forecastcomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrowthRate`

The growth rate computation configuration.

_Required_: No

_Type_: [GrowthRateComputation](aws-properties-quicksight-template-growthratecomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumMinimum`

The maximum and minimum computation configuration.

_Required_: No

_Type_: [MaximumMinimumComputation](aws-properties-quicksight-template-maximumminimumcomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricComparison`

The metric comparison computation configuration.

_Required_: No

_Type_: [MetricComparisonComputation](aws-properties-quicksight-template-metriccomparisoncomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodOverPeriod`

The period over period computation configuration.

_Required_: No

_Type_: [PeriodOverPeriodComputation](aws-properties-quicksight-template-periodoverperiodcomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodToDate`

The period to `DataSetIdentifier` computation configuration.

_Required_: No

_Type_: [PeriodToDateComputation](aws-properties-quicksight-template-periodtodatecomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopBottomMovers`

The top movers and bottom movers computation configuration.

_Required_: No

_Type_: [TopBottomMoversComputation](aws-properties-quicksight-template-topbottommoverscomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopBottomRanked`

The top ranked and bottom ranked computation configuration.

_Required_: No

_Type_: [TopBottomRankedComputation](aws-properties-quicksight-template-topbottomrankedcomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalAggregation`

The total aggregation computation configuration.

_Required_: No

_Type_: [TotalAggregationComputation](aws-properties-quicksight-template-totalaggregationcomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UniqueValues`

The unique values computation configuration.

_Required_: No

_Type_: [UniqueValuesComputation](aws-properties-quicksight-template-uniquevaluescomputation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComparisonFormatConfiguration

ConditionalFormattingColor

All content copied from https://docs.aws.amazon.com/.
