This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FilledMapConfiguration

The configuration for a `FilledMapVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldWells" : FilledMapFieldWells,
  "Interactions" : VisualInteractionOptions,
  "Legend" : LegendOptions,
  "MapStyleOptions" : GeospatialMapStyleOptions,
  "SortConfiguration" : FilledMapSortConfiguration,
  "Tooltip" : TooltipOptions,
  "WindowOptions" : GeospatialWindowOptions
}

```

### YAML

```yaml

  FieldWells:
    FilledMapFieldWells
  Interactions:
    VisualInteractionOptions
  Legend:
    LegendOptions
  MapStyleOptions:
    GeospatialMapStyleOptions
  SortConfiguration:
    FilledMapSortConfiguration
  Tooltip:
    TooltipOptions
  WindowOptions:
    GeospatialWindowOptions

```

## Properties

`FieldWells`

The field wells of the visual.

_Required_: No

_Type_: [FilledMapFieldWells](aws-properties-quicksight-dashboard-filledmapfieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-dashboard-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Legend`

The legend display setup of the visual.

_Required_: No

_Type_: [LegendOptions](aws-properties-quicksight-dashboard-legendoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapStyleOptions`

The map style options of the filled map visual.

_Required_: No

_Type_: [GeospatialMapStyleOptions](aws-properties-quicksight-dashboard-geospatialmapstyleoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a `FilledMapVisual`.

_Required_: No

_Type_: [FilledMapSortConfiguration](aws-properties-quicksight-dashboard-filledmapsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip display setup of the visual.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-dashboard-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WindowOptions`

The window options of the filled map visual.

_Required_: No

_Type_: [GeospatialWindowOptions](aws-properties-quicksight-dashboard-geospatialwindowoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilledMapConditionalFormattingOption

FilledMapFieldWells

All content copied from https://docs.aws.amazon.com/.
