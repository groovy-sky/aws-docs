This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ParameterDropDownControl

A control to display a dropdown list with buttons that are used to select a single value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CascadingControlConfiguration" : CascadingControlConfiguration,
  "CommitMode" : String,
  "DisplayOptions" : DropDownControlDisplayOptions,
  "ParameterControlId" : String,
  "SelectableValues" : ParameterSelectableValues,
  "SourceParameterName" : String,
  "Title" : String,
  "Type" : String
}

```

### YAML

```yaml

  CascadingControlConfiguration:
    CascadingControlConfiguration
  CommitMode: String
  DisplayOptions:
    DropDownControlDisplayOptions
  ParameterControlId: String
  SelectableValues:
    ParameterSelectableValues
  SourceParameterName: String
  Title: String
  Type: String

```

## Properties

`CascadingControlConfiguration`

The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.

_Required_: No

_Type_: [CascadingControlConfiguration](aws-properties-quicksight-analysis-cascadingcontrolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CommitMode`

The visibility configuration of the Apply button on a `ParameterDropDownControl`.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [DropDownControlDisplayOptions](aws-properties-quicksight-analysis-dropdowncontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterControlId`

The ID of the `ParameterDropDownControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectableValues`

A list of selectable values that are used in a control.

_Required_: No

_Type_: [ParameterSelectableValues](aws-properties-quicksight-analysis-parameterselectablevalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceParameterName`

The source parameter name of the `ParameterDropDownControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `ParameterDropDownControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type parameter name of the `ParameterDropDownControl`.

_Required_: No

_Type_: String

_Allowed values_: `MULTI_SELECT | SINGLE_SELECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterDeclaration

ParameterListControl

All content copied from https://docs.aws.amazon.com/.
