This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet CastColumnTypesOperation

A transform operation that changes the data types of one or more columns in the dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "CastColumnTypeOperations" : [ CastColumnTypeOperation, ... ],
  "Source" : TransformOperationSource
}

```

### YAML

```yaml

  Alias: String
  CastColumnTypeOperations:
    - CastColumnTypeOperation
  Source:
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

`CastColumnTypeOperations`

The list of column type casting operations to perform.

_Required_: Yes

_Type_: Array of [CastColumnTypeOperation](aws-properties-quicksight-dataset-castcolumntypeoperation.md)

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source transform operation that provides input data for the type casting.

_Required_: Yes

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CastColumnTypeOperation

ColumnDescription

All content copied from https://docs.aws.amazon.com/.
