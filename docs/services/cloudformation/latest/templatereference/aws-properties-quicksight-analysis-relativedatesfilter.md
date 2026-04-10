This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis RelativeDatesFilter

A `RelativeDatesFilter` filters relative dates values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnchorDateConfiguration" : AnchorDateConfiguration,
  "Column" : ColumnIdentifier,
  "DefaultFilterControlConfiguration" : DefaultFilterControlConfiguration,
  "ExcludePeriodConfiguration" : ExcludePeriodConfiguration,
  "FilterId" : String,
  "MinimumGranularity" : String,
  "NullOption" : String,
  "ParameterName" : String,
  "RelativeDateType" : String,
  "RelativeDateValue" : Number,
  "TimeGranularity" : String
}

```

### YAML

```yaml

  AnchorDateConfiguration:
    AnchorDateConfiguration
  Column:
    ColumnIdentifier
  DefaultFilterControlConfiguration:
    DefaultFilterControlConfiguration
  ExcludePeriodConfiguration:
    ExcludePeriodConfiguration
  FilterId: String
  MinimumGranularity: String
  NullOption: String
  ParameterName: String
  RelativeDateType: String
  RelativeDateValue: Number
  TimeGranularity: String

```

## Properties

`AnchorDateConfiguration`

The date configuration of the filter.

_Required_: Yes

_Type_: [AnchorDateConfiguration](aws-properties-quicksight-analysis-anchordateconfiguration.md)

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

`ExcludePeriodConfiguration`

The configuration for the exclude period of the filter.

_Required_: No

_Type_: [ExcludePeriodConfiguration](aws-properties-quicksight-analysis-excludeperiodconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterId`

An identifier that uniquely identifies a filter within a dashboard, analysis, or template.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumGranularity`

The minimum granularity (period granularity) of the relative dates filter.

_Required_: No

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

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

`ParameterName`

The parameter whose value should be used for the filter value.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelativeDateType`

The range date type of the filter. Choose one of the options below:

- `PREVIOUS`

- `THIS`

- `LAST`

- `NOW`

- `NEXT`

_Required_: Yes

_Type_: String

_Allowed values_: `PREVIOUS | THIS | LAST | NOW | NEXT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelativeDateValue`

The date value of the filter.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The level of time precision that is used to aggregate `DateTime` values.

_Required_: Yes

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceLineValueLabelConfiguration

RelativeDateTimeControlDisplayOptions

All content copied from https://docs.aws.amazon.com/.
