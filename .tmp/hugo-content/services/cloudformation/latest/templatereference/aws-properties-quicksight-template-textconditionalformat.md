This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TextConditionalFormat

The conditional formatting for the text.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : ConditionalFormattingColor,
  "Icon" : ConditionalFormattingIcon,
  "TextColor" : ConditionalFormattingColor
}

```

### YAML

```yaml

  BackgroundColor:
    ConditionalFormattingColor
  Icon:
    ConditionalFormattingIcon
  TextColor:
    ConditionalFormattingColor

```

## Properties

`BackgroundColor`

The conditional formatting for the text background color.

_Required_: No

_Type_: [ConditionalFormattingColor](aws-properties-quicksight-template-conditionalformattingcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Icon`

The conditional formatting for the icon.

_Required_: No

_Type_: [ConditionalFormattingIcon](aws-properties-quicksight-template-conditionalformattingicon.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextColor`

The conditional formatting for the text color.

_Required_: No

_Type_: [ConditionalFormattingColor](aws-properties-quicksight-template-conditionalformattingcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextAreaControlDisplayOptions

TextControlPlaceholderOptions

All content copied from https://docs.aws.amazon.com/.
