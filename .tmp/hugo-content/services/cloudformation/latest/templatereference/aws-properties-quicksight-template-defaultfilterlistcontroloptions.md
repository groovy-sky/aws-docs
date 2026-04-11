This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DefaultFilterListControlOptions

The default options that correspond to the `List` filter control type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisplayOptions" : ListControlDisplayOptions,
  "SelectableValues" : FilterSelectableValues,
  "Type" : String
}

```

### YAML

```yaml

  DisplayOptions:
    ListControlDisplayOptions
  SelectableValues:
    FilterSelectableValues
  Type: String

```

## Properties

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [ListControlDisplayOptions](aws-properties-quicksight-template-listcontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectableValues`

A list of selectable values that are used in a control.

_Required_: No

_Type_: [FilterSelectableValues](aws-properties-quicksight-template-filterselectablevalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the `DefaultFilterListControlOptions`. Choose one of the following options:

- `MULTI_SELECT`: The user can select multiple entries from the list.

- `SINGLE_SELECT`: The user can select a single entry from the list.

_Required_: No

_Type_: String

_Allowed values_: `MULTI_SELECT | SINGLE_SELECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultFilterDropDownControlOptions

DefaultFreeFormLayoutConfiguration

All content copied from https://docs.aws.amazon.com/.
