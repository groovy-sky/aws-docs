This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FilledMapShapeConditionalFormatting

The conditional formatting that determines the shape of the filled map.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldId" : String,
  "Format" : ShapeConditionalFormat
}

```

### YAML

```yaml

  FieldId: String
  Format:
    ShapeConditionalFormat

```

## Properties

`FieldId`

The field ID of the filled map shape.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

The conditional formatting that determines the background color of a filled map's shape.

_Required_: No

_Type_: [ShapeConditionalFormat](aws-properties-quicksight-dashboard-shapeconditionalformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilledMapFieldWells

FilledMapSortConfiguration

All content copied from https://docs.aws.amazon.com/.
