This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template GeospatialMapConfiguration

The configuration of a `GeospatialMapVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldWells" : GeospatialMapFieldWells,
  "Legend" : LegendOptions,
  "MapStyleOptions" : GeospatialMapStyleOptions,
  "PointStyleOptions" : GeospatialPointStyleOptions,
  "Tooltip" : TooltipOptions,
  "VisualPalette" : VisualPalette,
  "WindowOptions" : GeospatialWindowOptions
}

```

### YAML

```yaml

  FieldWells:
    GeospatialMapFieldWells
  Legend:
    LegendOptions
  MapStyleOptions:
    GeospatialMapStyleOptions
  PointStyleOptions:
    GeospatialPointStyleOptions
  Tooltip:
    TooltipOptions
  VisualPalette:
    VisualPalette
  WindowOptions:
    GeospatialWindowOptions

```

## Properties

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [GeospatialMapFieldWells](aws-properties-quicksight-template-geospatialmapfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

The legend display setup of the visual.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-template-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapStyleOptions`

The map style options of the geospatial map.

_Required_: No

_Type_: [GeospatialMapStyleOptions](aws-properties-quicksight-template-geospatialmapstyleoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PointStyleOptions`

The point style options of the geospatial map.

_Required_: No

_Type_: [GeospatialPointStyleOptions](aws-properties-quicksight-template-geospatialpointstyleoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-template-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPalette`

Property description not available.

_Required_: No

_Type_: [VisualPalette](aws-properties-quicksight-template-visualpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WindowOptions`

The window options of the geospatial map.

_Required_: No

_Type_: [GeospatialWindowOptions](aws-properties-quicksight-template-geospatialwindowoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialMapAggregatedFieldWells

GeospatialMapFieldWells

All content copied from https://docs.aws.amazon.com/.
