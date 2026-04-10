This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DataPathType

The type of the data path value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PivotTableDataPathType" : String
}

```

### YAML

```yaml

  PivotTableDataPathType: String

```

## Properties

`PivotTableDataPathType`

The type of data path value utilized in a pivot table. Choose one of the following options:

- `HIERARCHY_ROWS_LAYOUT_COLUMN` \- The type of data path for the rows layout column, when `RowsLayout` is set to `HIERARCHY`.

- `MULTIPLE_ROW_METRICS_COLUMN` \- The type of data path for the metric column when the row is set to Metric Placement.

- `EMPTY_COLUMN_HEADER` \- The type of data path for the column with empty column header, when there is no field in `ColumnsFieldWell` and the row is set to Metric Placement.

- `COUNT_METRIC_COLUMN` \- The type of data path for the column with `COUNT` as the metric, when there is no field in the `ValuesFieldWell`.

_Required_: No

_Type_: String

_Allowed values_: `HIERARCHY_ROWS_LAYOUT_COLUMN | MULTIPLE_ROW_METRICS_COLUMN | EMPTY_COLUMN_HEADER | COUNT_METRIC_COLUMN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPathSort

DataPathValue

All content copied from https://docs.aws.amazon.com/.
