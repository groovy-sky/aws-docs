This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ConditionalFormattingIcon

The formatting configuration for the icon.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomCondition" : ConditionalFormattingCustomIconCondition,
  "IconSet" : ConditionalFormattingIconSet
}

```

### YAML

```yaml

  CustomCondition:
    ConditionalFormattingCustomIconCondition
  IconSet:
    ConditionalFormattingIconSet

```

## Properties

`CustomCondition`

Determines the custom condition for an icon set.

_Required_: No

_Type_: [ConditionalFormattingCustomIconCondition](aws-properties-quicksight-analysis-conditionalformattingcustomiconcondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IconSet`

Formatting configuration for icon set.

_Required_: No

_Type_: [ConditionalFormattingIconSet](aws-properties-quicksight-analysis-conditionalformattingiconset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionalFormattingGradientColor

ConditionalFormattingIconDisplayConfiguration

All content copied from https://docs.aws.amazon.com/.
