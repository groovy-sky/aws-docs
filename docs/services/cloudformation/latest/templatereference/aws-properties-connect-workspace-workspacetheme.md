This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace WorkspaceTheme

Contains theme configuration for a workspace, supporting both light and dark modes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dark" : WorkspaceThemeConfig,
  "Light" : WorkspaceThemeConfig
}

```

### YAML

```yaml

  Dark:
    WorkspaceThemeConfig
  Light:
    WorkspaceThemeConfig

```

## Properties

`Dark`

The theme configuration for dark mode.

_Required_: No

_Type_: [WorkspaceThemeConfig](aws-properties-connect-workspace-workspacethemeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Light`

The theme configuration for light mode.

_Required_: No

_Type_: [WorkspaceThemeConfig](aws-properties-connect-workspace-workspacethemeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkspacePage

WorkspaceThemeConfig

All content copied from https://docs.aws.amazon.com/.
