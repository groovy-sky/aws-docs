This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FilterControl

The control of a filter that is used to interact with a dashboard or an analysis.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrossSheet" : FilterCrossSheetControl,
  "DateTimePicker" : FilterDateTimePickerControl,
  "Dropdown" : FilterDropDownControl,
  "List" : FilterListControl,
  "RelativeDateTime" : FilterRelativeDateTimeControl,
  "Slider" : FilterSliderControl,
  "TextArea" : FilterTextAreaControl,
  "TextField" : FilterTextFieldControl
}

```

### YAML

```yaml

  CrossSheet:
    FilterCrossSheetControl
  DateTimePicker:
    FilterDateTimePickerControl
  Dropdown:
    FilterDropDownControl
  List:
    FilterListControl
  RelativeDateTime:
    FilterRelativeDateTimeControl
  Slider:
    FilterSliderControl
  TextArea:
    FilterTextAreaControl
  TextField:
    FilterTextFieldControl

```

## Properties

`CrossSheet`

A control from a filter that is scoped across more than one sheet. This represents your filter control on a sheet

_Required_: No

_Type_: [FilterCrossSheetControl](aws-properties-quicksight-template-filtercrosssheetcontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateTimePicker`

A control from a date filter that is used to specify date and time.

_Required_: No

_Type_: [FilterDateTimePickerControl](aws-properties-quicksight-template-filterdatetimepickercontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dropdown`

A control to display a dropdown list with buttons that are used to select a single value.

_Required_: No

_Type_: [FilterDropDownControl](aws-properties-quicksight-template-filterdropdowncontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`List`

A control to display a list of buttons or boxes. This is used to select either a single value or multiple values.

_Required_: No

_Type_: [FilterListControl](aws-properties-quicksight-template-filterlistcontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelativeDateTime`

A control from a date filter that is used to specify the relative date.

_Required_: No

_Type_: [FilterRelativeDateTimeControl](aws-properties-quicksight-template-filterrelativedatetimecontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slider`

A control to display a horizontal toggle bar. This is used to change a value by sliding the toggle.

_Required_: No

_Type_: [FilterSliderControl](aws-properties-quicksight-template-filterslidercontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextArea`

A control to display a text box that is used to enter multiple entries.

_Required_: No

_Type_: [FilterTextAreaControl](aws-properties-quicksight-template-filtertextareacontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextField`

A control to display a text box that is used to enter a single entry.

_Required_: No

_Type_: [FilterTextFieldControl](aws-properties-quicksight-template-filtertextfieldcontrol.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

FilterCrossSheetControl

All content copied from https://docs.aws.amazon.com/.
