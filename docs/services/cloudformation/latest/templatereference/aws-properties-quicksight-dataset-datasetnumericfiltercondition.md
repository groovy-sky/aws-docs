This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetNumericFilterCondition

A filter condition for numeric columns, supporting both comparison and range-based filtering.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "ComparisonFilterCondition" : DataSetNumericComparisonFilterCondition,
  "RangeFilterCondition" : DataSetNumericRangeFilterCondition
}

```

### YAML

```yaml

  ColumnName: String
  ComparisonFilterCondition:
    DataSetNumericComparisonFilterCondition
  RangeFilterCondition:
    DataSetNumericRangeFilterCondition

```

## Properties

`ColumnName`

The name of the numeric column to filter.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComparisonFilterCondition`

A comparison-based filter condition for the numeric column.

_Required_: No

_Type_: [DataSetNumericComparisonFilterCondition](aws-properties-quicksight-dataset-datasetnumericcomparisonfiltercondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeFilterCondition`

A range-based filter condition for the numeric column, filtering values between minimum and maximum numbers.

_Required_: No

_Type_: [DataSetNumericRangeFilterCondition](aws-properties-quicksight-dataset-datasetnumericrangefiltercondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetNumericComparisonFilterCondition

DataSetNumericFilterValue

All content copied from https://docs.aws.amazon.com/.
