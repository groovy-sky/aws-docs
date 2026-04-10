This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DefaultFilterControlOptions

The option that corresponds to the control type of the filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultDateTimePickerOptions" : DefaultDateTimePickerControlOptions,
  "DefaultDropdownOptions" : DefaultFilterDropDownControlOptions,
  "DefaultListOptions" : DefaultFilterListControlOptions,
  "DefaultRelativeDateTimeOptions" : DefaultRelativeDateTimeControlOptions,
  "DefaultSliderOptions" : DefaultSliderControlOptions,
  "DefaultTextAreaOptions" : DefaultTextAreaControlOptions,
  "DefaultTextFieldOptions" : DefaultTextFieldControlOptions
}

```

### YAML

```yaml

  DefaultDateTimePickerOptions:
    DefaultDateTimePickerControlOptions
  DefaultDropdownOptions:
    DefaultFilterDropDownControlOptions
  DefaultListOptions:
    DefaultFilterListControlOptions
  DefaultRelativeDateTimeOptions:
    DefaultRelativeDateTimeControlOptions
  DefaultSliderOptions:
    DefaultSliderControlOptions
  DefaultTextAreaOptions:
    DefaultTextAreaControlOptions
  DefaultTextFieldOptions:
    DefaultTextFieldControlOptions

```

## Properties

`DefaultDateTimePickerOptions`

The default options that correspond to the filter control type of a `DateTimePicker`.

_Required_: No

_Type_: [DefaultDateTimePickerControlOptions](aws-properties-quicksight-dashboard-defaultdatetimepickercontroloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultDropdownOptions`

The default options that correspond to the `Dropdown` filter control type.

_Required_: No

_Type_: [DefaultFilterDropDownControlOptions](aws-properties-quicksight-dashboard-defaultfilterdropdowncontroloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultListOptions`

The default options that correspond to the `List` filter control type.

_Required_: No

_Type_: [DefaultFilterListControlOptions](aws-properties-quicksight-dashboard-defaultfilterlistcontroloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultRelativeDateTimeOptions`

The default options that correspond to the `RelativeDateTime` filter control type.

_Required_: No

_Type_: [DefaultRelativeDateTimeControlOptions](aws-properties-quicksight-dashboard-defaultrelativedatetimecontroloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultSliderOptions`

The default options that correspond to the `Slider` filter control type.

_Required_: No

_Type_: [DefaultSliderControlOptions](aws-properties-quicksight-dashboard-defaultslidercontroloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultTextAreaOptions`

The default options that correspond to the `TextArea` filter control type.

_Required_: No

_Type_: [DefaultTextAreaControlOptions](aws-properties-quicksight-dashboard-defaulttextareacontroloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultTextFieldOptions`

The default options that correspond to the `TextField` filter control type.

_Required_: No

_Type_: [DefaultTextFieldControlOptions](aws-properties-quicksight-dashboard-defaulttextfieldcontroloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultFilterControlConfiguration

DefaultFilterDropDownControlOptions

All content copied from https://docs.aws.amazon.com/.
