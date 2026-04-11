This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TimeBasedForecastProperties

The forecast properties setup of a forecast in the line chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LowerBoundary" : Number,
  "PeriodsBackward" : Number,
  "PeriodsForward" : Number,
  "PredictionInterval" : Number,
  "Seasonality" : Number,
  "UpperBoundary" : Number
}

```

### YAML

```yaml

  LowerBoundary: Number
  PeriodsBackward: Number
  PeriodsForward: Number
  PredictionInterval: Number
  Seasonality: Number
  UpperBoundary: Number

```

## Properties

`LowerBoundary`

The lower boundary setup of a forecast computation.

_Required_: No

_Type_: Number

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

- `NULL`: The input is set to `NULL`.

- `NON_NULL`: The input is set to a custom value.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `180`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpperBoundary`

The upper boundary setup of a forecast computation.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ThousandSeparatorOptions

TimeEqualityFilter

All content copied from https://docs.aws.amazon.com/.
