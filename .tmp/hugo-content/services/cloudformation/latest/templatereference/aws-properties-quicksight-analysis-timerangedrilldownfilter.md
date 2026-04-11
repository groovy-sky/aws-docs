This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TimeRangeDrillDownFilter

The time range drill down filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "RangeMaximum" : String,
  "RangeMinimum" : String,
  "TimeGranularity" : String
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  RangeMaximum: String
  RangeMinimum: String
  TimeGranularity: String

```

## Properties

`Column`

The column that the filter is applied to.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeMaximum`

The maximum value for the filter value range.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeMinimum`

The minimum value for the filter value range.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The level of time precision that is used to aggregate `DateTime` values.

_Required_: Yes

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeEqualityFilter

TimeRangeFilter

All content copied from https://docs.aws.amazon.com/.
