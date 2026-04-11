This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard MeasureField

The measure (metric) type field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CalculatedMeasureField" : CalculatedMeasureField,
  "CategoricalMeasureField" : CategoricalMeasureField,
  "DateMeasureField" : DateMeasureField,
  "NumericalMeasureField" : NumericalMeasureField
}

```

### YAML

```yaml

  CalculatedMeasureField:
    CalculatedMeasureField
  CategoricalMeasureField:
    CategoricalMeasureField
  DateMeasureField:
    DateMeasureField
  NumericalMeasureField:
    NumericalMeasureField

```

## Properties

`CalculatedMeasureField`

The calculated measure field only used in pivot tables.

_Required_: No

_Type_: [CalculatedMeasureField](aws-properties-quicksight-dashboard-calculatedmeasurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoricalMeasureField`

The measure type field with categorical type columns.

_Required_: No

_Type_: [CategoricalMeasureField](aws-properties-quicksight-dashboard-categoricalmeasurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateMeasureField`

The measure type field with date type columns.

_Required_: No

_Type_: [DateMeasureField](aws-properties-quicksight-dashboard-datemeasurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericalMeasureField`

The measure type field with numerical type columns.

_Required_: No

_Type_: [NumericalMeasureField](aws-properties-quicksight-dashboard-numericalmeasurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaximumMinimumComputation

MetricComparisonComputation

All content copied from https://docs.aws.amazon.com/.
