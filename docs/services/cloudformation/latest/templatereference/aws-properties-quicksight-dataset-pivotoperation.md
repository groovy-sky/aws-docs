This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet PivotOperation

A transform operation that pivots data by converting row values into columns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "GroupByColumnNames" : [ String, ... ],
  "PivotConfiguration" : PivotConfiguration,
  "Source" : TransformOperationSource,
  "ValueColumnConfiguration" : ValueColumnConfiguration
}

```

### YAML

```yaml

  Alias: String
  GroupByColumnNames:
    - String
  PivotConfiguration:
    PivotConfiguration
  Source:
    TransformOperationSource
  ValueColumnConfiguration:
    ValueColumnConfiguration

```

## Properties

`Alias`

Alias for this operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupByColumnNames`

The list of column names to group by when performing the pivot operation.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `127 | 128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PivotConfiguration`

Configuration that specifies which labels to pivot and how to structure the resulting columns.

_Required_: Yes

_Type_: [PivotConfiguration](aws-properties-quicksight-dataset-pivotconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source transform operation that provides input data for pivoting.

_Required_: Yes

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueColumnConfiguration`

Configuration for how to aggregate values when multiple rows map to the same pivoted column.

_Required_: Yes

_Type_: [ValueColumnConfiguration](aws-properties-quicksight-dataset-valuecolumnconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotedLabel

ProjectOperation

All content copied from https://docs.aws.amazon.com/.
