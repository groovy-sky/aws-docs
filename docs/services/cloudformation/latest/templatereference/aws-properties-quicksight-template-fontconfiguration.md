This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FontConfiguration

Configures the display properties of the given text.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FontColor" : String,
  "FontDecoration" : String,
  "FontFamily" : String,
  "FontSize" : FontSize,
  "FontStyle" : String,
  "FontWeight" : FontWeight
}

```

### YAML

```yaml

  FontColor: String
  FontDecoration: String
  FontFamily: String
  FontSize:
    FontSize
  FontStyle: String
  FontWeight:
    FontWeight

```

## Properties

`FontColor`

Determines the color of the text.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontDecoration`

Determines the appearance of decorative lines on the text.

_Required_: No

_Type_: String

_Allowed values_: `UNDERLINE | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontFamily`

The font family that you want to use.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontSize`

The option that determines the text display size.

_Required_: No

_Type_: [FontSize](aws-properties-quicksight-template-fontsize.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontStyle`

Determines the text display face that is inherited by the given font family.

_Required_: No

_Type_: String

_Allowed values_: `NORMAL | ITALIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontWeight`

The option that determines the text display weight, or boldness.

_Required_: No

_Type_: [FontWeight](aws-properties-quicksight-template-fontweight.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterTextFieldControl

FontSize

All content copied from https://docs.aws.amazon.com/.
