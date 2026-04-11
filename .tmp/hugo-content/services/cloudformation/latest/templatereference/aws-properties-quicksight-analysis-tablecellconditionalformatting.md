This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TableCellConditionalFormatting

The cell conditional formatting option for a table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldId" : String,
  "TextFormat" : TextConditionalFormat
}

```

### YAML

```yaml

  FieldId: String
  TextFormat:
    TextConditionalFormat

```

## Properties

`FieldId`

The field ID of the cell for conditional formatting.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextFormat`

The text format of the cell for conditional formatting.

_Required_: No

_Type_: [TextConditionalFormat](aws-properties-quicksight-analysis-textconditionalformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableBorderOptions

TableCellImageSizingConfiguration

All content copied from https://docs.aws.amazon.com/.
