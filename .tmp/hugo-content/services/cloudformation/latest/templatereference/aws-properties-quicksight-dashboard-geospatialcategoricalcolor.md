This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard GeospatialCategoricalColor

The definition for a categorical color.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryDataColors" : [ GeospatialCategoricalDataColor, ... ],
  "DefaultOpacity" : Number,
  "NullDataSettings" : GeospatialNullDataSettings,
  "NullDataVisibility" : String
}

```

### YAML

```yaml

  CategoryDataColors:
    - GeospatialCategoricalDataColor
  DefaultOpacity: Number
  NullDataSettings:
    GeospatialNullDataSettings
  NullDataVisibility: String

```

## Properties

`CategoryDataColors`

A list of categorical data colors for each category.

_Required_: Yes

_Type_: Array of [GeospatialCategoricalDataColor](aws-properties-quicksight-dashboard-geospatialcategoricaldatacolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOpacity`

The default opacity of a categorical color.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullDataSettings`

The null data visualization settings.

_Required_: No

_Type_: [GeospatialNullDataSettings](aws-properties-quicksight-dashboard-geospatialnulldatasettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullDataVisibility`

The state of visibility for null data.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GaugeChartVisual

GeospatialCategoricalDataColor

All content copied from https://docs.aws.amazon.com/.
