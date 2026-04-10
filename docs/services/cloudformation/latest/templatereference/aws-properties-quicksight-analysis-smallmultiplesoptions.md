This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis SmallMultiplesOptions

Options that determine the layout and display options of a chart's small multiples.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxVisibleColumns" : Number,
  "MaxVisibleRows" : Number,
  "PanelConfiguration" : PanelConfiguration,
  "XAxis" : SmallMultiplesAxisProperties,
  "YAxis" : SmallMultiplesAxisProperties
}

```

### YAML

```yaml

  MaxVisibleColumns: Number
  MaxVisibleRows: Number
  PanelConfiguration:
    PanelConfiguration
  XAxis:
    SmallMultiplesAxisProperties
  YAxis:
    SmallMultiplesAxisProperties

```

## Properties

`MaxVisibleColumns`

Sets the maximum number of visible columns to display in the grid of small multiples panels.

The default is `Auto`, which automatically adjusts the columns in the grid to fit the overall layout and size of the given chart.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxVisibleRows`

Sets the maximum number of visible rows to display in the grid of small multiples panels.

The default value is `Auto`,
which automatically adjusts the rows in the grid
to fit the overall layout and size of the given chart.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PanelConfiguration`

Configures the display options for each small multiples panel.

_Required_: No

_Type_: [PanelConfiguration](aws-properties-quicksight-analysis-panelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxis`

The properties of a small multiples X axis.

_Required_: No

_Type_: [SmallMultiplesAxisProperties](aws-properties-quicksight-analysis-smallmultiplesaxisproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`YAxis`

The properties of a small multiples Y axis.

_Required_: No

_Type_: [SmallMultiplesAxisProperties](aws-properties-quicksight-analysis-smallmultiplesaxisproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SmallMultiplesAxisProperties

Spacing

All content copied from https://docs.aws.amazon.com/.
