This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetNumericRangeFilterCondition

A filter condition that filters numeric values within a specified range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncludeMaximum" : Boolean,
  "IncludeMinimum" : Boolean,
  "RangeMaximum" : DataSetNumericFilterValue,
  "RangeMinimum" : DataSetNumericFilterValue
}

```

### YAML

```yaml

  IncludeMaximum: Boolean
  IncludeMinimum: Boolean
  RangeMaximum:
    DataSetNumericFilterValue
  RangeMinimum:
    DataSetNumericFilterValue

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

The maximum numeric value for the range filter.

_Required_: No

_Type_: [DataSetNumericFilterValue](aws-properties-quicksight-dataset-datasetnumericfiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeMinimum`

The minimum numeric value for the range filter.

_Required_: No

_Type_: [DataSetNumericFilterValue](aws-properties-quicksight-dataset-datasetnumericfiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetNumericFilterValue

DatasetParameter

All content copied from https://docs.aws.amazon.com/.
