This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TreeMapSortConfiguration

The sort configuration of a tree map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TreeMapGroupItemsLimitConfiguration" : ItemsLimitConfiguration,
  "TreeMapSort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  TreeMapGroupItemsLimitConfiguration:
    ItemsLimitConfiguration
  TreeMapSort:
    - FieldSortOptions

```

## Properties

`TreeMapGroupItemsLimitConfiguration`

The limit on the number of groups that are displayed.

_Required_: No

_Type_: [ItemsLimitConfiguration](aws-properties-quicksight-dashboard-itemslimitconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreeMapSort`

The sort configuration of group by fields.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-dashboard-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TreeMapFieldWells

TreeMapVisual

All content copied from https://docs.aws.amazon.com/.
