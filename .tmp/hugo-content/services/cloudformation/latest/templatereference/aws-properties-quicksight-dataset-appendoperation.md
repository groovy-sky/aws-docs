This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet AppendOperation

A transform operation that combines rows from two data sources by stacking them vertically (union operation).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "AppendedColumns" : [ AppendedColumn, ... ],
  "FirstSource" : TransformOperationSource,
  "SecondSource" : TransformOperationSource
}

```

### YAML

```yaml

  Alias: String
  AppendedColumns:
    - AppendedColumn
  FirstSource:
    TransformOperationSource
  SecondSource:
    TransformOperationSource

```

## Properties

`Alias`

Alias for this operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppendedColumns`

The list of columns to include in the appended result, mapping columns from both sources.

_Required_: Yes

_Type_: Array of [AppendedColumn](aws-properties-quicksight-dataset-appendedcolumn.md)

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirstSource`

The first data source to be included in the append operation.

_Required_: No

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondSource`

The second data source to be appended to the first source.

_Required_: No

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppendedColumn

CalculatedColumn

All content copied from https://docs.aws.amazon.com/.
