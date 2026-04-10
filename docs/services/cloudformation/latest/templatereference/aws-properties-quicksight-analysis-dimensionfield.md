This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DimensionField

The dimension type field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoricalDimensionField" : CategoricalDimensionField,
  "DateDimensionField" : DateDimensionField,
  "NumericalDimensionField" : NumericalDimensionField
}

```

### YAML

```yaml

  CategoricalDimensionField:
    CategoricalDimensionField
  DateDimensionField:
    DateDimensionField
  NumericalDimensionField:
    NumericalDimensionField

```

## Properties

`CategoricalDimensionField`

The dimension type field with categorical type columns.

_Required_: No

_Type_: [CategoricalDimensionField](aws-properties-quicksight-analysis-categoricaldimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateDimensionField`

The dimension type field with date type columns.

_Required_: No

_Type_: [DateDimensionField](aws-properties-quicksight-analysis-datedimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericalDimensionField`

The dimension type field with numerical type columns.

_Required_: No

_Type_: [NumericalDimensionField](aws-properties-quicksight-analysis-numericaldimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationParameterValueConfiguration

DonutCenterOptions

All content copied from https://docs.aws.amazon.com/.
