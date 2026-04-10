This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GeospatialLayerItem

The properties for a single geospatial layer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ LayerCustomAction, ... ],
  "DataSource" : GeospatialDataSourceItem,
  "JoinDefinition" : GeospatialLayerJoinDefinition,
  "Label" : String,
  "LayerDefinition" : GeospatialLayerDefinition,
  "LayerId" : String,
  "LayerType" : String,
  "Tooltip" : TooltipOptions,
  "Visibility" : String
}

```

### YAML

```yaml

  Actions:
    - LayerCustomAction
  DataSource:
    GeospatialDataSourceItem
  JoinDefinition:
    GeospatialLayerJoinDefinition
  Label: String
  LayerDefinition:
    GeospatialLayerDefinition
  LayerId: String
  LayerType: String
  Tooltip:
    TooltipOptions
  Visibility: String

```

## Properties

`Actions`

A list of custom actions for a layer.

_Required_: No

_Type_: Array of [LayerCustomAction](aws-properties-quicksight-analysis-layercustomaction.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSource`

The data source for the layer.

_Required_: No

_Type_: [GeospatialDataSourceItem](aws-properties-quicksight-analysis-geospatialdatasourceitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JoinDefinition`

The join definition properties for a layer.

_Required_: No

_Type_: [GeospatialLayerJoinDefinition](aws-properties-quicksight-analysis-geospatiallayerjoindefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

The label that is displayed for the layer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LayerDefinition`

The definition properties for a layer.

_Required_: No

_Type_: [GeospatialLayerDefinition](aws-properties-quicksight-analysis-geospatiallayerdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LayerId`

The ID of the layer.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LayerType`

The layer type.

_Required_: No

_Type_: String

_Allowed values_: `POINT | LINE | POLYGON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

Property description not available.

_Required_: No

_Type_: [TooltipOptions](aws-properties-quicksight-analysis-tooltipoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The state of visibility for the layer.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialLayerDefinition

GeospatialLayerJoinDefinition

All content copied from https://docs.aws.amazon.com/.
