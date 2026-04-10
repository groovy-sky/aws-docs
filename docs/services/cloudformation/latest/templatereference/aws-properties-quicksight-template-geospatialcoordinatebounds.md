This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template GeospatialCoordinateBounds

The bound
options (north, south, west, east) of the geospatial window options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "East" : Number,
  "North" : Number,
  "South" : Number,
  "West" : Number
}

```

### YAML

```yaml

  East: Number
  North: Number
  South: Number
  West: Number

```

## Properties

`East`

The longitude of the east bound of the geospatial coordinate bounds.

_Required_: Yes

_Type_: Number

_Minimum_: `-1800`

_Maximum_: `1800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`North`

The latitude of the north bound of the geospatial coordinate bounds.

_Required_: Yes

_Type_: Number

_Minimum_: `-90`

_Maximum_: `90`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`South`

The latitude of the south bound of the geospatial coordinate bounds.

_Required_: Yes

_Type_: Number

_Minimum_: `-90`

_Maximum_: `90`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`West`

The longitude of the west bound of the geospatial coordinate bounds.

_Required_: Yes

_Type_: Number

_Minimum_: `-1800`

_Maximum_: `1800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartVisual

GeospatialHeatmapColorScale

All content copied from https://docs.aws.amazon.com/.
