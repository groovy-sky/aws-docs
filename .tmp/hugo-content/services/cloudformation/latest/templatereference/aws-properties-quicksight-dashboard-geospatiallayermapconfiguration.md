This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialLayerMapConfiguration

The map definition that defines map state, map style, and geospatial layers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "MapLayers" : [ GeospatialLayerItem, ... ],
  "MapState" : GeospatialMapState,
  "MapStyle" : GeospatialMapStyle
}

```

### YAML

```yaml

  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  MapLayers:
    - GeospatialLayerItem
  MapState:
    GeospatialMapState
  MapStyle:
    GeospatialMapStyle

```

## Properties

`Interactions`

Property description not available.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-dashboard-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

Property description not available.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-dashboard-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapLayers`

The geospatial layers to visualize on the map.

_Required_: No

_Type_: Array of [GeospatialLayerItem](aws-properties-quicksight-dashboard-geospatiallayeritem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapState`

The map state properties for the map.

_Required_: No

_Type_: [GeospatialMapState](aws-properties-quicksight-dashboard-geospatialmapstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapStyle`

The map style properties for the map.

_Required_: No

_Type_: [GeospatialMapStyle](aws-properties-quicksight-dashboard-geospatialmapstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialLayerJoinDefinition

GeospatialLineLayer

All content copied from https://docs.aws.amazon.com/.
