This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PluginVisualConfiguration

The plugin visual configuration. This includes the field wells, sorting options, and persisted options of the plugin visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldWells" : [ PluginVisualFieldWell, ... ],
  "SortConfiguration" : PluginVisualSortConfiguration,
  "VisualOptions" : PluginVisualOptions
}

```

### YAML

```yaml

  FieldWells:
    - PluginVisualFieldWell
  SortConfiguration:
    PluginVisualSortConfiguration
  VisualOptions:
    PluginVisualOptions

```

## Properties

`FieldWells`

The field wells configuration of the plugin visual.

_Required_: No

_Type_: Array of [PluginVisualFieldWell](aws-properties-quicksight-template-pluginvisualfieldwell.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of the plugin visual.

_Required_: No

_Type_: [PluginVisualSortConfiguration](aws-properties-quicksight-template-pluginvisualsortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualOptions`

The persisted properties of the plugin visual.

_Required_: No

_Type_: [PluginVisualOptions](aws-properties-quicksight-template-pluginvisualoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PluginVisual

PluginVisualFieldWell

All content copied from https://docs.aws.amazon.com/.
