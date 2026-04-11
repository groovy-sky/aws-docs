This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Theme SheetStyle

The theme display options for sheets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Tile" : TileStyle,
  "TileLayout" : TileLayoutStyle
}

```

### YAML

```yaml

  Tile:
    TileStyle
  TileLayout:
    TileLayoutStyle

```

## Properties

`Tile`

The display options for tiles.

_Required_: No

_Type_: [TileStyle](aws-properties-quicksight-theme-tilestyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TileLayout`

The layout options for tiles.

_Required_: No

_Type_: [TileLayoutStyle](aws-properties-quicksight-theme-tilelayoutstyle.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourcePermission

Tag

All content copied from https://docs.aws.amazon.com/.
