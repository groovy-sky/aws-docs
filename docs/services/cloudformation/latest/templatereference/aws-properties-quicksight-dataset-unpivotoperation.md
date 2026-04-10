This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet UnpivotOperation

A transform operation that converts columns into rows, normalizing the data structure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "ColumnsToUnpivot" : [ ColumnToUnpivot, ... ],
  "Source" : TransformOperationSource,
  "UnpivotedLabelColumnId" : String,
  "UnpivotedLabelColumnName" : String,
  "UnpivotedValueColumnId" : String,
  "UnpivotedValueColumnName" : String
}

```

### YAML

```yaml

  Alias: String
  ColumnsToUnpivot:
    - ColumnToUnpivot
  Source:
    TransformOperationSource
  UnpivotedLabelColumnId: String
  UnpivotedLabelColumnName: String
  UnpivotedValueColumnId: String
  UnpivotedValueColumnName: String

```

## Properties

`Alias`

Alias for this operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnsToUnpivot`

The list of columns to unpivot from the source data.

_Required_: Yes

_Type_: Array of [ColumnToUnpivot](aws-properties-quicksight-dataset-columntounpivot.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source transform operation that provides input data for unpivoting.

_Required_: Yes

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnpivotedLabelColumnId`

A unique identifier for the new column that will contain the unpivoted column names.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnpivotedLabelColumnName`

The name for the new column that will contain the unpivoted column names.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnpivotedValueColumnId`

A unique identifier for the new column that will contain the unpivoted values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnpivotedValueColumnName`

The name for the new column that will contain the unpivoted values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UniqueKey

UntagColumnOperation

All content copied from https://docs.aws.amazon.com/.
