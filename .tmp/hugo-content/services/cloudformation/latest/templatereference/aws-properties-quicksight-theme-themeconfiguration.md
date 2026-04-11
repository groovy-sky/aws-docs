This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Theme ThemeConfiguration

The theme configuration. This configuration contains all of the display properties for
a theme.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataColorPalette" : DataColorPalette,
  "Sheet" : SheetStyle,
  "Typography" : Typography,
  "UIColorPalette" : UIColorPalette
}

```

### YAML

```yaml

  DataColorPalette:
    DataColorPalette
  Sheet:
    SheetStyle
  Typography:
    Typography
  UIColorPalette:
    UIColorPalette

```

## Properties

`DataColorPalette`

Color properties that apply to chart data colors.

_Required_: No

_Type_: [DataColorPalette](aws-properties-quicksight-theme-datacolorpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sheet`

Display options related to sheets.

_Required_: No

_Type_: [SheetStyle](aws-properties-quicksight-theme-sheetstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Typography`

Determines the typography options.

_Required_: No

_Type_: [Typography](aws-properties-quicksight-theme-typography.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UIColorPalette`

Color properties that apply to the UI and to charts, excluding the colors that apply
to data.

_Required_: No

_Type_: [UIColorPalette](aws-properties-quicksight-theme-uicolorpalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ThemeError

All content copied from https://docs.aws.amazon.com/.
