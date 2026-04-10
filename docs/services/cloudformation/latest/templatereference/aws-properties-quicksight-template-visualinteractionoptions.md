This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template VisualInteractionOptions

The general visual interactions setup for visual publish options

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContextMenuOption" : ContextMenuOption,
  "VisualMenuOption" : VisualMenuOption
}

```

### YAML

```yaml

  ContextMenuOption:
    ContextMenuOption
  VisualMenuOption:
    VisualMenuOption

```

## Properties

`ContextMenuOption`

The context menu options for a visual.

_Required_: No

_Type_: [ContextMenuOption](aws-properties-quicksight-template-contextmenuoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualMenuOption`

The on-visual menu options for a visual.

_Required_: No

_Type_: [VisualMenuOption](aws-properties-quicksight-template-visualmenuoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisualCustomActionOperation

VisualMenuOption

All content copied from https://docs.aws.amazon.com/.
