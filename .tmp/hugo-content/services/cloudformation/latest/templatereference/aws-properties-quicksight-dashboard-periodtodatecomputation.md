This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PeriodToDateComputation

The period to date computation configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputationId" : String,
  "Name" : String,
  "PeriodTimeGranularity" : String,
  "Time" : DimensionField,
  "Value" : MeasureField
}

```

### YAML

```yaml

  ComputationId: String
  Name: String
  PeriodTimeGranularity: String
  Time:
    DimensionField
  Value:
    MeasureField

```

## Properties

`ComputationId`

The ID for a computation.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a computation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodTimeGranularity`

The time granularity setup of period to date computation. Choose from the following options:

- YEAR: Year to date.

- MONTH: Month to date.

_Required_: No

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Time`

The time field that is used in a computation.

_Required_: No

_Type_: [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value field that is used in a computation.

_Required_: No

_Type_: [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PeriodOverPeriodComputation

PieChartAggregatedFieldWells

All content copied from https://docs.aws.amazon.com/.
