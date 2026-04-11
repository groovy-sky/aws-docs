This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TableBorderOptions

The border options for a table border.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "Style" : String,
  "Thickness" : Number
}

```

### YAML

```yaml

  Color: String
  Style: String
  Thickness: Number

```

## Properties

`Color`

The color of a table border.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Style`

The style (none, solid) of a table border.

_Required_: No

_Type_: String

_Allowed values_: `NONE | SOLID`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Thickness`

The thickness of a table border.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableAggregatedFieldWells

TableCellConditionalFormatting

All content copied from https://docs.aws.amazon.com/.
