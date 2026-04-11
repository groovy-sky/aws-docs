This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetDateFilterCondition

A filter condition for date columns, supporting both comparison and range-based filtering.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "ComparisonFilterCondition" : DataSetDateComparisonFilterCondition,
  "RangeFilterCondition" : DataSetDateRangeFilterCondition
}

```

### YAML

```yaml

  ColumnName: String
  ComparisonFilterCondition:
    DataSetDateComparisonFilterCondition
  RangeFilterCondition:
    DataSetDateRangeFilterCondition

```

## Properties

`ColumnName`

The name of the date column to filter.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComparisonFilterCondition`

A comparison-based filter condition for the date column.

_Required_: No

_Type_: [DataSetDateComparisonFilterCondition](aws-properties-quicksight-dataset-datasetdatecomparisonfiltercondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeFilterCondition`

A range-based filter condition for the date column, filtering values between minimum and maximum dates.

_Required_: No

_Type_: [DataSetDateRangeFilterCondition](aws-properties-quicksight-dataset-datasetdaterangefiltercondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetDateComparisonFilterCondition

DataSetDateFilterValue

All content copied from https://docs.aws.amazon.com/.
