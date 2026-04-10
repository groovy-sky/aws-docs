This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard PluginVisualTableQuerySort

The table query sorting options for the plugin visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ItemsLimitConfiguration" : PluginVisualItemsLimitConfiguration,
  "RowSort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  ItemsLimitConfiguration:
    PluginVisualItemsLimitConfiguration
  RowSort:
    - FieldSortOptions

```

## Properties

`ItemsLimitConfiguration`

The maximum amount of data to be returned by a query.

_Required_: No

_Type_: [PluginVisualItemsLimitConfiguration](aws-properties-quicksight-dashboard-pluginvisualitemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowSort`

Determines how data is sorted in the response.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-dashboard-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PluginVisualSortConfiguration

PredefinedHierarchy

All content copied from https://docs.aws.amazon.com/.
