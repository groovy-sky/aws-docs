This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Theme UIColorPalette

The theme colors that apply to UI and to charts, excluding data colors. The colors description is a hexadecimal
color code that consists of six alphanumerical characters, prefixed with `#`, for example #37BFF5. For
more information, see [Using Themes\
in Amazon Quick](../../../quicksight/latest/user/themes-in-quicksight.md) in the _Amazon Quick User Guide._

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Accent" : String,
  "AccentForeground" : String,
  "Danger" : String,
  "DangerForeground" : String,
  "Dimension" : String,
  "DimensionForeground" : String,
  "Measure" : String,
  "MeasureForeground" : String,
  "PrimaryBackground" : String,
  "PrimaryForeground" : String,
  "SecondaryBackground" : String,
  "SecondaryForeground" : String,
  "Success" : String,
  "SuccessForeground" : String,
  "Warning" : String,
  "WarningForeground" : String
}

```

### YAML

```yaml

  Accent: String
  AccentForeground: String
  Danger: String
  DangerForeground: String
  Dimension: String
  DimensionForeground: String
  Measure: String
  MeasureForeground: String
  PrimaryBackground: String
  PrimaryForeground: String
  SecondaryBackground: String
  SecondaryForeground: String
  Success: String
  SuccessForeground: String
  Warning: String
  WarningForeground: String

```

## Properties

`Accent`

This color is that applies to selected states and buttons.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccentForeground`

The foreground color that applies to any text or other elements that appear over the
accent color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Danger`

The color that applies to error messages.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DangerForeground`

The foreground color that applies to any text or other elements that appear over the
error color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dimension`

The color that applies to the names of fields that are identified as
dimensions.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DimensionForeground`

The foreground color that applies to any text or other elements that appear over the
dimension color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Measure`

The color that applies to the names of fields that are identified as measures.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureForeground`

The foreground color that applies to any text or other elements that appear over the
measure color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryBackground`

The background color that applies to visuals and other high emphasis UI.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryForeground`

The color of text and other foreground elements that appear over the primary
background regions, such as grid lines, borders, table banding, icons, and so on.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryBackground`

The background color that applies to the sheet background and sheet controls.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryForeground`

The foreground color that applies to any sheet title, sheet control text, or UI that
appears over the secondary background.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Success`

The color that applies to success messages, for example the check mark for a
successful download.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessForeground`

The foreground color that applies to any text or other elements that appear over the
success color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Warning`

This color that applies to warning and informational messages.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarningForeground`

The foreground color that applies to any text or other elements that appear over the
warning color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Typography

AWS::QuickSight::Topic

All content copied from https://docs.aws.amazon.com/.
