This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace WorkspaceThemeConfig

Contains detailed theme configuration for a workspace, including colors, images, and typography.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Palette" : WorkspaceThemePalette,
  "Typography" : WorkspaceThemeTypography
}

```

### YAML

```yaml

  Palette:
    WorkspaceThemePalette
  Typography:
    WorkspaceThemeTypography

```

## Properties

`Palette`

The color palette configuration for the workspace theme.

_Required_: No

_Type_: [WorkspaceThemePalette](aws-properties-connect-workspace-workspacethemepalette.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Typography`

The typography configuration for the workspace theme.

_Required_: No

_Type_: [WorkspaceThemeTypography](aws-properties-connect-workspace-workspacethemetypography.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkspaceTheme

WorkspaceThemePalette

All content copied from https://docs.aws.amazon.com/.
