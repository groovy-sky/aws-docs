This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard NumericalMeasureField

The measure type field with numerical type columns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationFunction" : NumericalAggregationFunction,
  "Column" : ColumnIdentifier,
  "FieldId" : String,
  "FormatConfiguration" : NumberFormatConfiguration
}

```

### YAML

```yaml

  AggregationFunction:
    NumericalAggregationFunction
  Column:
    ColumnIdentifier
  FieldId: String
  FormatConfiguration:
    NumberFormatConfiguration

```

## Properties

`AggregationFunction`

The aggregation function of the measure field.

_Required_: No

_Type_: [NumericalAggregationFunction](aws-properties-quicksight-dashboard-numericalaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Column`

The column that is used in the `NumericalMeasureField`.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-dashboard-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The custom field ID.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormatConfiguration`

The format configuration of the field.

_Required_: No

_Type_: [NumberFormatConfiguration](aws-properties-quicksight-dashboard-numberformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumericalDimensionField

NumericAxisOptions

All content copied from https://docs.aws.amazon.com/.
