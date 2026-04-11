This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template MetricComparisonComputation

The metric comparison computation configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputationId" : String,
  "FromValue" : MeasureField,
  "Name" : String,
  "TargetValue" : MeasureField,
  "Time" : DimensionField
}

```

### YAML

```yaml

  ComputationId: String
  FromValue:
    MeasureField
  Name: String
  TargetValue:
    MeasureField
  Time:
    DimensionField

```

## Properties

`ComputationId`

The ID for a computation.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FromValue`

The field that is used in a metric comparison from value setup.

_Required_: No

_Type_: [MeasureField](aws-properties-quicksight-template-measurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a computation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetValue`

The field that is used in a metric comparison to value setup.

_Required_: No

_Type_: [MeasureField](aws-properties-quicksight-template-measurefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Time`

The time field that is used in a computation.

_Required_: No

_Type_: [DimensionField](aws-properties-quicksight-template-dimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MeasureField

MinimumLabelType

All content copied from https://docs.aws.amazon.com/.
