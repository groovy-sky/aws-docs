This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ConditionalFormattingCustomIconOptions

Custom icon options for an icon set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Icon" : String,
  "UnicodeIcon" : String
}

```

### YAML

```yaml

  Icon: String
  UnicodeIcon: String

```

## Properties

`Icon`

Determines the type of icon.

_Required_: No

_Type_: String

_Allowed values_: `CARET_UP | CARET_DOWN | PLUS | MINUS | ARROW_UP | ARROW_DOWN | ARROW_LEFT | ARROW_UP_LEFT | ARROW_DOWN_LEFT | ARROW_RIGHT | ARROW_UP_RIGHT | ARROW_DOWN_RIGHT | FACE_UP | FACE_DOWN | FACE_FLAT | ONE_BAR | TWO_BAR | THREE_BAR | CIRCLE | TRIANGLE | SQUARE | FLAG | THUMBS_UP | THUMBS_DOWN | CHECKMARK | X`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnicodeIcon`

Determines the Unicode icon type.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000-\u00FF]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionalFormattingCustomIconCondition

ConditionalFormattingGradientColor

All content copied from https://docs.aws.amazon.com/.
