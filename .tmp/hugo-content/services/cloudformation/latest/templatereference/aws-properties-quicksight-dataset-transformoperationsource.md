This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet TransformOperationSource

Specifies the source of data for a transform operation, including the source operation and column mappings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnIdMappings" : [ DataSetColumnIdMapping, ... ],
  "TransformOperationId" : String
}

```

### YAML

```yaml

  ColumnIdMappings:
    - DataSetColumnIdMapping
  TransformOperationId: String

```

## Properties

`ColumnIdMappings`

The mappings between source column identifiers and target column identifiers for this transformation.

_Required_: No

_Type_: Array of [DataSetColumnIdMapping](aws-properties-quicksight-dataset-datasetcolumnidmapping.md)

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformOperationId`

The identifier of the transform operation that provides input data.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z-]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransformOperation

TransformStep

All content copied from https://docs.aws.amazon.com/.
