This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis BoxPlotOptions

The options of a box plot visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllDataPointsVisibility" : String,
  "OutlierVisibility" : String,
  "StyleOptions" : BoxPlotStyleOptions
}

```

### YAML

```yaml

  AllDataPointsVisibility: String
  OutlierVisibility: String
  StyleOptions:
    BoxPlotStyleOptions

```

## Properties

`AllDataPointsVisibility`

Determines the visibility of all data points of the box plot.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutlierVisibility`

Determines the visibility of the outlier in a box plot.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StyleOptions`

The style options of the box plot.

_Required_: No

_Type_: [BoxPlotStyleOptions](aws-properties-quicksight-analysis-boxplotstyleoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BoxPlotFieldWells

BoxPlotSortConfiguration

All content copied from https://docs.aws.amazon.com/.
