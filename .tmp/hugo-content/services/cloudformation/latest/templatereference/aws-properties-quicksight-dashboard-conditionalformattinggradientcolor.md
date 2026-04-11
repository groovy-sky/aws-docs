This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ConditionalFormattingGradientColor

Formatting configuration for gradient color.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : GradientColor,
  "Expression" : String
}

```

### YAML

```yaml

  Color:
    GradientColor
  Expression: String

```

## Properties

`Color`

Determines the color.

_Required_: Yes

_Type_: [GradientColor](aws-properties-quicksight-dashboard-gradientcolor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The expression that determines the formatting configuration for gradient color.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionalFormattingCustomIconOptions

ConditionalFormattingIcon

All content copied from https://docs.aws.amazon.com/.
