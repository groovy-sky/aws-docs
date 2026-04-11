This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis LineSeriesAxisDisplayOptions

The series axis configuration of a line chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AxisOptions" : AxisDisplayOptions,
  "MissingDataConfigurations" : [ MissingDataConfiguration, ... ]
}

```

### YAML

```yaml

  AxisOptions:
    AxisDisplayOptions
  MissingDataConfigurations:
    - MissingDataConfiguration

```

## Properties

`AxisOptions`

The options that determine the presentation of the line series axis.

_Required_: No

_Type_: [AxisDisplayOptions](aws-properties-quicksight-analysis-axisdisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MissingDataConfigurations`

The configuration options that determine how missing data is treated during the rendering of a line chart.

_Required_: No

_Type_: Array of [MissingDataConfiguration](aws-properties-quicksight-analysis-missingdataconfiguration.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LineChartVisual

ListControlDisplayOptions

All content copied from https://docs.aws.amazon.com/.
