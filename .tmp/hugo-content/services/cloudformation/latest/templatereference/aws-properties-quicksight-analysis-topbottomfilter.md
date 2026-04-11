This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TopBottomFilter

A `TopBottomFilter` filters values that are at the top or the bottom.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationSortConfigurations" : [ AggregationSortConfiguration, ... ],
  "Column" : ColumnIdentifier,
  "DefaultFilterControlConfiguration" : DefaultFilterControlConfiguration,
  "FilterId" : String,
  "Limit" : Number,
  "ParameterName" : String,
  "TimeGranularity" : String
}

```

### YAML

```yaml

  AggregationSortConfigurations:
    - AggregationSortConfiguration
  Column:
    ColumnIdentifier
  DefaultFilterControlConfiguration:
    DefaultFilterControlConfiguration
  FilterId: String
  Limit: Number
  ParameterName: String
  TimeGranularity: String

```

## Properties

`AggregationSortConfigurations`

The aggregation and sort configuration of the top bottom filter.

_Required_: Yes

_Type_: Array of [AggregationSortConfiguration](aws-properties-quicksight-analysis-aggregationsortconfiguration.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Column`

The column that the filter is applied to.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultFilterControlConfiguration`

The default configurations for the associated controls. This applies only for filters that are scoped to multiple sheets.

_Required_: No

_Type_: [DefaultFilterControlConfiguration](aws-properties-quicksight-analysis-defaultfiltercontrolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterId`

An identifier that uniquely identifies a filter within a dashboard, analysis, or template.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Limit`

The number of items to include in the top bottom filter results.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterName`

The parameter whose value should be used for the filter value.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The level of time precision that is used to aggregate `DateTime` values.

_Required_: No

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TooltipOptions

TopBottomMoversComputation

All content copied from https://docs.aws.amazon.com/.
