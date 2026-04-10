This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet TransformStep

A step in data preparation that performs a specific operation on the data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregateStep" : AggregateOperation,
  "AppendStep" : AppendOperation,
  "CastColumnTypesStep" : CastColumnTypesOperation,
  "CreateColumnsStep" : CreateColumnsOperation,
  "FiltersStep" : FiltersOperation,
  "ImportTableStep" : ImportTableOperation,
  "JoinStep" : JoinOperation,
  "PivotStep" : PivotOperation,
  "ProjectStep" : ProjectOperation,
  "RenameColumnsStep" : RenameColumnsOperation,
  "UnpivotStep" : UnpivotOperation
}

```

### YAML

```yaml

  AggregateStep:
    AggregateOperation
  AppendStep:
    AppendOperation
  CastColumnTypesStep:
    CastColumnTypesOperation
  CreateColumnsStep:
    CreateColumnsOperation
  FiltersStep:
    FiltersOperation
  ImportTableStep:
    ImportTableOperation
  JoinStep:
    JoinOperation
  PivotStep:
    PivotOperation
  ProjectStep:
    ProjectOperation
  RenameColumnsStep:
    RenameColumnsOperation
  UnpivotStep:
    UnpivotOperation

```

## Properties

`AggregateStep`

A transform step that groups data and applies aggregation functions to calculate summary values.

_Required_: No

_Type_: [AggregateOperation](aws-properties-quicksight-dataset-aggregateoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppendStep`

A transform step that combines rows from multiple sources by stacking them vertically.

_Required_: No

_Type_: [AppendOperation](aws-properties-quicksight-dataset-appendoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CastColumnTypesStep`

A transform step that changes the data types of one or more columns.

_Required_: No

_Type_: [CastColumnTypesOperation](aws-properties-quicksight-dataset-castcolumntypesoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateColumnsStep`

Property description not available.

_Required_: No

_Type_: [CreateColumnsOperation](aws-properties-quicksight-dataset-createcolumnsoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FiltersStep`

A transform step that applies filter conditions.

_Required_: No

_Type_: [FiltersOperation](aws-properties-quicksight-dataset-filtersoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImportTableStep`

A transform step that brings data from a source table.

_Required_: No

_Type_: [ImportTableOperation](aws-properties-quicksight-dataset-importtableoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JoinStep`

A transform step that combines data from two sources based on specified join conditions.

_Required_: No

_Type_: [JoinOperation](aws-properties-quicksight-dataset-joinoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PivotStep`

A transform step that converts row values into columns to reshape the data structure.

_Required_: No

_Type_: [PivotOperation](aws-properties-quicksight-dataset-pivotoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectStep`

Property description not available.

_Required_: No

_Type_: [ProjectOperation](aws-properties-quicksight-dataset-projectoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenameColumnsStep`

A transform step that changes the names of one or more columns.

_Required_: No

_Type_: [RenameColumnsOperation](aws-properties-quicksight-dataset-renamecolumnsoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnpivotStep`

A transform step that converts columns into rows to normalize the data structure.

_Required_: No

_Type_: [UnpivotOperation](aws-properties-quicksight-dataset-unpivotoperation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransformOperationSource

UniqueKey

All content copied from https://docs.aws.amazon.com/.
