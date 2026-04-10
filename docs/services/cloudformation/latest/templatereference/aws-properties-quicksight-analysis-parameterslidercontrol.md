This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ParameterSliderControl

A control to display a horizontal toggle bar. This is used to change a value by sliding the toggle.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisplayOptions" : SliderControlDisplayOptions,
  "MaximumValue" : Number,
  "MinimumValue" : Number,
  "ParameterControlId" : String,
  "SourceParameterName" : String,
  "StepSize" : Number,
  "Title" : String
}

```

### YAML

```yaml

  DisplayOptions:
    SliderControlDisplayOptions
  MaximumValue: Number
  MinimumValue: Number
  ParameterControlId: String
  SourceParameterName: String
  StepSize: Number
  Title: String

```

## Properties

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [SliderControlDisplayOptions](aws-properties-quicksight-analysis-slidercontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumValue`

The larger value that is displayed at the right of the slider.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumValue`

The smaller value that is displayed at the left of the slider.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterControlId`

The ID of the `ParameterSliderControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceParameterName`

The source parameter name of the `ParameterSliderControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StepSize`

The number of increments that the slider bar is divided into.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `ParameterSliderControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterSelectableValues

ParameterTextAreaControl

All content copied from https://docs.aws.amazon.com/.
