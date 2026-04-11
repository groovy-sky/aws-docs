This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FilledMapAggregatedFieldWells

The aggregated field well of the filled map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Geospatial" : [ DimensionField, ... ],
  "Values" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  Geospatial:
    - DimensionField
  Values:
    - MeasureField

```

## Properties

`Geospatial`

The aggregated location field well of the filled map. Values are grouped by location fields.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The aggregated color field well of a filled map. Values are aggregated based on location fields.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldTooltipItem

FilledMapConditionalFormatting

All content copied from https://docs.aws.amazon.com/.
