This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GeospatialSolidColor

The definition for a solid color.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "State" : String
}

```

### YAML

```yaml

  Color: String
  State: String

```

## Properties

`Color`

The color and opacity values for the color.

_Required_: Yes

_Type_: String

_Pattern_: `^#[A-F0-9]{6}(?:[A-F0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

Enables and disables the view state of the color.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialPolygonSymbolStyle

GeospatialStaticFileSource

All content copied from https://docs.aws.amazon.com/.
