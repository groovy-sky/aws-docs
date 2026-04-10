This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GaugeChartFieldWells

The field well configuration of a `GaugeChartVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetValues" : [ MeasureField, ... ],
  "Values" : [ MeasureField, ... ]
}

```

### YAML

```yaml

  TargetValues:
    - MeasureField
  Values:
    - MeasureField

```

## Properties

`TargetValues`

The target value field wells of a `GaugeChartVisual`.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value field wells of a `GaugeChartVisual`.

_Required_: No

_Type_: Array of [MeasureField](aws-properties-quicksight-dashboard-measurefield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartConfiguration

GaugeChartOptions

All content copied from https://docs.aws.amazon.com/.
