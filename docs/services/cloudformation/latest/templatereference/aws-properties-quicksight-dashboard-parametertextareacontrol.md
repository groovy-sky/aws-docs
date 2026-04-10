This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ParameterTextAreaControl

A control to display a text box that is used to enter multiple entries.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Delimiter" : String,
  "DisplayOptions" : TextAreaControlDisplayOptions,
  "ParameterControlId" : String,
  "SourceParameterName" : String,
  "Title" : String
}

```

### YAML

```yaml

  Delimiter: String
  DisplayOptions:
    TextAreaControlDisplayOptions
  ParameterControlId: String
  SourceParameterName: String
  Title: String

```

## Properties

`Delimiter`

The delimiter that is used to separate the lines in text.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [TextAreaControlDisplayOptions](aws-properties-quicksight-dashboard-textareacontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterControlId`

The ID of the `ParameterTextAreaControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceParameterName`

The source parameter name of the `ParameterTextAreaControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `ParameterTextAreaControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterSliderControl

ParameterTextFieldControl

All content copied from https://docs.aws.amazon.com/.
