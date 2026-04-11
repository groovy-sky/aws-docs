This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialPointStyleOptions

The point style of the geospatial map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterMarkerConfiguration" : ClusterMarkerConfiguration,
  "HeatmapConfiguration" : GeospatialHeatmapConfiguration,
  "SelectedPointStyle" : String
}

```

### YAML

```yaml

  ClusterMarkerConfiguration:
    ClusterMarkerConfiguration
  HeatmapConfiguration:
    GeospatialHeatmapConfiguration
  SelectedPointStyle: String

```

## Properties

`ClusterMarkerConfiguration`

The cluster marker configuration of the geospatial point style.

_Required_: No

_Type_: [ClusterMarkerConfiguration](aws-properties-quicksight-dashboard-clustermarkerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeatmapConfiguration`

The heatmap configuration of the geospatial point style.

_Required_: No

_Type_: [GeospatialHeatmapConfiguration](aws-properties-quicksight-dashboard-geospatialheatmapconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedPointStyle`

The selected point styles (point, cluster) of the geospatial map.

_Required_: No

_Type_: String

_Allowed values_: `POINT | CLUSTER | HEATMAP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialPointStyle

GeospatialPolygonLayer

All content copied from https://docs.aws.amazon.com/.
