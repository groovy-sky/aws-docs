This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DefaultDateTimePickerControlOptions

The default options that correspond to the filter control type of a `DateTimePicker`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommitMode" : String,
  "DisplayOptions" : DateTimePickerControlDisplayOptions,
  "Type" : String
}

```

### YAML

```yaml

  CommitMode: String
  DisplayOptions:
    DateTimePickerControlDisplayOptions
  Type: String

```

## Properties

`CommitMode`

The visibility configuration of the Apply button on a `DateTimePickerControl`.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [DateTimePickerControlDisplayOptions](aws-properties-quicksight-analysis-datetimepickercontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The date time picker type of the `DefaultDateTimePickerControlOptions`. Choose one of the following options:

- `SINGLE_VALUED`: The filter condition is a fixed date.

- `DATE_RANGE`: The filter condition is a date time range.

_Required_: No

_Type_: String

_Allowed values_: `SINGLE_VALUED | DATE_RANGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecimalValueWhenUnsetConfiguration

DefaultFilterControlConfiguration

All content copied from https://docs.aws.amazon.com/.
