This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis SheetImageTooltipConfiguration

The tooltip configuration for a sheet image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TooltipText" : SheetImageTooltipText,
  "Visibility" : String
}

```

### YAML

```yaml

  TooltipText:
    SheetImageTooltipText
  Visibility: String

```

## Properties

`TooltipText`

The text that appears in the tooltip.

_Required_: No

_Type_: [SheetImageTooltipText](aws-properties-quicksight-analysis-sheetimagetooltiptext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the tooltip.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SheetImageStaticFileSource

SheetImageTooltipText

All content copied from https://docs.aws.amazon.com/.
