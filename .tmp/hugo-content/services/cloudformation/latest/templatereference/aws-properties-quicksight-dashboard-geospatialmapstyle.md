This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialMapStyle

The map style properties for a map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : String,
  "BaseMapStyle" : String,
  "BaseMapVisibility" : String
}

```

### YAML

```yaml

  BackgroundColor: String
  BaseMapStyle: String
  BaseMapVisibility: String

```

## Properties

`BackgroundColor`

The background color and opacity values for a map.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}(?:[A-F0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseMapStyle`

The selected base map style.

_Required_: No

_Type_: String

_Allowed values_: `LIGHT_GRAY | DARK_GRAY | STREET | IMAGERY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseMapVisibility`

The state of visibility for the base map.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialMapState

GeospatialMapStyleOptions

All content copied from https://docs.aws.amazon.com/.
