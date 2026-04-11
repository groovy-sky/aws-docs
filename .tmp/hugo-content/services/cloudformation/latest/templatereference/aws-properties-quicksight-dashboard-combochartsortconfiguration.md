This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ComboChartSortConfiguration

The sort configuration of a `ComboChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryItemsLimit" : ItemsLimitConfiguration,
  "CategorySort" : [ FieldSortOptions, ... ],
  "ColorItemsLimit" : ItemsLimitConfiguration,
  "ColorSort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  CategoryItemsLimit:
    ItemsLimitConfiguration
  CategorySort:
    - FieldSortOptions
  ColorItemsLimit:
    ItemsLimitConfiguration
  ColorSort:
    - FieldSortOptions

```

## Properties

`CategoryItemsLimit`

The item limit configuration for the category field well of a combo chart.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-dashboard-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategorySort`

The sort configuration of the category field well in a combo chart.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-dashboard-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorItemsLimit`

The item limit configuration of the color field well in a combo chart.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-dashboard-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorSort`

The sort configuration of the color field well in a combo chart.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-dashboard-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComboChartFieldWells

ComboChartVisual

All content copied from https://docs.aws.amazon.com/.
