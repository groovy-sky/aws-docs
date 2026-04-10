This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FontSize

The option that determines the text display size.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Absolute" : String,
  "Relative" : String
}

```

### YAML

```yaml

  Absolute: String
  Relative: String

```

## Properties

`Absolute`

The font size that you want to use in px.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Relative`

The lexical name for the text size, proportional to its surrounding context.

_Required_: No

_Type_: String

_Allowed values_: `EXTRA_SMALL | SMALL | MEDIUM | LARGE | EXTRA_LARGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FontConfiguration

FontWeight

All content copied from https://docs.aws.amazon.com/.
