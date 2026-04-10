This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DateMeasureField

The measure type field with date type columns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationFunction" : String,
  "Column" : ColumnIdentifier,
  "FieldId" : String,
  "FormatConfiguration" : DateTimeFormatConfiguration
}

```

### YAML

```yaml

  AggregationFunction: String
  Column:
    ColumnIdentifier
  FieldId: String
  FormatConfiguration:
    DateTimeFormatConfiguration

```

## Properties

`AggregationFunction`

The aggregation function of the measure field.

_Required_: No

_Type_: [String](aws-properties-quicksight-analysis-aggregationfunction.md)

_Allowed values_: `COUNT | DISTINCT_COUNT | MIN | MAX`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Column`

The column that is used in the `DateMeasureField`.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

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

_Type_: [DateTimeFormatConfiguration](aws-properties-quicksight-analysis-datetimeformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateDimensionField

DateTimeDefaultValues

All content copied from https://docs.aws.amazon.com/.
