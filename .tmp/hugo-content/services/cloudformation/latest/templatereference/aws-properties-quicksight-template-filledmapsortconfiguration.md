This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FilledMapSortConfiguration

The sort configuration of a `FilledMapVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategorySort" : [ FieldSortOptions, ... ]
}

```

### YAML

```yaml

  CategorySort:
    - FieldSortOptions

```

## Properties

`CategorySort`

The sort configuration of the location fields.

_Required_: No

_Type_: Array of [FieldSortOptions](aws-properties-quicksight-template-fieldsortoptions.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilledMapShapeConditionalFormatting

FilledMapVisual

All content copied from https://docs.aws.amazon.com/.
