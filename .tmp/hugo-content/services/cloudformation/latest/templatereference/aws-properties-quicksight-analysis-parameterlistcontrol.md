This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ParameterListControl

A control to display a list with buttons or boxes that are used to select either a single value or multiple values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CascadingControlConfiguration" : CascadingControlConfiguration,
  "DisplayOptions" : ListControlDisplayOptions,
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
  DisplayOptions:
    ListControlDisplayOptions
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

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [ListControlDisplayOptions](aws-properties-quicksight-analysis-listcontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterControlId`

The ID of the `ParameterListControl`.

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

The source parameter name of the `ParameterListControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `ParameterListControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of `ParameterListControl`.

_Required_: No

_Type_: String

_Allowed values_: `MULTI_SELECT | SINGLE_SELECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterDropDownControl

Parameters

All content copied from https://docs.aws.amazon.com/.
