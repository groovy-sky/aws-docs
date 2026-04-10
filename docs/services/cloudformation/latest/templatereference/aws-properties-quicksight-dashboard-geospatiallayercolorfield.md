This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialLayerColorField

The color field that defines a gradient or categorical style.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorDimensionsFields" : [ DimensionField, ... ],
  "ColorValuesFields" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  ColorDimensionsFields:
    - DimensionField
  ColorValuesFields:
    - MeasureField

```

## Properties

`ColorDimensionsFields`

A list of color dimension fields.

_Required_: No

_Type_: Array of [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorValuesFields`

A list of color measure fields.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialHeatmapDataColor

GeospatialLayerDefinition

All content copied from https://docs.aws.amazon.com/.
