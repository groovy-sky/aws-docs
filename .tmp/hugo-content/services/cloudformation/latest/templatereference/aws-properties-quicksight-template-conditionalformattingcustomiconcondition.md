This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ConditionalFormattingCustomIconCondition

Determines the custom condition for an icon set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "DisplayConfiguration" : ConditionalFormattingIconDisplayConfiguration,
  "Expression" : String,
  "IconOptions" : ConditionalFormattingCustomIconOptions
}

```

### YAML

```yaml

  Color: String
  DisplayConfiguration:
    ConditionalFormattingIconDisplayConfiguration
  Expression: String
  IconOptions:
    ConditionalFormattingCustomIconOptions

```

## Properties

`Color`

Determines the color of the icon.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayConfiguration`

Determines the icon display configuration.

_Required_: No

_Type_: [ConditionalFormattingIconDisplayConfiguration](aws-properties-quicksight-template-conditionalformattingicondisplayconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The expression that determines the condition of the icon set.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IconOptions`

Custom icon options for an icon set.

_Required_: Yes

_Type_: [ConditionalFormattingCustomIconOptions](aws-properties-quicksight-template-conditionalformattingcustomiconoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionalFormattingColor

ConditionalFormattingCustomIconOptions

All content copied from https://docs.aws.amazon.com/.
