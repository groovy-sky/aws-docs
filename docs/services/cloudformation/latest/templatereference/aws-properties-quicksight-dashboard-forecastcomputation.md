This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ForecastComputation

The forecast computation configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputationId" : String,
  "CustomSeasonalityValue" : Number,
  "LowerBoundary" : Number,
  "Name" : String,
  "PeriodsBackward" : Number,
  "PeriodsForward" : Number,
  "PredictionInterval" : Number,
  "Seasonality" : String,
  "Time" : DimensionField,
  "UpperBoundary" : Number,
  "Value" : MeasureField
}

```

### YAML

```yaml

  ComputationId: String
  CustomSeasonalityValue: Number
  LowerBoundary: Number
  Name: String
  PeriodsBackward: Number
  PeriodsForward: Number
  PredictionInterval: Number
  Seasonality: String
  Time:
    DimensionField
  UpperBoundary: Number
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

`CustomSeasonalityValue`

The custom seasonality value setup of a forecast computation.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `180`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LowerBoundary`

The lower boundary setup of a forecast computation.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a computation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodsBackward`

The periods backward setup of a forecast computation.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeriodsForward`

The periods forward setup of a forecast computation.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredictionInterval`

The prediction interval setup of a forecast computation.

_Required_: No

_Type_: Number

_Minimum_: `50`

_Maximum_: `95`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Seasonality`

The seasonality setup of a forecast computation. Choose one of the following options:

- `AUTOMATIC`

- `CUSTOM`: Checks the custom seasonality value.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Time`

The time field that is used in a computation.

_Required_: No

_Type_: [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpperBoundary`

The upper boundary setup of a forecast computation.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value field that is used in a computation.

_Required_: No

_Type_: [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FontWeight

ForecastConfiguration

All content copied from https://docs.aws.amazon.com/.
