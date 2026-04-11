This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TimeRangeFilter

A `TimeRangeFilter` filters values that are between two specified values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "DefaultFilterControlConfiguration" : DefaultFilterControlConfiguration,
  "ExcludePeriodConfiguration" : ExcludePeriodConfiguration,
  "FilterId" : String,
  "IncludeMaximum" : Boolean,
  "IncludeMinimum" : Boolean,
  "NullOption" : String,
  "RangeMaximumValue" : TimeRangeFilterValue,
  "RangeMinimumValue" : TimeRangeFilterValue,
  "TimeGranularity" : String
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  DefaultFilterControlConfiguration:
    DefaultFilterControlConfiguration
  ExcludePeriodConfiguration:
    ExcludePeriodConfiguration
  FilterId: String
  IncludeMaximum: Boolean
  IncludeMinimum: Boolean
  NullOption: String
  RangeMaximumValue:
    TimeRangeFilterValue
  RangeMinimumValue:
    TimeRangeFilterValue
  TimeGranularity: String

```

## Properties

`Column`

The column that the filter is applied to.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-dashboard-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultFilterControlConfiguration`

The default configurations for the associated controls. This applies only for filters that are scoped to multiple sheets.

_Required_: No

_Type_: [DefaultFilterControlConfiguration](aws-properties-quicksight-dashboard-defaultfiltercontrolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludePeriodConfiguration`

The exclude period of the time range filter.

_Required_: No

_Type_: [ExcludePeriodConfiguration](aws-properties-quicksight-dashboard-excludeperiodconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterId`

An identifier that uniquely identifies a filter within a dashboard, analysis, or template.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeMaximum`

Determines whether the maximum value in the filter value range should be included in the filtered results.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeMinimum`

Determines whether the minimum value in the filter value range should be included in the filtered results.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullOption`

This option determines how null values should be treated when filtering data.

- `ALL_VALUES`: Include null values in filtered results.

- `NULLS_ONLY`: Only include null values in filtered results.

- `NON_NULLS_ONLY`: Exclude null values from filtered results.

_Required_: Yes

_Type_: String

_Allowed values_: `ALL_VALUES | NULLS_ONLY | NON_NULLS_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeMaximumValue`

The maximum value for the filter value range.

_Required_: No

_Type_: [TimeRangeFilterValue](aws-properties-quicksight-dashboard-timerangefiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeMinimumValue`

The minimum value for the filter value range.

_Required_: No

_Type_: [TimeRangeFilterValue](aws-properties-quicksight-dashboard-timerangefiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The level of time precision that is used to aggregate `DateTime` values.

_Required_: No

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeRangeDrillDownFilter

TimeRangeFilterValue

All content copied from https://docs.aws.amazon.com/.
