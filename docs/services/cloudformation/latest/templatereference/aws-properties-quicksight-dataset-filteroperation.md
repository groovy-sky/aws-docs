This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet FilterOperation

A transform operation that filters rows based on a condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionExpression" : String,
  "DateFilterCondition" : DataSetDateFilterCondition,
  "NumericFilterCondition" : DataSetNumericFilterCondition,
  "StringFilterCondition" : DataSetStringFilterCondition
}

```

### YAML

```yaml

  ConditionExpression: String
  DateFilterCondition:
    DataSetDateFilterCondition
  NumericFilterCondition:
    DataSetNumericFilterCondition
  StringFilterCondition:
    DataSetStringFilterCondition

```

## Properties

`ConditionExpression`

An expression that must evaluate to a Boolean value. Rows for which the expression
evaluates to true are kept in the dataset.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateFilterCondition`

A date-based filter condition within a filter operation.

_Required_: No

_Type_: [DataSetDateFilterCondition](aws-properties-quicksight-dataset-datasetdatefiltercondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericFilterCondition`

A numeric-based filter condition within a filter operation.

_Required_: No

_Type_: [DataSetNumericFilterCondition](aws-properties-quicksight-dataset-datasetnumericfiltercondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringFilterCondition`

A string-based filter condition within a filter operation.

_Required_: No

_Type_: [DataSetStringFilterCondition](aws-properties-quicksight-dataset-datasetstringfiltercondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldFolder

FiltersOperation

All content copied from https://docs.aws.amazon.com/.
