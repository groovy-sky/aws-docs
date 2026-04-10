This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetNumericComparisonFilterCondition

A filter condition that compares numeric values using operators like `EQUALS`, `GREATER_THAN`,
or `LESS_THAN`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operator" : String,
  "Value" : DataSetNumericFilterValue
}

```

### YAML

```yaml

  Operator: String
  Value:
    DataSetNumericFilterValue

```

## Properties

`Operator`

The comparison operator to use, such as `EQUALS`, `GREATER_THAN`, `LESS_THAN`,
or their variants.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | DOES_NOT_EQUAL | GREATER_THAN | GREATER_THAN_OR_EQUALS_TO | LESS_THAN | LESS_THAN_OR_EQUALS_TO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The numeric value to compare against.

_Required_: No

_Type_: [DataSetNumericFilterValue](aws-properties-quicksight-dataset-datasetnumericfiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetDateRangeFilterCondition

DataSetNumericFilterCondition

All content copied from https://docs.aws.amazon.com/.
