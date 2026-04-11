This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis AxisDisplayRange

The range setup of a numeric axis display range.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataDriven" : Json,
  "MinMax" : AxisDisplayMinMaxRange
}

```

### YAML

```yaml

  DataDriven: Json
  MinMax:
    AxisDisplayMinMaxRange

```

## Properties

`DataDriven`

The data-driven setup of an axis display range.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinMax`

The minimum and maximum setup of an axis display range.

_Required_: No

_Type_: [AxisDisplayMinMaxRange](aws-properties-quicksight-analysis-axisdisplayminmaxrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AxisDisplayOptions

AxisLabelOptions

All content copied from https://docs.aws.amazon.com/.
