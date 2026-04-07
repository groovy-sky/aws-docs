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

_Type_: [DataColorPalette](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-theme-datacolorpalette.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sheet`

Display options related to sheets.

_Required_: No

_Type_: [SheetStyle](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-theme-sheetstyle.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Typography`

Determines the typography options.

_Required_: No

_Type_: [Typography](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-theme-typography.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UIColorPalette`

Color properties that apply to the UI and to charts, excluding the colors that apply
to data.

_Required_: No

_Type_: [UIColorPalette](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-theme-uicolorpalette.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ThemeError
