This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DonutOptions

The options for configuring a donut chart or pie chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArcOptions" : ArcOptions,
  "DonutCenterOptions" : DonutCenterOptions
}

```

### YAML

```yaml

  ArcOptions:
    ArcOptions
  DonutCenterOptions:
    DonutCenterOptions

```

## Properties

`ArcOptions`

The option for define the arc of the chart shape. Valid values are as follows:

- `WHOLE` \- A pie chart

- `SMALL`\- A small-sized donut chart

- `MEDIUM`\- A medium-sized donut chart

- `LARGE`\- A large-sized donut chart

_Required_: No

_Type_: [ArcOptions](aws-properties-quicksight-dashboard-arcoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DonutCenterOptions`

The label options of the label that is displayed in the center of a donut chart. This option isn't available for pie charts.

_Required_: No

_Type_: [DonutCenterOptions](aws-properties-quicksight-dashboard-donutcenteroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DonutCenterOptions

DrillDownFilter

All content copied from https://docs.aws.amazon.com/.
