This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TopBottomMoversComputation

The top movers and bottom movers computation setup.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : DimensionField,
  "ComputationId" : String,
  "MoverSize" : Number,
  "Name" : String,
  "SortOrder" : String,
  "Time" : DimensionField,
  "Type" : String,
  "Value" : MeasureField
}

```

### YAML

```yaml

  Category:
    DimensionField
  ComputationId: String
  MoverSize: Number
  Name: String
  SortOrder: String
  Time:
    DimensionField
  Type: String
  Value:
    MeasureField

```

## Properties

`Category`

The category field that is used in a computation.

_Required_: No

_Type_: [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputationId`

The ID for a computation.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MoverSize`

The mover size setup of the top and bottom movers computation.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a computation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortOrder`

The sort order setup of the top and bottom movers computation.

_Required_: No

_Type_: String

_Allowed values_: `PERCENT_DIFFERENCE | ABSOLUTE_DIFFERENCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Time`

The time field that is used in a computation.

_Required_: No

_Type_: [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The computation type. Choose from the following options:

- TOP: Top movers computation.

- BOTTOM: Bottom movers computation.

_Required_: Yes

_Type_: String

_Allowed values_: `TOP | BOTTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value field that is used in a computation.

_Required_: No

_Type_: [MeasureField](aws-properties-quicksight-analysis-measurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopBottomFilter

TopBottomRankedComputation

All content copied from https://docs.aws.amazon.com/.
