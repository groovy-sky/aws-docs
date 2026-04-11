This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DefaultInteractiveLayoutConfiguration

The options that determine the default settings for interactive layout configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FreeForm" : DefaultFreeFormLayoutConfiguration,
  "Grid" : DefaultGridLayoutConfiguration
}

```

### YAML

```yaml

  FreeForm:
    DefaultFreeFormLayoutConfiguration
  Grid:
    DefaultGridLayoutConfiguration

```

## Properties

`FreeForm`

The options that determine the default settings of a free-form layout configuration.

_Required_: No

_Type_: [DefaultFreeFormLayoutConfiguration](aws-properties-quicksight-template-defaultfreeformlayoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Grid`

The options that determine the default settings for a grid layout configuration.

_Required_: No

_Type_: [DefaultGridLayoutConfiguration](aws-properties-quicksight-template-defaultgridlayoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultGridLayoutConfiguration

DefaultNewSheetConfiguration

All content copied from https://docs.aws.amazon.com/.
