This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace WorkspaceThemePalette

Contains color palette configuration for different areas of a workspace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Canvas" : PaletteCanvas,
  "Header" : PaletteHeader,
  "Navigation" : PaletteNavigation,
  "Primary" : PalettePrimary
}

```

### YAML

```yaml

  Canvas:
    PaletteCanvas
  Header:
    PaletteHeader
  Navigation:
    PaletteNavigation
  Primary:
    PalettePrimary

```

## Properties

`Canvas`

The color configuration for the canvas area.

_Required_: No

_Type_: [PaletteCanvas](aws-properties-connect-workspace-palettecanvas.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Header`

The color configuration for the header area.

_Required_: No

_Type_: [PaletteHeader](aws-properties-connect-workspace-paletteheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Navigation`

The color configuration for the navigation area.

_Required_: No

_Type_: [PaletteNavigation](aws-properties-connect-workspace-palettenavigation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Primary`

The primary color configuration used throughout the workspace.

_Required_: No

_Type_: [PalettePrimary](aws-properties-connect-workspace-paletteprimary.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkspaceThemeConfig

WorkspaceThemeTypography

All content copied from https://docs.aws.amazon.com/.
