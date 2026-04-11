This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ConditionalFormattingIconSet

Formatting configuration for icon set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "IconSetType" : String
}

```

### YAML

```yaml

  Expression: String
  IconSetType: String

```

## Properties

`Expression`

The expression that determines the formatting configuration for the icon set.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IconSetType`

Determines the icon set type.

_Required_: No

_Type_: String

_Allowed values_: `PLUS_MINUS | CHECK_X | THREE_COLOR_ARROW | THREE_GRAY_ARROW | CARET_UP_MINUS_DOWN | THREE_SHAPE | THREE_CIRCLE | FLAGS | BARS | FOUR_COLOR_ARROW | FOUR_GRAY_ARROW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionalFormattingIconDisplayConfiguration

ConditionalFormattingSolidColor

All content copied from https://docs.aws.amazon.com/.
