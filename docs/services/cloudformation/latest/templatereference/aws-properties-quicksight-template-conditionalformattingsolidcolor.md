This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ConditionalFormattingSolidColor

Formatting configuration for solid color.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "Expression" : String
}

```

### YAML

```yaml

  Color: String
  Expression: String

```

## Properties

`Color`

Determines the color.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The expression that determines the formatting configuration for solid color.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionalFormattingIconSet

ContextMenuOption

All content copied from https://docs.aws.amazon.com/.
