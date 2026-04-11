This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DefaultFilterDropDownControlOptions

The default options that correspond to the `Dropdown` filter control type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommitMode" : String,
  "DisplayOptions" : DropDownControlDisplayOptions,
  "SelectableValues" : FilterSelectableValues,
  "Type" : String
}

```

### YAML

```yaml

  CommitMode: String
  DisplayOptions:
    DropDownControlDisplayOptions
  SelectableValues:
    FilterSelectableValues
  Type: String

```

## Properties

`CommitMode`

The visibility configuration of the Apply button on a `FilterDropDownControl`.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [DropDownControlDisplayOptions](aws-properties-quicksight-dashboard-dropdowncontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectableValues`

A list of selectable values that are used in a control.

_Required_: No

_Type_: [FilterSelectableValues](aws-properties-quicksight-dashboard-filterselectablevalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the `FilterDropDownControl`. Choose one of the following options:

- `MULTI_SELECT`: The user can select multiple entries from a dropdown menu.

- `SINGLE_SELECT`: The user can select a single entry from a dropdown menu.

_Required_: No

_Type_: String

_Allowed values_: `MULTI_SELECT | SINGLE_SELECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultFilterControlOptions

DefaultFilterListControlOptions

All content copied from https://docs.aws.amazon.com/.
