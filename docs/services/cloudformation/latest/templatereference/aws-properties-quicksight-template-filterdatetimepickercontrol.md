This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FilterDateTimePickerControl

A control from a date filter that is used to specify date and time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommitMode" : String,
  "DisplayOptions" : DateTimePickerControlDisplayOptions,
  "FilterControlId" : String,
  "SourceFilterId" : String,
  "Title" : String,
  "Type" : String
}

```

### YAML

```yaml

  CommitMode: String
  DisplayOptions:
    DateTimePickerControlDisplayOptions
  FilterControlId: String
  SourceFilterId: String
  Title: String
  Type: String

```

## Properties

`CommitMode`

The visibility configurationof the Apply button on a `DateTimePickerControl`.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [DateTimePickerControlDisplayOptions](aws-properties-quicksight-template-datetimepickercontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterControlId`

The ID of the `FilterDateTimePickerControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFilterId`

The source filter ID of the `FilterDateTimePickerControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `FilterDateTimePickerControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the `FilterDropDownControl`. Choose one of the following options:

- `MULTI_SELECT`: The user can select multiple entries from a dropdown menu.

- `SINGLE_SELECT`: The user can select a single entry from a dropdown menu.

_Required_: No

_Type_: String

_Allowed values_: `SINGLE_VALUED | DATE_RANGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterCrossSheetControl

FilterDropDownControl

All content copied from https://docs.aws.amazon.com/.
