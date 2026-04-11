This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet JoinOperation

A transform operation that combines data from two sources based on specified join conditions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "LeftOperand" : TransformOperationSource,
  "LeftOperandProperties" : JoinOperandProperties,
  "OnClause" : String,
  "RightOperand" : TransformOperationSource,
  "RightOperandProperties" : JoinOperandProperties,
  "Type" : String
}

```

### YAML

```yaml

  Alias: String
  LeftOperand:
    TransformOperationSource
  LeftOperandProperties:
    JoinOperandProperties
  OnClause: String
  RightOperand:
    TransformOperationSource
  RightOperandProperties:
    JoinOperandProperties
  Type: String

```

## Properties

`Alias`

Alias for this operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LeftOperand`

The left operand for the join operation.

_Required_: Yes

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LeftOperandProperties`

Properties that control how the left operand's columns are handled in the join result.

_Required_: No

_Type_: [JoinOperandProperties](aws-properties-quicksight-dataset-joinoperandproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnClause`

The join condition that specifies how to match rows between the left and right operands.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RightOperand`

The right operand for the join operation.

_Required_: Yes

_Type_: [TransformOperationSource](aws-properties-quicksight-dataset-transformoperationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RightOperandProperties`

Properties that control how the right operand's columns are handled in the join result.

_Required_: No

_Type_: [JoinOperandProperties](aws-properties-quicksight-dataset-joinoperandproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of join to perform, such as `INNER`, `LEFT`, `RIGHT`, or `OUTER`.

_Required_: Yes

_Type_: String

_Allowed values_: `INNER | OUTER | LEFT | RIGHT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JoinOperandProperties

LogicalTableSource

All content copied from https://docs.aws.amazon.com/.
