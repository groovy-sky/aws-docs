This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ParameterControl

The control of a parameter that users can interact with in a dashboard or an analysis.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimePicker" : ParameterDateTimePickerControl,
  "Dropdown" : ParameterDropDownControl,
  "List" : ParameterListControl,
  "Slider" : ParameterSliderControl,
  "TextArea" : ParameterTextAreaControl,
  "TextField" : ParameterTextFieldControl
}

```

### YAML

```yaml

  DateTimePicker:
    ParameterDateTimePickerControl
  Dropdown:
    ParameterDropDownControl
  List:
    ParameterListControl
  Slider:
    ParameterSliderControl
  TextArea:
    ParameterTextAreaControl
  TextField:
    ParameterTextFieldControl

```

## Properties

`DateTimePicker`

A control from a date parameter that specifies date and time.

_Required_: No

_Type_: [ParameterDateTimePickerControl](aws-properties-quicksight-analysis-parameterdatetimepickercontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dropdown`

A control to display a dropdown list with buttons that are used to select a single value.

_Required_: No

_Type_: [ParameterDropDownControl](aws-properties-quicksight-analysis-parameterdropdowncontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`List`

A control to display a list with buttons or boxes that are used to select either a single value or multiple values.

_Required_: No

_Type_: [ParameterListControl](aws-properties-quicksight-analysis-parameterlistcontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slider`

A control to display a horizontal toggle bar. This is used to change a value by sliding the toggle.

_Required_: No

_Type_: [ParameterSliderControl](aws-properties-quicksight-analysis-parameterslidercontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextArea`

A control to display a text box that is used to enter multiple entries.

_Required_: No

_Type_: [ParameterTextAreaControl](aws-properties-quicksight-analysis-parametertextareacontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextField`

A control to display a text box that is used to enter a single entry.

_Required_: No

_Type_: [ParameterTextFieldControl](aws-properties-quicksight-analysis-parametertextfieldcontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PanelTitleOptions

ParameterDateTimePickerControl

All content copied from https://docs.aws.amazon.com/.
