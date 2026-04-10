This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FunnelChartSortConfiguration

The sort configuration of a `FunnelChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryItemsLimit" : ItemsLimitConfiguration,
  "CategorySort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  CategoryItemsLimit:
    ItemsLimitConfiguration
  CategorySort:
    - FieldSortOptions

```

## Properties

`CategoryItemsLimit`

The limit on the number of categories displayed.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-dashboard-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategorySort`

The sort configuration of the category fields.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-dashboard-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunnelChartFieldWells

FunnelChartVisual

All content copied from https://docs.aws.amazon.com/.
