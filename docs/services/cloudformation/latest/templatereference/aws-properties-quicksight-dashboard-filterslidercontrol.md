This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FilterSliderControl

A control to display a horizontal toggle bar. This is used to change a value by sliding the toggle.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisplayOptions" : SliderControlDisplayOptions,
  "FilterControlId" : String,
  "MaximumValue" : Number,
  "MinimumValue" : Number,
  "SourceFilterId" : String,
  "StepSize" : Number,
  "Title" : String,
  "Type" : String
}

```

### YAML

```yaml

  DisplayOptions:
    SliderControlDisplayOptions
  FilterControlId: String
  MaximumValue: Number
  MinimumValue: Number
  SourceFilterId: String
  StepSize: Number
  Title: String
  Type: String

```

## Properties

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [SliderControlDisplayOptions](aws-properties-quicksight-dashboard-slidercontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterControlId`

The ID of the `FilterSliderControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

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

`SourceFilterId`

The source filter ID of the `FilterSliderControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StepSize`

The number of increments that the slider bar is divided into.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `FilterSliderControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the `FilterSliderControl`. Choose one of the following options:

- `SINGLE_POINT`: Filter against(equals) a single data point.

- `RANGE`: Filter data that is in a specified range.

_Required_: No

_Type_: String

_Allowed values_: `SINGLE_POINT | RANGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterSelectableValues

FilterTextAreaControl

All content copied from https://docs.aws.amazon.com/.
