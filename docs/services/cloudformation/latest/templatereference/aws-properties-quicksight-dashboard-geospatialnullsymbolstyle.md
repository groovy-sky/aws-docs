This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialNullSymbolStyle

The symbol style for null data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FillColor" : String,
  "StrokeColor" : String,
  "StrokeWidth" : Number
}

```

### YAML

```yaml

  FillColor: String
  StrokeColor: String
  StrokeWidth: Number

```

## Properties

`FillColor`

The color and opacity values for the fill color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}(?:[A-F0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrokeColor`

The color and opacity values for the stroke color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}(?:[A-F0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrokeWidth`

The width of the border stroke.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialNullDataSettings

GeospatialPointLayer

All content copied from https://docs.aws.amazon.com/.
