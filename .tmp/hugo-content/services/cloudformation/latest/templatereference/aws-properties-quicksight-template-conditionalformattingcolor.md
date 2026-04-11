This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ConditionalFormattingColor

The formatting configuration for the color.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Gradient" : ConditionalFormattingGradientColor,
  "Solid" : ConditionalFormattingSolidColor
}

```

### YAML

```yaml

  Gradient:
    ConditionalFormattingGradientColor
  Solid:
    ConditionalFormattingSolidColor

```

## Properties

`Gradient`

Formatting configuration for gradient color.

_Required_: No

_Type_: [ConditionalFormattingGradientColor](aws-properties-quicksight-template-conditionalformattinggradientcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Solid`

Formatting configuration for solid color.

_Required_: No

_Type_: [ConditionalFormattingSolidColor](aws-properties-quicksight-template-conditionalformattingsolidcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Computation

ConditionalFormattingCustomIconCondition

All content copied from https://docs.aws.amazon.com/.
