This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet CreateColumnsOperation

A transform operation that creates calculated columns. Columns created in one such
operation form a lexical closure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "Columns" : [ CalculatedColumn, ... ],
  "Source" : TransformOperationSource
}

```

### YAML

```yaml

  Alias: String
  Columns:
    - CalculatedColumn
  Source:
    TransformOperationSource

```

## Properties

`Alias`

Alias for this operation.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Columns`

Calculated columns to create.

_Required_: Yes

_Type_: Array of [CalculatedColumn](aws-properties-quicksight-dataset-calculatedcolumn.md)

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source transform operation that provides input data for creating new calculated columns.

_Required_: No

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnToUnpivot

CustomSql

All content copied from https://docs.aws.amazon.com/.
