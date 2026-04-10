This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet RenameColumnsOperation

A transform operation that renames one or more columns in the dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "RenameColumnOperations" : [ RenameColumnOperation, ... ],
  "Source" : TransformOperationSource
}

```

### YAML

```yaml

  Alias: String
  RenameColumnOperations:
    - RenameColumnOperation
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

`RenameColumnOperations`

The list of column rename operations to perform, specifying old and new column names.

_Required_: Yes

_Type_: Array of [RenameColumnOperation](aws-properties-quicksight-dataset-renamecolumnoperation.md)

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source transform operation that provides input data for column renaming.

_Required_: Yes

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RenameColumnOperation

ResourcePermission

All content copied from https://docs.aws.amazon.com/.
