This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialMapState

The map state properties for a map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bounds" : GeospatialCoordinateBounds,
  "MapNavigation" : String
}

```

### YAML

```yaml

  Bounds:
    GeospatialCoordinateBounds
  MapNavigation: String

```

## Properties

`Bounds`

Property description not available.

_Required_: No

_Type_: [GeospatialCoordinateBounds](aws-properties-quicksight-dashboard-geospatialcoordinatebounds.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapNavigation`

Enables or disables map navigation for a map.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialMapFieldWells

GeospatialMapStyle

All content copied from https://docs.aws.amazon.com/.
