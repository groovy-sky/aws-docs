This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis FilterListControl

A control to display a list of buttons or boxes. This is used to select either a single value or multiple values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CascadingControlConfiguration" : CascadingControlConfiguration,
  "DisplayOptions" : ListControlDisplayOptions,
  "FilterControlId" : String,
  "SelectableValues" : FilterSelectableValues,
  "SourceFilterId" : String,
  "Title" : String,
  "Type" : String
}

```

### YAML

```yaml

  CascadingControlConfiguration:
    CascadingControlConfiguration
  DisplayOptions:
    ListControlDisplayOptions
  FilterControlId: String
  SelectableValues:
    FilterSelectableValues
  SourceFilterId: String
  Title: String
  Type: String

```

## Properties

`CascadingControlConfiguration`

The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.

_Required_: No

_Type_: [CascadingControlConfiguration](aws-properties-quicksight-analysis-cascadingcontrolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [ListControlDisplayOptions](aws-properties-quicksight-analysis-listcontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterControlId`

The ID of the `FilterListControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectableValues`

A list of selectable values that are used in a control.

_Required_: No

_Type_: [FilterSelectableValues](aws-properties-quicksight-analysis-filterselectablevalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFilterId`

The source filter ID of the `FilterListControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `FilterListControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the `FilterListControl`. Choose one of the following options:

- `MULTI_SELECT`: The user can select multiple entries from the list.

- `SINGLE_SELECT`: The user can select a single entry from the list.

_Required_: No

_Type_: String

_Allowed values_: `MULTI_SELECT | SINGLE_SELECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterListConfiguration

FilterOperationSelectedFieldsConfiguration

All content copied from https://docs.aws.amazon.com/.
