This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DefaultTextAreaControlOptions

The default options that correspond to the `TextArea` filter control type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Delimiter" : String,
  "DisplayOptions" : TextAreaControlDisplayOptions
}

```

### YAML

```yaml

  Delimiter: String
  DisplayOptions:
    TextAreaControlDisplayOptions

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

_Type_: [TextAreaControlDisplayOptions](aws-properties-quicksight-analysis-textareacontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultSliderControlOptions

DefaultTextFieldControlOptions

All content copied from https://docs.aws.amazon.com/.
