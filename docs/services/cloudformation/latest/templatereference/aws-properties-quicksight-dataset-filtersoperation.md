This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet FiltersOperation

A transform operation that applies one or more filter conditions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "FilterOperations" : [ FilterOperation, ... ],
  "Source" : TransformOperationSource
}

```

### YAML

```yaml

  Alias: String
  FilterOperations:
    - FilterOperation
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

`FilterOperations`

The list of filter operations to apply.

_Required_: Yes

_Type_: Array of [FilterOperation](aws-properties-quicksight-dataset-filteroperation.md)

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source transform operation that provides input data for filtering.

_Required_: Yes

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterOperation

GeoSpatialColumnGroup

All content copied from https://docs.aws.amazon.com/.
