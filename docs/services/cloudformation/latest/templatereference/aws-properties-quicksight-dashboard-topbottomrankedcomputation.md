This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TopBottomRankedComputation

The top ranked and bottom ranked computation configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : DimensionField,
  "ComputationId" : String,
  "Name" : String,
  "ResultSize" : Number,
  "Type" : String,
  "Value" : MeasureField
}

```

### YAML

```yaml

  Category:
    DimensionField
  ComputationId: String
  Name: String
  ResultSize: Number
  Type: String
  Value:
    MeasureField

```

## Properties

`Category`

The category field that is used in a computation.

_Required_: No

_Type_: [DimensionField](aws-properties-quicksight-dashboard-dimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputationId`

The ID for a computation.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a computation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResultSize`

The result size of a top and bottom ranked computation.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The computation type. Choose one of the following options:

- TOP: A top ranked computation.

- BOTTOM: A bottom ranked computation.

_Required_: Yes

_Type_: String

_Allowed values_: `TOP | BOTTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value field that is used in a computation.

_Required_: No

_Type_: [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopBottomMoversComputation

TotalAggregationComputation

All content copied from https://docs.aws.amazon.com/.
