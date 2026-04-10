This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TableCellStyle

The table cell style for a cell in pivot table or table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : String,
  "Border" : GlobalTableBorderOptions,
  "FontConfiguration" : FontConfiguration,
  "Height" : Number,
  "HorizontalTextAlignment" : String,
  "TextWrap" : String,
  "VerticalTextAlignment" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  BackgroundColor: String
  Border:
    GlobalTableBorderOptions
  FontConfiguration:
    FontConfiguration
  Height: Number
  HorizontalTextAlignment: String
  TextWrap: String
  VerticalTextAlignment: String
  Visibility: String

```

## Properties

`BackgroundColor`

The background color for the table cells.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Border`

The borders for the table cells.

_Required_: No

_Type_: [GlobalTableBorderOptions](aws-properties-quicksight-template-globaltableborderoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontConfiguration`

The font configuration of the table cells.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-template-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Height`

The height color for the table cells.

_Required_: No

_Type_: Number

_Minimum_: `8`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HorizontalTextAlignment`

The horizontal text alignment (left, center, right, auto) for the table cells.

_Required_: No

_Type_: String

_Allowed values_: `LEFT | CENTER | RIGHT | AUTO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextWrap`

The text wrap (none, wrap) for the table cells.

_Required_: No

_Type_: String

_Allowed values_: `NONE | WRAP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerticalTextAlignment`

The vertical text alignment (top, middle, bottom) for the table cells.

_Required_: No

_Type_: String

_Allowed values_: `TOP | MIDDLE | BOTTOM | AUTO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the table cells.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableCellImageSizingConfiguration

TableConditionalFormatting

All content copied from https://docs.aws.amazon.com/.
