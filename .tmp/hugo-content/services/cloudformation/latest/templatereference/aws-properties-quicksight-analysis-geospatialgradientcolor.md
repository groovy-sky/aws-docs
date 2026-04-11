This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GeospatialGradientColor

The definition for a gradient color.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultOpacity" : Number,
  "NullDataSettings" : GeospatialNullDataSettings,
  "NullDataVisibility" : String,
  "StepColors" : [ GeospatialGradientStepColor, ... ]
}

```

### YAML

```yaml

  DefaultOpacity: Number
  NullDataSettings:
    GeospatialNullDataSettings
  NullDataVisibility: String
  StepColors:
    - GeospatialGradientStepColor

```

## Properties

`DefaultOpacity`

The default opacity for the gradient color.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullDataSettings`

The null data visualization settings.

_Required_: No

_Type_: [GeospatialNullDataSettings](aws-properties-quicksight-analysis-geospatialnulldatasettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullDataVisibility`

The state of visibility for null data.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StepColors`

A list of gradient step colors for the gradient.

_Required_: Yes

_Type_: Array of [GeospatialGradientStepColor](aws-properties-quicksight-analysis-geospatialgradientstepcolor.md)

_Minimum_: `2`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeospatialDataSourceItem

GeospatialGradientStepColor

All content copied from https://docs.aws.amazon.com/.
