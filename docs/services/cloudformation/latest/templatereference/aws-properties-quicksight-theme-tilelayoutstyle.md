This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Theme TileLayoutStyle

The display options for the layout of tiles on a sheet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Gutter" : GutterStyle,
  "Margin" : MarginStyle
}

```

### YAML

```yaml

  Gutter:
    GutterStyle
  Margin:
    MarginStyle

```

## Properties

`Gutter`

The gutter settings that apply between tiles.

_Required_: No

_Type_: [GutterStyle](aws-properties-quicksight-theme-gutterstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Margin`

The margin settings that apply around the outside edge of sheets.

_Required_: No

_Type_: [MarginStyle](aws-properties-quicksight-theme-marginstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ThemeVersion

TileStyle

All content copied from https://docs.aws.amazon.com/.
