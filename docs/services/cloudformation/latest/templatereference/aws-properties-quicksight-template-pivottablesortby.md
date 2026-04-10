This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template PivotTableSortBy

The sort by field for the field sort options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnSort,
  "DataPath" : DataPathSort,
  "Field" : FieldSort
}

```

### YAML

```yaml

  Column:
    ColumnSort
  DataPath:
    DataPathSort
  Field:
    FieldSort

```

## Properties

`Column`

The column sort (field id, direction) for the pivot table sort by options.

_Required_: No

_Type_: [ColumnSort](aws-properties-quicksight-template-columnsort.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataPath`

The data path sort (data path value, direction) for the pivot table sort by options.

_Required_: No

_Type_: [DataPathSort](aws-properties-quicksight-template-datapathsort.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Field`

The field sort (field id, direction) for the pivot table sort by options.

_Required_: No

_Type_: [FieldSort](aws-properties-quicksight-template-fieldsort.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableRowsLabelOptions

PivotTableSortConfiguration

All content copied from https://docs.aws.amazon.com/.
