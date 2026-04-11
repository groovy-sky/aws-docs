This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template BarChartSortConfiguration

sort-configuration-description

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryItemsLimit" : ItemsLimitConfiguration,
  "CategorySort" : [ FieldSortOptions, ... ],
  "ColorItemsLimit" : ItemsLimitConfiguration,
  "ColorSort" : [ FieldSortOptions, ... ],
  "SmallMultiplesLimitConfiguration" : ItemsLimitConfiguration,
  "SmallMultiplesSort" : [ FieldSortOptions, ... ]
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
  SmallMultiplesLimitConfiguration:
    ItemsLimitConfiguration
  SmallMultiplesSort:
    - FieldSortOptions

```

## Properties

`CategoryItemsLimit`

The limit on the number of categories displayed in a bar chart.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-template-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategorySort`

The sort configuration of category fields.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-template-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorItemsLimit`

The limit on the number of values displayed in a bar chart.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-template-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorSort`

The sort configuration of color fields in a bar chart.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-template-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmallMultiplesLimitConfiguration`

The limit on the number of small multiples panels that are displayed.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-template-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmallMultiplesSort`

The sort configuration of the small multiples field.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-template-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BarChartFieldWells

BarChartVisual

All content copied from https://docs.aws.amazon.com/.
