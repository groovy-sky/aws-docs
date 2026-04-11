This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GeospatialMapAggregatedFieldWells

The aggregated field wells for a geospatial map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Colors" : [ DimensionField, ... ],
  "Geospatial" : [ DimensionField, ... ],
  "Values" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  Colors:
    - DimensionField
  Geospatial:
    - DimensionField
  Values:
    - MeasureField

```

## Properties

`Colors`

The color field wells of a geospatial map.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Geospatial`

The geospatial field wells of a geospatial map. Values are grouped by geospatial fields.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The size field wells of a geospatial map. Values are aggregated based on geospatial fields.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-analysis-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialLineWidth

GeospatialMapConfiguration

All content copied from https://docs.aws.amazon.com/.
