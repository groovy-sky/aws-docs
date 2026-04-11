This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis HeatMapSortConfiguration

The sort configuration of a heat map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeatMapColumnItemsLimitConfiguration" : ItemsLimitConfiguration,
  "HeatMapColumnSort" : [ FieldSortOptions, ... ],
  "HeatMapRowItemsLimitConfiguration" : ItemsLimitConfiguration,
  "HeatMapRowSort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  HeatMapColumnItemsLimitConfiguration:
    ItemsLimitConfiguration
  HeatMapColumnSort:
    - FieldSortOptions
  HeatMapRowItemsLimitConfiguration:
    ItemsLimitConfiguration
  HeatMapRowSort:
    - FieldSortOptions

```

## Properties

`HeatMapColumnItemsLimitConfiguration`

The limit on the number of columns that are displayed in a heat map.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-analysis-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeatMapColumnSort`

The column sort configuration for heat map for columns that aren't a part of a field well.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-analysis-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeatMapRowItemsLimitConfiguration`

The limit on the number of rows that are displayed in a heat map.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-analysis-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeatMapRowSort`

The field sort configuration of the rows fields.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-analysis-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeatMapFieldWells

HeatMapVisual

All content copied from https://docs.aws.amazon.com/.
