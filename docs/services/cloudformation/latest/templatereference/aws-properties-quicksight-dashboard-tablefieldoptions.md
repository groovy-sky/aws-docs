This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TableFieldOptions

The field options of a table visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Order" : [ String, ... ],
  "PinnedFieldOptions" : TablePinnedFieldOptions,
  "SelectedFieldOptions" : [ TableFieldOption, ... ],
  "TransposedTableOptions" : [ TransposedTableOption, ... ]
}

```

### YAML

```yaml

  Order:
    - String
  PinnedFieldOptions:
    TablePinnedFieldOptions
  SelectedFieldOptions:
    - TableFieldOption
  TransposedTableOptions:
    - TransposedTableOption

```

## Properties

`Order`

The order of the field IDs that are configured as field options for a table visual.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `512 | 200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PinnedFieldOptions`

The settings for the pinned columns of a table visual.

_Required_: No

_Type_: [TablePinnedFieldOptions](aws-properties-quicksight-dashboard-tablepinnedfieldoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedFieldOptions`

The field options to be configured to a table.

_Required_: No

_Type_: Array of [TableFieldOption](aws-properties-quicksight-dashboard-tablefieldoption.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransposedTableOptions`

The `TableOptions` of a transposed table.

_Required_: No

_Type_: Array of [TransposedTableOption](aws-properties-quicksight-dashboard-transposedtableoption.md)

_Minimum_: `0`

_Maximum_: `10001`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableFieldOption

TableFieldURLConfiguration

All content copied from https://docs.aws.amazon.com/.
