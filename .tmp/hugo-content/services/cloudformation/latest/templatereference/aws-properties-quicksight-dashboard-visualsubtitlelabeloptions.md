This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard VisualSubtitleLabelOptions

The subtitle label options for a visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FormatText" : LongFormatText,
  "Visibility" : String
}

```

### YAML

```yaml

  FormatText:
    LongFormatText
  Visibility: String

```

## Properties

`FormatText`

The long text format of the subtitle label, such as plain text or rich text.

_Required_: No

_Type_: [LongFormatText](aws-properties-quicksight-dashboard-longformattext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the subtitle label.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisualPalette

VisualTitleLabelOptions

All content copied from https://docs.aws.amazon.com/.
