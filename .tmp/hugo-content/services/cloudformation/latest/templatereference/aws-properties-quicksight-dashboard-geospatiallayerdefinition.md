This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialLayerDefinition

The definition properties for a geospatial layer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LineLayer" : GeospatialLineLayer,
  "PointLayer" : GeospatialPointLayer,
  "PolygonLayer" : GeospatialPolygonLayer
}

```

### YAML

```yaml

  LineLayer:
    GeospatialLineLayer
  PointLayer:
    GeospatialPointLayer
  PolygonLayer:
    GeospatialPolygonLayer

```

## Properties

`LineLayer`

The definition for a line layer.

_Required_: No

_Type_: [GeospatialLineLayer](aws-properties-quicksight-dashboard-geospatiallinelayer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PointLayer`

The definition for a point layer.

_Required_: No

_Type_: [GeospatialPointLayer](aws-properties-quicksight-dashboard-geospatialpointlayer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolygonLayer`

The definition for a polygon layer.

_Required_: No

_Type_: [GeospatialPolygonLayer](aws-properties-quicksight-dashboard-geospatialpolygonlayer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialLayerColorField

GeospatialLayerItem

All content copied from https://docs.aws.amazon.com/.
