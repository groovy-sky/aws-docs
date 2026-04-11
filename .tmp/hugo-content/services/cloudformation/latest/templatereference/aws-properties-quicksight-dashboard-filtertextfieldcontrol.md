This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FilterTextFieldControl

A control to display a text box that is used to enter a single entry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisplayOptions" : TextFieldControlDisplayOptions,
  "FilterControlId" : String,
  "SourceFilterId" : String,
  "Title" : String
}

```

### YAML

```yaml

  DisplayOptions:
    TextFieldControlDisplayOptions
  FilterControlId: String
  SourceFilterId: String
  Title: String

```

## Properties

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [TextFieldControlDisplayOptions](aws-properties-quicksight-dashboard-textfieldcontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterControlId`

The ID of the `FilterTextFieldControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFilterId`

The source filter ID of the `FilterTextFieldControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `FilterTextFieldControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterTextAreaControl

FontConfiguration

All content copied from https://docs.aws.amazon.com/.
