This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GeospatialWindowOptions

The window options of the geospatial map visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bounds" : GeospatialCoordinateBounds,
  "MapZoomMode" : String
}

```

### YAML

```yaml

  Bounds:
    GeospatialCoordinateBounds
  MapZoomMode: String

```

## Properties

`Bounds`

The bounds options (north, south, west, east) of the geospatial window options.

_Required_: No

_Type_: [GeospatialCoordinateBounds](aws-properties-quicksight-analysis-geospatialcoordinatebounds.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapZoomMode`

The map zoom modes (manual, auto) of the geospatial window options.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialStaticFileSource

GlobalTableBorderOptions

All content copied from https://docs.aws.amazon.com/.
