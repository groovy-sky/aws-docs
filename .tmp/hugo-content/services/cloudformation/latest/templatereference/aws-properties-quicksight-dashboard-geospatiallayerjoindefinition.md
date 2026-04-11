This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialLayerJoinDefinition

The custom actions for a layer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorField" : GeospatialLayerColorField,
  "DatasetKeyField" : UnaggregatedField,
  "ShapeKeyField" : String
}

```

### YAML

```yaml

  ColorField:
    GeospatialLayerColorField
  DatasetKeyField:
    UnaggregatedField
  ShapeKeyField: String

```

## Properties

`ColorField`

The geospatial color field for the join definition.

_Required_: No

_Type_: [GeospatialLayerColorField](aws-properties-quicksight-dashboard-geospatiallayercolorfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetKeyField`

Property description not available.

_Required_: No

_Type_: [UnaggregatedField](aws-properties-quicksight-dashboard-unaggregatedfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShapeKeyField`

The name of the field or property in the geospatial data source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialLayerItem

GeospatialLayerMapConfiguration

All content copied from https://docs.aws.amazon.com/.
