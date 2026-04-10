This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetDateRangeFilterCondition

A filter condition that filters date values within a specified range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncludeMaximum" : Boolean,
  "IncludeMinimum" : Boolean,
  "RangeMaximum" : DataSetDateFilterValue,
  "RangeMinimum" : DataSetDateFilterValue
}

```

### YAML

```yaml

  IncludeMaximum: Boolean
  IncludeMinimum: Boolean
  RangeMaximum:
    DataSetDateFilterValue
  RangeMinimum:
    DataSetDateFilterValue

```

## Properties

`IncludeMaximum`

Whether to include the maximum value in the filter range.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeMinimum`

Whether to include the minimum value in the filter range.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeMaximum`

The maximum date value for the range filter.

_Required_: No

_Type_: [DataSetDateFilterValue](aws-properties-quicksight-dataset-datasetdatefiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeMinimum`

The minimum date value for the range filter.

_Required_: No

_Type_: [DataSetDateFilterValue](aws-properties-quicksight-dataset-datasetdatefiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetDateFilterValue

DataSetNumericComparisonFilterCondition

All content copied from https://docs.aws.amazon.com/.
